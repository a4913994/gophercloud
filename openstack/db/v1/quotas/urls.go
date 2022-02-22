package quotas

import "github.com/lxdcc/gophercloud"

func quotaGetURL(c *gophercloud.ServiceClient, projectID string) string {
	return c.ServiceURL("mgmt", "quotas", projectID)
}

func quotaUpdateURL(c *gophercloud.ServiceClient, projectID string) string {
	return c.ServiceURL("mgmt", "quotas", projectID)
}
