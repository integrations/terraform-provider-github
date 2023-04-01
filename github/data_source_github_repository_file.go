package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubRepositoryFile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryFileRead,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository name",
			},
			"file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The file path to manage",
			},
			"branch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The branch name, defaults to the repository's default branch",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the commit/branch/tag",
			},
			"content": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The file's content",
			},
			"commit_sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SHA of the commit that modified the file",
			},
			"commit_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The commit message when creating or updating the file",
			},
			"commit_author": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The commit author name, defaults to the authenticated user's name",
			},
			"commit_email": {
				Type:        schema.TypeString,
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

func dataSourceGithubRepositoryFileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	owner := meta.(*Owner).name
	repo := d.Get("repository").(string)

	// checking if repo has a slash in it, which means that full_name was passed
	// split and replace owner and repo
	parts := strings.Split(repo, "/")
	if len(parts) == 2 {
		log.Printf("[DEBUG] repo has a slash, extracting owner from: %s", repo)
		owner = parts[0]
		repo = parts[1]

		log.Printf("[DEBUG] owner: %s repo:%s", owner, repo)
	}

	file := d.Get("file").(string)

	opts := &github.RepositoryContentGetOptions{}
	if branch, ok := d.GetOk("branch"); ok {
		opts.Ref = branch.(string)
	}

	fc, _, _, err := client.Repositories.GetContents(ctx, owner, repo, file, opts)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[DEBUG] Missing GitHub repository file %s/%s/%s", owner, repo, file)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	content, err := fc.GetContent()
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", repo, file))
	d.Set("content", content)
	d.Set("repository", repo)
	d.Set("file", file)
	d.Set("sha", fc.GetSHA())

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

	log.Printf("[DEBUG] Data Source fetching commit info for repository file: %s/%s/%s", owner, repo, file)
	commit, err := getFileCommit(client, owner, repo, file, ref)
	log.Printf("[DEBUG] Found file: %s/%s/%s, in commit SHA: %s ", owner, repo, file, commit.GetSHA())
	if err != nil {
		return err
	}

	d.Set("commit_sha", commit.GetSHA())
	d.Set("commit_author", commit.Commit.GetCommitter().GetName())
	d.Set("commit_email", commit.Commit.GetCommitter().GetEmail())
	d.Set("commit_message", commit.GetCommit().GetMessage())

	return nil
}
