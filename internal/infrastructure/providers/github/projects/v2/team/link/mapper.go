package link

import application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link"

func resultFromNode(projectID string, value teamNode) application.Result {
	return application.Result{ProjectID: projectID, TeamID: string(value.ID), Organization: string(value.Organization.Login), Slug: string(value.Slug)}
}
