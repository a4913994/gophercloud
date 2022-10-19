package notifications

import "github.com/lxdcc/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("notifications")
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("notifications", id)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("notifications")
}
