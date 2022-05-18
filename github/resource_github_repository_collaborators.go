package github

import (
	"context"
	"log"
	"sort"
	"strconv"

	"github.com/google/go-github/v43/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubRepositoryCollaborators() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryCollaboratorsCreate,
		Read:   resourceGithubRepositoryCollaboratorsRead,
		Update: resourceGithubRepositoryCollaboratorsUpdate,
		Delete: resourceGithubRepositoryCollaboratorsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of users",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "push",
							ValidateFunc: validateValueFunc([]string{"pull", "triage", "push", "maintain", "admin"}),
						},
						"username": {
							Type:             schema.TypeString,
							Description:      "(Required) The user to add to the repository as a collaborator.",
							Required:         true,
							DiffSuppressFunc: caseInsensitive(),
						},
					},
				},
				Set: func(val interface{}) int {
					return schema.HashString(val.(map[string]interface{})["username"])
				},
			},
			"team": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of teams",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "push",
							ValidateFunc: validateValueFunc([]string{"pull", "triage", "push", "maintain", "admin"}),
						},
						"team_id": {
							Type:         schema.TypeString,
							Description:  "(Required) Team ID to add to the repository as a collaborator.",
							Required:     true,
							ValidateFunc: validateTeamIDFunc,
						},
					},
				},
				Set: func(val interface{}) int {
					return schema.HashString(val.(map[string]interface{})["team_id"])
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
		},
	}
}

type userCollaborator struct {
	permission string
	username   string
}

func (c userCollaborator) Empty() bool {
	return c == userCollaborator{}
}

type invitedCollaborator struct {
	userCollaborator
	invitationID int64
}

func flattenUserCollaborator(obj userCollaborator) interface{} {
	if obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"permission": obj.permission,
		"username":   obj.username,
	}

	return transformed
}

func flattenUserCollaborators(objs []userCollaborator, invites []invitedCollaborator) []interface{} {
	if objs == nil {
		return nil
	}

	for _, invite := range invites {
		objs = append(objs, invite.userCollaborator)
	}

	sort.SliceStable(objs, func(i, j int) bool {
		return objs[i].username < objs[j].username
	})

	items := make([]interface{}, len(objs))
	for i, obj := range objs {
		items[i] = flattenUserCollaborator(obj)
	}

	return items
}

type teamCollaborator struct {
	permission string
	teamID     int64
}

func (c teamCollaborator) Empty() bool {
	return c == teamCollaborator{}
}

func flattenTeamCollaborator(obj teamCollaborator) interface{} {
	if obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"permission": obj.permission,
		"team_id":    strconv.FormatInt(obj.teamID, 10),
	}

	return transformed
}

func flattenTeamCollaborators(objs []teamCollaborator) []interface{} {
	if objs == nil {
		return nil
	}

	sort.SliceStable(objs, func(i, j int) bool {
		return objs[i].teamID < objs[j].teamID
	})

	items := make([]interface{}, len(objs))
	for i, obj := range objs {
		items[i] = flattenTeamCollaborator(obj)
	}

	return items
}

func listUserCollaborators(client *github.Client, ctx context.Context, owner, repoName string) ([]userCollaborator, error) {
	var userCollaborators []userCollaborator
	affiliations := []string{"direct", "outside"}
	for _, affiliation := range affiliations {
		opt := &github.ListCollaboratorsOptions{ListOptions: github.ListOptions{
			PerPage: maxPerPage,
		}, Affiliation: affiliation}

		for {
			collaborators, resp, err := client.Repositories.ListCollaborators(ctx,
				owner, repoName, opt)
			if err != nil {
				return nil, err
			}

			for _, c := range collaborators {
				permissionName, err := getRepoPermission(c.GetPermissions())
				if err != nil {
					return nil, err
				}

				userCollaborators = append(userCollaborators, userCollaborator{permissionName, c.GetLogin()})
			}

			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}
	return userCollaborators, nil
}

func listInvitations(client *github.Client, ctx context.Context, owner, repoName string) ([]invitedCollaborator, error) {
	var invitedCollaborators []invitedCollaborator

	opt := &github.ListOptions{PerPage: maxPerPage}
	for {
		invitations, resp, err := client.Repositories.ListInvitations(ctx, owner, repoName, opt)
		if err != nil {
			return nil, err
		}

		for _, i := range invitations {
			permissionName, err := getInvitationPermission(i)
			if err != nil {
				return nil, err
			}

			invitedCollaborators = append(invitedCollaborators, invitedCollaborator{
				userCollaborator{permissionName, i.GetInvitee().GetLogin()}, i.GetID()})
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return invitedCollaborators, nil
}

func listTeams(client *github.Client, ctx context.Context, owner, repoName string) ([]teamCollaborator, error) {
	var teamCollaborators []teamCollaborator

	opt := &github.ListOptions{PerPage: maxPerPage}
	for {
		repoTeams, resp, err := client.Repositories.ListTeams(ctx, owner, repoName, opt)
		if err != nil {
			return nil, err
		}

		for _, t := range repoTeams {
			permissionName, err := getRepoPermission(t.GetPermissions())
			if err != nil {
				return nil, err
			}

			teamCollaborators = append(teamCollaborators, teamCollaborator{permissionName, t.GetID()})
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return teamCollaborators, nil
}

func listAllCollaborators(client *github.Client, ctx context.Context, owner, repoName string) ([]userCollaborator, []invitedCollaborator, []teamCollaborator, error) {
	userCollaborators, err := listUserCollaborators(client, ctx, owner, repoName)
	if err != nil {
		return nil, nil, nil, err
	}
	invitations, err := listInvitations(client, ctx, owner, repoName)
	if err != nil {
		return nil, nil, nil, err
	}
	teamCollaborators, err := listTeams(client, ctx, owner, repoName)
	if err != nil {
		return nil, nil, nil, err
	}
	return userCollaborators, invitations, teamCollaborators, err
}

func matchUserCollaboratorsAndInvites(
	repoName string, users []interface{}, userCollaborators []userCollaborator, invitations []invitedCollaborator,
	meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	ctx := context.Background()

	for _, c := range userCollaborators {
		var permission string
		for _, u := range users {
			userData := u.(map[string]interface{})
			if userData["username"] == c.username {
				permission = userData["permission"].(string)
				break
			}
		}
		if permission == "" { // user should NOT have permission
			log.Printf("[DEBUG] Removing user %s from repo: %s.", c.username, repoName)
			_, err := client.Repositories.RemoveCollaborator(ctx, owner, repoName, c.username)
			if err != nil {
				return err
			}
		} else if permission != c.permission { // permission should be updated
			log.Printf("[DEBUG] Updating user %s permission from %s to %s for repo: %s.", c.username, c.permission, permission, repoName)
			_, _, err := client.Repositories.AddCollaborator(
				ctx, owner, repoName, c.username, &github.RepositoryAddCollaboratorOptions{
					Permission: permission,
				},
			)
			if err != nil {
				return err
			}
		}
	}

	for _, i := range invitations {
		var permission string
		for _, u := range users {
			userData := u.(map[string]interface{})
			if userData["username"] == i.username {
				permission = userData["permission"].(string)
				break
			}
		}
		if permission == "" { // user should NOT have permission
			log.Printf("[DEBUG] Deleting invite for user %s from repo: %s.", i.username, repoName)
			_, err := client.Repositories.DeleteInvitation(ctx, owner, repoName, i.invitationID)
			if err != nil {
				return err
			}
		} else if permission != i.permission { // permission should be updated
			log.Printf("[DEBUG] Updating invite for user %s permission from %s to %s for repo: %s.", i.username, i.permission, permission, repoName)
			_, _, err := client.Repositories.UpdateInvitation(ctx, owner, repoName, i.invitationID, permission)
			if err != nil {
				return err
			}
		}
	}

	for _, u := range users {
		userData := u.(map[string]interface{})
		username := userData["username"].(string)
		permission := userData["permission"].(string)
		var found bool
		for _, c := range userCollaborators {
			if username == c.username {
				found = true
				break
			}
		}
		if found {
			continue
		}
		for _, i := range invitations {
			if username == i.username {
				found = true
				break
			}
		}
		if found {
			continue
		}
		// user needs to be added
		log.Printf("[DEBUG] Inviting user %s with permission %s for repo: %s.", username, permission, repoName)
		_, _, err := client.Repositories.AddCollaborator(
			ctx, owner, repoName, username, &github.RepositoryAddCollaboratorOptions{
				Permission: permission,
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func matchTeamCollaborators(
	repoName string, teams []interface{}, teamCollaborators []teamCollaborator, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	orgId := meta.(*Owner).id
	ctx := context.Background()

	for _, tc := range teamCollaborators {
		var permission string
		for _, t := range teams {
			teamData := t.(map[string]interface{})
			teamIDString := teamData["team_id"].(string)
			teamId, err := strconv.ParseInt(teamIDString, 10, 64)
			if err != nil {
				return unconvertibleIdErr(teamIDString, err)
			}
			if teamId == tc.teamID {
				permission = teamData["permission"].(string)
				break
			}
		}
		if permission == "" { // user should NOT have permission
			log.Printf("[DEBUG] Removing team %d from repo: %s.", tc.teamID, repoName)
			_, err := client.Teams.RemoveTeamRepoByID(ctx, orgId, tc.teamID, owner, repoName)
			if err != nil {
				return err
			}
		} else if permission != tc.permission { // permission should be updated
			log.Printf("[DEBUG] Updating team %d permission from %s to %s for repo: %s.", tc.teamID, tc.permission, permission, repoName)
			_, err := client.Teams.AddTeamRepoByID(
				ctx, orgId, tc.teamID, owner, repoName, &github.TeamAddTeamRepoOptions{
					Permission: permission,
				},
			)
			if err != nil {
				return err
			}
		}
	}

	for _, t := range teams {
		teamData := t.(map[string]interface{})
		teamIDString := teamData["team_id"].(string)
		teamID, err := strconv.ParseInt(teamIDString, 10, 64)
		if err != nil {
			return unconvertibleIdErr(teamIDString, err)
		}
		permission := teamData["permission"].(string)
		var found bool
		for _, c := range teamCollaborators {
			if teamID == c.teamID {
				found = true
				break
			}
		}
		if found {
			continue
		}
		// team needs to be added
		log.Printf("[DEBUG] Adding team %d with permission %s for repo: %s.", teamID, permission, repoName)
		_, err = client.Teams.AddTeamRepoByID(
			ctx, orgId, teamID, owner, repoName, &github.TeamAddTeamRepoOptions{
				Permission: permission,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceGithubRepositoryCollaboratorsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	users := d.Get("user").(*schema.Set).List()
	teams := d.Get("team").(*schema.Set).List()
	repoName := d.Get("repository").(string)
	ctx := context.Background()

	userCollaborators, invitations, teamCollaborators, err := listAllCollaborators(client, ctx, owner, repoName)
	if err != nil {
		return handleAPIError(err, d, "repository collaborators (%s/%s)", owner, repoName)
	}

	err = matchUserCollaboratorsAndInvites(repoName, users, userCollaborators, invitations, meta)
	if err != nil {
		return err
	}

	err = matchTeamCollaborators(repoName, teams, teamCollaborators, meta)
	if err != nil {
		return err
	}

	d.SetId(repoName)

	return resourceGithubRepositoryCollaboratorsRead(d, meta)
}

func resourceGithubRepositoryCollaboratorsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Id()
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	userCollaborators, invitedCollaborators, teamCollaborators, err := listAllCollaborators(client, ctx, owner, repoName)
	if err != nil {
		return handleAPIError(err, d, "repository collaborators (%s/%s)", owner, repoName)
	}

	invitationIds := make(map[string]string, len(invitedCollaborators))
	for _, i := range invitedCollaborators {
		invitationIds[i.username] = strconv.FormatInt(i.invitationID, 10)
	}

	err = d.Set("repository", repoName)
	if err != nil {
		return err
	}
	err = d.Set("user", flattenUserCollaborators(userCollaborators, invitedCollaborators))
	if err != nil {
		return err
	}
	err = d.Set("team", flattenTeamCollaborators(teamCollaborators))
	if err != nil {
		return err
	}
	err = d.Set("invitation_ids", invitationIds)
	if err != nil {
		return err
	}

	return nil
}

func resourceGithubRepositoryCollaboratorsUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceGithubRepositoryCollaboratorsCreate(d, meta)
}

func resourceGithubRepositoryCollaboratorsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	ctx := context.Background()

	userCollaborators, invitations, teamCollaborators, err := listAllCollaborators(client, ctx, owner, repoName)
	if err != nil {
		return handleAPIError(err, d, "repository collaborators (%s/%s)", owner, repoName)
	}

	log.Printf("[DEBUG] Deleting all users, invites and collaborators for repo: %s.", repoName)

	// delete all users
	err = matchUserCollaboratorsAndInvites(repoName, nil, userCollaborators, invitations, meta)
	if err != nil {
		return err
	}

	// delete all teams
	err = matchTeamCollaborators(repoName, nil, teamCollaborators, meta)
	return err
}
