package datastores

import "github.com/lxdcc/gophercloud"

func baseURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("datastores")
}

func resourceURL(c *gophercloud.ServiceClient, dsID string) string {
	return c.ServiceURL("datastores", dsID)
}

func versionsURL(c *gophercloud.ServiceClient, dsID string) string {
	return c.ServiceURL("datastores", dsID, "versions")
}

func versionURL(c *gophercloud.ServiceClient, dsID, versionID string) string {
	return c.ServiceURL("datastores", dsID, "versions", versionID)
}

func datastoreVersionsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("mgmt", "datastore-versions")
}

func deleteDatastoreVersionsURL(c *gophercloud.ServiceClient, versionID string) string {
	return c.ServiceURL("mgmt", "datastore-versions", versionID)
}
