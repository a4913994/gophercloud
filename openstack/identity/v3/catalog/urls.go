package catalog

import "github.com/lxdcc/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("auth", "catalog")
}
