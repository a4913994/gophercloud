package extensions

import "github.com/lxdcc/gophercloud"

func ActionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}

func ListMigrations(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("os-migrations")
}
