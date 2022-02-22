package quotas

import (
	"encoding/json"

	"github.com/lxdcc/gophercloud"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a Quota resource.
func (r commonResult) Extract() (*Quota, error) {
	var s struct {
		Quota *Quota `json:"quota_set"`
	}
	err := r.ExtractInto(&s)
	return s.Quota, err
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Quota.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Quota.
type UpdateResult struct {
	commonResult
}

// Quota contains load balancer quotas for a project.
type Quota struct {
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

// UnmarshalJSON provides backwards compatibility to OpenStack APIs which still
// return the deprecated `load_balancer` or `health_monitor` as quota values
// instead of `loadbalancer` and `healthmonitor`.
func (r *Quota) UnmarshalJSON(b []byte) error {
	type tmp Quota

	// Support both underscore and non-underscore naming.
	var s struct {
		tmp
	}

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*r = Quota(s.tmp)

	return nil
}
