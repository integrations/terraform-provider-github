package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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
				ForceNew:    true,
				Description: "The repository name",
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
			},
			"autocreate_branch_source_branch": {
				Type:             schema.TypeString,
				Default:          "main",
				Optional:         true,
				Description:      "The branch name to start from, if 'autocreate_branch' is set. Defaults to 'main'.",
				RequiredWith:     []string{"autocreate_branch"},
				DiffSuppressFunc: autoBranchDiffSuppressFunc,
			},
			"autocreate_branch_source_sha": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Description:      "The commit hash to start from, if 'autocreate_branch' is set. Defaults to the tip of 'autocreate_branch_source_branch'. If provided, 'autocreate_branch_source_branch' is ignored.",
				RequiredWith:     []string{"autocreate_branch"},
				DiffSuppressFunc: autoBranchDiffSuppressFunc,
			},
		},
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

	if branch, ok := d.GetOk("branch"); ok {
		tflog.Debug(ctx, "Using explicitly set branch", map[string]any{
			"branch": branch.(string),
		})
		if err := checkRepositoryBranchExists(client, owner, repo, branch.(string)); err != nil {
			if d.Get("autocreate_branch").(bool) {
				if err := resourceGithubRepositoryFileCreateBranch(ctx, d, meta); err != nil {
					return diag.FromErr(err)
				}
			} else {
				return diag.FromErr(err)
			}
		}
		checkOpt.Ref = branch.(string)
	} else {
		repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("branch", repoInfo.GetDefaultBranch()); err != nil {
			return diag.FromErr(err)
		}
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

	branch := d.Get("branch").(string)

	d.SetId(fmt.Sprintf("%s/%s:%s", repo, file, branch))
	if err = d.Set("commit_sha", create.GetSHA()); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryFileRead(ctx, d, meta)
}

func resourceGithubRepositoryFileRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repo, file := splitRepoFilePath(d.Id())
	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "file", file)
	ctx = tflog.SetField(ctx, "owner", owner)

	opts := &github.RepositoryContentGetOptions{}

	if branch, ok := d.GetOk("branch"); ok {
		tflog.Debug(ctx, "Using explicitly set branch", map[string]any{
			"branch": branch.(string),
		})
		if err := checkRepositoryBranchExists(client, owner, repo, branch.(string)); err != nil {
			if d.Get("autocreate_branch").(bool) {
				branch = d.Get("autocreate_branch_source_branch").(string)
			} else {
				tflog.Info(ctx, "Removing repository path from state because the branch no longer exists in GitHub")
				d.SetId("")
				return nil
			}
		}
		opts.Ref = branch.(string)
	}

	fc, _, _, err := client.Repositories.GetContents(ctx, owner, repo, file, opts)
	if err != nil {
		var errorResponse *github.ErrorResponse
		if errors.As(err, &errorResponse) && errorResponse.Response.StatusCode == http.StatusTooManyRequests {
			return diag.FromErr(err)
		}
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
	if err = d.Set("repository", repo); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("file", file); err != nil {
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
		commit, _, err = client.Repositories.GetCommit(ctx, owner, repo, sha.(string), nil)
	} else {
		tflog.Debug(ctx, "Commit SHA unknown for file, looking for commit")
		commit, err = getFileCommit(ctx, client, owner, repo, file, ref)
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
	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "file", file)
	ctx = tflog.SetField(ctx, "owner", owner)

	if branch, ok := d.GetOk("branch"); ok {
		tflog.Debug(ctx, "Using explicitly set branch", map[string]any{
			"branch": branch.(string),
		})
		if err := checkRepositoryBranchExists(client, owner, repo, branch.(string)); err != nil {
			if d.Get("autocreate_branch").(bool) {
				if err := resourceGithubRepositoryFileCreateBranch(ctx, d, meta); err != nil {
					return diag.FromErr(err)
				}
			} else {
				return diag.FromErr(err)
			}
		}
	}

	opts := resourceGithubRepositoryFileOptions(d)

	if *opts.Message == fmt.Sprintf("Add %s", file) {
		opts.Message = github.Ptr(fmt.Sprintf("Update %s", file))
	}

	create, _, err := client.Repositories.CreateFile(ctx, owner, repo, file, opts)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("commit_sha", create.GetSHA()); err != nil {
		return diag.FromErr(err)
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

	if b, ok := d.GetOk("branch"); ok {
		tflog.Debug(ctx, "Using explicitly set branch", map[string]any{
			"branch": b.(string),
		})
		if err := checkRepositoryBranchExists(client, owner, repo, b.(string)); err != nil {
			if d.Get("autocreate_branch").(bool) {
				if err := resourceGithubRepositoryFileCreateBranch(ctx, d, meta); err != nil {
					return diag.FromErr(err)
				}
			} else {
				return diag.FromErr(err)
			}
		}
		branch := b.(string)
		opts.Branch = github.Ptr(branch)
	}

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
	importIDParts := strings.Split(d.Id(), idSeparator)

	if len(importIDParts) > 2 {
		return nil, fmt.Errorf("invalid ID specified. Supplied ID must be written as <repository>/<file path> (when branch is \"main\") or <repository>/<file path>:<branch>")
	}
	repoFilePath := importIDParts[0]

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repo, file := splitRepoFilePath(repoFilePath)

	opts := &github.RepositoryContentGetOptions{}

	if len(importIDParts) == 2 {
		branch := importIDParts[1]
		opts.Ref = branch
		if err := d.Set("branch", branch); err != nil {
			return nil, err
		}
	} else {
		repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
		if err != nil {
			return nil, err
		}
		defaultBranch := repoInfo.GetDefaultBranch()
		if err := d.Set("branch", defaultBranch); err != nil {
			return nil, err
		}
	}

	fc, _, _, err := client.Repositories.GetContents(ctx, owner, repo, file, opts)
	if err != nil {
		return nil, err
	}
	if fc == nil {
		return nil, fmt.Errorf("file %s is not a file in repository %s/%s or repository is not readable", file, owner, repo)
	}

	branch := d.Get("branch").(string)

	d.SetId(fmt.Sprintf("%s/%s:%s", repo, file, branch))
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
