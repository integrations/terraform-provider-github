package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v79/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubProjectFields() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubProjectFieldsRead,

		Schema: map[string]*schema.Schema{
			"organization": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"username"},
				Description:   "The organization name (for organization-owned projects).",
			},
			"username": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"organization"},
				Description:   "The username (for user-owned projects).",
			},
			"project_number": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The project number.",
			},
			"fields": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of fields for the project.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the field.",
						},
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node ID of the field.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the field.",
						},
						"data_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The data type of the field (text, number, date, single_select, iteration).",
						},
						"project_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the project this field belongs to.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp when the field was created.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp when the field was last updated.",
						},
						"options": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Options for single_select fields.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the option.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the option.",
									},
									"name_html": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The HTML-formatted name of the option.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the option.",
									},
									"description_html": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The HTML-formatted description of the option.",
									},
									"color": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The color of the option.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubProjectFieldsRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	projectNumber := d.Get("project_number").(int)
	organization := d.Get("organization").(string)
	username := d.Get("username").(string)

	if organization == "" && username == "" {
		return fmt.Errorf("either organization or username must be specified")
	}

	var allFields []*github.ProjectV2Field
	var opts *github.ListProjectsOptions
	var err error

	for {
		if opts == nil {
			opts = &github.ListProjectsOptions{
				ListProjectsPaginationOptions: github.ListProjectsPaginationOptions{PerPage: github.Ptr(100)},
			}
		}

		var fields []*github.ProjectV2Field
		var resp *github.Response

		if organization != "" {
			fields, resp, err = client.Projects.ListOrganizationProjectFields(ctx, organization, projectNumber, opts)
		} else {
			fields, resp, err = client.Projects.ListUserProjectFields(ctx, username, projectNumber, opts)
		}

		if err != nil {
			return fmt.Errorf("error listing project fields: %w", err)
		}

		allFields = append(allFields, fields...)

		if resp.After == "" {
			break
		}

		opts = &github.ListProjectsOptions{
			ListProjectsPaginationOptions: github.ListProjectsPaginationOptions{
				PerPage: github.Ptr(100),
				After:   github.Ptr(resp.After),
			},
		}
	}

	// Set ID as organization/username:project_number
	var resourceID string
	if organization != "" {
		resourceID = fmt.Sprintf("%s:%d", organization, projectNumber)
	} else {
		resourceID = fmt.Sprintf("%s:%d", username, projectNumber)
	}
	d.SetId(resourceID)

	fieldsData := make([]map[string]any, 0, len(allFields))

	for _, field := range allFields {
		fieldData := map[string]any{
			"id":          field.GetID(),
			"node_id":     field.GetNodeID(),
			"name":        field.GetName(),
			"data_type":   field.GetDataType(),
			"project_url": field.GetProjectURL(),
			"created_at":  field.GetCreatedAt().Format("2006-01-02T15:04:05Z"),
			"updated_at":  field.GetUpdatedAt().Format("2006-01-02T15:04:05Z"),
		}

		// Add options for single_select fields
		if len(field.Options) > 0 {
			optionsData := make([]map[string]any, 0, len(field.Options))
			for _, option := range field.Options {
				optionData := map[string]any{
					"id":    option.GetID(),
					"color": option.GetColor(),
				}

				// Handle name field - it may be nil
				if option.Name != nil {
					optionData["name"] = *option.Name
				}

				// Handle description field - it may be nil
				if option.Description != nil {
					optionData["description"] = *option.Description
				}

				optionsData = append(optionsData, optionData)
			}
			fieldData["options"] = optionsData
		}

		fieldsData = append(fieldsData, fieldData)
	}

	if err := d.Set("fields", fieldsData); err != nil {
		return fmt.Errorf("error setting fields: %w", err)
	}

	return nil
}
