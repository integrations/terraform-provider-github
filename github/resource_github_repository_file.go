package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryFile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryFileCreate,
		ReadContext:   resourceGithubRepositoryFileRead,
		UpdateContext: resourceGithubRepositoryFileUpdate,
		DeleteContext: resourceGithubRepositoryFileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryFileImport,
		},

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubRepositoryFileV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubRepositoryFileStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "This resource allows you to create and manage files within a GitHub repository.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository name",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The repository ID",
			},
			"file": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The file path to manage",
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The file's content",
			},
			"branch": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The branch name, defaults to the repository's default branch",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the commit/branch/tag",
				ForceNew:    true,
			},
			"commit_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SHA of the commit that modified the file",
			},
			"commit_message": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The commit message when creating, updating or deleting the file",
			},
			"commit_author": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     false,
				Description:  "The commit author name, defaults to the authenticated user's name. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. ",
				RequiredWith: []string{"commit_email"},
			},
			"commit_email": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     false,
				Description:  "The commit author email address, defaults to the authenticated user's email address. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.",
				RequiredWith: []string{"commit_author"},
			},
			"sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The blob SHA of the file",
			},
			"overwrite_on_create": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable overwriting existing files, defaults to \"false\"",
				Default:     false,
			},
			"autocreate_branch": {
				Type:             schema.TypeBool,
				Optional:         true,
				Description:      "Automatically create the branch if it could not be found. Subsequent reads if the branch is deleted will occur from 'autocreate_branch_source_branch'",
				Default:          false,
				DiffSuppressFunc: autoBranchDiffSuppressFunc,
				Deprecated:       "Use `github_branch` resource instead",
			},
			"autocreate_branch_source_branch": {
				Type:             schema.TypeString,
				Default:          "main",
				Optional:         true,
				Description:      "The branch name to start from, if 'autocreate_branch' is set. Defaults to 'main'.",
				RequiredWith:     []string{"autocreate_branch"},
				DiffSuppressFunc: autoBranchDiffSuppressFunc,
				Deprecated:       "Use `github_branch` resource instead",
			},
			"autocreate_branch_source_sha": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Description:      "The commit hash to start from, if 'autocreate_branch' is set. Defaults to the tip of 'autocreate_branch_source_branch'. If provided, 'autocreate_branch_source_branch' is ignored.",
				RequiredWith:     []string{"autocreate_branch"},
				DiffSuppressFunc: autoBranchDiffSuppressFunc,
				Deprecated:       "Use `github_branch` resource instead",
			},
		},
		CustomizeDiff: diffRepository,
	}
}

func resourceGithubRepositoryFileOptions(d *schema.ResourceData) *github.RepositoryContentFileOptions {
	opts := &github.RepositoryContentFileOptions{
		Content: []byte(d.Get("content").(string)),
	}

	if branch, ok := d.GetOk("branch"); ok {
		opts.Branch = github.Ptr(branch.(string))
	}

	if commitMessage, hasCommitMessage := d.GetOk("commit_message"); hasCommitMessage {
		opts.Message = github.Ptr(commitMessage.(string))
	}

	if SHA, hasSHA := d.GetOk("sha"); hasSHA {
		opts.SHA = github.Ptr(SHA.(string))
	}

	commitAuthor, hasCommitAuthor := d.GetOk("commit_author")
	commitEmail, hasCommitEmail := d.GetOk("commit_email")

	if hasCommitAuthor && hasCommitEmail {
		name := github.Ptr(commitAuthor.(string))
		mail := github.Ptr(commitEmail.(string))
		opts.Author = &github.CommitAuthor{Name: name, Email: mail}
		opts.Committer = &github.CommitAuthor{Name: name, Email: mail}
	}

	return opts
}

func resourceGithubRepositoryFileCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repo := d.Get("repository").(string)
	file := d.Get("file").(string)

	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "file", file)
	ctx = tflog.SetField(ctx, "owner", owner)

	checkOpt := github.RepositoryContentGetOptions{}

	repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return diag.FromErr(err)
	}
	var branch string

	if branchFieldVal, ok := d.GetOk("branch"); ok {
		branch = branchFieldVal.(string)
		tflog.Debug(ctx, "Using explicitly set branch", map[string]any{
			"branch": branch,
		})
		if err := checkRepositoryBranchExists(ctx, client, owner, repo, branch); err != nil {
			if d.Get("autocreate_branch").(bool) {
				if err := resourceGithubRepositoryFileCreateBranch(ctx, d, meta); err != nil {
					return diag.FromErr(err)
				}
			} else {
				return diag.FromErr(err)
			}
		}
		checkOpt.Ref = branch
	} else {
		branch = repoInfo.GetDefaultBranch()
	}

	opts := resourceGithubRepositoryFileOptions(d)

	if opts.Message == nil {
		opts.Message = github.Ptr(fmt.Sprintf("Add %s", file))
	}

	tflog.Debug(ctx, "Checking if overwriting a repository file")
	fileContent, _, resp, err := client.Repositories.GetContents(ctx, owner, repo, file, &checkOpt)
	if err != nil {
		if resp != nil {
			if resp.StatusCode != http.StatusNotFound {
				// 404 is a valid response if the file does not exist
				return diag.FromErr(err)
			}
		} else {
			// Response should be non-nil
			return diag.FromErr(err)
		}
	}

	if fileContent != nil {
		if d.Get("overwrite_on_create").(bool) {
			// Overwrite existing file if requested by configuring the options for
			// `client.Repositories.CreateFile` to match the existing file's SHA
			opts.SHA = fileContent.SHA
		} else {
			// Error if overwriting a file is not requested
			return diag.Errorf("refusing to overwrite existing file: configure `overwrite_on_create` to `true` to override")
		}
	}

	tflog.Debug(ctx, "Creating or overwriting a repository file")
	// Create a new or overwritten file
	create, _, err := client.Repositories.CreateFile(ctx, owner, repo, file, opts)
	if err != nil {
		return diag.FromErr(err)
	}

	newResourceID, err := buildID(repo, file, branch)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, "Setting ID", map[string]any{
		"id": newResourceID,
	})
	d.SetId(newResourceID)

	// Set computed values after the resource is created and in state
	if err = d.Set("commit_sha", create.GetSHA()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("branch", branch); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository_id", int(repoInfo.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryFileRead(ctx, d, meta)
}

func resourceGithubRepositoryFileRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	file := d.Get("file").(string)
	branch := d.Get("branch").(string)

	ctx = tflog.SetField(ctx, "repository", repoName)
	ctx = tflog.SetField(ctx, "file", file)
	ctx = tflog.SetField(ctx, "owner", owner)
	ctx = tflog.SetField(ctx, "owner", owner)

	opts := &github.RepositoryContentGetOptions{}

	opts.Ref = branch

	fc, _, _, err := client.Repositories.GetContents(ctx, owner, repoName, file, opts)
	if err != nil {
		var errorResponse *github.ErrorResponse
		if errors.As(err, &errorResponse) && errorResponse.Response.StatusCode == http.StatusTooManyRequests {
			return diag.FromErr(err)
		}
		return diag.FromErr(deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "repository file %s/%s:%s:%s", owner, repoName, file, branch))
	}
	if fc == nil {
		tflog.Info(ctx, "Removing repository path from state because it no longer exists in GitHub")
		d.SetId("")
		return nil
	}

	content, err := fc.GetContent()
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("content", content); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("sha", fc.GetSHA()); err != nil {
		return diag.FromErr(err)
	}

	var commit *github.RepositoryCommit

	parsedUrl, err := url.Parse(fc.GetURL())
	if err != nil {
		return diag.FromErr(err)
	}
	parsedQuery, err := url.ParseQuery(parsedUrl.RawQuery)
	if err != nil {
		return diag.FromErr(err)
	}
	ref := parsedQuery["ref"][0]
	if err = d.Set("ref", ref); err != nil {
		return diag.FromErr(err)
	}

	// Use the SHA to lookup the commit info if we know it, otherwise loop through commits
	if sha, ok := d.GetOk("commit_sha"); ok {
		tflog.Debug(ctx, "Using known commit SHA", map[string]any{
			"commit_sha": sha.(string),
		})
		commit, _, err = client.Repositories.GetCommit(ctx, owner, repoName, sha.(string), nil)
	} else {
		tflog.Debug(ctx, "Commit SHA unknown for file, looking for commit")
		commit, err = getFileCommit(ctx, client, owner, repoName, file, ref)
	}
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, "Found file in commit", map[string]any{
		"commit_sha": commit.GetSHA(),
	})

	if err = d.Set("commit_sha", commit.GetSHA()); err != nil {
		return diag.FromErr(err)
	}

	commit_author := commit.Commit.GetCommitter().GetName()
	commit_email := commit.Commit.GetCommitter().GetEmail()

	_, hasCommitAuthor := d.GetOk("commit_author")
	_, hasCommitEmail := d.GetOk("commit_email")

	// read from state if author+email is set explicitly, and if it was not github signing it for you previously
	if commit_author != "GitHub" && commit_email != "noreply@github.com" && hasCommitAuthor && hasCommitEmail {
		if err = d.Set("commit_author", commit_author); err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("commit_email", commit_email); err != nil {
			return diag.FromErr(err)
		}
	}
	if err = d.Set("commit_message", commit.GetCommit().GetMessage()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryFileUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repo := d.Get("repository").(string)
	file := d.Get("file").(string)
	branch := d.Get("branch").(string)

	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "file", file)
	ctx = tflog.SetField(ctx, "owner", owner)
	ctx = tflog.SetField(ctx, "branch", branch)

	opts := resourceGithubRepositoryFileOptions(d)

	if *opts.Message == fmt.Sprintf("Add %s", file) {
		opts.Message = github.Ptr(fmt.Sprintf("Update %s", file))
	}

	// TODO: Use UpdateFile if the file already exists
	create, _, err := client.Repositories.CreateFile(ctx, owner, repo, file, opts)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("commit_sha", create.GetSHA()); err != nil {
		return diag.FromErr(err)
	}

	if d.HasChanges("repository", "file", "branch") {
		newResourceID, err := buildID(repo, file, branch)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(newResourceID)
	}

	return resourceGithubRepositoryFileRead(ctx, d, meta)
}

func resourceGithubRepositoryFileDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repo := d.Get("repository").(string)
	file := d.Get("file").(string)

	opts := resourceGithubRepositoryFileOptions(d)

	if *opts.Message == fmt.Sprintf("Add %s", file) {
		opts.Message = github.Ptr(fmt.Sprintf("Delete %s", file))
	}

	branch := d.Get("branch").(string)
	opts.Branch = github.Ptr(branch)

	_, _, err := client.Repositories.DeleteFile(ctx, owner, repo, file, opts)
	return diag.FromErr(handleArchivedRepoDelete(err, "repository file", file, owner, repo))
}

func autoBranchDiffSuppressFunc(k, _, _ string, d *schema.ResourceData) bool {
	if !d.Get("autocreate_branch").(bool) {
		switch k {
		case "autocreate_branch", "autocreate_branch_source_branch", "autocreate_branch_source_sha":
			return true
		}
	}
	return false
}

func resourceGithubRepositoryFileImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	repo, filePath, branch, err := parseID3(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid ID specified. Supplied ID must be written as <repository>:<file path>: (when branch is default) or <repository>:<file path>:<branch>. %w", err)
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	opts := &github.RepositoryContentGetOptions{}

	repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	if branch == "" {
		branch = repoInfo.GetDefaultBranch()
	}

	opts.Ref = branch

	fc, _, _, err := client.Repositories.GetContents(ctx, owner, repo, filePath, opts)
	if err != nil {
		return nil, err
	}
	if fc == nil {
		return nil, fmt.Errorf("filePath %s is not a file in repository %s/%s or repository is not readable", filePath, owner, repo)
	}

	if err := d.Set("repository", repo); err != nil {
		return nil, err
	}
	if err := d.Set("file", filePath); err != nil {
		return nil, err
	}

	newResourceID, err := buildID(repo, filePath, branch)
	if err != nil {
		return nil, err
	}
	tflog.Debug(ctx, "Setting ID", map[string]any{
		"id": newResourceID,
	})
	d.SetId(newResourceID)

	if err := d.Set("branch", branch); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", int(repoInfo.GetID())); err != nil {
		return nil, err
	}
	if err = d.Set("overwrite_on_create", false); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func resourceGithubRepositoryFileCreateBranch(ctx context.Context, d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	branch := d.Get("branch").(string)
	repo := d.Get("repository").(string)

	branchRefName := "refs/heads/" + branch
	sourceBranchName := d.Get("autocreate_branch_source_branch").(string)
	sourceBranchRefName := "refs/heads/" + sourceBranchName

	if _, hasSourceSHA := d.GetOk("autocreate_branch_source_sha"); !hasSourceSHA {
		ref, _, err := client.Git.GetRef(ctx, owner, repo, sourceBranchRefName)
		if err != nil {
			return fmt.Errorf("error querying GitHub branch reference %s/%s (%s): %s",
				owner, repo, sourceBranchRefName, err.Error())
		}
		err = d.Set("autocreate_branch_source_sha", *ref.Object.SHA)
		if err != nil {
			return fmt.Errorf("error setting autocreate_branch_source_sha: %w", err)
		}
	}
	sourceBranchSHA := d.Get("autocreate_branch_source_sha").(string)
	branchRef := github.CreateRef{
		Ref: branchRefName,
		SHA: sourceBranchSHA,
	}
	if _, _, err := client.Git.CreateRef(ctx, owner, repo, branchRef); err != nil {
		return fmt.Errorf("error creating GitHub branch reference %s/%s (%s): %w",
			owner, repo, branchRefName, err)
	}
	return nil
}
