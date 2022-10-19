package hosts

import (
	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/openstack/instanceha/v1/segments"
	"github.com/lxdcc/gophercloud/pagination"
)

// Host represents a host list in masakari service.
type Host struct {
	// Name is the name of the host.
	Name string `json:"name"`

	// UUID The UUID of the host.
	UUID string `json:"uuid"`

	// FailoverSegmentID The UUID of the segment.
	FailoverSegmentID string

	// Deleted A boolean indicates whether this resource is deleted or not, if it has not been deleted, false will appear.
	Deleted bool `json:"deleted"`

	// OnMaintenance A boolean indicates whether this host is on maintenance or not, if it is not on maintenance mode,
	// false will appear.
	OnMaintenance bool `json:"on_maintenance"`

	// Reserved A boolean indicates whether this host is reserved or not, if it is not reserved, false will appear.
	Reserved bool `json:"reserved"`

	// CreatedAt The date and time stamp when the host was created.
	CreatedAt string `json:"created_at"`

	// UpdatedAt The date and time stamp when the host was updated.
	UpdatedAt string `json:"updated_at"`

	// ControlAttributes Attributes to control host.
	ControlAttributes string `json:"control_attributes"`

	// FailoverSegment A segment object.
	FailoverSegment segments.FailoverSegment `json:"failover_segment"`

	// Type Type of host.
	Type string `json:"type"`

	// ID ID of host.
	ID string `json:"id"`
}

// HostsPage is a single page of container results.
type HostsPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Container contains any results.
func (r HostsPage) IsEmpty() (bool, error) {
	containers, err := ExtractHosts(r)
	return len(containers) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r HostsPage) NextPageURL() (string, error) {
	var s struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Next, err
}

// ExtractHosts returns a slice of Hosts in a single page of results.
func ExtractHosts(r pagination.Page) ([]Host, error) {
	var s struct {
		Hosts []Host `json:"hosts"`
	}
	err := (r.(HostsPage)).ExtractInto(&s)
	return s.Hosts, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract interprets any commonResult as a Host.
func (r commonResult) Extract() (*Host, error) {
	var s *Host
	err := r.ExtractInto(&s)
	return s, err
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a host.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a host.
type CreateResult struct {
	commonResult
}

// UpdateResult is the response from a Create operation. Call its Extract method
// to interpret it as a host.
type UpdateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
