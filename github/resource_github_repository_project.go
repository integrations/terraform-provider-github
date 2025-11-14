package github

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryProject() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Repository projects have been replaced by Projects V2 which are organization or user-scoped. This resource is no longer functional and will be removed in a future version. Please use github_organization_project instead.",

		Create: resourceGithubRepositoryProjectCreate,
		Read:   resourceGithubRepositoryProjectRead,
		Update: resourceGithubRepositoryProjectUpdate,
		Delete: resourceGithubRepositoryProjectDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				parts := strings.Split(d.Id(), "/")
				if len(parts) != 2 {
					return nil, fmt.Errorf("invalid ID specified: supplied ID must be written as <repository>/<project_id>")
				}
				if err := d.Set("repository", parts[0]); err != nil {
					return nil, err
				}
				d.SetId(parts[1])
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the project.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository of the project.",
			},
			"body": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The body of the project.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the project",
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
				DiffSuppressOnRefresh: true,
			},
		},
	}
}

func resourceGithubRepositoryProjectCreate(d *schema.ResourceData, meta interface{}) error {
	// Repository projects have been replaced by Projects V2 which are organization or user-scoped
	// Projects cannot be created via the REST API
	return fmt.Errorf("Repository projects are no longer supported. Projects V2 are organization or user-scoped and cannot be created via the REST API. Please create the project through the GitHub web interface and use github_organization_project instead")
}

func resourceGithubRepositoryProjectRead(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Repository projects are no longer supported. Projects V2 are organization or user-scoped. Please migrate to github_organization_project")
}

func resourceGithubRepositoryProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Repository projects are no longer supported. Projects V2 are organization or user-scoped and cannot be updated via the REST API")
}

func resourceGithubRepositoryProjectDelete(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Repository projects are no longer supported. Projects V2 are organization or user-scoped and cannot be deleted via the REST API")
}
