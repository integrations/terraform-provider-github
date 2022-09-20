package github

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v47/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubIssue() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubIssueCreateOrUpdate,
		Read:   resourceGithubIssueRead,
		Update: resourceGithubIssueCreateOrUpdate,
		Delete: resourceGithubIssueDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"number": {
				Type:     schema.TypeInt,
				Required: false,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"body": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"labels": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Optional:    true,
				Description: "List of names of labels on the issue",
			},
			"assignees": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Optional:    true,
				Description: "List of Logins for Users to assign to this issue",
			},
			"milestone_number": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"issue_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubIssueCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	ctx := context.Background()
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	title := d.Get("title").(string)
	milestone := d.Get("milestone_number").(int)

	req := &github.IssueRequest{
		Title: github.String(title),
	}

	if v, ok := d.GetOk("body"); ok {
		req.Body = github.String(v.(string))
	}

	labels := expandStringList(d.Get("labels").(*schema.Set).List())
	req.Labels = &labels

	assignees := expandStringList(d.Get("assignees").(*schema.Set).List())
	req.Assignees = &assignees

	if milestone > 0 {
		req.Milestone = intPtr(milestone)
	}

	var issue *github.Issue
	var resp *github.Response
	var err error
	if d.IsNewResource() {
		log.Printf("[DEBUG] Creating issue: %s (%s/%s)",
			title, orgName, repoName)
		issue, resp, err = client.Issues.Create(ctx, orgName, repoName, req)
		if resp != nil {
			log.Printf("[DEBUG] Response from creating issue: %#v", *resp)
		}
	} else {
		number := d.Get("number").(int)
		log.Printf("[DEBUG] Updating issue: %d:%s (%s/%s)",
			number, title, orgName, repoName)
		issue, resp, err = client.Issues.Edit(ctx, orgName, repoName, number, req)
		if resp != nil {
			log.Printf("[DEBUG] Response from updating issue: %#v", *resp)
		}
	}
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repoName, strconv.Itoa(issue.GetNumber())))
	d.Set("issue_id", issue.GetID())
	return resourceGithubIssueRead(d, meta)
}

func resourceGithubIssueRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	repoName, idNumber, err := parseTwoPartID(d.Id(), "repository", "issue_number")
	if err != nil {
		return err
	}

	number, err := strconv.Atoi(idNumber)
	if err != nil {
		return err
	}

	orgName := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	log.Printf("[DEBUG] Reading issue: %d (%s/%s)", number, orgName, repoName)
	issue, resp, err := client.Issues.Get(ctx,
		orgName, repoName, number)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing issue %d (%s/%s) from state because it no longer exists in GitHub",
					number, orgName, repoName)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("repository", repoName)
	d.Set("number", number)
	d.Set("title", issue.GetTitle())
	d.Set("body", issue.GetBody())
	d.Set("milestone_number", issue.GetMilestone().GetNumber())

	var labels []string
	for _, v := range issue.Labels {
		labels = append(labels, v.GetName())
	}
	d.Set("labels", flattenStringList(labels))

	var assignees []string
	for _, v := range issue.Assignees {
		assignees = append(assignees, v.GetLogin())
	}
	d.Set("assignees", flattenStringList(assignees))

	d.Set("issue_id", issue.GetID())
	return nil
}

func resourceGithubIssueDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	number := d.Get("number").(int)
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting issue by closing: %d (%s/%s)", number, orgName, repoName)

	request := &github.IssueRequest{State: github.String("closed")}

	_, _, err := client.Issues.Edit(ctx, orgName, repoName, number, request)

	return err
}

func intPtr(i int) *int {
	return &i
}
