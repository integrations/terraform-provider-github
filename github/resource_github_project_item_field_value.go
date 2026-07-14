package github

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

var projectV2ItemFieldValuePaths = []string{"text", "number", "date", "single_select_option_id", "iteration_id"}

type projectV2ItemFieldValueNode struct {
	Typename githubv4.String `graphql:"__typename"`
	Text     struct {
		ID   githubv4.String
		Text githubv4.String
	} `graphql:"... on ProjectV2ItemFieldTextValue"`
	Number struct {
		ID     githubv4.String
		Number githubv4.Float
	} `graphql:"... on ProjectV2ItemFieldNumberValue"`
	Date struct {
		ID   githubv4.String
		Date githubv4.Date
	} `graphql:"... on ProjectV2ItemFieldDateValue"`
	SingleSelect struct {
		ID       githubv4.String
		OptionID githubv4.String
	} `graphql:"... on ProjectV2ItemFieldSingleSelectValue"`
	Iteration struct {
		ID          githubv4.String
		IterationID githubv4.String
	} `graphql:"... on ProjectV2ItemFieldIterationValue"`
}

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
	var mutation struct {
		UpdateProjectV2ItemFieldValue struct {
			Item struct {
				ID githubv4.String
			} `graphql:"projectV2Item"`
		} `graphql:"updateProjectV2ItemFieldValue(input: $input)"`
	}
	input := githubv4.UpdateProjectV2ItemFieldValueInput{
		ProjectID: githubv4.ID(d.Get("project_id").(string)),
		ItemID:    githubv4.ID(d.Get("item_id").(string)),
		FieldID:   githubv4.ID(d.Get("field_id").(string)),
		Value:     value,
	}
	if err := meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil); err != nil {
		return diag.FromErr(err)
	}
	id, err := buildID(d.Get("project_id").(string), d.Get("item_id").(string), d.Get("field_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	return resourceGithubProjectItemFieldValueRead(ctx, d, meta)
}

func resourceGithubProjectItemFieldValueRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	field, err := queryProjectV2Field(ctx, meta.(*Owner).v4client, d.Get("field_id").(string))
	if isProjectV2NotFound(err) {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	fieldName := projectV2FieldNodeName(field)
	if fieldName == "" {
		return diag.Errorf("project field %q has no supported type", d.Get("field_id").(string))
	}

	var query struct {
		Node struct {
			Typename githubv4.String `graphql:"__typename"`
			Item     struct {
				FieldValue projectV2ItemFieldValueNode `graphql:"fieldValueByName(name: $fieldName)"`
			} `graphql:"... on ProjectV2Item"`
		} `graphql:"node(id: $itemID)"`
	}
	variables := map[string]any{"itemID": githubv4.ID(d.Get("item_id").(string)), "fieldName": githubv4.String(fieldName)}
	err = meta.(*Owner).v4client.Query(ctx, &query, variables)
	if isProjectV2NotFound(err) {
		tflog.Info(ctx, "Removing project item field value from state because the item no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if err := setProjectV2ItemFieldValueState(d, query.Node.Item.FieldValue); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectItemFieldValueDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var mutation struct {
		ClearProjectV2ItemFieldValue struct {
			Item struct {
				ID githubv4.String
			} `graphql:"projectV2Item"`
		} `graphql:"clearProjectV2ItemFieldValue(input: $input)"`
	}
	input := githubv4.ClearProjectV2ItemFieldValueInput{
		ProjectID: githubv4.ID(d.Get("project_id").(string)),
		ItemID:    githubv4.ID(d.Get("item_id").(string)),
		FieldID:   githubv4.ID(d.Get("field_id").(string)),
	}
	err := meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil)
	if err != nil && !isProjectV2NotFound(err) {
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

func expandProjectV2ItemFieldValue(d *schema.ResourceData) (githubv4.ProjectV2FieldValue, error) {
	var value githubv4.ProjectV2FieldValue
	if raw, ok := projectV2ItemFieldValueConfigured(d, "text"); ok {
		text := githubv4.String(raw.(string))
		value.Text = &text
	}
	if raw, ok := projectV2ItemFieldValueConfigured(d, "number"); ok {
		number := githubv4.Float(raw.(float64))
		value.Number = &number
	}
	if raw, ok := projectV2ItemFieldValueConfigured(d, "date"); ok {
		date, err := time.Parse(time.DateOnly, raw.(string))
		if err != nil {
			return value, err
		}
		parsed := githubv4.Date{Time: date}
		value.Date = &parsed
	}
	if raw, ok := projectV2ItemFieldValueConfigured(d, "single_select_option_id"); ok {
		optionID := githubv4.String(raw.(string))
		value.SingleSelectOptionID = &optionID
	}
	if raw, ok := projectV2ItemFieldValueConfigured(d, "iteration_id"); ok {
		iterationID := githubv4.String(raw.(string))
		value.IterationID = &iterationID
	}
	return value, nil
}

func projectV2ItemFieldValueConfigured(d *schema.ResourceData, key string) (any, bool) {
	// GetOk discards valid zero values, including an empty string and the number zero.
	return d.GetOkExists(key) //nolint:staticcheck // The SDK has no non-deprecated equivalent that preserves zero values.
}

func projectV2FieldNodeName(field projectV2FieldNode) string {
	switch field.Typename {
	case "ProjectV2Field":
		return string(field.Field.Name)
	case "ProjectV2SingleSelectField":
		return string(field.SingleSelect.Name)
	case "ProjectV2IterationField":
		return string(field.Iteration.Name)
	}
	return ""
}

func setProjectV2ItemFieldValueState(d *schema.ResourceData, value projectV2ItemFieldValueNode) error {
	for _, key := range projectV2ItemFieldValuePaths {
		if err := d.Set(key, nil); err != nil {
			return err
		}
	}
	switch value.Typename {
	case "ProjectV2ItemFieldTextValue":
		return d.Set("text", value.Text.Text)
	case "ProjectV2ItemFieldNumberValue":
		return d.Set("number", value.Number.Number)
	case "ProjectV2ItemFieldDateValue":
		return d.Set("date", value.Date.Date.Format(time.DateOnly))
	case "ProjectV2ItemFieldSingleSelectValue":
		return d.Set("single_select_option_id", value.SingleSelect.OptionID)
	case "ProjectV2ItemFieldIterationValue":
		return d.Set("iteration_id", value.Iteration.IterationID)
	default:
		d.SetId("")
		return nil
	}
}
