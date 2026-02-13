package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func Test_getResourceAttr(t *testing.T) {
	t.Parallel()

	resourceName := "github_dummy.test"
	attrKey := "foo"
	want := "bar"

	for _, d := range []struct {
		name    string
		state   *terraform.State
		wantErr bool
	}{
		{
			name: "attribute_exists",
			state: &terraform.State{
				Modules: []*terraform.ModuleState{
					{
						Path: []string{"root"},
						Resources: map[string]*terraform.ResourceState{
							resourceName: {
								Primary: &terraform.InstanceState{
									Attributes: map[string]string{
										attrKey: want,
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "resource_missing",
			state: &terraform.State{
				Modules: []*terraform.ModuleState{
					{
						Path: []string{"root"},
						Resources: map[string]*terraform.ResourceState{
							"github_dummy.testx": {
								Primary: &terraform.InstanceState{
									Attributes: map[string]string{
										attrKey: want,
									},
								},
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "attribute_missing",
			state: &terraform.State{
				Modules: []*terraform.ModuleState{
					{
						Path: []string{"root"},
						Resources: map[string]*terraform.ResourceState{
							resourceName: {
								Primary: &terraform.InstanceState{
									Attributes: map[string]string{
										"attrx": want,
									},
								},
							},
						},
					},
				},
			},
			wantErr: true,
		},
	} {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()

			var got string
			f := getResourceAttr(resourceName, attrKey, &got)
			err := f(d.state)

			if (err != nil) != d.wantErr {
				t.Fatalf("unexpected error state")
			}

			if !d.wantErr {
				if got != want {
					t.Fatalf("unexpected value: got %s, want %s", got, want)
				}
			}
		})
	}
}
