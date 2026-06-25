package github

import (
	"context"

	"github.com/google/go-github/v88/github"
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

func resourceGithubIssueLabelsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repository, ok := d.Get("repository").(string)
	if !ok {
		return diag.Errorf(`expected "repository" to be string`)
	}

	wantLabels := d.Get("label").(*schema.Set).List()

	wantLabelsMap := make(map[string]any, len(wantLabels))
	for _, label := range wantLabels {
		name := label.(map[string]any)["name"].(string)
		if _, found := wantLabelsMap[name]; found {
			return diag.Errorf("duplicate set label: %s", name)
		}
		wantLabelsMap[name] = label
	}

	hasLabels, err := listLabels(client, ctx, owner, repository)
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
				tflog.Debug(ctx, "Updating GitHub issue label", map[string]any{"owner": owner, "repository": repository, "name": name})

				_, _, err := client.Issues.EditLabel(ctx, owner, repository, name, &github.Label{
					Name:        new(name),
					Description: new(description),
					Color:       new(color),
				})
				if err != nil {
					return diag.FromErr(err)
				}
			}
		} else {
			tflog.Debug(ctx, "Deleting GitHub issue label", map[string]any{"owner": owner, "repository": repository, "name": name})

			_, err := client.Issues.DeleteLabel(ctx, owner, repository, name)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		hasLabelsMap[name] = struct{}{}
	}

	for _, l := range wantLabels {
		labelData := l.(map[string]any)
		name := labelData["name"].(string)

		_, found := hasLabelsMap[name]
		if !found {
			tflog.Debug(ctx, "Creating GitHub issue label", map[string]any{"owner": owner, "repository": repository, "name": name})

			_, _, err := client.Issues.CreateLabel(ctx, owner, repository, &github.Label{
				Name:        new(name),
				Description: new(labelData["description"].(string)),
				Color:       new(labelData["color"].(string)),
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

func resourceGithubIssueLabelsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repository := d.Id()

	tflog.Debug(ctx, "Reading GitHub issue labels", map[string]any{"owner": owner, "repository": repository})

	labels, err := listLabels(client, ctx, owner, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("repository", repository); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("label", flattenLabels(labels)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubIssueLabelsDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repository, ok := d.Get("repository").(string)
	if !ok {
		return diag.Errorf(`expected "repository" to be string`)
	}

	labels := d.Get("label").(*schema.Set).List()

	tflog.Debug(ctx, "Deleting GitHub issue labels", map[string]any{"owner": owner, "repository": repository})

	// delete
	for _, raw := range labels {
		label := raw.(map[string]any)
		name := label["name"].(string)

		tflog.Debug(ctx, "Deleting GitHub issue label", map[string]any{"owner": owner, "repository": repository, "name": name})

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
