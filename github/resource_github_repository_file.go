package github

import (
	"context"
	"log"
	"net/http"
	"strings"

	"fmt"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubRepositoryFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryFileCreate,
		Read:   resourceGithubRepositoryFileRead,
		Update: resourceGithubRepositoryFileUpdate,
		Delete: resourceGithubRepositoryFileDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				parts := strings.Split(d.Id(), ":")
				branch := "master"

				if len(parts) > 2 {
					return nil, fmt.Errorf("Invalid ID specified. Supplied ID must be written as <repository>/<file path> (when branch is \"master\") or <repository>/<file path>:<branch>")
				}

				if len(parts) == 2 {
					branch = parts[1]
				}

				client := meta.(*Owner).v3client
				owner := meta.(*Owner).name
				repo, file := splitRepoFilePath(parts[0])
				if err := checkRepositoryFileExists(client, owner, repo, file, branch); err != nil {
					return nil, err
				}

				d.SetId(fmt.Sprintf("%s/%s", repo, file))
				d.Set("branch", branch)

				return []*schema.ResourceData{d}, nil
			},
		},

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
				Description: "The branch name, defaults to \"master\"",
				Default:     "master",
			},
			"commit_message": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The commit message when creating or updating the file",
			},
			"commit_author": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The commit author name, defaults to the authenticated user's name",
			},
			"commit_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The commit author email address, defaults to the authenticated user's email address",
			},
			"sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The blob SHA of the file",
			},
		},
	}
}

func resourceGithubRepositoryFileOptions(d *schema.ResourceData) (*github.RepositoryContentFileOptions, error) {
	opts := &github.RepositoryContentFileOptions{
		Content: []byte(*github.String(d.Get("content").(string))),
		Branch:  github.String(d.Get("branch").(string)),
	}

	if commitMessage, hasCommitMessage := d.GetOk("commit_message"); hasCommitMessage {
		opts.Message = new(string)
		*opts.Message = commitMessage.(string)
	}

	if SHA, hasSHA := d.GetOk("sha"); hasSHA {
		opts.SHA = new(string)
		*opts.SHA = SHA.(string)
	}

	commitAuthor, hasCommitAuthor := d.GetOk("commit_author")
	commitEmail, hasCommitEmail := d.GetOk("commit_email")

	if hasCommitAuthor && !hasCommitEmail {
		return nil, fmt.Errorf("Cannot set commit_author without setting commit_email")
	}

	if hasCommitEmail && !hasCommitAuthor {
		return nil, fmt.Errorf("Cannot set commit_email without setting commit_author")
	}

	if hasCommitAuthor && hasCommitEmail {
		name := commitAuthor.(string)
		mail := commitEmail.(string)
		opts.Author = &github.CommitAuthor{Name: &name, Email: &mail}
		opts.Committer = &github.CommitAuthor{Name: &name, Email: &mail}
	}

	return opts, nil
}

func resourceGithubRepositoryFileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	file := d.Get("file").(string)
	branch := d.Get("branch").(string)

	if err := checkRepositoryBranchExists(client, owner, repo, branch); err != nil {
		return err
	}

	opts, err := resourceGithubRepositoryFileOptions(d)
	if err != nil {
		return err
	}

	if opts.Message == nil {
		m := fmt.Sprintf("Add %s", file)
		opts.Message = &m
	}

	log.Printf("[DEBUG] Creating repository file: %s/%s/%s in branch: %s", owner, repo, file, branch)
	resp, _, err := client.Repositories.CreateFile(ctx, owner, repo, file, opts)
	if err != nil {
		return err
	}

	d.Set("commit_author", resp.GetCommitter().GetName())
	d.Set("commit_email", resp.GetCommitter().GetEmail())
	d.Set("commit_message", resp.GetMessage())
	d.SetId(fmt.Sprintf("%s/%s", repo, file))

	return resourceGithubRepositoryFileRead(d, meta)
}

func resourceGithubRepositoryFileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	repo, file := splitRepoFilePath(d.Id())
	branch := d.Get("branch").(string)

	if err := checkRepositoryBranchExists(client, owner, repo, branch); err != nil {
		return err
	}

	log.Printf("[DEBUG] Reading repository file: %s/%s/%s, branch: %s", owner, repo, file, branch)
	opts := &github.RepositoryContentGetOptions{Ref: branch}
	fc, _, _, _ := client.Repositories.GetContents(ctx, owner, repo, file, opts)
	if fc == nil {
		log.Printf("[WARN] Removing repository path %s/%s/%s from state because it no longer exists in GitHub",
			owner, repo, file)
		d.SetId("")
		return nil
	}

	content, err := fc.GetContent()
	if err != nil {
		return err
	}

	d.Set("content", content)
	d.Set("repository", repo)
	d.Set("file", file)
	d.Set("sha", fc.GetSHA())

	return nil
}

func resourceGithubRepositoryFileUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	file := d.Get("file").(string)
	branch := d.Get("branch").(string)

	if err := checkRepositoryBranchExists(client, owner, repo, branch); err != nil {
		return err
	}

	opts, err := resourceGithubRepositoryFileOptions(d)
	if err != nil {
		return err
	}

	if *opts.Message == fmt.Sprintf("Add %s", file) {
		m := fmt.Sprintf("Update %s", file)
		opts.Message = &m
	}

	log.Printf("[DEBUG] Updating content in repository file: %s/%s/%s", owner, repo, file)
	resp, _, err := client.Repositories.CreateFile(ctx, owner, repo, file, opts)
	if err != nil {
		return err
	}

	d.Set("commit_author", resp.GetCommitter().GetName())
	d.Set("commit_email", resp.GetCommitter().GetEmail())
	d.Set("commit_message", resp.GetMessage())

	return resourceGithubRepositoryFileRead(d, meta)
}

func resourceGithubRepositoryFileDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	file := d.Get("file").(string)
	branch := d.Get("branch").(string)

	message := fmt.Sprintf("Delete %s", file)
	sha := d.Get("sha").(string)
	opts := &github.RepositoryContentFileOptions{
		Message: &message,
		SHA:     &sha,
		Branch:  &branch,
	}

	log.Printf("[DEBUG] Deleting repository file: %s/%s/%s", owner, repo, file)
	_, _, err := client.Repositories.DeleteFile(ctx, owner, repo, file, opts)
	if err != nil {
		return nil
	}

	return nil
}

// checkRepositoryBranchExists tests if a branch exists in a repository.
func checkRepositoryBranchExists(client *github.Client, owner, repo, branch string) error {
	ctx := context.WithValue(context.Background(), ctxId, buildTwoPartID(repo, branch))
	_, _, err := client.Repositories.GetBranch(ctx, owner, repo, branch)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				return fmt.Errorf("Branch %s not found in repository or repository is not readable", branch)
			}
		}
		return err
	}

	return nil
}

// checkRepositoryFileExists tests if a file exists in a repository.
func checkRepositoryFileExists(client *github.Client, owner, repo, file, branch string) error {
	ctx := context.WithValue(context.Background(), ctxId, fmt.Sprintf("%s/%s", repo, file))
	fc, _, _, err := client.Repositories.GetContents(ctx, owner, repo, file, &github.RepositoryContentGetOptions{Ref: branch})
	if err != nil {
		return nil
	}
	if fc == nil {
		return fmt.Errorf("File %s not a file in in repository %s/%s or repository is not readable", file, owner, repo)
	}

	return nil
}
