package hosts

import "github.com/lxdcc/gophercloud"

func listURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("segments", id, "hosts")
}

func getURL(client *gophercloud.ServiceClient, segmentID, HostID string) string {
	return client.ServiceURL("segments", segmentID, "hosts", HostID)
}

func createURL(client *gophercloud.ServiceClient, segmentID string) string {
	return client.ServiceURL("segments", segmentID, "hosts")
}

func updateURL(client *gophercloud.ServiceClient, segmentID, HostID string) string {
	return client.ServiceURL("segments", segmentID, "hosts", HostID)
}

func deleteURL(client *gophercloud.ServiceClient, segmentID, HostID string) string {
	return client.ServiceURL("segments", segmentID, "hosts", HostID)
}
