package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	repositoryFilesBlobMode         = "100644"
	repositoryFilesBlobType         = "blob"
	repositoryFilesMaxRebaseRetries = 3
	repositoryFilesBlobConcurrency  = 8
)

func resourceGithubRepositoryFiles() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryFilesCreate,
		ReadContext:   resourceGithubRepositoryFilesRead,
		UpdateContext: resourceGithubRepositoryFilesUpdate,
		DeleteContext: resourceGithubRepositoryFilesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryFilesImport,
		},

		Description: "Manages a set of files within a GitHub repository, writing all changes in a single commit per apply via the Git Data API. Use this resource instead of multiple `github_repository_file` resources when you need atomic multi-file commits and want to avoid 409 conflicts caused by parallel per-file writes to the same branch.",

		CustomizeDiff: diffRepository,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository name. Renaming the repository in GitHub is detected via `repository_id` and treated as a rename (not a recreate).",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The repository ID. Used to distinguish a repository rename from a recreate.",
			},
			"branch": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The branch to commit to. Defaults to the repository's default branch.",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The fully-qualified ref (`refs/heads/<branch>`) that this resource commits to.",
			},
			"commit_message": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The commit message used when creating, updating, or deleting files. Auto-generated if empty.",
			},
			"commit_author": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The commit author name. Defaults to the authenticated user's name. GitHub App users may omit author and email so GitHub can verify commits as the GitHub App.",
				RequiredWith: []string{"commit_email"},
			},
			"commit_email": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The commit author email address. Defaults to the authenticated user's email address. GitHub App users may omit author and email so GitHub can verify commits as the GitHub App.",
				RequiredWith: []string{"commit_author"},
			},
			"file": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Description: "The set of files this resource manages. Each block contributes one entry to the commit's tree.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The path of the file in the repository, relative to the repo root.",
						},
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The file's content. Stored verbatim; encoded as base64 when sent to the GitHub API so any byte sequence is supported.",
						},
						"sha": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The blob SHA of the file's current content on the branch.",
						},
					},
				},
				Set: func(v any) int {
					m := v.(map[string]any)
					return schema.HashString(m["path"].(string))
				},
			},
			"commit_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SHA of the most recent commit created by this resource.",
			},
			"tree_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The tree SHA of the most recent commit created by this resource.",
			},
		},
	}
}

// computeDesiredFiles returns a map of path -> content from the file set in d,
// erroring on duplicate paths.
func computeDesiredFiles(d *schema.ResourceData) (map[string]string, error) {
	raw := d.Get("file").(*schema.Set).List()
	out := make(map[string]string, len(raw))
	for _, item := range raw {
		m := item.(map[string]any)
		path := m["path"].(string)
		content := m["content"].(string)
		if _, dup := out[path]; dup {
			return nil, fmt.Errorf("duplicate file path %q in file blocks; each path must appear only once", path)
		}
		out[path] = content
	}
	return out, nil
}

// computeOldFiles returns the prior file set as a map of path -> content from d.GetChange.
func computeOldFiles(d *schema.ResourceData) map[string]string {
	old, _ := d.GetChange("file")
	set, ok := old.(*schema.Set)
	if !ok {
		return map[string]string{}
	}
	items := set.List()
	out := make(map[string]string, len(items))
	for _, item := range items {
		m := item.(map[string]any)
		out[m["path"].(string)] = m["content"].(string)
	}
	return out
}

// commitAuthorFrom returns a CommitAuthor for the resource, or nil if not set.
func commitAuthorFrom(d *schema.ResourceData) *github.CommitAuthor {
	authorRaw, hasAuthor := d.GetOk("commit_author")
	emailRaw, hasEmail := d.GetOk("commit_email")
	if !hasAuthor || !hasEmail {
		return nil
	}
	name := authorRaw.(string)
	email := emailRaw.(string)
	return &github.CommitAuthor{Name: &name, Email: &email}
}

// defaultBatchCommitMessage returns a generated commit message describing the
// scope of changes in this commit.
func defaultBatchCommitMessage(addCount, updateCount, deleteCount int) string {
	if addCount == 0 && updateCount == 0 && deleteCount == 0 {
		return "Terraform: managed repository files"
	}
	parts := []string{}
	if addCount > 0 {
		parts = append(parts, fmt.Sprintf("%d added", addCount))
	}
	if updateCount > 0 {
		parts = append(parts, fmt.Sprintf("%d updated", updateCount))
	}
	if deleteCount > 0 {
		parts = append(parts, fmt.Sprintf("%d removed", deleteCount))
	}
	return "Terraform: " + strings.Join(parts, ", ")
}

// isRefUpdateConflict reports whether err is a 409/422 from UpdateRef, indicating
// the branch advanced between read and write and we should rebase and retry.
func isRefUpdateConflict(err error) bool {
	var ghErr *github.ErrorResponse
	if !errors.As(err, &ghErr) || ghErr.Response == nil {
		return false
	}
	switch ghErr.Response.StatusCode {
	case http.StatusConflict, http.StatusUnprocessableEntity:
		return true
	}
	return false
}

// createBlobsParallel creates blobs concurrently (bounded by
// repositoryFilesBlobConcurrency) and populates out with path -> blob SHA.
// Returns the first error encountered.
func createBlobsParallel(
	ctx context.Context,
	client *github.Client,
	owner, repo string,
	paths []string,
	desiredFiles map[string]string,
) (map[string]string, error) {
	out := make(map[string]string, len(paths))
	if len(paths) == 0 {
		return out, nil
	}
	type result struct {
		path string
		sha  string
		err  error
	}
	sem := make(chan struct{}, repositoryFilesBlobConcurrency)
	results := make(chan result, len(paths))
	var wg sync.WaitGroup
	for _, p := range paths {
		path := p
		wg.Go(func() {
			sem <- struct{}{}
			defer func() { <-sem }()
			encoded := base64.StdEncoding.EncodeToString([]byte(desiredFiles[path]))
			enc := "base64"
			blob, _, err := client.Git.CreateBlob(ctx, owner, repo, github.Blob{
				Content:  &encoded,
				Encoding: &enc,
			})
			if err != nil {
				results <- result{path: path, err: err}
				return
			}
			results <- result{path: path, sha: blob.GetSHA()}
		})
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	var firstErr error
	for r := range results {
		if r.err != nil {
			if firstErr == nil {
				firstErr = fmt.Errorf("failed to create blob for %s: %w", r.path, r.err)
			}
			continue
		}
		out[r.path] = r.sha
	}
	return out, firstErr
}

// commitFiles writes a single commit to branch containing the desired file
// changes and deletions, with a bounded rebase-retry loop on 409/422 from
// UpdateRef. Blob creation happens once and is reused across retries because
// blobs are content-addressed and parent-independent.
func commitFiles(
	ctx context.Context,
	client *github.Client,
	owner, repo, branch string,
	desiredFiles map[string]string,
	deletePaths []string,
	message string,
	author *github.CommitAuthor,
) (commitSHA, treeSHA string, err error) {
	desiredPaths := make([]string, 0, len(desiredFiles))
	for p := range desiredFiles {
		desiredPaths = append(desiredPaths, p)
	}
	sort.Strings(desiredPaths)
	sortedDeletes := append([]string(nil), deletePaths...)
	sort.Strings(sortedDeletes)

	blobSHAs, err := createBlobsParallel(ctx, client, owner, repo, desiredPaths, desiredFiles)
	if err != nil {
		return "", "", err
	}

	branchRef := "heads/" + branch
	backoff := 250 * time.Millisecond
	var lastErr error
	for attempt := 0; attempt <= repositoryFilesMaxRebaseRetries; attempt++ {
		if attempt > 0 {
			tflog.Debug(ctx, "Retrying commit after branch advanced", map[string]any{
				"attempt": attempt,
				"backoff": backoff.String(),
			})
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return "", "", ctx.Err()
			}
			backoff *= 2
		}

		baseRef, _, err := client.Git.GetRef(ctx, owner, repo, branchRef)
		if err != nil {
			return "", "", fmt.Errorf("failed to get branch %s ref: %w", branch, err)
		}
		baseCommitSHA := baseRef.GetObject().GetSHA()

		baseCommit, _, err := client.Git.GetCommit(ctx, owner, repo, baseCommitSHA)
		if err != nil {
			return "", "", fmt.Errorf("failed to get base commit %s: %w", baseCommitSHA, err)
		}
		baseTreeSHA := baseCommit.GetTree().GetSHA()

		entries := make([]*github.TreeEntry, 0, len(desiredPaths)+len(sortedDeletes))
		blobMode := repositoryFilesBlobMode
		blobType := repositoryFilesBlobType
		for _, p := range desiredPaths {
			path := p
			sha := blobSHAs[p]
			entries = append(entries, &github.TreeEntry{
				Path: &path,
				Mode: &blobMode,
				Type: &blobType,
				SHA:  &sha,
			})
		}
		for _, p := range sortedDeletes {
			path := p
			entries = append(entries, &github.TreeEntry{
				Path: &path,
				Mode: &blobMode,
				Type: &blobType,
				SHA:  nil,
			})
		}

		newTree, _, err := client.Git.CreateTree(ctx, owner, repo, baseTreeSHA, entries)
		if err != nil {
			return "", "", fmt.Errorf("failed to create tree: %w", err)
		}

		commitMessage := message
		commitInput := github.Commit{
			Message: &commitMessage,
			Tree:    &github.Tree{SHA: newTree.SHA},
			Parents: []*github.Commit{{SHA: &baseCommitSHA}},
		}
		if author != nil {
			commitInput.Author = author
			commitInput.Committer = author
		}
		newCommit, _, err := client.Git.CreateCommit(ctx, owner, repo, commitInput, nil)
		if err != nil {
			return "", "", fmt.Errorf("failed to create commit: %w", err)
		}

		_, _, err = client.Git.UpdateRef(ctx, owner, repo, branchRef, github.UpdateRef{
			SHA: newCommit.GetSHA(),
		})
		if err == nil {
			return newCommit.GetSHA(), newTree.GetSHA(), nil
		}
		lastErr = err
		if !isRefUpdateConflict(err) {
			return "", "", fmt.Errorf("failed to update ref %s: %w", branchRef, err)
		}
		tflog.Info(ctx, "Branch advanced during commit; will rebase and retry", map[string]any{
			"branch": branch,
		})
	}
	return "", "", fmt.Errorf("exceeded %d rebase retries for branch %s: %w",
		repositoryFilesMaxRebaseRetries, branch, lastErr)
}

func resourceGithubRepositoryFilesCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)

	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "owner", owner)

	repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return diag.FromErr(err)
	}

	branch, ok := d.GetOk("branch")
	branchName := repoInfo.GetDefaultBranch()
	if ok {
		branchName = branch.(string)
		if err := checkRepositoryBranchExists(ctx, client, owner, repo, branchName); err != nil {
			return diag.FromErr(err)
		}
	}
	ctx = tflog.SetField(ctx, "branch", branchName)

	desiredFiles, err := computeDesiredFiles(d)
	if err != nil {
		return diag.FromErr(err)
	}

	msg := d.Get("commit_message").(string)
	if msg == "" {
		msg = defaultBatchCommitMessage(len(desiredFiles), 0, 0)
	}

	commitSHA, treeSHA, err := commitFiles(
		ctx, client, owner, repo, branchName,
		desiredFiles, nil,
		msg, commitAuthorFrom(d),
	)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repo, branchName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("branch", branchName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository_id", int(repoInfo.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("commit_message", msg); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("commit_sha", commitSHA); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tree_sha", treeSHA); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryFilesRead(ctx, d, meta)
}

func resourceGithubRepositoryFilesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)
	branch := d.Get("branch").(string)

	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "owner", owner)
	ctx = tflog.SetField(ctx, "branch", branch)

	repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return diag.FromErr(deleteResourceOn404AndSwallow304OtherwiseReturnError(
			err, d, "repository %s/%s", owner, repo))
	}

	branchRef := "heads/" + branch
	refObj, _, err := client.Git.GetRef(ctx, owner, repo, branchRef)
	if err != nil {
		return diag.FromErr(deleteResourceOn404AndSwallow304OtherwiseReturnError(
			err, d, "repository files %s/%s:%s", owner, repo, branch))
	}
	currentCommitSHA := refObj.GetObject().GetSHA()

	desiredFiles, err := computeDesiredFiles(d)
	if err != nil {
		return diag.FromErr(err)
	}
	managedPaths := make([]string, 0, len(desiredFiles))
	for p := range desiredFiles {
		managedPaths = append(managedPaths, p)
	}
	sort.Strings(managedPaths)

	currentCommit, _, err := client.Git.GetCommit(ctx, owner, repo, currentCommitSHA)
	if err != nil {
		return diag.FromErr(err)
	}
	currentTreeSHA := currentCommit.GetTree().GetSHA()

	tree, _, err := client.Git.GetTree(ctx, owner, repo, currentTreeSHA, true)
	if err != nil {
		return diag.FromErr(err)
	}

	pathToSHA := make(map[string]string, len(tree.Entries))
	for _, e := range tree.Entries {
		if e.GetType() != repositoryFilesBlobType {
			continue
		}
		pathToSHA[e.GetPath()] = e.GetSHA()
	}

	priorSHAs := make(map[string]string)
	if priorSet, ok := d.Get("file").(*schema.Set); ok {
		for _, item := range priorSet.List() {
			m, ok := item.(map[string]any)
			if !ok {
				continue
			}
			if shaVal, ok := m["sha"].(string); ok && shaVal != "" {
				priorSHAs[m["path"].(string)] = shaVal
			}
		}
	}

	newFileSet := make([]any, 0, len(managedPaths))
	for _, p := range managedPaths {
		sha, present := pathToSHA[p]
		content := desiredFiles[p]
		if !present {
			// File missing from remote: signal drift via empty content/sha.
			content = ""
			sha = ""
		} else if priorSHAs[p] != sha {
			raw, _, blobErr := client.Git.GetBlobRaw(ctx, owner, repo, sha)
			if blobErr != nil {
				return diag.FromErr(blobErr)
			}
			content = string(raw)
		}
		newFileSet = append(newFileSet, map[string]any{
			"path":    p,
			"content": content,
			"sha":     sha,
		})
	}

	if err := d.Set("file", newFileSet); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tree_sha", currentTreeSHA); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("commit_sha", currentCommitSHA); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository", repo); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository_id", int(repoInfo.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("branch", branch); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ref", "refs/heads/"+branch); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryFilesUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)
	branch := d.Get("branch").(string)

	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "owner", owner)
	ctx = tflog.SetField(ctx, "branch", branch)

	desiredFiles, err := computeDesiredFiles(d)
	if err != nil {
		return diag.FromErr(err)
	}
	oldFiles := computeOldFiles(d)

	uploads := make(map[string]string)
	adds, modifies := 0, 0
	for p, c := range desiredFiles {
		oldContent, present := oldFiles[p]
		switch {
		case !present:
			uploads[p] = c
			adds++
		case oldContent != c:
			uploads[p] = c
			modifies++
		}
	}
	var deletes []string
	for p := range oldFiles {
		if _, kept := desiredFiles[p]; !kept {
			deletes = append(deletes, p)
		}
	}

	if len(uploads) == 0 && len(deletes) == 0 {
		// No file changes; e.g. only commit_author/commit_message edited.
		return resourceGithubRepositoryFilesRead(ctx, d, meta)
	}

	msg := d.Get("commit_message").(string)
	if !d.HasChange("commit_message") || msg == "" {
		msg = defaultBatchCommitMessage(adds, modifies, len(deletes))
	}

	commitSHA, treeSHA, err := commitFiles(
		ctx, client, owner, repo, branch,
		uploads, deletes,
		msg, commitAuthorFrom(d),
	)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("commit_message", msg); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("commit_sha", commitSHA); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tree_sha", treeSHA); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryFilesRead(ctx, d, meta)
}

func resourceGithubRepositoryFilesDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)
	branch := d.Get("branch").(string)

	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "owner", owner)
	ctx = tflog.SetField(ctx, "branch", branch)

	desiredFiles, err := computeDesiredFiles(d)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(desiredFiles) == 0 {
		return nil
	}
	deletes := make([]string, 0, len(desiredFiles))
	for p := range desiredFiles {
		deletes = append(deletes, p)
	}

	msg := fmt.Sprintf("Terraform: removed %d files", len(deletes))

	_, _, err = commitFiles(
		ctx, client, owner, repo, branch,
		nil, deletes,
		msg, commitAuthorFrom(d),
	)
	if err != nil {
		return diag.FromErr(handleArchivedRepoDelete(err, "repository files", strings.Join(deletes, ","), owner, repo))
	}
	return nil
}

func resourceGithubRepositoryFilesImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	parts := strings.SplitN(d.Id(), ":", 2)
	if len(parts) == 0 || parts[0] == "" {
		return nil, fmt.Errorf("invalid import ID %q; expected <repository> or <repository>:<branch>", d.Id())
	}
	repo := parts[0]
	var branch string
	if len(parts) == 2 {
		branch = parts[1]
	}
	if branch == "" {
		repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
		if err != nil {
			return nil, err
		}
		branch = repoInfo.GetDefaultBranch()
	}
	if err := checkRepositoryBranchExists(ctx, client, owner, repo, branch); err != nil {
		return nil, err
	}

	id, err := buildID(repo, branch)
	if err != nil {
		return nil, err
	}
	d.SetId(id)
	if err := d.Set("repository", repo); err != nil {
		return nil, err
	}
	if err := d.Set("branch", branch); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
