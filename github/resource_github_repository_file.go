package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/v92/github"
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

		CustomizeDiff: resourceGithubRepositoryFileDiff,

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubRepositoryFileV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubRepositoryFileStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "Resource to manage a file within a GitHub repository.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the repository.",
			},
			"branch": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "Name of the branch; this defaults to the repository's default branch. The branch must already exist, unless `autocreate_branch` is set.",
			},
			"file": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Path of the file to manage.",
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Content of the file.",
			},
			"commit_message": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true, // TODO: Remove Computed in next major version.
				Description: "Commit message used when creating, updating or deleting the file; this defaults to `Add <file>`, `Update <file>` or `Delete <file>` respectively. It is only used when the file is written and is not read back from GitHub, so changing it on its own will not update the file.",
			},
			"commit_author": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Name of the commit author; this defaults to the authenticated user's name and must be set together with `commit_email`. It is only used when the file is written and is not read back from GitHub. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App, which may be useful when a branch protection rule requires signed commits.",
				RequiredWith: []string{"commit_email"},
			},
			"commit_email": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Email address of the commit author; this defaults to the authenticated user's email address and must be set together with `commit_author`. It is only used when the file is written and is not read back from GitHub. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App, which may be useful when a branch protection rule requires signed commits.",
				RequiredWith: []string{"commit_author"},
			},
			"overwrite_on_create": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable overwriting existing files; this defaults to `false`.",
			},
			"autocreate_branch": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				DiffSuppressFunc: autoBranchDiffSuppressFunc,
				Deprecated:       "Use `github_branch` resource instead",
				Description:      "Automatically create the branch if it could not be found. Subsequent reads if the branch is deleted will occur from `autocreate_branch_source_branch`.",
			},
			"autocreate_branch_source_branch": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "main",
				RequiredWith:     []string{"autocreate_branch"},
				DiffSuppressFunc: autoBranchDiffSuppressFunc,
				Deprecated:       "Use `github_branch` resource instead",
				Description:      "The branch name to start from, if `autocreate_branch` is set; this defaults to `main`.",
			},
			"autocreate_branch_source_sha": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				RequiredWith:     []string{"autocreate_branch"},
				DiffSuppressFunc: autoBranchDiffSuppressFunc,
				Deprecated:       "Use `github_branch` resource instead",
				Description:      "The commit hash to start from, if `autocreate_branch` is set; this defaults to the tip of `autocreate_branch_source_branch`. If provided, `autocreate_branch_source_branch` is ignored.",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the commit, branch or tag.",
			},
			"sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Blob SHA of the file.",
			},
			"commit_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SHA of the latest commit that modified the file; this is set when the file is written and refreshed if the file is changed outside of Terraform.",
			},
		},
	}
}

// resourceGithubRepositoryFileDiff customises the diff for the resource. It first checks the repository, then suppresses
// the diff when the commit message is the only change as it's only used when writing the file, so changing it alone must
// not trigger an update with identical content.
func resourceGithubRepositoryFileDiff(ctx context.Context, diff *schema.ResourceDiff, m any) error {
	if err := diffRepository(ctx, diff, m); err != nil {
		return err
	}

	if diff.Id() == "" || !diff.HasChange("commit_message") {
		return nil
	}

	if len(diff.GetChangedKeysPrefix("")) != 1 {
		return nil
	}

	return diff.Clear("commit_message")
}

func resourceGithubRepositoryFileCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repo, _ := d.Get("repository").(string)
	file, _ := d.Get("file").(string)

	ctx = tflog.SetField(ctx, "owner", owner)
	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "file", file)

	repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return diag.FromErr(err)
	}

	branch, _ := d.Get("branch").(string)
	if branch != "" {
		tflog.Debug(ctx, "Using explicitly set branch", map[string]any{"branch": branch})
		if err := checkRepositoryBranchExists(ctx, client, owner, repo, branch); err != nil {
			if autocreate, _ := d.Get("autocreate_branch").(bool); !autocreate {
				return diag.FromErr(err)
			}

			branchRefName := "refs/heads/" + branch
			sourceBranchName, _ := d.Get("autocreate_branch_source_branch").(string)
			sourceBranchRefName := "refs/heads/" + sourceBranchName

			if _, hasSourceSHA := d.GetOk("autocreate_branch_source_sha"); !hasSourceSHA {
				ref, _, err := client.Git.GetRef(ctx, owner, repo, sourceBranchRefName)
				if err != nil {
					return diag.FromErr(fmt.Errorf("error querying GitHub branch reference %s/%s (%s): %w", owner, repo, sourceBranchRefName, err))
				}
				if err := d.Set("autocreate_branch_source_sha", ref.Object.GetSHA()); err != nil {
					return diag.FromErr(err)
				}
			}

			sourceBranchSHA, _ := d.Get("autocreate_branch_source_sha").(string)
			tflog.Debug(ctx, "Creating branch", map[string]any{"source_branch": sourceBranchName, "source_sha": sourceBranchSHA})
			if _, _, err := client.Git.CreateRef(ctx, owner, repo, github.CreateRef{Ref: branchRefName, SHA: sourceBranchSHA}); err != nil {
				return diag.FromErr(fmt.Errorf("error creating GitHub branch reference %s/%s (%s): %w", owner, repo, branchRefName, err))
			}
		}
	} else {
		branch = repoInfo.GetDefaultBranch()
	}

	ctx = tflog.SetField(ctx, "branch", branch)

	tflog.Debug(ctx, "Checking if overwriting a repository file")
	fileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, file, &github.RepositoryContentGetOptions{Ref: branch})
	if err != nil {
		// 404 is a valid response if the file does not exist
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); !ok || ghErr.Response == nil || ghErr.Response.StatusCode != http.StatusNotFound {
			return diag.FromErr(err)
		}
	}

	opts := resourceGithubRepositoryFileOptions(d, "Add")

	if fileContent != nil {
		if overwrite, _ := d.Get("overwrite_on_create").(bool); overwrite {
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

	newResourceID, err := buildID(repo, escapeIDPart(file), branch)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, "Setting ID", map[string]any{"id": newResourceID})
	d.SetId(newResourceID)

	// Set the computed values from the create response as the file is not read back after being written.
	// The commit message is not stored when it is not configured as it's only used when writing.
	for k, v := range map[string]any{
		"repository_id": int(repoInfo.GetID()),
		"branch":        branch,
		"ref":           branch,
		"sha":           create.Content.GetSHA(),
		"commit_sha":    create.GetSHA(),
	} {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubRepositoryFileRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repo, _ := d.Get("repository").(string)
	branch, _ := d.Get("branch").(string)
	file, _ := d.Get("file").(string)

	ctx = tflog.SetField(ctx, "owner", owner)
	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "branch", branch)

	ctx = tflog.SetField(ctx, "file", file)
	fc, _, _, err := client.Repositories.GetContents(ctx, owner, repo, file, &github.RepositoryContentGetOptions{Ref: branch})
	if err != nil {
		// A 404 leaves fc nil, which removes the file from state below
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); !ok || ghErr.Response.StatusCode != http.StatusNotFound {
			return diag.FromErr(err)
		}
	}
	if fc == nil {
		tflog.Info(ctx, "Removing repository file from state because it no longer exists in GitHub")
		d.SetId("")
		return nil
	}

	content, err := fc.GetContent()
	if err != nil {
		return diag.FromErr(err)
	}

	oldSHA, _ := d.Get("sha").(string)

	if err = d.Set("content", content); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("sha", fc.GetSHA()); err != nil {
		return diag.FromErr(err)
	}

	if oldSHA == fc.GetSHA() {
		return nil
	}

	tflog.Debug(ctx, "File SHA changed, looking for commit that modified the file", map[string]any{"sha": fc.GetSHA()})
	commit, err := fileCommit(ctx, client, owner, repo, branch, file)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("commit_sha", commit.GetSHA()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryFileUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repo, _ := d.Get("repository").(string)
	branch, _ := d.Get("branch").(string)
	file, _ := d.Get("file").(string)

	ctx = tflog.SetField(ctx, "owner", owner)
	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "branch", branch)
	ctx = tflog.SetField(ctx, "file", file)

	opts := resourceGithubRepositoryFileOptions(d, "Update")

	update, _, err := client.Repositories.UpdateFile(ctx, owner, repo, file, opts)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set the computed values from the update response as the file is not read back after being written.
	// The commit message is not stored when it is not configured as it's only used when writing.
	for k, v := range map[string]any{
		"sha":        update.Content.GetSHA(),
		"commit_sha": update.GetSHA(),
	} {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubRepositoryFileDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repo, _ := d.Get("repository").(string)
	branch, _ := d.Get("branch").(string)
	file, _ := d.Get("file").(string)

	ctx = tflog.SetField(ctx, "owner", owner)
	ctx = tflog.SetField(ctx, "repository", repo)
	ctx = tflog.SetField(ctx, "branch", branch)
	ctx = tflog.SetField(ctx, "file", file)

	opts := resourceGithubRepositoryFileOptions(d, "Delete")

	if _, _, err := client.Repositories.DeleteFile(ctx, owner, repo, file, opts); err != nil {
		if isArchivedRepositoryError(err) {
			tflog.Info(ctx, "Skipping deletion of repository file from archived repository")
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryFileImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	repo, filePathPart, branch, err := parseID3(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid ID specified. Supplied ID must be written as <repository>:<file path>: (when branch is default) or <repository>:<file path>:<branch>. %w", err)
	}

	filePath := unescapeIDPart(filePathPart)

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}
	if branch == "" {
		branch = repoInfo.GetDefaultBranch()
	}

	fc, _, _, err := client.Repositories.GetContents(ctx, owner, repo, filePath, &github.RepositoryContentGetOptions{Ref: branch})
	if err != nil {
		return nil, err
	}
	if fc == nil {
		return nil, fmt.Errorf("filePath %s is not a file in repository %s/%s or repository is not readable", filePath, owner, repo)
	}

	commit, err := fileCommit(ctx, client, owner, repo, branch, filePath)
	if err != nil {
		return nil, err
	}

	newResourceID, err := buildID(repo, escapeIDPart(filePath), branch)
	if err != nil {
		return nil, err
	}
	d.SetId(newResourceID)

	if err := d.Set("repository", repo); err != nil {
		return nil, err
	}

	if err := d.Set("repository_id", int(repoInfo.GetID())); err != nil {
		return nil, err
	}

	if err := d.Set("branch", branch); err != nil {
		return nil, err
	}

	if err := d.Set("file", filePath); err != nil {
		return nil, err
	}

	if err = d.Set("overwrite_on_create", false); err != nil {
		return nil, err
	}

	if err := d.Set("ref", branch); err != nil {
		return nil, err
	}

	if err := d.Set("sha", fc.GetSHA()); err != nil {
		return nil, err
	}

	if err = d.Set("commit_sha", commit.GetSHA()); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func autoBranchDiffSuppressFunc(k, _, _ string, d *schema.ResourceData) bool {
	if autocreate, _ := d.Get("autocreate_branch").(bool); !autocreate {
		switch k {
		case "autocreate_branch", "autocreate_branch_source_branch", "autocreate_branch_source_sha":
			return true
		}
	}
	return false
}

// resourceGithubRepositoryFileOptions builds the content options for a write. The commit message is
// defaulted to "<action> <file>" when it is unset or is still the default message from creation.
func resourceGithubRepositoryFileOptions(d *schema.ResourceData, action string) *github.RepositoryContentFileOptions {
	content, _ := d.Get("content").(string)
	file, _ := d.Get("file").(string)
	commitMessage, _ := d.Get("commit_message").(string)

	opts := &github.RepositoryContentFileOptions{
		Content: []byte(content),
		Message: new(commitMessage),
	}

	if commitMessage == "" || commitMessage == fmt.Sprintf("Add %s", file) {
		opts.Message = new(fmt.Sprintf("%s %s", action, file))
	}

	if v, ok := d.GetOk("branch"); ok {
		s, _ := v.(string)
		opts.Branch = new(s)
	}

	if v, ok := d.GetOk("sha"); ok {
		s, _ := v.(string)
		opts.SHA = new(s)
	}

	commitAuthor, hasCommitAuthor := d.GetOk("commit_author")
	commitEmail, hasCommitEmail := d.GetOk("commit_email")

	if hasCommitAuthor && hasCommitEmail {
		name, _ := commitAuthor.(string)
		mail, _ := commitEmail.(string)
		opts.Author = &github.CommitAuthor{Name: new(name), Email: new(mail)}
		opts.Committer = &github.CommitAuthor{Name: new(name), Email: new(mail)}
	}

	return opts
}
