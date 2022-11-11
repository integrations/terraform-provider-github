package github

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubTeamSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTeamSettingsCreate,
		Read:   resourceGithubTeamSettingsRead,
		Update: resourceGithubTeamSettingsUpdate,
		Delete: resourceGithubTeamSettingsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID or slug of team",
			},
			"review_request_algorithm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Algorithm to use when determining team members to be assigned to a pull request. Allowed values are ROUND_ROBIN and LOAD_BALANCE",
				Default:     "ROUND_ROBIN",
				ValidateFunc: func(v interface{}, key string) (we []string, errs []error) {
					algorithm, ok := v.(string)
					if !ok {
						return nil, []error{fmt.Errorf("expected type of %s to be string", key)}
					}

					if !(algorithm == "ROUND_ROBIN" || algorithm == "LOAD_BALANCE") {
						errs = append(errs, errors.New("review request delegation algorithm must be one of [\"ROUND_ROBIN\", \"LOAD_BALANCE\"]"))
					}

					return we, errs
				},
			},
			"review_request_delegation": {
				Type:         schema.TypeBool,
				Default:      false,
				Optional:     true,
				Description:  "Enable review request delegation for the Team",
				RequiredWith: []string{"review_request_algorithm", "review_request_count", "review_request_notify"},
			},
			"review_request_count": {
				Type:         schema.TypeInt,
				Optional:     true,
				RequiredWith: []string{"review_request_delegation"},
				Description:  "The number of reviewers to be assigned to a pull request from this team",
				ValidateFunc: func(v interface{}, key string) (we []string, errs []error) {
					count, ok := v.(int)
					if !ok {
						return nil, []error{fmt.Errorf("expected type of %s to be an integer", key)}
					}
					if count <= 0 {
						errs = append(errs, errors.New("review request delegation reviewer count must be a positive number"))
					}
					return we, errs
				},
			},
			"review_request_notify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Notify the entire team when a pull request is assigned to a member of the team",
			},
		},
	}
}

func resourceGithubTeamSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	return nil

}

func resourceGithubTeamSettingsRead(d *schema.ResourceData, meta interface{}) error {
	return nil

}

func resourceGithubTeamSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil

}

func resourceGithubTeamSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
