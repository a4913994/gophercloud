package bootfromvolume

import "github.com/lxdcc/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("servers")
}
