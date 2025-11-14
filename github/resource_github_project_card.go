package github

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubProjectCard() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This resource is deprecated as GitHub Classic Projects have been sunset. Use the 'github_project_item' resource for GitHub Projects V2 instead.",

		Create: resourceGithubProjectCardCreate,
		Read:   resourceGithubProjectCardRead,
		Update: resourceGithubProjectCardUpdate,
		Delete: resourceGithubProjectCardDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubProjectCardImport,
		},
		Schema: map[string]*schema.Schema{
			"column_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the project column.",
			},
			"note": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The note contents of the card. Markdown supported.",
			},
			"content_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "'github_issue.issue_id'.",
			},
			"content_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Must be either 'Issue' or 'PullRequest'.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"card_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the card.",
			},
		},
	}
}

func resourceGithubProjectCardCreate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("github_project_card has been deprecated as GitHub Classic Projects were sunset on May 23, 2024. Please migrate to the 'github_project_item' resource for GitHub Projects V2. See: https://github.blog/changelog/2024-05-23-sunset-notice-projects-classic/")
}

func resourceGithubProjectCardRead(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("github_project_card has been deprecated as GitHub Classic Projects were sunset on May 23, 2024. Please migrate to the 'github_project_item' resource for GitHub Projects V2. See: https://github.blog/changelog/2024-05-23-sunset-notice-projects-classic/")
}

func resourceGithubProjectCardUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("github_project_card has been deprecated as GitHub Classic Projects were sunset on May 23, 2024. Please migrate to the 'github_project_item' resource for GitHub Projects V2. See: https://github.blog/changelog/2024-05-23-sunset-notice-projects-classic/")
}

func resourceGithubProjectCardDelete(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("github_project_card has been deprecated as GitHub Classic Projects were sunset on May 23, 2024. Please migrate to the 'github_project_item' resource for GitHub Projects V2. See: https://github.blog/changelog/2024-05-23-sunset-notice-projects-classic/")
}

func resourceGithubProjectCardImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return nil, fmt.Errorf("github_project_card has been deprecated as GitHub Classic Projects were sunset on May 23, 2024. Please migrate to the 'github_project_item' resource for GitHub Projects V2. See: https://github.blog/changelog/2024-05-23-sunset-notice-projects-classic/")
}
