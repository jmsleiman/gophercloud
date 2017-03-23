package projects

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

type response struct {
	Project Project `json:"project"`
}

// List enumerates the projects available to a specific user.
func List(client *gophercloud.ServiceClient, uid string) pagination.Pager {
	u := listURL(client, uid)

	createPage := func(r pagination.PageResult) pagination.Page {
		return ProjectPage{
			pagination.LinkedPageBase{
				PageResult: r,
			},
		}
	}

	return pagination.NewPager(client, u, createPage)
}
