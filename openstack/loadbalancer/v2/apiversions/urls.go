package apiversions

import (
	"strings"

	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/openstack/utils"
)

func listURL(c *gophercloud.ServiceClient) string {
	baseEndpoint, _ := utils.BaseEndpoint(c.Endpoint)
	endpoint := strings.TrimRight(baseEndpoint, "/") + "/"
	return endpoint
}
