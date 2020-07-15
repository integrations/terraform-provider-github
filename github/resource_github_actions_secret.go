package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"golang.org/x/crypto/nacl/box"
	"log"
	"net/http"
)

func resourceGithubActionsSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsSecretCreateOrUpdate,
		Read:   resourceGithubActionsSecretRead,
		Update: resourceGithubActionsSecretCreateOrUpdate,
		Delete: resourceGithubActionsSecretDelete,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"plaintext_value": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
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

func resourceGithubActionsSecretCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	secretName := d.Get("secret_name").(string)
	plaintextValue := d.Get("plaintext_value").(string)

	keyId, publicKey, err := getPublicKeyDetails(owner, repo, meta)
	if err != nil {
		return err
	}

	encryptedText, err := encryptPlaintext(plaintextValue, publicKey)
	if err != nil {
		return err
	}

	// Create an EncryptedSecret and encrypt the plaintext value into it
	eSecret := &github.EncryptedSecret{
		Name:           secretName,
		KeyID:          keyId,
		EncryptedValue: base64.StdEncoding.EncodeToString(encryptedText),
	}

	_, err = client.Actions.CreateOrUpdateSecret(ctx, owner, repo, eSecret)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repo, secretName))
	return resourceGithubActionsSecretRead(d, meta)
}

func resourceGithubActionsSecretRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repoName, secretName, err := parseTwoPartID(d.Id(), "repository", "secret_name")
	if err != nil {
		return err
	}

	secret, _, err := client.Actions.GetSecret(ctx, owner, repoName, secretName)
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
	d.Set("updated_at", secret.UpdatedAt.Format("default"))
	d.Set("created_at", secret.CreatedAt.Format("default"))

	return nil
}

func resourceGithubActionsSecretDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	repoName, secretName, err := parseTwoPartID(d.Id(), "repository", "secret_name")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting secret: %s", d.Id())
	_, err = client.Actions.DeleteSecret(ctx, orgName, repoName, secretName)

	return err
}

func getPublicKeyDetails(owner, repository string, meta interface{}) (keyId, pkValue string, err error) {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	publicKey, _, err := client.Actions.GetPublicKey(ctx, owner, repository)
	if err != nil {
		return keyId, pkValue, err
	}

	return publicKey.GetKeyID(), publicKey.GetKey(), err
}

func encryptPlaintext(plaintext, publicKeyB64 string) ([]byte, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyB64)
	if err != nil {
		return nil, err
	}

	var publicKeyBytes32 [32]byte
	copiedLen := copy(publicKeyBytes32[:], publicKeyBytes)
	if copiedLen == 0 {
		return nil, fmt.Errorf("could not convert publicKey to bytes")
	}

	plaintextBytes := []byte(plaintext)
	var encryptedBytes []byte

	cipherText, err := box.SealAnonymous(encryptedBytes, plaintextBytes, &publicKeyBytes32, nil)
	if err != nil {
		return nil, err
	}

	return cipherText, nil
}
