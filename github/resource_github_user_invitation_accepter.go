package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v49/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubUserInvitationAccepter() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubUserInvitationAccepterCreate,
		Read:   schema.Noop, // Nothing to read as invitation is removed after it's accepted
		Delete: schema.Noop, // Nothing to remove as invitation is removed after it's accepted

		Schema: map[string]*schema.Schema{
			"invitation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ignore_not_found": {
				Type:     schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceGithubUserInvitationAccepterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	ignore_not_found := d.Get("ignore_not_found").(bool)
	invitationIdString := d.Get("invitation_id").(string)
	invitationId, err := strconv.Atoi(invitationIdString)
	if err != nil {
		return fmt.Errorf("failed to parse invitation ID: %s", err)
	}
	ctx := context.Background()

	_, err = client.Users.AcceptInvitation(ctx, int64(invitationId))
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound && ignore_not_found {
				log.Printf("[DEBUG] Ignoring non-existing invitation with ID %d", invitationId)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(invitationIdString)

	return nil
}
