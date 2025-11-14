package github

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubProjectColumn() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This resource is deprecated as the API endpoints for classic projects have been removed. This resource no longer works and will be removed in a future version.",

		Create: resourceGithubProjectColumnCreate,
		Read:   resourceGithubProjectColumnRead,
		Update: resourceGithubProjectColumnUpdate,
		Delete: resourceGithubProjectColumnDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of an existing project that the column will be created in.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the column.",
			},
			"column_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the column.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubProjectColumnCreate(d *schema.ResourceData, meta interface{}) error {
	// Classic Project columns are not supported in Projects V2 API
	// Projects V2 uses custom fields instead of columns
	return fmt.Errorf("Classic project columns are no longer supported. GitHub Projects V2 uses custom fields instead of columns. Please migrate to Projects V2 and use the GitHub web interface to manage project fields")
}

func resourceGithubProjectColumnRead(d *schema.ResourceData, meta interface{}) error {
	// Classic Project columns are not supported in Projects V2 API
	return fmt.Errorf("Classic project columns are no longer supported. GitHub Projects V2 uses custom fields instead of columns. Please migrate to Projects V2 and use the GitHub web interface to manage project fields")
}

func resourceGithubProjectColumnUpdate(d *schema.ResourceData, meta interface{}) error {
	// Classic Project columns are not supported in Projects V2 API
	return fmt.Errorf("Classic project columns are no longer supported. GitHub Projects V2 uses custom fields instead of columns. Please migrate to Projects V2 and use the GitHub web interface to manage project fields")
}

func resourceGithubProjectColumnDelete(d *schema.ResourceData, meta interface{}) error {
	// Classic Project columns are not supported in Projects V2 API
	return fmt.Errorf("Classic project columns are no longer supported. GitHub Projects V2 uses custom fields instead of columns. Please migrate to Projects V2 and use the GitHub web interface to manage project fields")
}
