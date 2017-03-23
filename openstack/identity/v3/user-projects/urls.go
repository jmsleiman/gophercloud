package projects

import "github.com/rackspace/gophercloud"

func listURL(client *gophercloud.ServiceClient, uid string) string {
	return client.ServiceURL("users/" + uid + "/projects")
}
