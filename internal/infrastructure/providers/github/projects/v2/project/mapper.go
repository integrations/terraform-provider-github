package project

import (
	"fmt"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

func resultFromNode(value node) (application.Result, error) {
	if value.ID == "" {
		return application.Result{}, fmt.Errorf("GitHub returned a Projects V2 project without an ID")
	}
	result := application.Result{
		ID: string(value.ID), Number: int(value.Number), Title: string(value.Title),
		ShortDescription: string(value.ShortDescription), Readme: string(value.Readme),
		Public: bool(value.Public), Closed: bool(value.Closed),
	}
	if value.URL.URL != nil {
		result.URL = value.URL.String()
	}
	switch value.Owner.Typename {
	case "Organization":
		result.OwnerKind = application.OwnerOrganization
		result.Owner = string(value.Owner.Organization.Login)
		result.OwnerID = int(value.Owner.Organization.DatabaseID)
	case "User":
		result.OwnerKind = application.OwnerUser
		result.Owner = string(value.Owner.User.Login)
		result.OwnerID = int(value.Owner.User.DatabaseID)
	default:
		return application.Result{}, fmt.Errorf("GitHub returned unsupported Projects V2 owner type %q", value.Owner.Typename)
	}
	return result, nil
}
