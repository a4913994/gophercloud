package segments

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

	// RecoveryMethod Filter the segment list result by recovery_method.
	RecoveryMethod bool `q:"recovery_method"`

	// ServiceType Filter the segment list result by service_type.
	ServiceType string `q:"service_type"`

	// Enabled Boolean whether this segment is enabled or not.
	Enabled bool `q:"enabled"`

	// SortDir Sort direction. A valid value is asc (ascending) or desc (descending). Default is desc.
	// You can specify multiple pairs of sort key and sort direction query parameters. If you omit the sort direction in
	// a pair, the API uses the natural sorting direction of the direction of the segment sort_key attribute.
	SortDir string `q:"sort_dir"`

	// SortKey Sorts by a hosts attribute. Default attribute is created_at. You can specify multiple pairs of sort key
	// and sort direction query parameters. If you omit the sort direction in a pair, the API uses the natural sorting
	// direction of the segment sort_key attribute. The sort keys are limited to:
	// created_at, description, name, updated_at, uuid, recovery_method, service_type
	SortKey string `q:"sort_key"`
}

// ToContainerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToContainerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List retrieves a list segment.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToContainerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SegmentsPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details of a segment.
func Get(client *gophercloud.ServiceClient, segmentID string) (r GetResult) {
	resp, err := client.Get(getURL(client, segmentID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToContainerCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to add a segment.
type CreateOpts struct {
	// Name The segment name.
	Name string `json:"name" required:"true"`

	// RecoveryMethod Type of recovery if any host in this segment goes down. User can mention either
	// ‘auto’, ‘reserved_host’, ‘auto_priority’ or ‘rh_priority’.
	RecoveryMethod string `json:"recovery_method" required:"true"`

	// ServiceType The name of service which will be deployed in this segment.
	// As of now user can mention ‘COMPUTE’ as service_type.
	ServiceType string `json:"service_type" required:"true"`

	// Enabled Boolean whether this segment is enabled or not.
	Enabled bool `json:"enabled,omitempty"`

	// Description The segment description.
	Description string `json:"description,omitempty"`
}

// ToContainerCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToContainerCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "segment")
}

// Create creates a new segment.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToContainerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(createURL(client), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{201, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete deletes a segment.
func Delete(client *gophercloud.ServiceClient, segmentID string) (r DeleteResult) {
	resp, err := client.Delete(deleteURL(client, segmentID), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOpts provides options used to update a segment.
type UpdateOpts struct {
	// Name The segment name.
	Name string `json:"name" required:"true"`

	// RecoveryMethod Type of recovery if any host in this segment goes down. User can mention either
	// ‘auto’, ‘reserved_host’, ‘auto_priority’ or ‘rh_priority’.
	RecoveryMethod string `json:"recovery_method" required:"true"`

	// ServiceType The name of service which will be deployed in this segment.
	// As of now user can mention ‘COMPUTE’ as service_type.
	ServiceType string `json:"service_type" required:"true"`

	// Enabled Boolean whether this segment is enabled or not.
	Enabled bool `json:"enabled,omitempty"`

	// Description The segment description.
	Description string `json:"description,omitempty"`
}

//ToUpdateRequest formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToUpdateRequest() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "segment")
}

// Update will update a segment.
func Update(client *gophercloud.ServiceClient, segmentID string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToUpdateRequest()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(updateURL(client, segmentID), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
