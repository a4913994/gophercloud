package segments

import "github.com/lxdcc/gophercloud"

const (
	apiVersion = "v1"
	apiName    = "segments"
)

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func getURL(client *gophercloud.ServiceClient, segmentID string) string {
	return client.ServiceURL(apiVersion, apiName, segmentID)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func updateURL(client *gophercloud.ServiceClient, segmentID string) string {
	return client.ServiceURL(apiVersion, apiName, segmentID)
}

func deleteURL(client *gophercloud.ServiceClient, segmentID string) string {
	return client.ServiceURL(apiVersion, apiName, segmentID)
}
