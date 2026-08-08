package value

import (
	"fmt"

	"github.com/shurcooL/githubv4"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
	projectgraphql "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/graphql"
)

func resultFromNode(value node) (application.Result, error) {
	switch value.Typename {
	case "ProjectV2ItemFieldTextValue":
		return application.Result{Kind: application.KindText, Text: string(value.Text.Text)}, nil
	case "ProjectV2ItemFieldNumberValue":
		return application.Result{Kind: application.KindNumber, Number: float64(value.Number.Number)}, nil
	case "ProjectV2ItemFieldDateValue":
		date, err := projectgraphql.ParseDate(string(value.Date.Date), "Projects V2 item field")
		if err != nil {
			return application.Result{}, err
		}
		return application.Result{Kind: application.KindDate, Date: date}, nil
	case "ProjectV2ItemFieldSingleSelectValue":
		return application.Result{Kind: application.KindSingleSelect, OptionID: string(value.SingleSelect.OptionID)}, nil
	case "ProjectV2ItemFieldIterationValue":
		return application.Result{Kind: application.KindIteration, IterationID: string(value.Iteration.IterationID)}, nil
	case "":
		return application.Result{}, nil
	default:
		return application.Result{}, fmt.Errorf("GitHub returned unsupported Projects V2 item field value type %q", value.Typename)
	}
}

func nodeFieldID(value node) string {
	ids := []githubv4.String{
		value.Field.Field.ID,
		value.Field.SingleSelect.ID,
		value.Field.Iteration.ID,
	}
	for _, id := range ids {
		if id != "" {
			return string(id)
		}
	}
	return ""
}

func fieldValueInput(value application.Result) githubv4.ProjectV2FieldValue {
	var result githubv4.ProjectV2FieldValue
	switch value.Kind {
	case application.KindText:
		text := githubv4.String(value.Text)
		result.Text = &text
	case application.KindNumber:
		number := githubv4.Float(value.Number)
		result.Number = &number
	case application.KindDate:
		date := githubv4.Date{Time: value.Date}
		result.Date = &date
	case application.KindSingleSelect:
		optionID := githubv4.String(value.OptionID)
		result.SingleSelectOptionID = &optionID
	case application.KindIteration:
		iterationID := githubv4.String(value.IterationID)
		result.IterationID = &iterationID
	}
	return result
}
