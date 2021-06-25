package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v35/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubActionsOrganizationSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsOrganizationSecretCreateOrUpdate,
		Read:   resourceGithubActionsOrganizationSecretRead,
		Update: resourceGithubActionsOrganizationSecretCreateOrUpdate,
		Delete: resourceGithubActionsOrganizationSecretDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				d.Set("secret_name", d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateSecretNameFunc,
			},
			"encrypted_value": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"plaintext_value"},
				ValidateFunc:  validation.StringIsBase64,
			},
			"plaintext_value": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"encrypted_value"},
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
	var encryptedValue string

	visibility := d.Get("visibility").(string)
	selectedRepositories, hasSelectedRepositories := d.GetOk("selected_repository_ids")

	if visibility != "selected" && hasSelectedRepositories {
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

	if encryptedText, ok := d.GetOk("encrypted_value"); ok {
		encryptedValue = encryptedText.(string)
	} else {
		encryptedBytes, err := encryptPlaintext(plaintextValue, publicKey)
		if err != nil {
			return err
		}
		encryptedValue = base64.StdEncoding.EncodeToString(encryptedBytes)
	}

	// Create an EncryptedSecret and encrypt the plaintext value into it
	eSecret := &github.EncryptedSecret{
		Name:                  secretName,
		KeyID:                 keyId,
		Visibility:            visibility,
		SelectedRepositoryIDs: selectedRepositoryIDs,
		EncryptedValue:        encryptedValue,
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

	d.Set("encrypted_value", d.Get("encrypted_value"))
	d.Set("plaintext_value", d.Get("plaintext_value"))
	d.Set("created_at", secret.CreatedAt.String())
	d.Set("visibility", secret.Visibility)

	selectedRepositoryIDs := []int64{}

	if secret.Visibility == "selected" {
		selectedRepoList, _, err := client.Actions.ListSelectedReposForOrgSecret(ctx, owner, d.Id())
		if err != nil {
			return err
		}

		selectedRepositories := selectedRepoList.Repositories

		for _, repo := range selectedRepositories {
			selectedRepositoryIDs = append(selectedRepositoryIDs, repo.GetID())
		}
	}

	d.Set("selected_repository_ids", selectedRepositoryIDs)

	// This is a drift detection mechanism based on timestamps.
	//
	// If we do not currently store the "updated_at" field, it means we've only
	// just created the resource and the value is most likely what we want it to
	// be.
	//
	// If the resource is changed externally in the meantime then reading back
	// the last update timestamp will return a result different than the
	// timestamp we've persisted in the state. In that case, we can no longer
	// trust that the value (which we don't see) is equal to what we've declared
	// previously.
	//
	// The only solution to enforce consistency between is to mark the resource
	// as deleted (unset the ID) in order to fix potential drift by recreating
	// the resource.
	if updatedAt, ok := d.GetOk("updated_at"); ok && updatedAt != secret.UpdatedAt.String() {
		log.Printf("[WARN] The secret %s has been externally updated in GitHub", d.Id())
		d.SetId("")
	} else if !ok {
		d.Set("updated_at", secret.UpdatedAt.String())
	}

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
