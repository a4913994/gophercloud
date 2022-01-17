package services

import "github.com/lxdcc/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-services")
}
