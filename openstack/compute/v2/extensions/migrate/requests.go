package migrate

import (
	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/openstack/compute/v2/extensions"
)

// MigrateOptsBuilder allows extensions to add additional parameters to the
// Migrate request.
type MigrateOptsBuilder interface {
	ToMigrateMap() (map[string]interface{}, error)
}

// MigrateOpts specifies parameters of live migrate action.
type MigrateOpts struct {
	// The host to which to migrate the server.
	// If this parameter is None, the scheduler chooses a host.
	// Migrate Server (migrate Action) (v2.56)
	Host *string `json:"host"`
}

// ToMigrateMap constructs a request body from MigrateOpts.
func (opts LiveMigrateOpts) ToMigrateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "migrate")
}

// Migrate will initiate a migration of the instance to another host.
func Migrate(client *gophercloud.ServiceClient, id string, opts MigrateOptsBuilder) (r MigrateResult) {
	b, err := opts.ToMigrateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(extensions.ActionURL(client, id), b, nil, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// LiveMigrateOptsBuilder allows extensions to add additional parameters to the
// LiveMigrate request.
type LiveMigrateOptsBuilder interface {
	ToLiveMigrateMap() (map[string]interface{}, error)
}

// LiveMigrateOpts specifies parameters of live migrate action.
type LiveMigrateOpts struct {
	// The host to which to migrate the server.
	// If this parameter is None, the scheduler chooses a host.
	Host *string `json:"host"`

	// Set to True to migrate local disks by using block migration.
	// If the source or destination host uses shared storage and you set
	// this value to True, the live migration fails.
	BlockMigration *bool `json:"block_migration,omitempty"`

	// Set to True to enable over commit when the destination host is checked
	// for available disk space. Set to False to disable over commit. This setting
	// affects only the libvirt virt driver.
	DiskOverCommit *bool `json:"disk_over_commit,omitempty"`
}

// ToLiveMigrateMap constructs a request body from LiveMigrateOpts.
func (opts LiveMigrateOpts) ToLiveMigrateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "os-migrateLive")
}

// LiveMigrate will initiate a live-migration (without rebooting) of the instance to another host.
func LiveMigrate(client *gophercloud.ServiceClient, id string, opts LiveMigrateOptsBuilder) (r MigrateResult) {
	b, err := opts.ToLiveMigrateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(extensions.ActionURL(client, id), b, nil, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
