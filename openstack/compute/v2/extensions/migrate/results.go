package migrate

import (
	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/pagination"
)

// MigrateResult is the response from a Migrate operation. Call its ExtractErr
// method to determine if the request suceeded or failed.
type MigrateResult struct {
	gophercloud.ErrResult
}

// NetworkPage stores a single page of all Network results from a List call.
type NetworkPage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a NetworkPage is empty.
func (page NetworkPage) IsEmpty() (bool, error) {
	va, err := ExtractMigration(page)
	return len(va) == 0, err
}

type Migration struct {
	// ID The ID of the server migration.
	ID int `json:"id"`

	// CreatedAt The date and time when the resource was created.
	CreatedAt string `json:"created_at"`

	// DestCompute The target compute for a migration.
	DestCompute string `json:"dest_compute"`

	// DestHost The target node for a migration.
	DestHost string `json:"dest_host"`

	// InstanceUUID The UUID of the server.
	InstanceUUID string `json:"instance_uuid"`

	// NewInstanceTypeID In resize case, the flavor ID for resizing the server.
	// In the other cases, this parameter is same as the flavor ID of the server when the migration was started.
	NewInstanceTypeID int `json:"new_instance_type_id"`

	// OldInstanceTypeID The flavor ID of the server when the migration was started.
	OldInstanceTypeID int `json:"old_instance_type_id"`

	// SourceCompute The source compute for a migration.
	SourceCompute string `json:"source_compute"`

	// SourceNode The source node for a migration.
	SourceNode string `json:"source_node"`

	// Status The current status of the migration.
	Status string `json:"status"`

	// UpdatedAt The date and time when the resource was updated.
	UpdatedAt string `json:"updated_at"`

	// MigrationType The type of the server migration. This is one of live-migration, migration, resize and evacuation.
	MigrationType string `json:"migration_type"`

	// Links Links to the migration. This parameter is returned if the migration type is live-migration and the
	// migration status is one of queued, preparing, running and post-migrating.
	Links interface{} `json:"links"`

	// Uuid The UUID of the migration.
	Uuid string `json:"uuid"`

	// MigrationsLinks Links pertaining to the migration. This parameter is returned when paging and more data is available.
	MigrationsLinks interface{} `json:"migrations_links"`

	// UserID The ID of the user which initiated the server migration. The value may be null for older migration records.
	UserID string `json:"user_id"`

	// ProjectID The ID of the project which initiated the server migration. The value may be null for older migration records.
	ProjectID string `json:"project_id"`
}

// ExtractMigration interprets a page of results as a slice of Migration.
func ExtractMigration(r pagination.Page) ([]Migration, error) {
	var s struct {
		Migrations []Migration `json:"migrations"`
	}
	err := (r.(NetworkPage)).ExtractInto(&s)
	return s.Migrations, err
}
