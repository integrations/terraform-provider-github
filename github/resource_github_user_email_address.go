package github

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubUserEmailAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubUserEmailCreate,
		Read:   resourceGithubUserEmailRead,
		Delete: resourceGithubUserEmailDelete,

		Schema: map[string]*schema.Schema{
			"email_address": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The email address to add to the user account.",
			},
		},
	}
}

func resourceGithubUserEmailCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	emailAddress := d.Get("email_address").(string)
	ctx := context.Background()

	emails, _, err := client.Users.AddEmails(ctx, []string{emailAddress})
	if err != nil {
		return err
	}

	for _, email := range emails {
		if *email.Email == emailAddress {
			d.SetId(*email.Email)
			return nil
		}
	}

	return resourceGithubUserEmailRead(d, meta)
}

func resourceGithubUserEmailRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	emailAddress := d.Id()
	ctx := context.Background()

	emails, _, err := client.Users.ListEmails(ctx, nil)
	if err != nil {
		return err
	}

	for _, email := range emails {
		if *email.Email == emailAddress {
			d.Set("email_address", emailAddress)
			return nil
		}
	}

	log.Printf("[INFO] Email address %s no longer exists in GitHub", emailAddress)
	d.SetId("")
	return nil
}

func resourceGithubUserEmailDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	emailAddress := d.Id()
	ctx := context.Background()

	_, err := client.Users.DeleteEmails(ctx, []string{emailAddress})

	return err
}
