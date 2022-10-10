package shareinstances

import (
	"time"

	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/pagination"
)

// ShareInstances contains all the information associated with an OpenStack
type ShareInstances struct {
	// The Share Network ID
	ID string `json:"id"`

	Status string `json:"status"`

	ShareID string `json:"share_id"`

	AccessRulesStatus string `json:"access_rules_status"`

	Progress string `json:"progress"`

	AvailabilityZone bool `json:"availability_zone"`

	ReplicaState string `json:"replica_state"`

	ExportLocation string `json:"export_location"`

	ExportLocations []string `json:"export_locations"`

	CastRulesToReadOnly bool `json:"cast_rules_to_readonly"`

	ShareNetworkID string `json:"share_network_id"`

	ShareServerID string `json:"share_server_id"`

	Host string `json:"host"`

	ShareTypeID string `json:"share_type_id"`

	// The date and time stamp when the Share Network was created
	CreatedAt time.Time `json:"created_at"`
	// The date and time stamp when the Share Network was updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ShareInstancesPage is a pagination.pager that is returned from a call to the List function.
type ShareInstancesPage struct {
	pagination.MarkerPageBase
}

type commonResult struct {
	gophercloud.Result
}

// ListResult contains the response body and error from a Get request.
type ListResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}
