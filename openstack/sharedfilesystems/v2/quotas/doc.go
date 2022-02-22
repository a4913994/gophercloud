/*
Package quotas provides the ability to retrieve and manage Trove quotas

Example to Get project quotas

    projectID = "23d5d3f79dfa4f73b72b8b0b0063ec55"
    quotasInfo, err := quotas.Get(networkClient, projectID).Extract()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("quotas: %#v\n", quotasInfo)

Example to Update project quotas

    projectID = "23d5d3f79dfa4f73b72b8b0b0063ec55"

    updateOpts := quotas.UpdateOpts{
		Backups: gophercloud.IntToPointer(20),
		Instances: gophercloud.IntToPointer(20),
		Ram: gophercloud.IntToPointer(20),
		Volumes: gophercloud.IntToPointer(20),
    }
    quotasInfo, err := quotas.Update(networkClient, projectID)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("quotas: %#v\n", quotasInfo)
*/
package quotas
