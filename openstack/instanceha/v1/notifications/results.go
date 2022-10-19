package notifications

import (
	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/pagination"
)

type Notification struct {
	// ID ID of notification.
	ID string `json:"id"`

	// NotificationUUID The UUID of the notification.
	NotificationUUID string `json:"notification_uuid"`

	// CreatedAt The date and time stamp when the segment was created.
	CreatedAt string `json:"created_at"`

	// UpdatedAt The date and time stamp when the segment was updated.
	UpdatedAt string `json:"updated_at"`

	// Delete A boolean indicates whether this resource is deleted or not, if it has not been deleted, false will appear.
	Delete bool `json:"delete"`

	// Status Status of notification.
	Status string `json:"status"`

	// UUID The UUID of the notification.
	UUID string `json:"uuid"`

	// SourceHostUUID The UUID of host for which notification is generated.
	SourceHostUUID string `json:"source_host_uuid"`

	// GeneratedTime The date and time stamp when the notification was created.
	GeneratedTime string `json:"generated_time"`

	// Type Type of notification. can be either PROCESS, COMPUTE_HOST or VM.
	Type string `json:"type"`

	// Payload Payload of notification.
	Payload string `json:"payload"`
}

// NotificationPage is a single page of notification results.
type NotificationPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Container contains any results.
func (r NotificationPage) IsEmpty() (bool, error) {
	containers, err := ExtractNotifications(r)
	return len(containers) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r NotificationPage) NextPageURL() (string, error) {
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

// ExtractNotifications returns a slice of Notification in a single page of results.
func ExtractNotifications(r pagination.Page) ([]Notification, error) {
	var s struct {
		Notifications []Notification `json:"notifications"`
	}
	err := (r.(NotificationPage)).ExtractInto(&s)
	return s.Notifications, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract interprets any commonResult as a Segment.
func (r commonResult) Extract() (*Notification, error) {
	var s *Notification
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
