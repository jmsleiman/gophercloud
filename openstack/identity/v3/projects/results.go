package projects

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

// Project the object to hold a project.
type Project struct {
	ID          string `json:"id,omitempty"`
	IsDomain    bool   `json:"is_domain,omitempty"`
	Description string `json:"description,omitempty"`
	DomainID    string `json:"domain_id,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Name        string `json:"name"`
	ParentID    string `json:"parent_id,omitempty"`
	Links       Link   `json:"links,omitempty"`
}

// Link the object to hold a project link.
type Link struct {
	Self string `json:"self,omitempty"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a concrete Service.
// An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*Project, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		Project `json:"project"`
	}

	err := mapstructure.Decode(r.Body, &res)

	return &res.Project, err
}

// CreateResult the object to hold a project link.
type CreateResult struct {
	commonResult
}

type ListReturns struct {
	gophercloud.Result
	Links struct {
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
		Self     string      `json:"self"`
	} `json:"links"`
	Projects []struct {
		IsDomain    bool        `json:"is_domain"`
		Description interface{} `json:"description"`
		DomainID    string      `json:"domain_id"`
		Enabled     bool        `json:"enabled"`
		ID          string      `json:"id"`
		Links       struct {
			Self string `json:"self"`
		} `json:"links"`
		Name     string      `json:"name"`
		ParentID interface{} `json:"parent_id"`
	} `json:"projects"`
}
