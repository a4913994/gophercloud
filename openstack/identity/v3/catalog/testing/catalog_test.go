package testing

import (
	"testing"

	"github.com/lxdcc/gophercloud/openstack/identity/v3/catalog"
	"github.com/lxdcc/gophercloud/pagination"
	th "github.com/lxdcc/gophercloud/testhelper"
	"github.com/lxdcc/gophercloud/testhelper/client"
)

func TestListCatalog(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListCatalogSuccessfully(t)

	count := 0
	err := catalog.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := catalog.ExtractServiceCatalog(page)
		th.AssertNoErr(t, err)

		th.CheckDeepEquals(t, ExpectedCatalogSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, count, 1)
}
