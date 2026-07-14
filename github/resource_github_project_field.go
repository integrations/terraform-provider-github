package github

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"
)

var projectV2FieldTypes = []string{"TEXT", "SINGLE_SELECT", "NUMBER", "DATE", "ITERATION"}

var projectV2FieldColors = []string{"GRAY", "BLUE", "GREEN", "YELLOW", "ORANGE", "RED", "PINK", "PURPLE"}

type projectV2FieldFragment struct {
	ID       githubv4.String
	Name     githubv4.String
	DataType githubv4.ProjectV2FieldType
	Project  struct {
		ID githubv4.String
	}
}

type projectV2SingleSelectFieldFragment struct {
	projectV2FieldFragment
	Options []struct {
		ID          githubv4.String
		Name        githubv4.String
		Description githubv4.String
		Color       githubv4.ProjectV2SingleSelectFieldOptionColor
	}
}

type projectV2IterationFieldFragment struct {
	projectV2FieldFragment
	Configuration struct {
		Duration            githubv4.Int
		Iterations          []projectV2IterationFragment
		CompletedIterations []projectV2IterationFragment
	}
}

type projectV2IterationFragment struct {
	ID        githubv4.String
	Title     githubv4.String
	StartDate githubv4.Date
	Duration  githubv4.Int
}

type projectV2FieldNode struct {
	Typename githubv4.String `graphql:"__typename"`
	Field    struct {
		projectV2FieldFragment
	} `graphql:"... on ProjectV2Field"`
	SingleSelect struct {
		projectV2SingleSelectFieldFragment
	} `graphql:"... on ProjectV2SingleSelectField"`
	Iteration struct {
		projectV2IterationFieldFragment
	} `graphql:"... on ProjectV2IterationField"`
}

func resourceGithubProjectField() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages a custom field in a GitHub Projects V2 project.",
		CreateContext: resourceGithubProjectFieldCreate,
		ReadContext:   resourceGithubProjectFieldRead,
		UpdateContext: resourceGithubProjectFieldUpdate,
		DeleteContext: resourceGithubProjectFieldDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CustomizeDiff: customdiff.Sequence(resourceGithubProjectFieldValidate),

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node ID of the Projects V2 project.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the field.",
			},
			"data_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(projectV2FieldTypes, false)),
				Description:      "Field data type: `TEXT`, `SINGLE_SELECT`, `NUMBER`, `DATE`, or `ITERATION`.",
			},
			"single_select_option": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Ordered options for a `SINGLE_SELECT` field.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"description": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"color": {
						Type:             schema.TypeString,
						Required:         true,
						ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(projectV2FieldColors, false)),
					},
				}},
			},
			"iteration_configuration": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Configuration for an `ITERATION` field.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"start_date": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validateProjectV2Date,
					},
					"duration": {
						Type:         schema.TypeInt,
						Required:     true,
						ValidateFunc: validation.IntAtLeast(1),
					},
					"iteration": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"id": {
								Type:     schema.TypeString,
								Computed: true,
							},
							"title": {
								Type:     schema.TypeString,
								Required: true,
							},
							"start_date": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validateProjectV2Date,
							},
							"duration": {
								Type:         schema.TypeInt,
								Required:     true,
								ValidateFunc: validation.IntAtLeast(1),
							},
						}},
					},
				}},
			},
		},
	}
}

func resourceGithubProjectFieldValidate(_ context.Context, d *schema.ResourceDiff, _ any) error {
	dataType := d.Get("data_type").(string)
	options := d.Get("single_select_option").([]any)
	iterations := d.Get("iteration_configuration").([]any)
	if dataType == "SINGLE_SELECT" && len(options) == 0 {
		return fmt.Errorf("single_select_option must contain at least one option for a SINGLE_SELECT field")
	}
	if dataType != "SINGLE_SELECT" && len(options) != 0 {
		return fmt.Errorf("single_select_option is only valid for a SINGLE_SELECT field")
	}
	if dataType == "ITERATION" && len(iterations) == 0 {
		return fmt.Errorf("iteration_configuration is required for an ITERATION field")
	}
	if dataType != "ITERATION" && len(iterations) != 0 {
		return fmt.Errorf("iteration_configuration is only valid for an ITERATION field")
	}
	return nil
}

func resourceGithubProjectFieldCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	input := githubv4.CreateProjectV2FieldInput{
		ProjectID: githubv4.ID(d.Get("project_id").(string)),
		Name:      githubv4.String(d.Get("name").(string)),
		DataType:  githubv4.ProjectV2CustomFieldType(d.Get("data_type").(string)),
	}
	if err := expandProjectV2FieldConfiguration(d, &input.SingleSelectOptions, &input.IterationConfiguration); err != nil {
		return diag.FromErr(err)
	}
	var mutation struct {
		CreateProjectV2Field struct {
			Field projectV2FieldNode `graphql:"projectV2Field"`
		} `graphql:"createProjectV2Field(input: $input)"`
	}
	if err := meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil); err != nil {
		return diag.FromErr(err)
	}
	id := projectV2FieldNodeID(mutation.CreateProjectV2Field.Field)
	if id == "" {
		return diag.Errorf("GitHub returned a Projects V2 field without an ID")
	}
	d.SetId(id)
	return resourceGithubProjectFieldRead(ctx, d, meta)
}

func resourceGithubProjectFieldRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	field, err := queryProjectV2Field(ctx, meta.(*Owner).v4client, d.Id())
	if isProjectV2NotFound(err) {
		tflog.Info(ctx, "Removing project field from state because it no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if err := setProjectV2FieldState(d, field); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectFieldUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	name := githubv4.String(d.Get("name").(string))
	input := githubv4.UpdateProjectV2FieldInput{FieldID: githubv4.ID(d.Id()), Name: &name}
	var singleSelectOptions *[]githubv4.ProjectV2SingleSelectFieldOptionInput
	var iterationConfiguration *githubv4.ProjectV2IterationFieldConfigurationInput
	if d.HasChange("single_select_option") || d.HasChange("iteration_configuration") {
		if err := expandProjectV2FieldConfiguration(d, &singleSelectOptions, &iterationConfiguration); err != nil {
			return diag.FromErr(err)
		}
	}
	input.SingleSelectOptions = singleSelectOptions
	input.IterationConfiguration = iterationConfiguration
	var mutation struct {
		UpdateProjectV2Field struct {
			Field projectV2FieldNode `graphql:"projectV2Field"`
		} `graphql:"updateProjectV2Field(input: $input)"`
	}
	if err := meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil); err != nil {
		return diag.FromErr(err)
	}
	return resourceGithubProjectFieldRead(ctx, d, meta)
}

func resourceGithubProjectFieldDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var mutation struct {
		DeleteProjectV2Field struct {
			ClientMutationID githubv4.String
		} `graphql:"deleteProjectV2Field(input: $input)"`
	}
	err := meta.(*Owner).v4client.Mutate(ctx, &mutation, githubv4.DeleteProjectV2FieldInput{FieldID: githubv4.ID(d.Id())}, nil)
	if err != nil && !isProjectV2NotFound(err) {
		return diag.FromErr(err)
	}
	return nil
}

func queryProjectV2Field(ctx context.Context, client *githubv4.Client, id string) (projectV2FieldNode, error) {
	var query struct {
		Node projectV2FieldNode `graphql:"node(id: $id)"`
	}
	err := client.Query(ctx, &query, map[string]any{"id": githubv4.ID(id)})
	return query.Node, err
}

func projectV2FieldNodeID(field projectV2FieldNode) string {
	switch field.Typename {
	case "ProjectV2Field":
		return string(field.Field.ID)
	case "ProjectV2SingleSelectField":
		return string(field.SingleSelect.ID)
	case "ProjectV2IterationField":
		return string(field.Iteration.ID)
	}
	return ""
}

func expandProjectV2FieldConfiguration(d *schema.ResourceData, options **[]githubv4.ProjectV2SingleSelectFieldOptionInput, configuration **githubv4.ProjectV2IterationFieldConfigurationInput) error {
	switch d.Get("data_type").(string) {
	case "SINGLE_SELECT":
		expanded := make([]githubv4.ProjectV2SingleSelectFieldOptionInput, 0)
		for _, raw := range d.Get("single_select_option").([]any) {
			option := raw.(map[string]any)
			expanded = append(expanded, githubv4.ProjectV2SingleSelectFieldOptionInput{
				Name:        githubv4.String(option["name"].(string)),
				Description: githubv4.String(option["description"].(string)),
				Color:       githubv4.ProjectV2SingleSelectFieldOptionColor(option["color"].(string)),
			})
		}
		*options = &expanded
	case "ITERATION":
		raw := d.Get("iteration_configuration").([]any)[0].(map[string]any)
		startDate, err := time.Parse(time.DateOnly, raw["start_date"].(string))
		if err != nil {
			return err
		}
		iterations := make([]githubv4.ProjectV2Iteration, 0)
		for _, value := range raw["iteration"].([]any) {
			iteration := value.(map[string]any)
			iterationStart, err := time.Parse(time.DateOnly, iteration["start_date"].(string))
			if err != nil {
				return err
			}
			iterations = append(iterations, githubv4.ProjectV2Iteration{
				Title:     githubv4.String(iteration["title"].(string)),
				StartDate: githubv4.Date{Time: iterationStart},
				Duration:  githubv4.Int(iteration["duration"].(int)),
			})
		}
		*configuration = &githubv4.ProjectV2IterationFieldConfigurationInput{
			StartDate:  githubv4.Date{Time: startDate},
			Duration:   githubv4.Int(raw["duration"].(int)),
			Iterations: iterations,
		}
	}
	return nil
}

func setProjectV2FieldState(d *schema.ResourceData, field projectV2FieldNode) error {
	var base projectV2FieldFragment
	var options []map[string]any
	var configuration []map[string]any
	switch field.Typename {
	case "ProjectV2Field":
		base = field.Field.projectV2FieldFragment
	case "ProjectV2SingleSelectField":
		base = field.SingleSelect.projectV2FieldFragment
		for _, option := range field.SingleSelect.Options {
			options = append(options, map[string]any{"id": option.ID, "name": option.Name, "description": option.Description, "color": option.Color})
		}
	case "ProjectV2IterationField":
		base = field.Iteration.projectV2FieldFragment
		iterations := make([]map[string]any, 0, len(field.Iteration.Configuration.Iterations))
		for _, iteration := range field.Iteration.Configuration.Iterations {
			iterations = append(iterations, map[string]any{
				"id": iteration.ID, "title": iteration.Title, "start_date": iteration.StartDate.Format(time.DateOnly), "duration": int(iteration.Duration),
			})
		}
		startDate := projectV2IterationStartDate(d, field.Iteration.Configuration.CompletedIterations, field.Iteration.Configuration.Iterations)
		configuration = []map[string]any{{"start_date": startDate, "duration": int(field.Iteration.Configuration.Duration), "iteration": iterations}}
	default:
		return fmt.Errorf("GitHub returned an unsupported Projects V2 field type")
	}

	values := map[string]any{"project_id": base.Project.ID, "name": base.Name, "data_type": base.DataType, "single_select_option": options, "iteration_configuration": configuration}
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("setting %s: %w", key, err)
		}
	}
	return nil
}

func projectV2IterationStartDate(d *schema.ResourceData, completed, current []projectV2IterationFragment) string {
	if configured := d.Get("iteration_configuration").([]any); len(configured) > 0 && configured[0] != nil {
		if startDate, ok := configured[0].(map[string]any)["start_date"].(string); ok && startDate != "" {
			return startDate
		}
	}

	var earliest time.Time
	for _, iteration := range append(completed, current...) {
		if earliest.IsZero() || iteration.StartDate.Before(earliest) {
			earliest = iteration.StartDate.Time
		}
	}
	if earliest.IsZero() {
		return ""
	}
	return earliest.Format(time.DateOnly)
}
