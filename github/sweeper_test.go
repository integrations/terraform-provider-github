package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func sharedConfigForRegion(region string) (interface{}, error) {
	if os.Getenv("GITHUB_TOKEN") == "" {
		return nil, fmt.Errorf("empty GITHUB_TOKEN")
	}

	if os.Getenv("GITHUB_ORGANIZATION") == "" {
		return nil, fmt.Errorf("empty GITHUB_ORGANIZATION")
	}

	config := Config{
		Token:        os.Getenv("GITHUB_TOKEN"),
		Organization: os.Getenv("GITHUB_ORGANIZATION"),
		BaseURL:      "",
	}

	client, err := config.Clients()
	if err != nil {
		return nil, fmt.Errorf("error getting Github client")
	}

	return client, nil
}
