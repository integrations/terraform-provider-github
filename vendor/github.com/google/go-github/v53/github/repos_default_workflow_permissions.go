// Copyright 2023 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
)

// DefaultWorkflowPermissionsRepository represents the default workflow permissions granted to the GITHUB_TOKEN when running workflows in a repository, and sets if GitHub Actions can submit approving pull request reviews.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/permissions?apiVersion=2022-11-28#set-default-workflow-permissions-for-a-repository
type DefaultWorkflowPermissionsRepository struct {
	DefaultWorkflowPermissions   *string `json:"default_workflow_permissions,omitempty"`
	CanApprovePullRequestReviews *bool   `json:"can_approve_pull_request_reviews,omitempty"`
}

func (a DefaultWorkflowPermissionsRepository) String() string {
	return Stringify(a)
}

// GetDefaultWorkflowPermissions gets the default workflow permissions granted to the GITHUB_TOKEN when running workflows in a repository.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/permissions#get-github-actions-permissions-for-a-repository
func (s *RepositoriesService) GetDefaultWorkflowPermissions(ctx context.Context, owner, repo string) (*DefaultWorkflowPermissionsRepository, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/actions/permissions/workflow", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	permissions := new(DefaultWorkflowPermissionsRepository)
	resp, err := s.client.Do(ctx, req, permissions)
	if err != nil {
		return nil, resp, err
	}

	return permissions, resp, nil
}

// EditDefaultWorkflowPermissions sets the default workflow permissions granted to the GITHUB_TOKEN when running workflows in a repository.
//
// GitHub API docs: https://docs.github.com/en/rest/actions/permissions?apiVersion=2022-11-28#set-default-workflow-permissions-for-a-repository
func (s *RepositoriesService) EditDefaultWorkflowPermissions(ctx context.Context, owner, repo string, defaultWorkflowPermissionsRepository DefaultWorkflowPermissionsRepository) (*DefaultWorkflowPermissionsRepository, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/actions/permissions/workflow", owner, repo)
	req, err := s.client.NewRequest("PUT", u, defaultWorkflowPermissionsRepository)
	if err != nil {
		return nil, nil, err
	}

	permissions := new(DefaultWorkflowPermissionsRepository)
	resp, err := s.client.Do(ctx, req, permissions)
	if err != nil {
		return nil, resp, err
	}

	return permissions, resp, nil
}
