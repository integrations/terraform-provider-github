package github

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryCollaboratorsV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of users.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "push",
						},
						"username": {
							Type:             schema.TypeString,
							Description:      "(Required) The user to add to the repository as a collaborator.",
							Required:         true,
							DiffSuppressFunc: caseInsensitive(),
						},
					},
				},
			},
			"team": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of teams.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "push",
						},
						"team_id": {
							Type:        schema.TypeString,
							Description: "Team ID or slug to add to the repository as a collaborator.",
							Required:    true,
						},
					},
				},
			},
			"invitation_ids": {
				Type:        schema.TypeMap,
				Description: "Map of usernames to invitation ID for any users added",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"ignore_team": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of teams to ignore.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"team_id": {
							Type:        schema.TypeString,
							Description: "ID or slug of the team to ignore.",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func resourceGithubRepositoryCollaboratorsStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	log.Printf("[DEBUG] GitHub Repository Collaborators Attributes before migration: %#v", rawState)

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	rawState["id"] = strconv.FormatInt(repo.GetID(), 10)
	rawState["repository_id"] = int(repo.GetID())

	log.Printf("[DEBUG] GitHub Repository Collaborators Attributes after migration: %#v", rawState)

	return rawState, nil
}
