package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubIssueLabels() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubIssueLabelsCreateOrUpdate,
		ReadContext:   resourceGithubIssueLabelsRead,
		UpdateContext: resourceGithubIssueLabelsCreateOrUpdate,
		DeleteContext: resourceGithubIssueLabelsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub repository.",
			},
			"label": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of labels",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the label.",
						},
						"color": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "A 6 character hex code, without the leading '#', identifying the color of the label.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A short description of the label.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to the issue label.",
						},
					},
				},
			},
		},
	}
}

func resourceGithubIssueLabelsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	owner := meta.name
	repository := d.Id()
	ctx = context.WithValue(ctx, ctxId, repository)
	tflog.Debug(ctx, "Reading GitHub issue labels", map[string]any{"owner": owner, "repository": repository})

	labels, err := listLabels(meta, ctx, owner, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("repository", repository)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("label", flattenLabels(labels))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubIssueLabelsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repository := d.Get("repository").(string)
	ctx = context.WithValue(ctx, ctxId, repository)

	wantLabels := d.Get("label").(*schema.Set).List()

	wantLabelsMap := make(map[string]any, len(wantLabels))
	for _, label := range wantLabels {
		name := label.(map[string]any)["name"].(string)
		if _, found := wantLabelsMap[name]; found {
			return diag.FromErr(fmt.Errorf("duplicate set label: %s", name))
		}
		wantLabelsMap[name] = label
	}

	hasLabels, err := listLabels(meta, ctx, owner, repository)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, "Updating GitHub issue labels", map[string]any{"owner": owner, "repository": repository})

	hasLabelsMap := make(map[string]struct{}, len(hasLabels))
	for _, hasLabel := range hasLabels {
		name := hasLabel.GetName()
		wantLabel, found := wantLabelsMap[name]
		if found {
			labelData := wantLabel.(map[string]any)
			description := labelData["description"].(string)
			color := labelData["color"].(string)
			if hasLabel.GetDescription() != description || hasLabel.GetColor() != color {
				tflog.Debug(ctx, "Updating GitHub issue label", map[string]any{"owner": owner, "repository": repository, "label_name": name})

				_, _, err := client.Issues.UpdateLabel(ctx, owner, repository, name, github.UpdateIssueLabelRequest{
					NewName:     new(name),
					Description: new(description),
					Color:       new(color),
				})
				if err != nil {
					return diag.FromErr(err)
				}
			}
		} else {
			tflog.Debug(ctx, "Deleting GitHub issue label", map[string]any{"owner": owner, "repository": repository, "label_name": name})

			_, err := client.Issues.DeleteLabel(ctx, owner, repository, name)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		hasLabelsMap[name] = struct{}{}
	}

	for _, l := range wantLabels {
		labelData := l.(map[string]any)
		name, _ := labelData["name"].(string)

		_, found := hasLabelsMap[name]
		if !found {
			tflog.Debug(ctx, "Creating GitHub issue label", map[string]any{"owner": owner, "repository": repository, "label_name": name})

			description, _ := labelData["description"].(string)
			color, _ := labelData["color"].(string)

			_, _, err := client.Issues.CreateLabel(ctx, owner, repository, github.CreateIssueLabelRequest{
				Name:        name,
				Description: new(description),
				Color:       new(color),
			})
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId(repository)

	err = d.Set("label", wantLabels)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubIssueLabelsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repository := d.Get("repository").(string)
	ctx = context.WithValue(ctx, ctxId, repository)

	labels := d.Get("label").(*schema.Set).List()
	tflog.Debug(ctx, "Deleting GitHub issue labels", map[string]any{"owner": owner, "repository": repository})

	// delete
	for _, raw := range labels {
		label := raw.(map[string]any)
		name := label["name"].(string)
		tflog.Debug(ctx, "Deleting GitHub issue label", map[string]any{"owner": owner, "repository": repository, "label_name": name})

		_, err := client.Issues.DeleteLabel(ctx, owner, repository, name)
		if err != nil {
			if isArchivedRepositoryError(err) {
				tflog.Info(ctx, "Skipping deletion of remaining issue labels from archived repository", map[string]any{"owner": owner, "repository": repository})
				break // Skip deleting remaining labels
			}
			return diag.FromErr(err)
		}
	}

	d.SetId(repository)

	err := d.Set("label", make([]map[string]any, 0))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
