// Copyright 2019 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"errors"
	"fmt"
)

// createOrgRequest is a subset of Organization and is used internally
// by CreateOrg to pass only the known fields for the endpoint.
type createOrgRequest struct {
	Login *string `json:"login,omitempty"`
	Admin *string `json:"admin,omitempty"`
}

// CreateOrg creates a new organization in GitHub Enterprise.
//
// Note that only a subset of the org fields are used and org must
// not be nil.
//
// GitHub API docs: https://docs.github.com/enterprise-server@3.17/rest/enterprise-admin/orgs#create-an-organization
//
//meta:operation POST /admin/organizations
func (s *AdminService) CreateOrg(ctx context.Context, org *Organization, admin string) (*Organization, *Response, error) {
	u := "admin/organizations"

	orgReq := &createOrgRequest{
		Login: org.Login,
		Admin: &admin,
	}

	req, err := s.client.NewRequest("POST", u, orgReq)
	if err != nil {
		return nil, nil, err
	}

	o := new(Organization)
	resp, err := s.client.Do(ctx, req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}

// renameOrgRequest is a subset of Organization and is used internally
// by RenameOrg and RenameOrgByName to pass only the known fields for the endpoint.
type renameOrgRequest struct {
	Login *string `json:"login,omitempty"`
}

// RenameOrgResponse is the response given when renaming an Organization.
type RenameOrgResponse struct {
	Message *string `json:"message,omitempty"`
	URL     *string `json:"url,omitempty"`
}

// RenameOrg renames an organization in GitHub Enterprise.
//
// GitHub API docs: https://docs.github.com/enterprise-server@3.17/rest/enterprise-admin/orgs#update-an-organization-name
//
//meta:operation PATCH /admin/organizations/{org}
func (s *AdminService) RenameOrg(ctx context.Context, org *Organization, newName string) (*RenameOrgResponse, *Response, error) {
	if org == nil {
		return nil, nil, errors.New("organization must be provided")
	}
	if org.Login == nil {
		return nil, nil, errors.New("login must be provided")
	}

	return s.RenameOrgByName(ctx, *org.Login, newName)
}

// RenameOrgByName renames an organization in GitHub Enterprise using its current name.
//
// GitHub API docs: https://docs.github.com/enterprise-server@3.17/rest/enterprise-admin/orgs#update-an-organization-name
//
//meta:operation PATCH /admin/organizations/{org}
func (s *AdminService) RenameOrgByName(ctx context.Context, org, newName string) (*RenameOrgResponse, *Response, error) {
	u := fmt.Sprintf("admin/organizations/%v", org)

	orgReq := &renameOrgRequest{
		Login: &newName,
	}

	req, err := s.client.NewRequest("PATCH", u, orgReq)
	if err != nil {
		return nil, nil, err
	}

	o := new(RenameOrgResponse)
	resp, err := s.client.Do(ctx, req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}
