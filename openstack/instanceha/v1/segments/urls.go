package segments

import "github.com/lxdcc/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("segments")
}

func getURL(client *gophercloud.ServiceClient, segmentID string) string {
	return client.ServiceURL("segments", segmentID)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("segments")
}

func updateURL(client *gophercloud.ServiceClient, segmentID string) string {
	return client.ServiceURL("segments", segmentID)
}

func deleteURL(client *gophercloud.ServiceClient, segmentID string) string {
	return client.ServiceURL("segments", segmentID)
}
