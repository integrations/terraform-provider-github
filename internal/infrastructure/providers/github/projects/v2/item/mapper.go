package item

import (
	"fmt"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

func resultFromNode(value node) (application.Result, error) {
	if value.ID == "" {
		return application.Result{}, fmt.Errorf("GitHub returned a Projects V2 item without an ID")
	}
	contentID := value.Content.Issue.ID
	if contentID == "" {
		contentID = value.Content.PullRequest.ID
	}
	if contentID == "" {
		contentID = value.Content.DraftIssue.ID
	}
	if contentID == "" {
		return application.Result{}, fmt.Errorf("project item %q has no supported content", value.ID)
	}
	return application.Result{ID: string(value.ID), ProjectID: string(value.Project.ID), ContentID: string(contentID), Archived: bool(value.IsArchived)}, nil
}
