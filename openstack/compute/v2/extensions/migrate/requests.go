package migrate

import (
	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/openstack/compute/v2/extensions"
	pagination "github.com/lxdcc/gophercloud/pagination"
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
func (opts MigrateOpts) ToMigrateMap() (map[string]interface{}, error) {
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

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToContainerListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Limit is the amount of containers to retrieve.
	// New in version 2.59
	Limit int `q:"limit"`

	// Marker The ID of the last-seen item. Use the limit parameter to make an initial limited request and use the
	// ID of the last-seen item from the response as the marker parameter value in a subsequent limited request.
	// New in version 2.59
	Marker string `q:"marker"`

	// Hidden The ‘hidden’ setting of migration to filter. The ‘hidden’ flag is set if the value is 1.
	// The ‘hidden’ flag is not set if the value is 0. But the ‘hidden’ setting of migration is always 0, so this parameter is useless to filter migrations.
	Hidden int `q:"hidden"`

	// Host The source/destination compute node of migration to filter.
	Host string `q:"host"`

	// MigrationType The type of migration to filter. Valid values are:  evacuation,live-migration, migration, resize
	MigrationType string `q:"migration_type"`

	// Status The status of migration to filter.
	Status string `q:"status"`

	// ChangesSince filters the response by a date and time stamp when the migration last changed. New in version 2.59
	ChangesSince string `q:"changes-since"`

	// ChangesBefore Filters the response by a date and time stamp when the migration last changed. New in version 2.66
	ChangesBefore string `q:"changes-before "`

	// UserID Filter the migrations by the given user ID. New in version 2.80
	UserID string `q:"user_id"`

	// ProjectID Filter the migrations by the given project ID. New in version 2.80
	ProjectID string `q:"project_id"`
}

// ToContainerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToContainerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager that allows you to iterate over a collection of Migrations.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := extensions.ListMigrations(client)
	if opts != nil {
		query, err := opts.ToContainerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return MigrationsPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Delete deletes a migration action.
func Delete(client *gophercloud.ServiceClient, serverID, migrationID string) (r DeleteResult) {
	resp, err := client.Delete(extensions.DeleteMigration(client, serverID, migrationID), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
