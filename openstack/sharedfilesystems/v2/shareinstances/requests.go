package shareinstances

import (
	"github.com/lxdcc/gophercloud"
)

// ListOpts holds options for listing ShareInstances. It is passed to the
// shareInstances.List function.
type ListOpts struct {
	// The UUID of the project where the share network was created
	ProjectID string `q:"project_id"`
	// The export Location id
	ExportLocationID string `q:"export_location_id"`
	// The export Location path
	ExportLocationPath string `q:"export_location_path"`
}

// ToShareInstancesListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToShareInstancesListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns ShareInstances optionally limited by the conditions provided in ListOpts.
func List(client *gophercloud.ServiceClient, opts *ListOpts) (r ListResult) {
	url := listInstancesURL(client)
	if opts != nil {
		query, err := opts.ToShareInstancesListQuery()
		if err != nil {
			r.Err = err
			return r
		}
		url += query
	}

	resp, err := client.Get(url, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getInstancesURL(client, id), &r.Body, nil)
	return
}

func GetServerDetail(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getServerDetailURL(client, id), &r.Body, nil)
	return
}
