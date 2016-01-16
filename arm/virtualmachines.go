package arm

import (
	"crypto/rand"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/arm/compute"
	"github.com/Azure/azure-sdk-for-go/arm/network"
	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	"github.com/Azure/azure-sdk-for-go/arm/storage"

	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest/to"
)

// VMParameters contains the required and optional (pointer types) arguments used to create
// a virtual machine in Azure.
type VMParameters struct {
	compute.ImageReference
	Name            *string
	User            string
	Password        string
	Version         *string
	StorageAccount  *string
	AvailabilitySet *string
}

// CreateSimpleVM sets up a single virtual machine in the specified resource group within Azure. The
// function does not accommodate more than a basic amount of customization.
func (client Client) CreateSimpleVM(
	groupName, location *string,
	params VMParameters) (result compute.VirtualMachine, err error) {

	if err = validateStringParameter("userName", &params.User); err != nil {
		return
	}
	if err = validateStringParameter("password", &params.Password); err != nil {
		return
	}

	var group resources.ResourceGroup
	if groupName == nil {
		group, err = createResourceGroup("grp"+randName(8), *location, client)
		if err != nil {
			err = fmt.Errorf("ERROR:'%s'\n", err.Error())
			return
		}
	} else {
		response, e := client.ResourceGroups().CheckExistence(*groupName)
		if e != nil {
			err = fmt.Errorf("ERROR:'%s'\n", e.Error())
			return
		}

		if response.StatusCode == 404 {
			group, err = createResourceGroup(*groupName, *location, client)
			if err != nil {
				err = fmt.Errorf("ERROR:'%s'\n", err.Error())
				return
			}
		} else {
			group, err = client.ResourceGroups().Get(*groupName)
			if err != nil {
				err = fmt.Errorf("ERROR:'%s'\n", err.Error())
				return
			}
		}
	}

	account, err := createStorageAccount(group, params.StorageAccount, client)
	if err != nil {
		err = fmt.Errorf("ERROR: '%s'\n", err.Error())
		return
	}

	params.StorageAccount = to.StringPtr(account)

	var subnet network.Subnet
	var nic network.Interface
	var avset compute.AvailabilitySet

	if params.AvailabilitySet == nil || *params.AvailabilitySet == "" {
		params.AvailabilitySet = to.StringPtr("av-" + randName(8))
	}

	if avset, err = createAvailabilitySet(group, *params.AvailabilitySet, client); err != nil {
		fmt.Printf("ERROR: '%s'\n", err.Error())
		return
	}

	if subnet, err = createNetwork(group, client); err != nil {
		err = fmt.Errorf("ERROR: '%s'\n", err.Error())
		return
	}

	if nic, err = createNetworkInterface(group, subnet, client); err != nil {
		err = fmt.Errorf("ERROR: '%s'\n", err.Error())
		return
	}

	if params.Name == nil || *params.Name == "" {
		params.Name = to.StringPtr("vm" + randName(8))
	}

	if result, err = createVirtualMachine(group, params, avset, nic, client); err != nil {
		err = fmt.Errorf("ERROR: '%s'\n", err.Error())
		return
	}

	return
}

func createResourceGroup(
	name, location string,
	arm Client) (group resources.ResourceGroup, err error) {

	rgc := arm.ResourceGroups()
	rpc := arm.Providers()

	params := resources.ResourceGroup{Name: &name, Location: &location}

	group, err = rgc.CreateOrUpdate(name, params)
	if err != nil {
		err = fmt.Errorf("Failed to create resource group '%s' in location '%s': '%s'\n", name, location, err.Error())
		return
	}

	fmt.Printf("Created resource group '%s'\n", *group.Name)

	if _, err1 := rpc.Register("Microsoft.Storage"); err != nil {
		err = fmt.Errorf("Failed to register resource provider 'Microsoft.Storage': '%s'\n", err1.Error())
	}
	if _, err1 := rpc.Register("Microsoft.Network"); err != nil {
		err = fmt.Errorf("Failed to register resource provider 'Microsoft.Network': '%s'\n", err1.Error())
	}
	if _, err1 := rpc.Register("Microsoft.Compute"); err != nil {
		err = fmt.Errorf("Failed to register resource provider 'Microsoft.Compute': '%s'\n", err1.Error())
	}

	return
}

func createStorageAccount(
	group resources.ResourceGroup,
	accountName *string,
	arm Client) (string, error) {

	ac := arm.StorageAccounts()

	var name string
	if accountName == nil {
		name = *group.Name + "accnt" + randName(8)
	} else {
		name = *accountName
	}

	cna, err := ac.CheckNameAvailability(
		storage.AccountCheckNameAvailabilityParameters{
			Name: &name,
			Type: to.StringPtr("Microsoft.Storage/storageAccounts")})

	if err != nil {
		return name, err
	}

	if to.Bool(cna.NameAvailable) {

		props := storage.AccountPropertiesCreateParameters{AccountType: storage.StandardLRS}

		_, err = ac.Create(*group.Name, name,
			storage.AccountCreateParameters{
				Location:   group.Location,
				Properties: &props,
			})

		if err != nil {
			return name, fmt.Errorf("Failed to create storage account '%s' in location '%s': '%s'\n", name, *group.Location, err.Error())
		}
	}

	return name, nil
}

func createAvailabilitySet(
	group resources.ResourceGroup,
	name string,
	arm Client) (result compute.AvailabilitySet, err error) {

	avsc := arm.AvailabilitySets()

	result, err = avsc.CreateOrUpdate(*group.Name, name, compute.AvailabilitySet{Location: group.Location})
	if err != nil {
		err = fmt.Errorf("Failed to create availability set '%s' in location '%s': '%s'\n", name, *group.Location, err.Error())
		return
	}

	return result, nil
}

func createNetwork(
	group resources.ResourceGroup,
	arm Client) (snetResult network.Subnet, err error) {

	vnetc := arm.VirtualNetworks()
	snetc := arm.Subnets()

	name := *group.Name
	vnet := name + "vnet"
	subnet := name + "subnet"

	snet := network.Subnet{
		Name:       &subnet,
		Properties: &network.SubnetPropertiesFormat{AddressPrefix: to.StringPtr("10.0.0.0/24")}}
	snets := make([]network.Subnet, 1, 1)
	snets[0] = snet

	addrPrefixes := make([]string, 1, 1)
	addrPrefixes[0] = "10.0.0.0/16"
	address := network.AddressSpace{AddressPrefixes: &addrPrefixes}

	nwkProps := network.VirtualNetworkPropertiesFormat{AddressSpace: &address, Subnets: &snets}

	_, err = vnetc.CreateOrUpdate(name, vnet, network.VirtualNetwork{Location: group.Location, Properties: &nwkProps})
	if err != nil {
		err = fmt.Errorf("Failed to create virtual network '%s' in location '%s': '%s'\n", vnet, *group.Location, err.Error())
		return
	}

	snetResult, err = snetc.CreateOrUpdate(name, vnet, subnet, snet)
	if err != nil {
		err = fmt.Errorf("Failed to create subnet '%s' in location '%s': '%s'\n", subnet, *group.Location, err.Error())
	}

	return
}

func createNetworkInterface(
	group resources.ResourceGroup,
	subnet network.Subnet,
	arm Client) (networkInterface network.Interface, err error) {

	pipc := arm.PublicIPAddresses()
	nicc := arm.NetworkInterfaces()

	suffix := randName(4)

	groupName := *group.Name
	ipName := "ip0" + suffix
	nicName := "nic0" + suffix

	pipResult, err := pipc.CreateOrUpdate(
		groupName,
		ipName,
		network.PublicIPAddress{
			Location: group.Location,
			Properties: &network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.Dynamic,
			},
		})

	if err != nil {
		err = fmt.Errorf("Failed to create public ip address '%s' in location '%s': '%s'\n", ipName, *group.Location, err.Error())
		return
	}

	nicProps := network.InterfaceIPConfigurationPropertiesFormat{
		PublicIPAddress: &pipResult,
		Subnet:          &subnet}

	ipConfigs := make([]network.InterfaceIPConfiguration, 1, 1)
	ipConfigs[0] = network.InterfaceIPConfiguration{
		Name:       to.StringPtr(nicName + "Config"),
		Properties: &nicProps,
	}
	props := network.InterfacePropertiesFormat{IPConfigurations: &ipConfigs}

	networkInterface, err = nicc.CreateOrUpdate(
		groupName,
		nicName,
		network.Interface{
			Location:   group.Location,
			Properties: &props,
		})
	if err != nil {
		err = fmt.Errorf("Failed to create network interface '%s' in location '%s': '%s'\n", nicName, *group.Location, err.Error())
	}

	return
}

func createVirtualMachine(
	group resources.ResourceGroup,
	params VMParameters,
	availSet compute.AvailabilitySet,
	networkInterface network.Interface,
	arm Client) (vm compute.VirtualMachine, err error) {

	vmc := arm.VirtualMachines()

	netRefs := make([]compute.NetworkInterfaceReference, 1, 1)
	netRefs[0] = compute.NetworkInterfaceReference{ID: networkInterface.ID}

	groupName := *group.Name
	accountName := *params.StorageAccount

	osDiskName := "osdisk" + randName(4)
	dataDiskName := "ddisk" + randName(4)

	dataDisks := make([]compute.DataDisk, 1, 1)
	dataDisks[0] = compute.DataDisk{
		Name:         &dataDiskName,
		CreateOption: "Empty",
		Vhd:          &compute.VirtualHardDisk{URI: to.StringPtr("http://" + accountName + ".blob.core.windows.net/vhds/" + dataDiskName + ".vhd")},
		DiskSizeGB:   to.IntPtr(100),
		Lun:          to.IntPtr(0),
	}

	vmParams := compute.VirtualMachine{
		Location: group.Location,
		Properties: &compute.VirtualMachineProperties{
			AvailabilitySet: &compute.SubResource{ID: availSet.ID},
			HardwareProfile: &compute.HardwareProfile{VMSize: compute.StandardA0},
			NetworkProfile:  &compute.NetworkProfile{NetworkInterfaces: &netRefs},
			StorageProfile: &compute.StorageProfile{
				ImageReference: &params.ImageReference,
				OsDisk: &compute.OSDisk{
					Name:         &osDiskName,
					CreateOption: compute.FromImage,
					Vhd: &compute.VirtualHardDisk{
						URI: to.StringPtr("http://" + accountName + ".blob.core.windows.net/vhds/" + osDiskName + ".vhd"),
					},
				},
				DataDisks: &dataDisks,
			},
			OsProfile: &compute.OSProfile{
				AdminUsername: to.StringPtr(params.User),
				AdminPassword: to.StringPtr(params.Password),
				ComputerName:  params.Name,
			},
		},
	}

	if vm, err = vmc.CreateOrUpdate(groupName, *params.Name, vmParams); err != nil {
		err = fmt.Errorf("Failed to create virtual machine '%s' in location '%s': '%s'\n", *params.Name, *group.Location, err.Error())
		return
	}

	return
}

func randName(n int) string {
	if n <= 0 {
		panic("negative number")
	}
	const alphanum = "0123456789abcdef"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func validateStringParameter(name string, param *string) error {

	if param == nil || *param == "" {
		return fmt.Errorf("The value passed for '%s' cannot be nil or empty", name)
	}
	return nil
}
