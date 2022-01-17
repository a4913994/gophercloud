//go:build acceptance || blockstorage
// +build acceptance blockstorage

package extensions

import (
	"testing"

	"github.com/lxdcc/gophercloud/acceptance/clients"
	blockstorage "github.com/lxdcc/gophercloud/acceptance/openstack/blockstorage/v3"
	"github.com/lxdcc/gophercloud/openstack/blockstorage/extensions/volumetenants"
	"github.com/lxdcc/gophercloud/openstack/blockstorage/v3/volumes"
	th "github.com/lxdcc/gophercloud/testhelper"
)

func TestVolumeTenants(t *testing.T) {
	type volumeWithTenant struct {
		volumes.Volume
		volumetenants.VolumeTenantExt
	}

	var allVolumes []volumeWithTenant

	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	listOpts := volumes.ListOpts{
		Name: "I SHOULD NOT EXIST",
	}
	allPages, err := volumes.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	err = volumes.ExtractVolumesInto(allPages, &allVolumes)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(allVolumes))

	volume1, err := blockstorage.CreateVolume(t, client)
	th.AssertNoErr(t, err)
	defer blockstorage.DeleteVolume(t, client, volume1)

	allPages, err = volumes.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	err = volumes.ExtractVolumesInto(allPages, &allVolumes)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(allVolumes) > 0)
}
