package rules

import "github.com/lxdcc/gophercloud"

const rootPath = "security-group-rules"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, id)
}
