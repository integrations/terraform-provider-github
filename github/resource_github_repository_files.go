package github

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"slices"
	"strings"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

		CustomizeDiff: customdiff.All(diffRepository, diffRepositoryFiles),

		Description: "Manages a set of files on a branch of a GitHub repository, committing all changes together in a single commit.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository name.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The repository ID.",
			},
			"branch": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The branch to commit to. Defaults to the repository's default branch. The branch must already exist.",
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
				Description: "The commit message used when creating, updating, or deleting files. Generated from the number of files added, updated, and removed if not set.",
			},
			"commit_author": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"commit_email"},
				Description:  "The commit author name. Defaults to the authenticated user's name. GitHub App users may omit author and email so GitHub can verify commits as the GitHub App.",
			},
			"commit_email": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"commit_author"},
				Description:  "The commit author email address. Defaults to the authenticated user's email address. GitHub App users may omit author and email so GitHub can verify commits as the GitHub App.",
			},
			"file": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Description: "The set of files this resource manages. Files in the repository that are not listed here are never modified.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The path of the file in the repository, relative to the repository root.",
						},
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The file's content.",
						},
						"sha": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The blob SHA of the file's current content on the branch.",
						},
					},
				},
			},
			"commit_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SHA of the most recent commit created by this resource, or the branch head when imported.",
			},
			"tree_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The tree SHA of the most recent commit created by this resource, or the branch head when imported.",
			},
		},
	}
}

func diffRepositoryFiles(_ context.Context, diff *schema.ResourceDiff, _ any) error {
	_, err := expandRepositoryFiles(diff.Get("file"))
	return err
}

func expandRepositoryFiles(v any) (map[string]string, error) {
	set, _ := v.(*schema.Set)
	files := make(map[string]string, set.Len())
	for _, item := range set.List() {
		file, _ := item.(map[string]any)
		path, _ := file["path"].(string)
		if _, duplicate := files[path]; duplicate {
			return nil, fmt.Errorf("file path %q is declared more than once", path)
		}
		content, _ := file["content"].(string)
		files[path] = content
	}
	return files, nil
}

func flattenRepositoryFiles(files, shas map[string]string) []any {
	flattened := make([]any, 0, len(files))
	for path, content := range files {
		flattened = append(flattened, map[string]any{"path": path, "content": content, "sha": shas[path]})
	}
	return flattened
}

func expandRepositoryFilesAuthor(d *schema.ResourceData) *github.CommitAuthor {
	name, hasName := d.GetOk("commit_author")
	email, hasEmail := d.GetOk("commit_email")
	if !hasName || !hasEmail {
		return nil
	}
	authorName, _ := name.(string)
	authorEmail, _ := email.(string)
	return &github.CommitAuthor{Name: new(authorName), Email: new(authorEmail)}
}

func repositoryFilesCommitMessage(d *schema.ResourceData, added, updated, removed int) string {
	message, _ := d.Get("commit_message").(string)
	if config := d.GetRawConfig(); message != "" && !config.IsNull() && !config.GetAttr("commit_message").IsNull() {
		return message
	}

	var changes []string
	for _, change := range []struct {
		count int
		verb  string
	}{{added, "added"}, {updated, "updated"}, {removed, "removed"}} {
		if change.count > 0 {
			changes = append(changes, fmt.Sprintf("%d %s", change.count, change.verb))
		}
	}
	return "Terraform: " + strings.Join(changes, ", ")
}

func setRepositoryFilesCommit(d *schema.ResourceData, commitSHA, treeSHA string, files []any) error {
	for key, value := range map[string]any{
		"commit_sha": commitSHA,
		"tree_sha":   treeSHA,
		"file":       files,
	} {
		if err := d.Set(key, value); err != nil {
			return err
		}
	}
	return nil
}

func resourceGithubRepositoryFilesCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	branch, _ := d.Get("branch").(string)

	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	if branch == "" {
		branch = repository.GetDefaultBranch()
	}

	files, err := expandRepositoryFiles(d.Get("file"))
	if err != nil {
		return diag.FromErr(err)
	}
	message := repositoryFilesCommitMessage(d, len(files), 0, 0)

	tflog.Debug(ctx, "Committing repository files", map[string]any{"owner": owner, "repository": repoName, "branch": branch, "files": len(files)})

	commitSHA, treeSHA, err := commitRepositoryFiles(ctx, client, owner, repoName, branch, message, expandRepositoryFilesAuthor(d), files, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	shas, err := getRepositoryBlobSHAs(ctx, client, owner, repoName, treeSHA)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, branch, commitSHA)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	for key, value := range map[string]any{
		"repository_id":  int(repository.GetID()),
		"branch":         branch,
		"ref":            "refs/heads/" + branch,
		"commit_message": message,
	} {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}
	return diag.FromErr(setRepositoryFilesCommit(d, commitSHA, treeSHA, flattenRepositoryFiles(files, shas)))
}

func resourceGithubRepositoryFilesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	branch, _ := d.Get("branch").(string)

	head, _, err := client.Repositories.GetBranch(ctx, owner, repoName, branch, 0)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing repository files from state because the branch no longer exists in GitHub", map[string]any{"owner": owner, "repository": repoName, "branch": branch})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	shas, err := getRepositoryBlobSHAs(ctx, client, owner, repoName, head.GetCommit().GetCommit().GetTree().GetSHA())
	if err != nil {
		return diag.FromErr(err)
	}

	state, _ := d.Get("file").(*schema.Set)
	files := make([]any, 0, state.Len())
	for _, item := range state.List() {
		file, _ := item.(map[string]any)
		path, _ := file["path"].(string)
		sha, exists := shas[path]
		if !exists {
			tflog.Debug(ctx, "Managed file no longer exists on the branch", map[string]any{"path": path})
			continue
		}
		if stateSHA, _ := file["sha"].(string); stateSHA != sha {
			content, _, err := client.Git.GetBlobRaw(ctx, owner, repoName, sha)
			if err != nil {
				return diag.FromErr(err)
			}
			file["content"] = string(content)
			file["sha"] = sha
		}
		files = append(files, file)
	}

	return diag.FromErr(d.Set("file", files))
}

func resourceGithubRepositoryFilesUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	branch, _ := d.Get("branch").(string)

	previous, desired := d.GetChange("file")
	previousFiles, err := expandRepositoryFiles(previous)
	if err != nil {
		return diag.FromErr(err)
	}
	files, err := expandRepositoryFiles(desired)
	if err != nil {
		return diag.FromErr(err)
	}

	upserts := make(map[string]string)
	added := 0
	for path, content := range files {
		previousContent, exists := previousFiles[path]
		if exists && previousContent == content {
			continue
		}
		upserts[path] = content
		if !exists {
			added++
		}
	}
	var deletes []string
	for path := range previousFiles {
		if _, kept := files[path]; !kept {
			deletes = append(deletes, path)
		}
	}

	if len(upserts) > 0 || len(deletes) > 0 {
		message := repositoryFilesCommitMessage(d, added, len(upserts)-added, len(deletes))

		tflog.Debug(ctx, "Committing repository files", map[string]any{"owner": owner, "repository": repoName, "branch": branch, "upserts": len(upserts), "deletes": len(deletes)})

		commitSHA, treeSHA, err := commitRepositoryFiles(ctx, client, owner, repoName, branch, message, expandRepositoryFilesAuthor(d), upserts, deletes)
		if err != nil {
			return diag.FromErr(err)
		}

		shas, err := getRepositoryBlobSHAs(ctx, client, owner, repoName, treeSHA)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("commit_message", message); err != nil {
			return diag.FromErr(err)
		}
		if err := setRepositoryFilesCommit(d, commitSHA, treeSHA, flattenRepositoryFiles(files, shas)); err != nil {
			return diag.FromErr(err)
		}
	}

	commitSHA, _ := d.Get("commit_sha").(string)
	id, err := buildID(repoName, branch, commitSHA)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return nil
}

func resourceGithubRepositoryFilesDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	branch, _ := d.Get("branch").(string)

	files, err := expandRepositoryFiles(d.Get("file"))
	if err != nil {
		return diag.FromErr(err)
	}
	deletes := slices.Sorted(maps.Keys(files))
	if len(deletes) == 0 {
		return nil
	}

	tflog.Debug(ctx, "Deleting repository files", map[string]any{"owner": owner, "repository": repoName, "branch": branch, "deletes": len(deletes)})

	_, _, err = commitRepositoryFiles(ctx, client, owner, repoName, branch, repositoryFilesCommitMessage(d, 0, 0, len(deletes)), expandRepositoryFilesAuthor(d), nil, deletes)
	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
		tflog.Info(ctx, "Repository files already removed because the branch no longer exists in GitHub", map[string]any{"owner": owner, "repository": repoName, "branch": branch})
		return nil
	}
	return diag.FromErr(handleArchivedRepoDelete(err, "repository files", strings.Join(deletes, ","), owner, repoName))
}

func resourceGithubRepositoryFilesImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, branch, _ := strings.Cut(d.Id(), ":")

	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	if branch == "" {
		branch = repository.GetDefaultBranch()
	}

	head, _, err := client.Repositories.GetBranch(ctx, owner, repoName, branch, 0)
	if err != nil {
		return nil, err
	}

	id, err := buildID(repoName, branch, head.GetCommit().GetSHA())
	if err != nil {
		return nil, err
	}
	d.SetId(id)

	for key, value := range map[string]any{
		"repository":    repoName,
		"repository_id": int(repository.GetID()),
		"branch":        branch,
		"ref":           "refs/heads/" + branch,
		"commit_sha":    head.GetCommit().GetSHA(),
		"tree_sha":      head.GetCommit().GetCommit().GetTree().GetSHA(),
	} {
		if err := d.Set(key, value); err != nil {
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}
