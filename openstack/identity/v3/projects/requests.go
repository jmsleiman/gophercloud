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
