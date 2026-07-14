package github

import (
	"context"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	valueapplication "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
	valueusecases "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value/use-cases"
	valuegithub "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/item/field/value"
)

var projectV2ItemFieldValuePaths = []string{"text", "number", "date", "single_select_option_id", "iteration_id"}

func resourceGithubProjectItemFieldValue() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages one custom field value for a GitHub Projects V2 item.",
		CreateContext: resourceGithubProjectItemFieldValueCreateOrUpdate,
		ReadContext:   resourceGithubProjectItemFieldValueRead,
		UpdateContext: resourceGithubProjectItemFieldValueCreateOrUpdate,
		DeleteContext: resourceGithubProjectItemFieldValueDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubProjectItemFieldValueImport,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node ID of the Projects V2 project.",
			},
			"item_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node ID of the Projects V2 item.",
			},
			"field_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node ID of the Projects V2 field.",
			},
			"text": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: projectV2ItemFieldValuePaths,
			},
			"number": {
				Type:         schema.TypeFloat,
				Optional:     true,
				ExactlyOneOf: projectV2ItemFieldValuePaths,
			},
			"date": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: projectV2ItemFieldValuePaths,
				ValidateFunc: validateProjectV2Date,
			},
			"single_select_option_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: projectV2ItemFieldValuePaths,
			},
			"iteration_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: projectV2ItemFieldValuePaths,
			},
		},
	}
}

func resourceGithubProjectItemFieldValueCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	value, err := expandProjectV2ItemFieldValue(d)
	if err != nil {
		return diag.FromErr(err)
	}
	result, err := valueusecases.NewSet(valuegithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, valueapplication.SetInput{
		ProjectID: projectV2Get[string](d, "project_id"), ItemID: projectV2Get[string](d, "item_id"), FieldID: projectV2Get[string](d, "field_id"), Value: value,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	id, err := buildID(projectV2Get[string](d, "project_id"), projectV2Get[string](d, "item_id"), projectV2Get[string](d, "field_id"))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	if err := setProjectV2ItemFieldValueState(d, result); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectItemFieldValueRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	value, err := valueusecases.NewGet(valuegithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, projectV2Get[string](d, "item_id"), projectV2Get[string](d, "field_id"))
	if errors.Is(err, projects.ErrNotFound) {
		tflog.Info(ctx, "Removing project item field value from state because the item no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if err := setProjectV2ItemFieldValueState(d, value); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectItemFieldValueDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := valueusecases.NewClear(valuegithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, projectV2Get[string](d, "project_id"), projectV2Get[string](d, "item_id"), projectV2Get[string](d, "field_id"))
	if err != nil && !errors.Is(err, projects.ErrNotFound) {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectItemFieldValueImport(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	projectID, itemID, fieldID, err := parseID3(d.Id())
	if err != nil {
		return nil, err
	}
	for key, value := range map[string]string{"project_id": projectID, "item_id": itemID, "field_id": fieldID} {
		if err := d.Set(key, value); err != nil {
			return nil, err
		}
	}
	return []*schema.ResourceData{d}, nil
}

func expandProjectV2ItemFieldValue(d *schema.ResourceData) (valueapplication.Result, error) {
	var value valueapplication.Result
	if raw, ok := projectV2ItemFieldValueConfigured(d, "text"); ok {
		value.Kind = valueapplication.KindText
		value.Text = projectV2As[string](raw, "text")
	}
	if raw, ok := projectV2ItemFieldValueConfigured(d, "number"); ok {
		value.Kind = valueapplication.KindNumber
		value.Number = projectV2As[float64](raw, "number")
	}
	if raw, ok := projectV2ItemFieldValueConfigured(d, "date"); ok {
		date, err := time.Parse(time.DateOnly, projectV2As[string](raw, "date"))
		if err != nil {
			return value, err
		}
		value.Kind = valueapplication.KindDate
		value.Date = date
	}
	if raw, ok := projectV2ItemFieldValueConfigured(d, "single_select_option_id"); ok {
		value.Kind = valueapplication.KindSingleSelect
		value.OptionID = projectV2As[string](raw, "single_select_option_id")
	}
	if raw, ok := projectV2ItemFieldValueConfigured(d, "iteration_id"); ok {
		value.Kind = valueapplication.KindIteration
		value.IterationID = projectV2As[string](raw, "iteration_id")
	}
	return value, nil
}

func projectV2ItemFieldValueConfigured(d *schema.ResourceData, key string) (any, bool) {
	// GetOk discards valid zero values, including an empty string and the number zero.
	return d.GetOkExists(key) //nolint:staticcheck // The SDK has no non-deprecated equivalent that preserves zero values.
}

func setProjectV2ItemFieldValueState(d *schema.ResourceData, value valueapplication.Result) error {
	for _, key := range projectV2ItemFieldValuePaths {
		if err := d.Set(key, nil); err != nil {
			return err
		}
	}
	switch value.Kind {
	case valueapplication.KindText:
		return d.Set("text", value.Text)
	case valueapplication.KindNumber:
		return d.Set("number", value.Number)
	case valueapplication.KindDate:
		return d.Set("date", value.Date.Format(time.DateOnly))
	case valueapplication.KindSingleSelect:
		return d.Set("single_select_option_id", value.OptionID)
	case valueapplication.KindIteration:
		return d.Set("iteration_id", value.IterationID)
	default:
		d.SetId("")
		return nil
	}
}
