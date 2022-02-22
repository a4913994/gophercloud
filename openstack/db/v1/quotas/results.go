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
		Quota *Quota `json:"quotas"`
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
	// Backups represents the number of  backups. A "-1" value means no limit.
	Backups int `json:"backups,omitempty"`

	// Instances represents the number of instances. A "-1" value means no limit.
	Instances int `json:"instances"`

	// Ram represents the number of ram. A "-1" value means no limit.
	Ram int `json:"ram"`

	// Volumes represents the number of volumes. A "-1" value means no limit.
	Volumes int `json:"volumes"`
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
