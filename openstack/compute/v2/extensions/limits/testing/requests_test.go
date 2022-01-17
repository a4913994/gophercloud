package testing

import (
	"testing"

	"github.com/lxdcc/gophercloud/openstack/compute/v2/extensions/limits"
	th "github.com/lxdcc/gophercloud/testhelper"
	"github.com/lxdcc/gophercloud/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	getOpts := limits.GetOpts{
		TenantID: TenantID,
	}

	actual, err := limits.Get(client.ServiceClient(), getOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &LimitsResult, actual)
}
