package projects

import (
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

type ListOpts struct {
	DomainID string `json:"domain_id"` // Filters the response by a domain ID.
	Enabled  bool   `json:"enabled"`   // If set to true, then only enabled projects will be returned. Any value other than 0 (including no value) will be interpreted as true.
	IsDomain bool   `json:"is_domain"` // If this is specified as true, then only projects acting as a domain are included. Otherwise, only projects that are not acting as a domain are included.
	Name     string `json:"name"`      //	Filters the response by a project name.
	ParentID string `json:"parent_id"` // Filters the response by a parent ID.
}

// ListAllProjects lists all projects using the provieded client.
func ListAllProjects(client *gophercloud.ServiceClient, opts ListOpts) ListReturns {
	var result ListReturns
	_, result.Err = client.Post(listURL(client), opts, &result.Body, nil)
	return result
}
