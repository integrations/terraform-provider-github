package github

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	fieldapplication "github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
	fieldusecases "github.com/integrations/terraform-provider-github/v6/internal/application/projects/field/use-cases"
	fieldgithub "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/field"
)

var projectV2FieldTypes = []string{"TEXT", "SINGLE_SELECT", "NUMBER", "DATE", "ITERATION"}

var projectV2FieldColors = []string{"GRAY", "BLUE", "GREEN", "YELLOW", "ORANGE", "RED", "PINK", "PURPLE"}

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
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Stable node ID of the single-select option.",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Display name of the single-select option.",
					},
					"description": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Description of the single-select option.",
					},
					"color": {
						Type:             schema.TypeString,
						Required:         true,
						ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(projectV2FieldColors, false)),
						Description:      "Color of the single-select option.",
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
						Type:             schema.TypeString,
						Required:         true,
						ValidateDiagFunc: validateProjectV2Date,
						Description:      "Start date of the iteration schedule in YYYY-MM-DD format.",
					},
					"duration": {
						Type:             schema.TypeInt,
						Required:         true,
						ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
						Description:      "Default iteration duration in days.",
					},
					"iteration": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Configured iterations in schedule order.",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"id": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "Stable node ID of the iteration.",
							},
							"title": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "Display title of the iteration.",
							},
							"start_date": {
								Type:             schema.TypeString,
								Required:         true,
								ValidateDiagFunc: validateProjectV2Date,
								Description:      "Start date of the iteration in YYYY-MM-DD format.",
							},
							"duration": {
								Type:             schema.TypeInt,
								Required:         true,
								ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
								Description:      "Iteration duration in days.",
							},
						}},
					},
				}},
			},
		},
	}
}

func resourceGithubProjectFieldValidate(_ context.Context, d *schema.ResourceDiff, _ any) error {
	dataType := projectV2Get[string](d, "data_type")
	options := projectV2Get[[]any](d, "single_select_option")
	iterations := projectV2Get[[]any](d, "iteration_configuration")
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
	configuration, err := expandProjectV2FieldConfiguration(d)
	if err != nil {
		return diag.FromErr(err)
	}
	input := fieldapplication.CreateInput{
		ProjectID: projectV2Get[string](d, "project_id"), Name: projectV2Get[string](d, "name"), DataType: projectV2Get[string](d, "data_type"), Configuration: configuration,
	}
	field, err := fieldusecases.NewCreate(fieldgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, input)
	if err != nil {
		return diag.FromErr(err)
	}
	if field.ID == "" {
		return diag.Errorf("GitHub returned a Projects V2 field without an ID")
	}
	d.SetId(field.ID)
	if err := setProjectV2FieldState(d, field); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectFieldRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	field, err := fieldusecases.NewGet(fieldgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, d.Id())
	if errors.Is(err, projects.ErrNotFound) {
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
	input := fieldapplication.UpdateInput{ID: d.Id(), Name: projectV2Get[string](d, "name")}
	if d.HasChange("single_select_option") || d.HasChange("iteration_configuration") {
		configuration, err := expandProjectV2FieldConfiguration(d)
		if err != nil {
			return diag.FromErr(err)
		}
		input.Configuration = &configuration
	}
	field, err := fieldusecases.NewUpdate(fieldgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, input)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := setProjectV2FieldState(d, field); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectFieldDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := fieldusecases.NewDelete(fieldgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, d.Id())
	if err != nil && !errors.Is(err, projects.ErrNotFound) {
		return diag.FromErr(err)
	}
	return nil
}

func expandProjectV2FieldConfiguration(d *schema.ResourceData) (fieldapplication.Configuration, error) {
	var configuration fieldapplication.Configuration
	switch projectV2Get[string](d, "data_type") {
	case "SINGLE_SELECT":
		configuration.SingleSelectOptions = make([]fieldapplication.SingleSelectOptionInput, 0)
		for _, raw := range projectV2Get[[]any](d, "single_select_option") {
			option := projectV2As[map[string]any](raw, "single_select_option")
			configuration.SingleSelectOptions = append(configuration.SingleSelectOptions, fieldapplication.SingleSelectOptionInput{
				Name: projectV2MapGet[string](option, "name"), Description: projectV2MapGet[string](option, "description"), Color: projectV2MapGet[string](option, "color"),
			})
		}
	case "ITERATION":
		raw := projectV2As[map[string]any](projectV2Get[[]any](d, "iteration_configuration")[0], "iteration_configuration")
		startDate, err := time.Parse(time.DateOnly, projectV2MapGet[string](raw, "start_date"))
		if err != nil {
			return configuration, err
		}
		iterations := make([]fieldapplication.IterationInput, 0)
		for _, value := range projectV2MapGet[[]any](raw, "iteration") {
			iteration := projectV2As[map[string]any](value, "iteration")
			iterationStart, err := time.Parse(time.DateOnly, projectV2MapGet[string](iteration, "start_date"))
			if err != nil {
				return configuration, err
			}
			iterations = append(iterations, fieldapplication.IterationInput{
				Title: projectV2MapGet[string](iteration, "title"), StartDate: iterationStart, Duration: projectV2MapGet[int](iteration, "duration"),
			})
		}
		configuration.Iteration = &fieldapplication.IterationConfigurationInput{
			StartDate: startDate, Duration: projectV2MapGet[int](raw, "duration"), Iterations: iterations,
		}
	}
	return configuration, nil
}

func setProjectV2FieldState(d *schema.ResourceData, field fieldapplication.Result) error {
	var options []map[string]any
	var configuration []map[string]any
	if field.SingleSelectOptions != nil {
		for _, option := range field.SingleSelectOptions {
			options = append(options, map[string]any{"id": option.ID, "name": option.Name, "description": option.Description, "color": option.Color})
		}
	}
	if field.IterationConfiguration != nil {
		iterations := make([]map[string]any, 0, len(field.IterationConfiguration.Iterations))
		for _, iteration := range field.IterationConfiguration.Iterations {
			iterations = append(iterations, map[string]any{
				"id": iteration.ID, "title": iteration.Title, "start_date": iteration.StartDate.Format(time.DateOnly), "duration": iteration.Duration,
			})
		}
		startDate := projectV2IterationStartDate(d, field.IterationConfiguration.CompletedIterations, field.IterationConfiguration.Iterations)
		configuration = []map[string]any{{"start_date": startDate, "duration": field.IterationConfiguration.Duration, "iteration": iterations}}
	}

	values := map[string]any{"project_id": field.ProjectID, "name": field.Name, "data_type": field.DataType, "single_select_option": options, "iteration_configuration": configuration}
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("setting %s: %w", key, err)
		}
	}
	return nil
}

func projectV2IterationStartDate(d *schema.ResourceData, completed, current []fieldapplication.Iteration) string {
	if configured := projectV2Get[[]any](d, "iteration_configuration"); len(configured) > 0 && configured[0] != nil {
		configuration := projectV2As[map[string]any](configured[0], "iteration_configuration")
		if startDate, ok := configuration["start_date"].(string); ok && startDate != "" {
			return startDate
		}
	}

	var earliest time.Time
	for _, iteration := range append(completed, current...) {
		if earliest.IsZero() || iteration.StartDate.Before(earliest) {
			earliest = iteration.StartDate
		}
	}
	if earliest.IsZero() {
		return ""
	}
	return earliest.Format(time.DateOnly)
}
