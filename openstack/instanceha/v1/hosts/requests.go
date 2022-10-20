package hosts

import (
	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToContainerListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Limit is the amount of containers to retrieve.
	Limit int `q:"limit"`

	// Marker The ID of the last-seen item. Use the limit parameter to make an initial limited request and use the
	// ID of the last-seen item from the response as the marker parameter value in a subsequent limited request.
	Marker string `q:"marker"`

	// OnMaintenance Filter the host list result by on_maintenance.
	OnMaintenance bool `q:"on_maintenance"`

	// Reserved Filter the host list result by reserved flag.
	Reserved bool `q:"reserved"`

	// SortDir Sort direction. A valid value is asc (ascending) or desc (descending). Default is desc.
	// You can specify multiple pairs of sort key and sort direction query parameters. If you omit the sort direction in
	// a pair, the API uses the natural sorting direction of the direction of the segment sort_key attribute.
	SortDir string `q:"sort_dir"`

	// SortKey Sorts by a hosts attribute. Default attribute is created_at. You can specify multiple pairs of sort key
	// and sort direction query parameters. If you omit the sort direction in a pair, the API uses the natural sorting
	// direction of the segment sort_key attribute. The sort keys are limited to:
	// created_at, type, name, updated_at, uuid, reserved, on_maintenance
	SortKey string `q:"sort_key"`

	// Status Filter the host list result by status.
	Type bool `q:"type"`
}

// ToContainerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToContainerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List retrieves a list of hosts in segment.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder, id string) pagination.Pager {
	url := listURL(client, id)
	if opts != nil {
		query, err := opts.ToContainerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return HostsPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details of a host in segment.
func Get(client *gophercloud.ServiceClient, segmentID, hostID string) (r GetResult) {
	resp, err := client.Get(getURL(client, segmentID, hostID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToContainerCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to add a host to segment.
type CreateOpts struct {
	// Host A host object.
	Host Host `json:"host" required:"true"`

	// Type Type of host.
	Type string `json:"type" required:"true"`

	// Name The host name.
	Name string `json:"name" required:"true"`

	// ControlAttributes Attributes to control host.
	ControlAttributes string `json:"control_attributes" required:"true"`

	// Reserved A boolean indicates whether this host is reserved or not, if it is not reserved, false will appear.
	Reserved bool `json:"reserved,omitempty"`

	// OnMaintenance A boolean indicates whether this host is on maintenance or not, if it is not on maintenance mode,
	// false will appear.
	OnMaintenance bool `json:"on_maintenance,omitempty"`
}

// ToContainerCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToContainerCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "host")
}

// Create creates a host in segment.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder, segmentID string) (r CreateResult) {
	b, err := opts.ToContainerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client, segmentID), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete deletes a host in segment.
func Delete(client *gophercloud.ServiceClient, segmentID, hostID string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, segmentID, hostID), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOpts provides options used to update a host to segment.
type UpdateOpts struct {
	// Type Type of host.
	Type string `json:"type" required:"true"`

	// Name The host name.
	Name string `json:"name" required:"true"`

	// ControlAttributes Attributes to control host.
	ControlAttributes string `json:"control_attributes" required:"true"`

	// Reserved A boolean indicates whether this host is reserved or not, if it is not reserved, false will appear.
	Reserved bool `json:"reserved,omitempty"`

	// OnMaintenance A boolean indicates whether this host is on maintenance or not, if it is not on maintenance mode,
	// false will appear.
	OnMaintenance bool `json:"on_maintenance,omitempty"`
}

//ToUpdateRequest formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToUpdateRequest() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "host")
}

// Update will update a host in segment.
func Update(client *gophercloud.ServiceClient, segmentID, hostID string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToUpdateRequest()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(updateURL(client, segmentID, hostID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
