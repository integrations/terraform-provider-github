package github

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/integrations/terraform-provider-github/v6/internal/tfpluginv2util"
)

func resourceGithubIssue() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubIssueCreateOrUpdate,
		ReadContext:   resourceGithubIssueRead,
		UpdateContext: resourceGithubIssueCreateOrUpdate,
		DeleteContext: resourceGithubIssueDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CustomizeDiff: diffETag,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub repository name.",
			},
			"number": {
				Type:        schema.TypeInt,
				Required:    false,
				Computed:    true,
				Description: "The issue number.",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Title of the issue.",
			},
			"body": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Body of the issue.",
			},
			"labels": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Optional:    true,
				Description: "List of labels to attach to the issue.",
			},
			"assignees": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Optional:    true,
				Description: "List of Logins to assign to the issue.",
			},
			"milestone_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Milestone number to assign to the issue.",
			},
			"issue_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The issue id.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the issue.",
			},
		},
	}
}

func resourceGithubIssueCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	if err := d.Set("etag", nil); err != nil {
		return diag.FromErr(err)
	}

	repoName := tfpluginv2util.Get[string](d, "repository")
	title := tfpluginv2util.Get[string](d, "title")
	body := tfpluginv2util.Get[string](d, "body")
	milestone := tfpluginv2util.Get[int](d, "milestone_number")

	labels := tfpluginv2util.GetSet[string](d, "labels", true)
	asignees := tfpluginv2util.GetSet[string](d, "assignees", true)

	var issue *github.Issue
	var resp *github.Response
	var err error
	if d.IsNewResource() {
		req := github.CreateIssueRequest{
			Title:     title,
			Body:      new(body),
			Labels:    labels,
			Assignees: asignees,
		}

		if milestone > 0 {
			req.Milestone = new(milestone)
		}
		tflog.Debug(ctx, "Creating issue", map[string]any{"owner": orgName, "repository": repoName})
		issue, resp, err = client.Issues.Create(ctx, orgName, repoName, req)
		if resp != nil {
			tflog.Debug(ctx, "Response from creating issue", map[string]any{"status_code": resp.StatusCode})
		}
	} else {
		req := github.UpdateIssueRequest{
			Title:     new(title),
			Body:      new(body),
			Labels:    labels,
			Assignees: asignees,
		}

		if milestone > 0 {
			req.Milestone = new(milestone)
		}

		number, _ := d.Get("number").(int)
		tflog.Debug(ctx, "Updating issue", map[string]any{"issue_number": number, "owner": orgName, "repository": repoName})
		issue, resp, err = client.Issues.Update(ctx, orgName, repoName, number, req)
		if resp != nil {
			tflog.Debug(ctx, "Response from updating issue", map[string]any{"status_code": resp.StatusCode})
		}
	}
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(repoName, strconv.Itoa(issue.GetNumber())))
	if err = d.Set("issue_id", issue.GetID()); err != nil {
		return diag.FromErr(err)
	}
	return resourceGithubIssueRead(ctx, d, meta)
}

func resourceGithubIssueRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	repoName, idNumber, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	number, err := strconv.Atoi(idNumber)
	if err != nil {
		return diag.FromErr(err)
	}

	orgName := meta.(*Owner).name
	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}
	tflog.Debug(ctx, "Reading issue", map[string]any{"issue_number": number, "owner": orgName, "repository": repoName})
	issue, resp, err := client.Issues.Get(ctx,
		orgName, repoName, number)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Warn(ctx, "Removing issue from state because it no longer exists in GitHub", map[string]any{"issue_number": number, "owner": orgName, "repository": repoName})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("number", number); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("title", issue.GetTitle()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("body", issue.GetBody()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("milestone_number", issue.GetMilestone().GetNumber()); err != nil {
		return diag.FromErr(err)
	}

	var labels []string
	for _, v := range issue.Labels {
		labels = append(labels, v.GetName())
	}
	if err = d.Set("labels", flattenStringList(labels)); err != nil {
		return diag.FromErr(err)
	}

	var assignees []string
	for _, v := range issue.Assignees {
		assignees = append(assignees, v.GetLogin())
	}
	if err = d.Set("assignees", flattenStringList(assignees)); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("issue_id", issue.GetID()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubIssueDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	number := d.Get("number").(int)
	ctx = context.WithValue(ctx, ctxId, d.Id())
	tflog.Debug(ctx, "Deleting issue by closing", map[string]any{"issue_number": number, "owner": orgName, "repository": repoName})

	request := github.UpdateIssueRequest{State: new("closed")}

	_, _, err := client.Issues.Update(ctx, orgName, repoName, number, request)

	return diag.FromErr(err)
}
