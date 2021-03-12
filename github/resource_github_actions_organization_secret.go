package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v32/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubActionsOrganizationSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsOrganizationSecretCreateOrUpdate,
		Read:   resourceGithubActionsOrganizationSecretRead,
		Update: resourceGithubActionsOrganizationSecretCreateOrUpdate,
		Delete: resourceGithubActionsOrganizationSecretDelete,

		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateSecretNameFunc,
			},
			"plaintext_value": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"visibility": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateValueFunc([]string{"all", "private", "selected"}),
				ForceNew:     true,
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:      schema.HashInt,
				Optional: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubActionsOrganizationSecretCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	secretName := d.Get("secret_name").(string)
	plaintextValue := d.Get("plaintext_value").(string)

	visibility := d.Get("visibility").(string)
	selectedRepositories, hasSelectedRepositories := d.GetOk("selected_repository_ids")

	if visibility == "selected" && !hasSelectedRepositories {
		return fmt.Errorf("Cannot use visbility set to selected without selected_repository_ids")
	} else if visibility != "selected" && hasSelectedRepositories {
		return fmt.Errorf("Cannot use selected_repository_ids without visibility being set to selected")
	}

	selectedRepositoryIDs := []int64{}

	if hasSelectedRepositories {
		ids := selectedRepositories.(*schema.Set).List()

		for _, id := range ids {
			selectedRepositoryIDs = append(selectedRepositoryIDs, int64(id.(int)))
		}
	}

	keyId, publicKey, err := getOrganizationPublicKeyDetails(owner, meta)
	if err != nil {
		return err
	}

	encryptedText, err := encryptPlaintext(plaintextValue, publicKey)
	if err != nil {
		return err
	}

	// Create an EncryptedSecret and encrypt the plaintext value into it
	eSecret := &github.EncryptedSecret{
		Name:                  secretName,
		KeyID:                 keyId,
		Visibility:            visibility,
		SelectedRepositoryIDs: selectedRepositoryIDs,
		EncryptedValue:        base64.StdEncoding.EncodeToString(encryptedText),
	}

	_, err = client.Actions.CreateOrUpdateOrgSecret(ctx, owner, eSecret)
	if err != nil {
		return err
	}

	d.SetId(secretName)
	return resourceGithubActionsOrganizationSecretRead(d, meta)
}

func resourceGithubActionsOrganizationSecretRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	secret, _, err := client.Actions.GetOrgSecret(ctx, owner, d.Id())
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing actions secret %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("plaintext_value", d.Get("plaintext_value"))
	d.Set("updated_at", secret.UpdatedAt.String())
	d.Set("created_at", secret.CreatedAt.String())

	return nil
}

func resourceGithubActionsOrganizationSecretDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting secret: %s", d.Id())
	_, err := client.Actions.DeleteOrgSecret(ctx, orgName, d.Id())
	return err
}

func getOrganizationPublicKeyDetails(owner string, meta interface{}) (keyId, pkValue string, err error) {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	publicKey, _, err := client.Actions.GetOrgPublicKey(ctx, owner)
	if err != nil {
		return keyId, pkValue, err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), err
}
