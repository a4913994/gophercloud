package shareinstances

import "github.com/lxdcc/gophercloud"

func listInstancesURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("share_instances")
}

func getInstancesURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("share_instances", id)
}
