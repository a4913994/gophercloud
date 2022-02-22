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
	Force bool `json:"force,omitempty"`

	Gigabytes int `json:"gigabytes,omitempty"`

	Snapshots int `json:"snapshots,omitempty"`

	SnapshotGigabytes int `json:"snapshot_gigabytes,omitempty"`

	Shares int `json:"shares,omitempty"`

	ShareNetworks int `json:"share_networks,omitempty"`

	// version 2.40
	ShareGroups int `json:"share_groups,omitempty"`
	// version 2.40
	ShareGroupSnapshots int `json:"share_group_snapshots,omitempty"`
	// version 2.53
	ShareReplicas int `json:"share_replicas,omitempty"`
	// version 2.53
	ReplicaGigabytes int `json:"replica_gigabytes,omitempty"`
	// version 2.53
	PerShareGigabytes int `json:"per_share_gigabytes,omitempty"`
}

// ToQuotaUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToQuotaUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "quota_set")
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
