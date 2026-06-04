package github

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

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

	log.Printf("[DEBUG] GitHub Repository Collaborators Attributes before migration to v1: %#v", rawState)

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

	log.Printf("[DEBUG] GitHub Repository Collaborators Attributes after migration to v1: %#v", rawState)

	return rawState, nil
}

func resourceGithubRepositoryCollaboratorsV1() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the repository.",
			},
			"user": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Users to grant access to the repository.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:             schema.TypeString,
							Description:      "Login for the user to add to the repository as a collaborator.",
							Required:         true,
							DiffSuppressFunc: caseInsensitive(),
						},
						"permission": {
							Type:        schema.TypeString,
							Description: "Permission to grant to the user. Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization. Must be `push` for personal repositories. Defaults to `push`.",
							Optional:    true,
							Default:     "push",
						},
					},
				},
			},
			"team": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Teams to grant access to the repository.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"team_id": {
							Type:        schema.TypeString,
							Description: "ID or slug of the team to add to the repository as a collaborator.",
							Required:    true,
						},
						"permission": {
							Type:        schema.TypeString,
							Description: "Permission to grant to the team. Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization. Defaults to `push`.",
							Optional:    true,
							Default:     "push",
						},
					},
				},
			},
			"invitation_ids": {
				Type:        schema.TypeMap,
				Description: "Map of usernames to invitation ID for users that haven't yet accepted their invitation to become a collaborator. This is only set on read, and is used internally to track pending invitations for users that aren't yet collaborators.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"ignore_team": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Teams to ignore when managing repository collaborators.",
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

func resourceGithubRepositoryCollaboratorsStateUpgradeV1(_ context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta, _ := m.(*Owner)

	log.Printf("[DEBUG] GitHub Repository Collaborators Attributes before migration to v2: %#v", rawState)

	if meta.IsOrganization {
		// If the repository belongs to an organization the owner cannot be a,
		// collaborator, so owner_configured is always false.

		rawState["owner_configured"] = false
	} else {
		// If the repository belongs to a user and we know the new value of user
		// we can determine the value of owner_configured by checking if the owner
		// is included in the list of users.

		ownerConfigured := false
		owner := strings.ToLower(meta.name)

		if usersVal, ok := rawState["user"]; ok {
			if users, ok := usersVal.([]any); ok {
				for _, userVal := range users {
					user, ok := userVal.(map[string]any)
					if !ok {
						continue
					}

					usernameVal, ok := user["username"]
					if !ok {
						continue
					}

					username, ok := usernameVal.(string)
					if !ok {
						continue
					}

					if strings.ToLower(username) == owner {
						ownerConfigured = true
						break
					}
				}
			}
		}

		rawState["owner_configured"] = ownerConfigured
	}

	log.Printf("[DEBUG] GitHub Repository Collaborators Attributes after migration to v2: %#v", rawState)

	return rawState, nil
}
