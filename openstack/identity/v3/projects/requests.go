package projects

import (
	"errors"
	"github.com/rackspace/gophercloud"
)

type response struct {
	Project Project `json:"project"`
}

// CreateOpts allows you to create a project
type CreateOpts struct {
	IsDomain    bool   `json:"is_domain,omitempty"`
	Description string `json:"description,omitempty"`
	DomainID    string `json:"domain_id,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Name        string `json:"name"`
	ParentID    string `json:"parent_id,omitempty"`
}

// Create adds a new project using the provieded client.
func Create(client *gophercloud.ServiceClient, opts CreateOpts) CreateResult {
	type request struct {
		Project CreateOpts `json:"project"`
	}

	req := request{Project: opts}

	var result CreateResult
	_, result.Err = client.Post(listURL(client), req, &result.Body, nil)
	return result
}

// PairOpts allows you to pair roles, groups, on a project
type PairOpts struct {
	ID      string `json:"id"`
	GroupID string `json:"group_id"`
	RoleID  string `json:"role_id"`
}

// Pair creates a relationship between a role, group, and project
func Pair(client *gophercloud.ServiceClient, opts PairOpts) error {
	if opts.ID == "" || opts.GroupID == "" || opts.RoleID == "" {
		return errors.New("Project, Role, and Group ids are required.")
	}

	reqOpts := &gophercloud.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Content-Type": ""},
	}

	var result PairResult
	_, result.Err = client.Put(pairProjectGroupAndRoleURL(client, opts.ID, opts.GroupID, opts.RoleID), nil, nil, reqOpts)
	return result.Err
}
