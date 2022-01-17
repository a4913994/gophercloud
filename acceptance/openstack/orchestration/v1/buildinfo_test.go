//go:build acceptance
// +build acceptance

package v1

import (
	"testing"

	"github.com/lxdcc/gophercloud/acceptance/clients"
	"github.com/lxdcc/gophercloud/openstack/orchestration/v1/buildinfo"
	th "github.com/lxdcc/gophercloud/testhelper"
)

func TestBuildInfo(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")

	client, err := clients.NewOrchestrationV1Client()
	th.AssertNoErr(t, err)

	bi, err := buildinfo.Get(client).Extract()
	th.AssertNoErr(t, err)
	t.Logf("retrieved build info: %+v\n", bi)
}
