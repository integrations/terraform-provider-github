package github

import (
	"context"
	"log"
	"net/url"
	"strings"

	"fmt"

	"github.com/google/go-github/v50/github"
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

				if len(parts) > 2 {
					return nil, fmt.Errorf("invalid ID specified. Supplied ID must be written as <repository>/<file path> (when branch is \"main\") or <repository>/<file path>:<branch>")
				}

				client := meta.(*Owner).v3client
				owner := meta.(*Owner).name
				repo, file := splitRepoFilePath(parts[0])
				// test if a file exists in a repository.
				ctx := context.WithValue(context.Background(), ctxId, fmt.Sprintf("%s/%s", repo, file))
				opts := &github.RepositoryContentGetOptions{}
				if len(parts) == 2 {
					opts.Ref = parts[1]
					d.Set("branch", parts[1])
				}
				fc, _, _, err := client.Repositories.GetContents(ctx, owner, repo, file, opts)
				if err != nil {
					return nil, err
				}
				if fc == nil {
					return nil, fmt.Errorf("file %s is not a file in repository %s/%s or repository is not readable", file, owner, repo)
				}

				d.SetId(fmt.Sprintf("%s/%s", repo, file))
				d.Set("overwrite_on_create", false)

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
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    false,
				Description: "The commit author name, defaults to the authenticated user's name. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. ",
			},
			"commit_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    false,
				Description: "The commit author email address, defaults to the authenticated user's email address. GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.",
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
		},
	}
}

func resourceGithubRepositoryFileOptions(d *schema.ResourceData) (*github.RepositoryContentFileOptions, error) {
	opts := &github.RepositoryContentFileOptions{
		Content: []byte(*github.String(d.Get("content").(string))),
	}

	if branch, ok := d.GetOk("branch"); ok {
		opts.Branch = github.String(branch.(string))
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
		return nil, fmt.Errorf("cannot set commit_author without setting commit_email")
	}

	if hasCommitEmail && !hasCommitAuthor {
		return nil, fmt.Errorf("cannot set commit_email without setting commit_author")
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

	checkOpt := github.RepositoryContentGetOptions{}

	if branch, ok := d.GetOk("branch"); ok {
		log.Printf("[DEBUG] Using explicitly set branch: %s", branch.(string))
		if err := checkRepositoryBranchExists(client, owner, repo, branch.(string)); err != nil {
			return err
		}
		checkOpt.Ref = branch.(string)
	}

	opts, err := resourceGithubRepositoryFileOptions(d)
	if err != nil {
		return err
	}

	if opts.Message == nil {
		m := fmt.Sprintf("Add %s", file)
		opts.Message = &m
	}

	log.Printf("[DEBUG] Checking if overwriting a repository file: %s/%s/%s", owner, repo, file)
	fileContent, _, resp, err := client.Repositories.GetContents(ctx, owner, repo, file, &checkOpt)
	if err != nil {
		if resp != nil {
			if resp.StatusCode != 404 {
				// 404 is a valid response if the file does not exist
				return err
			}
		} else {
			// Response should be non-nil
			return err
		}
	}

	if fileContent != nil {
		if d.Get("overwrite_on_create").(bool) {
			// Overwrite existing file if requested by configuring the options for
			// `client.Repositories.CreateFile` to match the existing file's SHA
			opts.SHA = fileContent.SHA
		} else {
			// Error if overwriting a file is not requested
			return fmt.Errorf("refusing to overwrite existing file: configure `overwrite_on_create` to `true` to override")
		}
	}

	// Create a new or overwritten file
	create, _, err := client.Repositories.CreateFile(ctx, owner, repo, file, opts)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", repo, file))
	d.Set("commit_sha", create.Commit.GetSHA())

	return resourceGithubRepositoryFileRead(d, meta)
}

func resourceGithubRepositoryFileRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	repo, file := splitRepoFilePath(d.Id())

	opts := &github.RepositoryContentGetOptions{}

	if branch, ok := d.GetOk("branch"); ok {
		log.Printf("[DEBUG] Using explicitly set branch: %s", branch.(string))
		if err := checkRepositoryBranchExists(client, owner, repo, branch.(string)); err != nil {
			return err
		}
		opts.Ref = branch.(string)
	}

	fc, _, _, _ := client.Repositories.GetContents(ctx, owner, repo, file, opts)
	if fc == nil {
		log.Printf("[INFO] Removing repository path %s/%s/%s from state because it no longer exists in GitHub",
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

	var commit *github.RepositoryCommit

	parsedUrl, err := url.Parse(fc.GetURL())
	if err != nil {
		return err
	}
	parsedQuery, err := url.ParseQuery(parsedUrl.RawQuery)
	if err != nil {
		return err
	}
	ref := parsedQuery["ref"][0]
	d.Set("ref", ref)

	// Use the SHA to lookup the commit info if we know it, otherwise loop through commits
	if sha, ok := d.GetOk("commit_sha"); ok {
		log.Printf("[DEBUG] Using known commit SHA: %s", sha.(string))
		commit, _, err = client.Repositories.GetCommit(ctx, owner, repo, sha.(string), nil)
	} else {
		log.Printf("[DEBUG] Commit SHA unknown for file: %s/%s/%s, looking for commit...", owner, repo, file)
		commit, err = getFileCommit(client, owner, repo, file, ref)
		log.Printf("[DEBUG] Found file: %s/%s/%s, in commit SHA: %s ", owner, repo, file, commit.GetSHA())
	}
	if err != nil {
		return err
	}

	d.Set("commit_sha", commit.GetSHA())

	commit_author := commit.Commit.GetCommitter().GetName()
	commit_email := commit.Commit.GetCommitter().GetEmail()

	_, hasCommitAuthor := d.GetOk("commit_author")
	_, hasCommitEmail := d.GetOk("commit_email")

	//read from state if author+email is set explicitly, and if it was not github signing it for you previously
	if commit_author != "GitHub" && commit_email != "noreply@github.com" && hasCommitAuthor && hasCommitEmail {
		d.Set("commit_author", commit_author)
		d.Set("commit_email", commit_email)
	}
	d.Set("commit_message", commit.GetCommit().GetMessage())

	return nil
}

func resourceGithubRepositoryFileUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	file := d.Get("file").(string)

	if branch, ok := d.GetOk("branch"); ok {
		log.Printf("[DEBUG] Using explicitly set branch: %s", branch.(string))
		if err := checkRepositoryBranchExists(client, owner, repo, branch.(string)); err != nil {
			return err
		}
	}

	opts, err := resourceGithubRepositoryFileOptions(d)
	if err != nil {
		return err
	}

	if *opts.Message == fmt.Sprintf("Add %s", file) {
		m := fmt.Sprintf("Update %s", file)
		opts.Message = &m
	}

	create, _, err := client.Repositories.CreateFile(ctx, owner, repo, file, opts)
	if err != nil {
		return err
	}

	d.Set("commit_sha", create.GetSHA())

	return resourceGithubRepositoryFileRead(d, meta)
}

func resourceGithubRepositoryFileDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	file := d.Get("file").(string)

	var branch string

	message := fmt.Sprintf("Delete %s", file)

	if commitMessage, hasCommitMessage := d.GetOk("commit_message"); hasCommitMessage {
		message = commitMessage.(string)
	}

	sha := d.Get("sha").(string)
	opts := &github.RepositoryContentFileOptions{
		Message: &message,
		SHA:     &sha,
	}

	if b, ok := d.GetOk("branch"); ok {
		log.Printf("[DEBUG] Using explicitly set branch: %s", b.(string))
		branch = b.(string)
		opts.Branch = &branch
	}

	_, _, err := client.Repositories.DeleteFile(ctx, owner, repo, file, opts)
	if err != nil {
		return nil
	}

	return nil
}
