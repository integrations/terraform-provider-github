package github

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/google/go-github/v54/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
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
				Description: "List of teams",
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
		},

		CustomizeDiff: customdiff.Sequence(
			// If there was a new user added to the list of collaborators,
			// it's possible a new invitation id will be created in GitHub.
			customdiff.ComputedIf("invitation_ids", func(d *schema.ResourceDiff, meta interface{}) bool {
				return d.HasChange("user")
			}),
		),
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
	if objs == nil && invites == nil {
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
	teamSlug   string
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
		"team_id":    obj.teamSlug,
	}

	return transformed
}

func flattenTeamCollaborators(objs []teamCollaborator) []interface{} {
	if objs == nil {
		return nil
	}

	sort.SliceStable(objs, func(i, j int) bool {
		return objs[i].teamSlug < objs[j].teamSlug
	})

	items := make([]interface{}, len(objs))
	for i, obj := range objs {
		items[i] = flattenTeamCollaborator(obj)
	}

	return items
}

func listUserCollaborators(client *github.Client, isOrg bool, ctx context.Context, owner, repoName string) ([]userCollaborator, error) {
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
				// owners are listed in the collaborators list even though they don't have direct permissions
				if !isOrg && c.GetLogin() == owner {
					continue
				}
				permissionName := getPermission(c.GetRoleName())

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
			permissionName := getPermission(i.GetPermissions())

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

func listTeams(client *github.Client, isOrg bool, ctx context.Context, owner, repoName string) ([]teamCollaborator, error) {
	var teamCollaborators []teamCollaborator

	if !isOrg {
		return teamCollaborators, nil
	}

	opt := &github.ListOptions{PerPage: maxPerPage}
	for {
		repoTeams, resp, err := client.Repositories.ListTeams(ctx, owner, repoName, opt)
		if err != nil {
			return nil, err
		}

		for _, t := range repoTeams {
			permissionName := getPermission(t.GetPermission())

			teamCollaborators = append(teamCollaborators, teamCollaborator{permissionName, t.GetSlug()})
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return teamCollaborators, nil
}

func listAllCollaborators(client *github.Client, isOrg bool, ctx context.Context, owner, repoName string) ([]userCollaborator, []invitedCollaborator, []teamCollaborator, error) {
	userCollaborators, err := listUserCollaborators(client, isOrg, ctx, owner, repoName)
	if err != nil {
		return nil, nil, nil, err
	}
	invitations, err := listInvitations(client, ctx, owner, repoName)
	if err != nil {
		return nil, nil, nil, err
	}
	teamCollaborators, err := listTeams(client, isOrg, ctx, owner, repoName)
	if err != nil {
		return nil, nil, nil, err
	}
	return userCollaborators, invitations, teamCollaborators, err
}

func matchUserCollaboratorsAndInvites(
	repoName string, want []interface{}, hasUsers []userCollaborator, hasInvites []invitedCollaborator,
	meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	ctx := context.Background()

	for _, has := range hasUsers {
		var wantPermission string
		for _, w := range want {
			userData := w.(map[string]interface{})
			if userData["username"] == has.username {
				wantPermission = userData["permission"].(string)
				break
			}
		}
		if wantPermission == "" { // user should NOT have permission
			log.Printf("[DEBUG] Removing user %s from repo: %s.", has.username, repoName)
			_, err := client.Repositories.RemoveCollaborator(ctx, owner, repoName, has.username)
			if err != nil {
				return err
			}
		} else if wantPermission != has.permission { // permission should be updated
			log.Printf("[DEBUG] Updating user %s permission from %s to %s for repo: %s.", has.username, has.permission, wantPermission, repoName)
			_, _, err := client.Repositories.AddCollaborator(
				ctx, owner, repoName, has.username, &github.RepositoryAddCollaboratorOptions{
					Permission: wantPermission,
				},
			)
			if err != nil {
				return err
			}
		}
	}

	for _, has := range hasInvites {
		var wantPermission string
		for _, u := range want {
			userData := u.(map[string]interface{})
			if userData["username"] == has.username {
				wantPermission = userData["permission"].(string)
				break
			}
		}
		if wantPermission == "" { // user should NOT have permission
			log.Printf("[DEBUG] Deleting invite for user %s from repo: %s.", has.username, repoName)
			_, err := client.Repositories.DeleteInvitation(ctx, owner, repoName, has.invitationID)
			if err != nil {
				return err
			}
		} else if wantPermission != has.permission { // permission should be updated
			log.Printf("[DEBUG] Updating invite for user %s permission from %s to %s for repo: %s.", has.username, has.permission, wantPermission, repoName)
			_, _, err := client.Repositories.UpdateInvitation(ctx, owner, repoName, has.invitationID, wantPermission)
			if err != nil {
				return err
			}
		}
	}

	for _, w := range want {
		userData := w.(map[string]interface{})
		username := userData["username"].(string)
		permission := userData["permission"].(string)
		var found bool
		for _, has := range hasUsers {
			if username == has.username {
				found = true
				break
			}
		}
		if found {
			continue
		}
		for _, has := range hasInvites {
			if username == has.username {
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
	repoName string, want []interface{}, has []teamCollaborator, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	for _, hasTeam := range has {
		var wantPerm string
		for _, w := range want {
			teamData := w.(map[string]interface{})
			teamIDString := teamData["team_id"].(string)
			teamSlug, err := getTeamSlug(teamIDString, meta)
			if err != nil {
				return err
			}
			if teamSlug == hasTeam.teamSlug {
				wantPerm = teamData["permission"].(string)
				break
			}
		}
		if wantPerm == "" { // user should NOT have permission
			log.Printf("[DEBUG] Removing team %s from repo: %s.", hasTeam.teamSlug, repoName)
			_, err := client.Teams.RemoveTeamRepoBySlug(ctx, owner, hasTeam.teamSlug, owner, repoName)
			if err != nil {
				return err
			}
		} else if wantPerm != hasTeam.permission { // permission should be updated
			log.Printf("[DEBUG] Updating team %s permission from %s to %s for repo: %s.", hasTeam.teamSlug, hasTeam.permission, wantPerm, repoName)
			_, err := client.Teams.AddTeamRepoBySlug(
				ctx, owner, hasTeam.teamSlug, owner, repoName, &github.TeamAddTeamRepoOptions{
					Permission: wantPerm,
				},
			)
			if err != nil {
				return err
			}
		}
	}

	for _, t := range want {
		teamData := t.(map[string]interface{})
		teamIDString := teamData["team_id"].(string)
		teamSlug, err := getTeamSlug(teamIDString, meta)
		if err != nil {
			return err
		}
		permission := teamData["permission"].(string)
		var found bool
		for _, c := range has {
			if teamSlug == c.teamSlug {
				found = true
				break
			}
		}
		if found {
			continue
		}
		// team needs to be added
		log.Printf("[DEBUG] Adding team %s with permission %s for repo: %s.", teamSlug, permission, repoName)
		_, err = client.Teams.AddTeamRepoBySlug(
			ctx, owner, teamSlug, owner, repoName, &github.TeamAddTeamRepoOptions{
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
	isOrg := meta.(*Owner).IsOrganization
	users := d.Get("user").(*schema.Set).List()
	teams := d.Get("team").(*schema.Set).List()
	repoName := d.Get("repository").(string)
	ctx := context.Background()

	usersMap := make(map[string]struct{})
	for _, user := range users {
		username := user.(map[string]interface{})["username"].(string)
		if _, found := usersMap[username]; found {
			return fmt.Errorf("duplicate set member found: %s", username)
		}
		usersMap[username] = struct{}{}
	}
	teamsMap := make(map[string]struct{})
	for _, team := range teams {
		teamID := team.(map[string]interface{})["team_id"].(string)
		if _, found := teamsMap[teamID]; found {
			return fmt.Errorf("duplicate set member: %s", teamID)
		}
		teamsMap[teamID] = struct{}{}
	}

	userCollaborators, invitations, teamCollaborators, err := listAllCollaborators(client, isOrg, ctx, owner, repoName)
	if err != nil {
		return deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "repository collaborators (%s/%s)", owner, repoName)
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
	isOrg := meta.(*Owner).IsOrganization
	repoName := d.Id()
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	userCollaborators, invitedCollaborators, teamCollaborators, err := listAllCollaborators(client, isOrg, ctx, owner, repoName)
	if err != nil {
		return deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "repository collaborators (%s/%s)", owner, repoName)
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
	isOrg := meta.(*Owner).IsOrganization
	repoName := d.Get("repository").(string)
	ctx := context.Background()

	userCollaborators, invitations, teamCollaborators, err := listAllCollaborators(client, isOrg, ctx, owner, repoName)
	if err != nil {
		return deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "repository collaborators (%s/%s)", owner, repoName)
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
