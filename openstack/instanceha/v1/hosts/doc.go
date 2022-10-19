/*
Package hosts and retrieves hosts in the OpenStack masakari Service.

Example to List Hosts

	createdQuery := &hosts.List{}

	listOpts := hosts.ListOpts{}

	allPages, err := hosts.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allHosts, err := hosts.ExtractHosts(allPages)
	if err != nil {
		panic(err)
	}

	for _, v := range allSecrets {
		fmt.Printf("%v\n", v)
	}

Example to Get a Host

	host, err := hosts.Get(client, segmentID, hostID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", host)

Example to Create a Host

	createOpts := hosts.CreateOpts{
	}

	host, err := hosts.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Println(host.Name)

Example to Delete a Secrets

	err := hosts.Delete(client, segmentID, hostID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Update of a  Host in segment

	opts := hosts.UpdateOpts{
		Key:   "foo",
		Value: "bar",
	}

	r, err := hosts.Update(client, segmentID, hostID, opts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", r)
*/
package hosts
