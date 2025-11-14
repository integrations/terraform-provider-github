package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v77/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationProjects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationProjectsRead,

		Schema: map[string]*schema.Schema{
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The organization name.",
			},
			"projects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Projects V2 for the organization.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the project.",
						},
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node ID of the project.",
						},
						"number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of the project.",
						},
						"title": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The title of the project.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the project.",
						},
						"short_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The short description of the project.",
						},
						"public": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the project is public.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the project.",
						},
						"html_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTML URL of the project.",
						},
						"owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The login of the project owner.",
						},
						"creator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The login of the project creator.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp when the project was created.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp when the project was last updated.",
						},
						"closed_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp when the project was closed (if applicable).",
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp when the project was deleted (if applicable).",
						},
						"deleted_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The login of the user who deleted the project (if applicable).",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the project (open or closed).",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationProjectsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	orgName := d.Get("organization").(string)

	var allProjects []*github.ProjectV2
	var opts *github.ListProjectsOptions

	for {
		if opts == nil {
			opts = &github.ListProjectsOptions{
				ListProjectsPaginationOptions: github.ListProjectsPaginationOptions{PerPage: github.Int(100)},
			}
		}

		projects, resp, err := client.Projects.ListOrganizationProjects(ctx, orgName, opts)
		if err != nil {
			return fmt.Errorf("error listing organization Projects V2: %v", err)
		}

		allProjects = append(allProjects, projects...)

		if resp.After == "" {
			break
		}

		opts = &github.ListProjectsOptions{
			ListProjectsPaginationOptions: github.ListProjectsPaginationOptions{
				PerPage: github.Int(100),
				After:   github.String(resp.After),
			},
		}
	}

	d.SetId(orgName)

	projectsData := make([]map[string]interface{}, 0, len(allProjects))

	for _, project := range allProjects {
		projectData := map[string]interface{}{
			"id":                project.GetID(),
			"node_id":           project.GetNodeID(),
			"number":            project.GetNumber(),
			"title":             project.GetTitle(),
			"description":       project.GetDescription(),
			"short_description": project.GetShortDescription(),
			"public":            project.GetPublic(),
			"url":               project.GetURL(),
			"html_url":          project.GetHTMLURL(),
			"created_at":        project.GetCreatedAt().Format("2006-01-02T15:04:05Z"),
			"updated_at":        project.GetUpdatedAt().Format("2006-01-02T15:04:05Z"),
		}

		if project.GetClosedAt() != (github.Timestamp{}) {
			projectData["closed_at"] = project.GetClosedAt().Format("2006-01-02T15:04:05Z")
		}

		if project.GetDeletedAt() != (github.Timestamp{}) {
			projectData["deleted_at"] = project.GetDeletedAt().Format("2006-01-02T15:04:05Z")
		}

		if project.GetState() != "" {
			projectData["state"] = project.GetState()
		}

		if project.GetOwner() != nil {
			projectData["owner"] = project.GetOwner().GetLogin()
		}

		if project.GetCreator() != nil {
			projectData["creator"] = project.GetCreator().GetLogin()
		}

		if project.GetDeletedBy() != nil {
			projectData["deleted_by"] = project.GetDeletedBy().GetLogin()
		}

		projectsData = append(projectsData, projectData)
	}

	if err := d.Set("projects", projectsData); err != nil {
		return fmt.Errorf("error setting projects: %v", err)
	}

	return nil
}
