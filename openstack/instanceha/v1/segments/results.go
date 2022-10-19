package segments

import (
	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/pagination"
)

type FailoverSegment struct {

	// ID The Id of the segment.
	ID string `json:"id"`

	// Name The segment name.
	Name string `json:"name"`

	// Description A free form description of the segment. Limited to 255 characters in length.
	Description string `json:"description"`

	// Created The date and time stamp when the segment was created.
	Created string `json:"created"`

	// Updated The date and time stamp when the segment was updated.
	Updated string `json:"updated"`

	// RecoveryMethod Type of recovery if any host in this segment goes down. User can mention either
	// ‘auto’, ‘reserved_host’, ‘auto_priority’ or ‘rh_priority’.
	RecoveryMethod string `json:"recovery_method"`

	// ServiceType The name of service which will be deployed in this segment.
	// As of now user can mention ‘COMPUTE’ as service_type.
	ServiceType string `json:"service_type"`

	// Enabled Boolean whether this segment is enabled or not.
	Enabled bool `json:"enabled"`

	// UUID The UUID of the segment.
	UUID string `json:"uuid"`
}

// SegmentsPage is a single page of container results.
type SegmentsPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Container contains any results.
func (r SegmentsPage) IsEmpty() (bool, error) {
	containers, err := ExtractSegments(r)
	return len(containers) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r SegmentsPage) NextPageURL() (string, error) {
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

// ExtractSegments returns a slice of Segments in a single page of results.
func ExtractSegments(r pagination.Page) ([]FailoverSegment, error) {
	var s struct {
		Segments []FailoverSegment `json:"segments"`
	}
	err := (r.(SegmentsPage)).ExtractInto(&s)
	return s.Segments, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract interprets any commonResult as a Segment.
func (r commonResult) Extract() (*FailoverSegment, error) {
	var s *FailoverSegment
	err := r.ExtractInto(&s)
	return s, err
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a Segment.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Segment.
type CreateResult struct {
	commonResult
}

// UpdateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Segment.
type UpdateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
