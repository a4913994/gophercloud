package testing

import (
	"testing"

	"github.com/lxdcc/gophercloud/openstack/clustering/v1/policytypes"
	"github.com/lxdcc/gophercloud/pagination"
	th "github.com/lxdcc/gophercloud/testhelper"
	fake "github.com/lxdcc/gophercloud/testhelper/client"
)

func TestListPolicyTypes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandlePolicyTypeList(t)

	count := 0
	err := policytypes.List(fake.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := policytypes.ExtractPolicyTypes(page)
		if err != nil {
			t.Errorf("Failed to extract policy types: %v", err)
			return false, err
		}
		th.AssertDeepEquals(t, ExpectedPolicyTypes, actual)

		return true, nil
	})

	th.AssertNoErr(t, err)

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetPolicyType(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandlePolicyTypeGet(t)

	actual, err := policytypes.Get(fake.ServiceClient(), FakePolicyTypetoGet).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, ExpectedPolicyTypeDetail, actual)
}
