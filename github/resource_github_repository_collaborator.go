package github

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubRepositoryCollaborator() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryCollaboratorCreate,
		Read:   resourceGithubRepositoryCollaboratorRead,
		Delete: resourceGithubRepositoryCollaboratorDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubRepositoryCollaboratorImport,
		},
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubRepositoryCollaboratorV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubRepositoryCollaboratorStateUpgradeV0,
				Version: 0,
			},
		},

		SchemaVersion: 1,
		// editing repository collaborators are not supported by github api so forcing new on any changes
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateNumericIDFunc,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permission": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "push",
				ValidateFunc: validateValueFunc([]string{"pull", "push", "admin"}),
			},
			"invitation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubRepositoryCollaboratorCreate(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	userIDString := d.Get("user_id").(string)
	repoName := d.Get("repository").(string)

	ctx := prepareResourceContext(d)

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(userIDString, err)
	}

	username, ok := meta.(*Organization).UserMap.GetUsername(userID, client)
	if !ok {
		return fmt.Errorf("Unable to get GitHub user %d", userID)
	}

	log.Printf("[DEBUG] Creating repository collaborator: %s (%s/%s)", userIDString, orgName, repoName)
	_, err = client.Repositories.AddCollaborator(ctx,
		orgName,
		repoName,
		username,
		&github.RepositoryAddCollaboratorOptions{
			Permission: d.Get("permission").(string),
		})

	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repoName, userIDString))

	return resourceGithubRepositoryCollaboratorRead(d, meta)
}

func resourceGithubRepositoryCollaboratorRead(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	repoName, userIDString, err := parseTwoPartID(d.Id(), "repository", "user_id")
	if err != nil {
		return err
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(userIDString, err)
	}

	username, ok := meta.(*Organization).UserMap.GetUsername(userID, client)
	if !ok {
		return fmt.Errorf("Unable to get GitHub user %d", userID)
	}

	ctx := prepareResourceContext(d)

	// First, check if the user has been invited but has not yet accepted
	invitation, err := findRepoInvitation(client, ctx, orgName, repoName, username)
	if err != nil {
		return err
	} else if invitation != nil {
		log.Printf("[DEBUG] Found invitation for %q", username)

		permissionName, err := getInvitationPermission(invitation)
		if err != nil {
			return err
		}

		d.Set("repository", repoName)
		d.Set("permission", permissionName)
		d.Set("invitation_id", fmt.Sprintf("%d", invitation.GetID()))
		return nil
	}

	// Next, check if the user has accepted the invite and is a full collaborator
	opt := &github.ListCollaboratorsOptions{ListOptions: github.ListOptions{
		PerPage: maxPerPage,
	}}

	for {
		collaborators, resp, err := client.Repositories.ListCollaborators(ctx,
			orgName, repoName, opt)
		if err != nil {
			return err
		}
		log.Printf("[DEBUG] Found %d collaborators, checking if any matches %q", len(collaborators), username)

		for _, c := range collaborators {
			if strings.EqualFold(*c.Login, username) {
				log.Printf("[DEBUG] Matching collaborator found for %q", username)
				permissionName, err := getRepoPermission(c.Permissions)
				if err != nil {
					return err
				}

				d.Set("repository", repoName)
				d.Set("permission", permissionName)
				return nil
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	// The user is neither invited nor a collaborator
	log.Printf("[WARN] Removing repository collaborator %s (%s/%s) from state because it no longer exists in GitHub",
		username, orgName, repoName)
	d.SetId("")

	return nil
}

func resourceGithubRepositoryCollaboratorDelete(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	repoName, userIDString, err := parseTwoPartID(d.Id(), "repository", "user_id")
	if err != nil {
		return err
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(userIDString, err)
	}

	username, ok := meta.(*Organization).UserMap.GetUsername(userID, client)
	if !ok {
		return fmt.Errorf("Unable to get GitHub user %d", userID)
	}

	ctx := prepareResourceContext(d)

	// Delete any pending invitations
	invitation, err := findRepoInvitation(client, ctx, orgName, repoName, username)
	if err != nil {
		return err
	} else if invitation != nil {
		_, err = client.Repositories.DeleteInvitation(ctx, orgName, repoName, *invitation.ID)
		return err
	}

	log.Printf("[DEBUG] Deleting repository collaborator: %s (%s/%s)",
		username, orgName, repoName)
	_, err = client.Repositories.RemoveCollaborator(ctx, orgName, repoName, username)
	return err
}

func findRepoInvitation(client *github.Client, ctx context.Context, owner, repo, collaborator string) (*github.RepositoryInvitation, error) {
	opt := &github.ListOptions{PerPage: maxPerPage}
	for {
		invitations, resp, err := client.Repositories.ListInvitations(ctx, owner, repo, opt)
		if err != nil {
			return nil, err
		}

		for _, i := range invitations {
			if strings.EqualFold(*i.Invitee.Login, collaborator) {
				return i, nil
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return nil, nil
}

func resourceGithubRepositoryCollaboratorImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*Organization).client
	ctx := prepareResourceContext(d)

	repository, userString, err := parseTwoPartID(d.Id(), "repository", "user_id_or_name")
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Reading user: %s", userString)
	// Attempt to parse the string as a numeric ID
	userID, err := strconv.ParseInt(userString, 10, 64)
	if err != nil {
		// It wasn't a numeric ID, try to use it as a username
		user, _, err := client.Users.Get(ctx, userString)
		if err != nil {
			return nil, err
		}
		userID = *user.ID
	}

	d.SetId(buildTwoPartID(repository, strconv.FormatInt(userID, 10)))

	return []*schema.ResourceData{d}, nil
}
