package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func sharedConfigForRegion(region string) (any, error) {
	if os.Getenv("GITHUB_TOKEN") == "" {
		return nil, fmt.Errorf("empty GITHUB_TOKEN")
	}

	if os.Getenv("GITHUB_OWNER") == "" {
		return nil, fmt.Errorf("empty GITHUB_OWNER")
	}

	config := Config{
		Token:   os.Getenv("GITHUB_TOKEN"),
		Owner:   os.Getenv("GITHUB_OWNER"),
		BaseURL: "",
	}

	meta, err := config.Meta()
	if err != nil {
		return nil, fmt.Errorf("error getting GitHub meta parameter")
	}

	return meta, nil
}
