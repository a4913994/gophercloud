package testing

import (
	"testing"

	"github.com/lxdcc/gophercloud/openstack/baremetalintrospection/noauth"
	th "github.com/lxdcc/gophercloud/testhelper"
)

func TestNoAuth(t *testing.T) {
	noauthClient, err := noauth.NewBareMetalIntrospectionNoAuth(noauth.EndpointOpts{
		IronicInspectorEndpoint: "http://ironic:5050/v1",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "", noauthClient.TokenID)
}
