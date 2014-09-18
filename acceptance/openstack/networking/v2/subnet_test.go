// +build acceptance networking

package v2

import (
	"testing"

	"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
	"github.com/rackspace/gophercloud/openstack/networking/v2/subnets"
	"github.com/rackspace/gophercloud/pagination"
	th "github.com/rackspace/gophercloud/testhelper"
)

func TestList(t *testing.T) {
	Setup(t)
	defer Teardown()

	pager := subnets.List(Client, subnets.ListOpts{Limit: 2})
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		t.Logf("--- Page ---")

		subnets, err := subnets.ExtractSubnets(page)
		th.AssertNoErr(t, err)

		for _, s := range subnets {
			t.Logf("Subnet: ID [%s] Name [%s] IP Version [%d] CIDR [%s] GatewayIP [%s]",
				s.ID, s.Name, s.IPVersion, s.CIDR, s.GatewayIP)
		}

		return true, nil
	})
	th.CheckNoErr(t, err)
}

func TestCRUD(t *testing.T) {
	Setup(t)
	defer Teardown()

	// Setup network
	t.Log("Setting up network")
	res, err := networks.Create(Client, networks.NetworkOpts{Name: "tmp_network", AdminStateUp: true})
	th.AssertNoErr(t, err)
	networkID := res.ID
	defer networks.Delete(Client, networkID)

	// Create subnet
	t.Log("Create subnet")
	enable := false
	opts := subnets.SubnetOpts{
		NetworkID:  networkID,
		CIDR:       "192.168.199.0/24",
		IPVersion:  subnets.IPv4,
		Name:       "my_subnet",
		EnableDHCP: &enable,
	}
	s, err := subnets.Create(Client, opts)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, s.NetworkID, networkID)
	th.AssertEquals(t, s.CIDR, "192.168.199.0/24")
	th.AssertEquals(t, s.IPVersion, 4)
	th.AssertEquals(t, s.Name, "my_subnet")
	th.AssertEquals(t, s.EnableDHCP, false)
	subnetID := s.ID

	// Get subnet
	t.Log("Getting subnet")
	s, err = subnets.Get(Client, subnetID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, s.ID, subnetID)

	// Update subnet
	t.Log("Update subnet")
	s, err = subnets.Update(Client, subnetID, subnets.SubnetOpts{Name: "new_subnet_name"})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, s.Name, "new_subnet_name")

	// Delete subnet
	t.Log("Delete subnet")
	err = subnets.Delete(Client, subnetID)
	th.AssertNoErr(t, err)
}

func TestBatchCreate(t *testing.T) {
	// todo
}
