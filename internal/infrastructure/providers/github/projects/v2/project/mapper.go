package project

import application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"

func resultFromNode(value node) application.Result {
	result := application.Result{
		ID: string(value.ID), Number: int(value.Number), Title: string(value.Title),
		ShortDescription: string(value.ShortDescription), Readme: string(value.Readme),
		Public: bool(value.Public), Closed: bool(value.Closed),
	}
	if value.URL.URL != nil {
		result.URL = value.URL.String()
	}
	if value.Owner.Typename == "Organization" {
		result.OwnerKind = application.OwnerOrganization
		result.Owner = string(value.Owner.Organization.Login)
	} else {
		result.OwnerKind = application.OwnerUser
		result.Owner = string(value.Owner.User.Login)
	}
	return result
}
