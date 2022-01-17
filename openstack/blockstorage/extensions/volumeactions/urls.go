package volumeactions

import "github.com/lxdcc/gophercloud"

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
