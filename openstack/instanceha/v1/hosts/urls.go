package hosts

import "github.com/lxdcc/gophercloud"

const (
	apiVersion = "v1"
	apiPrefix  = "segments"
	apiName    = "hosts"
)

func listURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiPrefix, id, apiName)
}

func getURL(client *gophercloud.ServiceClient, segmentID, HostID string) string {
	return client.ServiceURL(apiVersion, apiPrefix, segmentID, apiName, HostID)
}

func createURL(client *gophercloud.ServiceClient, segmentID string) string {
	return client.ServiceURL(apiVersion, apiPrefix, segmentID, apiName)
}

func updateURL(client *gophercloud.ServiceClient, segmentID, HostID string) string {
	return client.ServiceURL(apiVersion, apiPrefix, segmentID, apiName, HostID)
}

func deleteURL(client *gophercloud.ServiceClient, segmentID, HostID string) string {
	return client.ServiceURL(apiVersion, apiPrefix, segmentID, apiName, HostID)
}
