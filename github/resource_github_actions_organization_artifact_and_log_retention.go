package github

import (
	"context"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsOrganizationArtifactAndLogRetention() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsOrganizationArtifactAndLogRetentionCreate,
		ReadContext:   resourceGithubActionsOrganizationArtifactAndLogRetentionRead,
		UpdateContext: resourceGithubActionsOrganizationArtifactAndLogRetentionUpdate,
		DeleteContext: resourceGithubActionsOrganizationArtifactAndLogRetentionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Description: "Resource to manage how long GitHub Actions artifacts and logs are retained for an organization. The value also caps what each repository in the organization may set.",

		Schema: map[string]*schema.Schema{
			"days": {
				Type:             schema.TypeInt,
				Required:         true,
				Description:      "Number of days to retain artifacts and logs. Must not exceed 'maximum_allowed_days', which the enterprise imposes.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1, 400)),
			},
			"maximum_allowed_days": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Longest retention period the organization is allowed to set, as limited by the enterprise.",
			},
		},
	}
}

func resourceGithubActionsOrganizationArtifactAndLogRetentionCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	client := meta.v3client
	orgName := meta.name
	days, _ := d.Get("days").(int)

	if _, err := client.Actions.UpdateArtifactAndLogRetentionPeriodInOrganization(ctx, orgName, github.ArtifactPeriodOpt{
		Days: new(days),
	}); err != nil {
		return artifactAndLogRetentionError(err)
	}

	d.SetId(orgName)

	return resourceGithubActionsOrganizationArtifactAndLogRetentionRead(ctx, d, m)
}

func resourceGithubActionsOrganizationArtifactAndLogRetentionRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	client := meta.v3client

	retention, _, err := client.Actions.GetArtifactAndLogRetentionPeriodInOrganization(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("days", retention.GetDays()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("maximum_allowed_days", retention.GetMaximumAllowedDays()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationArtifactAndLogRetentionUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	client := meta.v3client
	days, _ := d.Get("days").(int)

	if _, err := client.Actions.UpdateArtifactAndLogRetentionPeriodInOrganization(ctx, d.Id(), github.ArtifactPeriodOpt{
		Days: new(days),
	}); err != nil {
		return artifactAndLogRetentionError(err)
	}

	return resourceGithubActionsOrganizationArtifactAndLogRetentionRead(ctx, d, m)
}

func resourceGithubActionsOrganizationArtifactAndLogRetentionDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	client := meta.v3client

	retention, _, err := client.Actions.GetArtifactAndLogRetentionPeriodInOrganization(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// An organization may not exceed the period its enterprise allows, so the reset target is
	// capped by that ceiling rather than always being GitHub's default.
	days := min(githubDefaultArtifactAndLogRetentionDays, retention.GetMaximumAllowedDays())

	if _, err := client.Actions.UpdateArtifactAndLogRetentionPeriodInOrganization(ctx, d.Id(), github.ArtifactPeriodOpt{
		Days: new(days),
	}); err != nil {
		return artifactAndLogRetentionError(err)
	}

	return nil
}
