package notifications

import "github.com/lxdcc/gophercloud"

const (
	apiVersion = "v1"
	apiName    = "notifications"
)

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}
