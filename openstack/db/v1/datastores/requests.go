package datastores

import (
	"github.com/lxdcc/gophercloud"
	"github.com/lxdcc/gophercloud/pagination"
)

// List will list all available datastore types that instances can use.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, baseURL(client), func(r pagination.PageResult) pagination.Page {
		return DatastorePage{pagination.SinglePageBase(r)}
	})
}

// Get will retrieve the details of a specified datastore type.
func Get(client *gophercloud.ServiceClient, datastoreID string) (r GetResult) {
	resp, err := client.Get(resourceURL(client, datastoreID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListVersions will list all of the available versions for a specified
// datastore type.
func ListVersions(client *gophercloud.ServiceClient, datastoreID string) pagination.Pager {
	return pagination.NewPager(client, versionsURL(client, datastoreID), func(r pagination.PageResult) pagination.Page {
		return VersionPage{pagination.SinglePageBase(r)}
	})
}

// GetVersion will retrieve the details of a specified datastore version.
func GetVersion(client *gophercloud.ServiceClient, datastoreID, versionID string) (r GetVersionResult) {
	resp, err := client.Get(versionURL(client, datastoreID, versionID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// CreateOpts represents the attributes used when creating a new subnet.
type CreateOpts struct {
	// Name is a human-readable name of the datastore versions.
	Name string `json:"name,omitempty" required:"true"`

	// DatastoreName The name of a datastore.
	DatastoreName string `json:"datastore_name,omitempty" required:"true"`

	// DatastoreManager The type of datastore.
	DatastoreManager string `json:"datastore_manager,omitempty" required:"true"`

	// Image The ID of an image.
	//Either image or image_tags needs to be specified when creating datastore version.
	Image string `json:"image,omitempty"`

	// ImageTags Either image or image_tags needs to be specified when creating datastore version.
	// If the image ID is not provided, the image can be retrieved by the image tags. The tags are used for filtering as a whole rather than separately. Using image tags is more flexible than ID especially when a new guest image is uploaded to Glance, Trove can pick up the latest image automatically for creating instances.
	// When updating, only specifying image_tags could remove image from the datastore version.
	ImageTags []string `json:"image_tags,omitempty"`

	// The UUID of the project who owns the Subnet. Only administrative users
	// can specify a project UUID other than their own.
	ProjectID string `json:"project_id,omitempty"`

	// Active Whether the database version is enabled.
	Active bool `json:"active,omitempty" required:"true"`

	// Default When true this datastore version is created as the default in the datastore. If not specified, for creating, default is false, for updating, it’s ignored.
	Default bool `json:"default,omitempty"`

	// Version The version number for the database. In container based trove instance deployment, the version number is the same as the container image tag, e.g. for MySQL, a valid version number is 5.7.30
	Version string `json:"version,omitempty"`
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// List request.
type CreateOptsBuilder interface {
	ToDataStoreVersionCreateMap() (map[string]interface{}, error)
}

// ToDataStoreVersionCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToDataStoreVersionCreateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "version")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Create Admin only API. Register a datastore version, the datastore is created automatically if it doesn’t exist.
// It’s allowed to create datastore versions with the same name but different version numbers, or vice versa.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToDataStoreVersionCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(datastoreVersionsURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete accepts a unique ID and deletes the subnet associated with it.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := c.Delete(deleteDatastoreVersionsURL(c, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
