package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v54/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubRepositoryCollaborator() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryCollaboratorCreate,
		Read:   resourceGithubRepositoryCollaboratorRead,
		Update: resourceGithubRepositoryCollaboratorUpdate,
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
				Description:      "The user to add to the repository as a collaborator.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub repository",
			},
			"permission": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "push",
				Description: "The permission of the outside collaborator for the repository. Must be one of 'pull', 'push', 'maintain', 'triage' or 'admin' or the name of an existing custom repository role within the organization for organization-owned repositories. Must be 'push' for personal repositories. Defaults to 'push'.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("permission_diff_suppression").(bool) {
						if new == "triage" || new == "maintain" {
							return true
						}
					}
					return false
				},
			},
			"permission_diff_suppression": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Suppress plan diffs for triage and maintain. Defaults to 'false'.",
			},
			"invitation_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the invitation to be used in 'github_user_invitation_accepter'",
			},
		},
	}
}

func resourceGithubRepositoryCollaboratorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	username := d.Get("username").(string)
	repoName := d.Get("repository").(string)
	ctx := context.Background()

	_, _, err := client.Repositories.AddCollaborator(ctx,
		owner,
		repoName,
		username,
		&github.RepositoryAddCollaboratorOptions{
			Permission: d.Get("permission").(string),
		})

	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repoName, username))

	return resourceGithubRepositoryCollaboratorRead(d, meta)
}

func resourceGithubRepositoryCollaboratorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName, username, err := parseTwoPartID(d.Id(), "repository", "username")
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	// First, check if the user has been invited but has not yet accepted
	invitation, err := findRepoInvitation(client, ctx, owner, repoName, username)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				// this short circuits the rest of the code because if the
				// repo is 404, no reason to try to list existing collaborators
				log.Printf("[INFO] Removing repository collaborator %s/%s %s from state because it no longer exists in GitHub",
					owner, repoName, username)
				d.SetId("")
				return nil
			}
		}
		return err
	}
	if invitation != nil {
		username = invitation.GetInvitee().GetLogin()

		permissionName := getPermission(invitation.GetPermissions())

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
			owner, repoName, opt)
		if err != nil {
			return err
		}

		for _, c := range collaborators {
			if strings.EqualFold(c.GetLogin(), username) {
				d.Set("repository", repoName)
				d.Set("username", c.GetLogin())
				d.Set("permission", getPermission(c.GetRoleName()))
				return nil
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	// The user is neither invited nor a collaborator
	log.Printf("[INFO] Removing repository collaborator %s (%s/%s) from state because it no longer exists in GitHub",
		username, owner, repoName)
	d.SetId("")

	return nil
}

func resourceGithubRepositoryCollaboratorUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceGithubRepositoryCollaboratorRead(d, meta)
}

func resourceGithubRepositoryCollaboratorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	username := d.Get("username").(string)
	repoName := d.Get("repository").(string)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	// Delete any pending invitations
	invitation, err := findRepoInvitation(client, ctx, owner, repoName, username)
	if err != nil {
		return err
	} else if invitation != nil {
		_, err = client.Repositories.DeleteInvitation(ctx, owner, repoName, invitation.GetID())
		return err
	}

	_, err = client.Repositories.RemoveCollaborator(ctx, owner, repoName, username)
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
			if strings.EqualFold(i.GetInvitee().GetLogin(), collaborator) {
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
