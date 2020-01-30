package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
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
			State: schema.ImportStatePassthrough,
		},

		// editing repository collaborators are not supported by github api so forcing new on any changes
		Schema: map[string]*schema.Schema{
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: caseInsensitive(),
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
				ValidateFunc: validateValueFunc([]string{"pull", "triage", "push", "maintain", "admin"}),
			},
			"invitation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubRepositoryCollaboratorCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client

	orgName := meta.(*Organization).name
	username := d.Get("username").(string)
	repoName := d.Get("repository").(string)
	ctx := context.Background()

	log.Printf("[DEBUG] Creating repository collaborator: %s (%s/%s)",
		username, orgName, repoName)
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

	d.SetId(buildTwoPartID(&repoName, &username))

	return resourceGithubRepositoryCollaboratorRead(d, meta)
}

func resourceGithubRepositoryCollaboratorRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client

	orgName := meta.(*Organization).name
	repoName, username, err := parseTwoPartID(d.Id(), "repository", "username")
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	// First, check if the user has been invited but has not yet accepted
	invitation, err := findRepoInvitation(client, ctx, orgName, repoName, username)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				// this short circuits the rest of the code because if the
				// repo is 404, no reason to try to list existing collaborators
				log.Printf("[WARN] Removing repository collaborator %s/%s %s from state because it no longer exists in GitHub",
					orgName, repoName, username)
				d.SetId("")
				return nil
			}
		}
		return err
	}
	if invitation != nil {
		username = *invitation.Invitee.Login
		log.Printf("[DEBUG] Found invitation for %q", username)

		permissionName, err := getInvitationPermission(invitation)
		if err != nil {
			return err
		}

		d.Set("repository", repoName)
		d.Set("username", username)
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
				d.Set("username", c.Login)
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
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client

	orgName := meta.(*Organization).name
	username := d.Get("username").(string)
	repoName := d.Get("repository").(string)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

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
