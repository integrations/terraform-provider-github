package github

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
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
		},
	}
}

func resourceGithubUserInvitationAccepterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).v3client

	invitationIdString := d.Get("invitation_id").(string)
	invitationId, err := strconv.Atoi(invitationIdString)
	if err != nil {
		return fmt.Errorf("Failed to parse invitation ID: %s", err)
	}
	ctx := context.Background()

	log.Printf("[DEBUG] Accepting invitation: %d", invitationId)
	_, err = client.Users.AcceptInvitation(ctx, int64(invitationId))
	if err != nil {
		return err
	}

	d.SetId(invitationIdString)

	return nil
}
