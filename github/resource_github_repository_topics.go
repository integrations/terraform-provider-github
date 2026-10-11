package github

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryTopics() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryTopicsCreateOrUpdate,
		ReadContext:   resourceGithubRepositoryTopicsRead,
		UpdateContext: resourceGithubRepositoryTopicsCreateOrUpdate,
		DeleteContext: resourceGithubRepositoryTopicsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
				_ = d.Set("repository", d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringMatch(regexp.MustCompile(`^[-a-zA-Z0-9_.]{1,100}$`), "must include only alphanumeric characters, underscores or hyphens and consist of 100 characters or less")),
				Description:      "The name of the repository. The name is not case sensitive.",
			},
			"topics": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "An array of topics to add to the repository. Pass one or more topics to replace the set of existing topics. Send an empty array ([]) to clear all topics from the repository. Note: Topic names cannot contain uppercase letters.",
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validation.ToDiagFunc(validation.StringMatch(regexp.MustCompile(`^[a-z0-9][a-z0-9-]{0,49}$`), "must include only lowercase alphanumeric characters or hyphens and cannot start with a hyphen and consist of 50 characters or less")),
				},
			},
		},
	}
}

func resourceGithubRepositoryTopicsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner := meta.name
	repoName := d.Get("repository").(string)
	topics := expandStringList(d.Get("topics").(*schema.Set).List())

	if len(topics) > 0 {
		_, _, err := client.Repositories.ReplaceAllTopics(ctx, owner, repoName, topics)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(repoName)
	return resourceGithubRepositoryTopicsRead(ctx, d, meta)
}

func resourceGithubRepositoryTopicsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	ctx = context.WithValue(ctx, ctxId, d.Id())

	owner := meta.name
	repoName := d.Get("repository").(string)

	results := make([]string, 0)
	listOptions := &github.ListOptions{PerPage: meta.maxPerPage}
	for {
		topics, resp, err := client.Repositories.ListAllTopics(ctx, owner, repoName, listOptions)
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
				if ghErr.Response.StatusCode == http.StatusNotModified {
					return nil
				}
				if ghErr.Response.StatusCode == http.StatusNotFound {
					tflog.Info(ctx, "Removing topics from repository from state because it no longer exists in GitHub", map[string]any{"owner": owner, "repository": repoName})
					d.SetId("")
					return nil
				}
			}
			return diag.FromErr(err)
		}

		results = append(results, topics...)

		if resp.NextPage == 0 {
			break
		}

		listOptions.Page = resp.NextPage
	}

	if err := d.Set("topics", flattenStringList(results)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryTopicsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	ctx = context.WithValue(ctx, ctxId, d.Id())

	owner := meta.name
	repoName := d.Get("repository").(string)

	_, _, err := client.Repositories.ReplaceAllTopics(ctx, owner, repoName, []string{})
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
