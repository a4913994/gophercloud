package quotas

import (
	"github.com/lxdcc/gophercloud"
)

// Get returns load balancer Quotas for a project.
func Get(client *gophercloud.ServiceClient, projectID string) (r GetResult) {
	resp, err := client.Get(quotaGetURL(client, projectID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToQuotaUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update the load balancer Quotas.
type UpdateOpts struct {
	// Backups represents the number of  backups. A "-1" value means no limit.
	Backups int `json:"backups,omitempty"`

	// Instances represents the number of instances. A "-1" value means no limit.
	Instances int `json:"instances"`

	// Ram represents the number of ram. A "-1" value means no limit.
	Ram int `json:"ram"`

	// Volumes represents the number of volumes. A "-1" value means no limit.
	Volumes int `json:"volumes"`
}

// ToQuotaUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToQuotaUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "quotas")
}

// Update accepts a UpdateOpts struct and updates an existing load balancer Quotas using the
// values provided.
func Update(c *gophercloud.ServiceClient, projectID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToQuotaUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(quotaUpdateURL(c, projectID), b, &r.Body, &gophercloud.RequestOpts{
		// allow 200 (neutron/lbaasv2) and 202 (octavia)
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
