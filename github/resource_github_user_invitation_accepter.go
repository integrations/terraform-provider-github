package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubUserInvitationAccepter() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubUserInvitationAccepterCreate,
		Read:   schema.Noop, // Nothing to read as invitation is removed after it's accepted
		Delete: schema.Noop, // Nothing to remove as invitation is removed after it's accepted

		Schema: map[string]*schema.Schema{
			"invitation_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ID of the invitation to accept. Must be set when 'allow_empty_id' is 'false'.",
			},
			"allow_empty_id": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.",
			},
		},
	}
}

func resourceGithubUserInvitationAccepterCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client

	invitationIdString := d.Get("invitation_id").(string)
	allowEmptyId := d.Get("allow_empty_id").(bool)

	if invitationIdString == "" {
		if allowEmptyId {
			// We're setting a random UUID as resource ID since every resource needs an ID
			// and we can't destroy the resource while we create it.
			d.SetId(uuid.NewString())
			return nil
		} else {
			return fmt.Errorf("invitation_id is not set and allow_empty_id is false")
		}
	}

	invitationId, err := strconv.Atoi(invitationIdString)
	if err != nil {
		return fmt.Errorf("failed to parse invitation ID: %s", err)
	}
	ctx := context.Background()

	_, err = client.Users.AcceptInvitation(ctx, int64(invitationId))
	if err != nil {
		return err
	}

	d.SetId(invitationIdString)

	return nil
}
