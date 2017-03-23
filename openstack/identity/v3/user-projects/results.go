package projects

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"

	"github.com/mitchellh/mapstructure"
)

// Project holds the project response.
type Project struct {
	Description string `json:"description"`
	DomainID    string `json:"domain_id"`
	Enabled     bool   `json:"enabled"`
	ID          string `json:"id"`
	Links       struct {
		Self string `json:"self"`
	} `json:"links"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract pops out the list of projects from a request
// An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() ([]Project, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		Projects []Project `json:"projects"`
	}

	err := mapstructure.Decode(r.Body, &res)

	return res.Projects, err
}

// ProjectPage is a single page of Project results.
type ProjectPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the page contains no results.
func (p ProjectPage) IsEmpty() (bool, error) {
	projects, err := ExtractProjects(p)
	if err != nil {
		return true, err
	}
	return len(projects) == 0, nil
}

// ExtractProjects extracts a slice of Projects from a Collection acquired from List.
func ExtractProjects(page pagination.Page) ([]Project, error) {
	var response struct {
		Projects []Project `mapstructure:"projects"`
	}

	err := mapstructure.Decode(page.(ProjectPage).Body, &response)
	return response.Projects, err
}
