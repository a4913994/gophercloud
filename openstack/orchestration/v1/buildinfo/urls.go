package buildinfo

import "github.com/lxdcc/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
