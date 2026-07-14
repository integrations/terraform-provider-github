package link

import application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/repository/link"

func resultFromNode(projectID string, value repositoryNode) application.Result {
	return application.Result{ProjectID: projectID, RepositoryID: string(value.ID), Owner: string(value.Owner.Login), Name: string(value.Name)}
}
