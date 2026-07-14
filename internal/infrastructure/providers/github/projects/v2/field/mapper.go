package field

import (
	"fmt"

	"github.com/shurcooL/githubv4"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
	projectgraphql "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/graphql"
)

func resultFromNode(value node) (application.Result, error) {
	var base baseNode
	result := application.Result{}
	switch value.Typename {
	case "ProjectV2Field":
		base = value.Field.baseNode
	case "ProjectV2SingleSelectField":
		base = value.SingleSelect.baseNode
		for _, option := range value.SingleSelect.Options {
			result.SingleSelectOptions = append(result.SingleSelectOptions, application.SingleSelectOption{ID: string(option.ID), Name: string(option.Name), Description: string(option.Description), Color: string(option.Color)})
		}
	case "ProjectV2IterationField":
		base = value.Iteration.baseNode
		iterations, err := iterationsFromNodes(value.Iteration.Configuration.Iterations)
		if err != nil {
			return application.Result{}, err
		}
		completed, err := iterationsFromNodes(value.Iteration.Configuration.CompletedIterations)
		if err != nil {
			return application.Result{}, err
		}
		result.IterationConfiguration = &application.IterationConfiguration{Duration: int(value.Iteration.Configuration.Duration), Iterations: iterations, CompletedIterations: completed}
	default:
		return application.Result{}, fmt.Errorf("GitHub returned unsupported Projects V2 field type %q", value.Typename)
	}
	result.ID, result.ProjectID, result.Name, result.DataType = string(base.ID), string(base.Project.ID), string(base.Name), string(base.DataType)
	return result, nil
}

func iterationsFromNodes(values []iterationValue) ([]application.Iteration, error) {
	results := make([]application.Iteration, 0, len(values))
	for _, value := range values {
		date, err := projectgraphql.ParseDate(string(value.StartDate), fmt.Sprintf("Projects V2 iteration %q start", value.ID))
		if err != nil {
			return nil, err
		}
		results = append(results, application.Iteration{ID: string(value.ID), Title: string(value.Title), StartDate: date, Duration: int(value.Duration)})
	}
	return results, nil
}

func nodeID(value node) string {
	switch value.Typename {
	case "ProjectV2Field":
		return string(value.Field.ID)
	case "ProjectV2SingleSelectField":
		return string(value.SingleSelect.ID)
	case "ProjectV2IterationField":
		return string(value.Iteration.ID)
	default:
		return ""
	}
}

func configurationInput(configuration application.Configuration) (*[]githubv4.ProjectV2SingleSelectFieldOptionInput, *githubv4.ProjectV2IterationFieldConfigurationInput) {
	var options *[]githubv4.ProjectV2SingleSelectFieldOptionInput
	if configuration.SingleSelectOptions != nil {
		expanded := make([]githubv4.ProjectV2SingleSelectFieldOptionInput, 0, len(configuration.SingleSelectOptions))
		for _, option := range configuration.SingleSelectOptions {
			expanded = append(expanded, githubv4.ProjectV2SingleSelectFieldOptionInput{Name: githubv4.String(option.Name), Description: githubv4.String(option.Description), Color: githubv4.ProjectV2SingleSelectFieldOptionColor(option.Color)})
		}
		options = &expanded
	}
	var iteration *githubv4.ProjectV2IterationFieldConfigurationInput
	if configuration.Iteration != nil {
		iterations := make([]githubv4.ProjectV2Iteration, 0, len(configuration.Iteration.Iterations))
		for _, value := range configuration.Iteration.Iterations {
			iterations = append(iterations, githubv4.ProjectV2Iteration{Title: githubv4.String(value.Title), StartDate: githubv4.Date{Time: value.StartDate}, Duration: githubv4.Int(value.Duration)})
		}
		iteration = &githubv4.ProjectV2IterationFieldConfigurationInput{StartDate: githubv4.Date{Time: configuration.Iteration.StartDate}, Duration: githubv4.Int(configuration.Iteration.Duration), Iterations: iterations}
	}
	return options, iteration
}
