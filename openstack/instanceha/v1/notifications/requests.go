package notifications

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

	// SortDir Sort direction. A valid value is asc (ascending) or desc (descending). Default is desc.
	// You can specify multiple pairs of sort key and sort direction query parameters.
	// If you omit the sort direction in a pair, the API uses the natural sorting direction of the direction of the segment sort_key attribute.
	SortDir string `q:"sort_dir"`

	// SortKey Sorts by a notification attribute. Default attribute is created_at.
	// You can specify multiple pairs of sort key and sort direction query parameters.
	// If you omit the sort direction in a pair, the API uses the natural sorting direction of the segment sort_key attribute.
	// The sort keys are limited to: created_at,type,generated_time,updated_at,uuid,payload,status,source_host_uuid
	SortKey string `q:"sort_key"`

	// GeneratedSince Filter the notifications list result by notification generated time.
	GeneratedSince string `json:"generated_since"`

	// SourceHostUUID Filter the notifications list result by source_host_uuid.
	SourceHostUUID string `json:"source_host_uuid"`

	// Type Filter the notifications list result by notification type.
	Type string `json:"type"`
}

// ToContainerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToContainerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List retrieves a list notification.
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
		return NotificationPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details of a notification.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(getURL(client, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToContainerCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to add a notification.
type CreateOpts struct {

	// Type of notification, can be either PROCESS, COMPUTE_HOST or VM.
	Type string `json:"type" required:"true"`

	// GeneratedTime The date and time when the notification was created
	GeneratedTime string `json:"generated_time,omitempty"`

	// Payload for notification. This is a JSON string.
	Payload string `json:"payload" required:"true"`

	// HostName A name of host for which notification is created.
	HostName string `json:"host_name" required:"true"`
}

// ToContainerCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToContainerCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "notification")
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
