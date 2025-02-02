// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `AzureAvailabilitySet` manages logical groupings of VMs that enhance reliability by placing VMs in different fault domains to minimize correlated failures, offering improved VM-to-VM latency and high availability, with no extra cost beyond the VM instances themselves, though they may still be affected by shared infrastructure failures.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myapp, err := duplocloud.NewTenant(ctx, "myapp", &duplocloud.TenantArgs{
//				AccountName: pulumi.String("myapp"),
//				PlanId:      pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = duplocloud.NewAzureAvailabilitySet(ctx, "st", &duplocloud.AzureAvailabilitySetArgs{
//				TenantId:                  myapp.TenantId,
//				Name:                      pulumi.String("availset"),
//				UseManagedDisk:            pulumi.String("Aligned"),
//				PlatformUpdateDomainCount: pulumi.Int(5),
//				PlatformFaultDomainCount:  pulumi.Int(2),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Example: Importing an existing Azure Availablitu set
//
//   - *TENANT_ID* is the tenant GUID
//
//   - *NAME* is the  name of the Azure Availability set
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/azureAvailabilitySet:AzureAvailabilitySet this *TENANT_ID*/availability-set/*NAME*
// ```
type AzureAvailabilitySet struct {
	pulumi.CustomResourceState

	AvailabilitySetId pulumi.StringOutput `pulumi:"availabilitySetId"`
	Location          pulumi.StringOutput `pulumi:"location"`
	// The name for availability set
	Name pulumi.StringOutput `pulumi:"name"`
	// Specify platform fault domain count betweem 1-3, for availability set. Virtual machines in the same fault domain share a common power source and physical network switch. Defaults to `2`.
	PlatformFaultDomainCount pulumi.IntPtrOutput `pulumi:"platformFaultDomainCount"`
	// Specify platform update domain count between 1-20, for availability set. Virtual machines in the same update domain will be restarted together during planned maintenance. Azure never restarts more than one update domain at a time. Defaults to `5`.
	PlatformUpdateDomainCount pulumi.IntPtrOutput    `pulumi:"platformUpdateDomainCount"`
	Tags                      pulumi.StringMapOutput `pulumi:"tags"`
	// The GUID of the tenant that the host will be created in.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	Type     pulumi.StringOutput `pulumi:"type"`
	// Set this to `Aligned` if you plan to create virtual machines in this availability set with managed disks. Defaults to `Classic`.
	UseManagedDisk  pulumi.StringPtrOutput                        `pulumi:"useManagedDisk"`
	VirtualMachines AzureAvailabilitySetVirtualMachineArrayOutput `pulumi:"virtualMachines"`
}

// NewAzureAvailabilitySet registers a new resource with the given unique name, arguments, and options.
func NewAzureAvailabilitySet(ctx *pulumi.Context,
	name string, args *AzureAvailabilitySetArgs, opts ...pulumi.ResourceOption) (*AzureAvailabilitySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AzureAvailabilitySet
	err := ctx.RegisterResource("duplocloud:index/azureAvailabilitySet:AzureAvailabilitySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureAvailabilitySet gets an existing AzureAvailabilitySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureAvailabilitySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureAvailabilitySetState, opts ...pulumi.ResourceOption) (*AzureAvailabilitySet, error) {
	var resource AzureAvailabilitySet
	err := ctx.ReadResource("duplocloud:index/azureAvailabilitySet:AzureAvailabilitySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureAvailabilitySet resources.
type azureAvailabilitySetState struct {
	AvailabilitySetId *string `pulumi:"availabilitySetId"`
	Location          *string `pulumi:"location"`
	// The name for availability set
	Name *string `pulumi:"name"`
	// Specify platform fault domain count betweem 1-3, for availability set. Virtual machines in the same fault domain share a common power source and physical network switch. Defaults to `2`.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Specify platform update domain count between 1-20, for availability set. Virtual machines in the same update domain will be restarted together during planned maintenance. Azure never restarts more than one update domain at a time. Defaults to `5`.
	PlatformUpdateDomainCount *int              `pulumi:"platformUpdateDomainCount"`
	Tags                      map[string]string `pulumi:"tags"`
	// The GUID of the tenant that the host will be created in.
	TenantId *string `pulumi:"tenantId"`
	Type     *string `pulumi:"type"`
	// Set this to `Aligned` if you plan to create virtual machines in this availability set with managed disks. Defaults to `Classic`.
	UseManagedDisk  *string                              `pulumi:"useManagedDisk"`
	VirtualMachines []AzureAvailabilitySetVirtualMachine `pulumi:"virtualMachines"`
}

type AzureAvailabilitySetState struct {
	AvailabilitySetId pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	// The name for availability set
	Name pulumi.StringPtrInput
	// Specify platform fault domain count betweem 1-3, for availability set. Virtual machines in the same fault domain share a common power source and physical network switch. Defaults to `2`.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Specify platform update domain count between 1-20, for availability set. Virtual machines in the same update domain will be restarted together during planned maintenance. Azure never restarts more than one update domain at a time. Defaults to `5`.
	PlatformUpdateDomainCount pulumi.IntPtrInput
	Tags                      pulumi.StringMapInput
	// The GUID of the tenant that the host will be created in.
	TenantId pulumi.StringPtrInput
	Type     pulumi.StringPtrInput
	// Set this to `Aligned` if you plan to create virtual machines in this availability set with managed disks. Defaults to `Classic`.
	UseManagedDisk  pulumi.StringPtrInput
	VirtualMachines AzureAvailabilitySetVirtualMachineArrayInput
}

func (AzureAvailabilitySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureAvailabilitySetState)(nil)).Elem()
}

type azureAvailabilitySetArgs struct {
	// The name for availability set
	Name *string `pulumi:"name"`
	// Specify platform fault domain count betweem 1-3, for availability set. Virtual machines in the same fault domain share a common power source and physical network switch. Defaults to `2`.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Specify platform update domain count between 1-20, for availability set. Virtual machines in the same update domain will be restarted together during planned maintenance. Azure never restarts more than one update domain at a time. Defaults to `5`.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// The GUID of the tenant that the host will be created in.
	TenantId string `pulumi:"tenantId"`
	// Set this to `Aligned` if you plan to create virtual machines in this availability set with managed disks. Defaults to `Classic`.
	UseManagedDisk *string `pulumi:"useManagedDisk"`
}

// The set of arguments for constructing a AzureAvailabilitySet resource.
type AzureAvailabilitySetArgs struct {
	// The name for availability set
	Name pulumi.StringPtrInput
	// Specify platform fault domain count betweem 1-3, for availability set. Virtual machines in the same fault domain share a common power source and physical network switch. Defaults to `2`.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Specify platform update domain count between 1-20, for availability set. Virtual machines in the same update domain will be restarted together during planned maintenance. Azure never restarts more than one update domain at a time. Defaults to `5`.
	PlatformUpdateDomainCount pulumi.IntPtrInput
	// The GUID of the tenant that the host will be created in.
	TenantId pulumi.StringInput
	// Set this to `Aligned` if you plan to create virtual machines in this availability set with managed disks. Defaults to `Classic`.
	UseManagedDisk pulumi.StringPtrInput
}

func (AzureAvailabilitySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureAvailabilitySetArgs)(nil)).Elem()
}

type AzureAvailabilitySetInput interface {
	pulumi.Input

	ToAzureAvailabilitySetOutput() AzureAvailabilitySetOutput
	ToAzureAvailabilitySetOutputWithContext(ctx context.Context) AzureAvailabilitySetOutput
}

func (*AzureAvailabilitySet) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureAvailabilitySet)(nil)).Elem()
}

func (i *AzureAvailabilitySet) ToAzureAvailabilitySetOutput() AzureAvailabilitySetOutput {
	return i.ToAzureAvailabilitySetOutputWithContext(context.Background())
}

func (i *AzureAvailabilitySet) ToAzureAvailabilitySetOutputWithContext(ctx context.Context) AzureAvailabilitySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureAvailabilitySetOutput)
}

// AzureAvailabilitySetArrayInput is an input type that accepts AzureAvailabilitySetArray and AzureAvailabilitySetArrayOutput values.
// You can construct a concrete instance of `AzureAvailabilitySetArrayInput` via:
//
//	AzureAvailabilitySetArray{ AzureAvailabilitySetArgs{...} }
type AzureAvailabilitySetArrayInput interface {
	pulumi.Input

	ToAzureAvailabilitySetArrayOutput() AzureAvailabilitySetArrayOutput
	ToAzureAvailabilitySetArrayOutputWithContext(context.Context) AzureAvailabilitySetArrayOutput
}

type AzureAvailabilitySetArray []AzureAvailabilitySetInput

func (AzureAvailabilitySetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureAvailabilitySet)(nil)).Elem()
}

func (i AzureAvailabilitySetArray) ToAzureAvailabilitySetArrayOutput() AzureAvailabilitySetArrayOutput {
	return i.ToAzureAvailabilitySetArrayOutputWithContext(context.Background())
}

func (i AzureAvailabilitySetArray) ToAzureAvailabilitySetArrayOutputWithContext(ctx context.Context) AzureAvailabilitySetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureAvailabilitySetArrayOutput)
}

// AzureAvailabilitySetMapInput is an input type that accepts AzureAvailabilitySetMap and AzureAvailabilitySetMapOutput values.
// You can construct a concrete instance of `AzureAvailabilitySetMapInput` via:
//
//	AzureAvailabilitySetMap{ "key": AzureAvailabilitySetArgs{...} }
type AzureAvailabilitySetMapInput interface {
	pulumi.Input

	ToAzureAvailabilitySetMapOutput() AzureAvailabilitySetMapOutput
	ToAzureAvailabilitySetMapOutputWithContext(context.Context) AzureAvailabilitySetMapOutput
}

type AzureAvailabilitySetMap map[string]AzureAvailabilitySetInput

func (AzureAvailabilitySetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureAvailabilitySet)(nil)).Elem()
}

func (i AzureAvailabilitySetMap) ToAzureAvailabilitySetMapOutput() AzureAvailabilitySetMapOutput {
	return i.ToAzureAvailabilitySetMapOutputWithContext(context.Background())
}

func (i AzureAvailabilitySetMap) ToAzureAvailabilitySetMapOutputWithContext(ctx context.Context) AzureAvailabilitySetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureAvailabilitySetMapOutput)
}

type AzureAvailabilitySetOutput struct{ *pulumi.OutputState }

func (AzureAvailabilitySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureAvailabilitySet)(nil)).Elem()
}

func (o AzureAvailabilitySetOutput) ToAzureAvailabilitySetOutput() AzureAvailabilitySetOutput {
	return o
}

func (o AzureAvailabilitySetOutput) ToAzureAvailabilitySetOutputWithContext(ctx context.Context) AzureAvailabilitySetOutput {
	return o
}

func (o AzureAvailabilitySetOutput) AvailabilitySetId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) pulumi.StringOutput { return v.AvailabilitySetId }).(pulumi.StringOutput)
}

func (o AzureAvailabilitySetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name for availability set
func (o AzureAvailabilitySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specify platform fault domain count betweem 1-3, for availability set. Virtual machines in the same fault domain share a common power source and physical network switch. Defaults to `2`.
func (o AzureAvailabilitySetOutput) PlatformFaultDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) pulumi.IntPtrOutput { return v.PlatformFaultDomainCount }).(pulumi.IntPtrOutput)
}

// Specify platform update domain count between 1-20, for availability set. Virtual machines in the same update domain will be restarted together during planned maintenance. Azure never restarts more than one update domain at a time. Defaults to `5`.
func (o AzureAvailabilitySetOutput) PlatformUpdateDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) pulumi.IntPtrOutput { return v.PlatformUpdateDomainCount }).(pulumi.IntPtrOutput)
}

func (o AzureAvailabilitySetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The GUID of the tenant that the host will be created in.
func (o AzureAvailabilitySetOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o AzureAvailabilitySetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Set this to `Aligned` if you plan to create virtual machines in this availability set with managed disks. Defaults to `Classic`.
func (o AzureAvailabilitySetOutput) UseManagedDisk() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) pulumi.StringPtrOutput { return v.UseManagedDisk }).(pulumi.StringPtrOutput)
}

func (o AzureAvailabilitySetOutput) VirtualMachines() AzureAvailabilitySetVirtualMachineArrayOutput {
	return o.ApplyT(func(v *AzureAvailabilitySet) AzureAvailabilitySetVirtualMachineArrayOutput { return v.VirtualMachines }).(AzureAvailabilitySetVirtualMachineArrayOutput)
}

type AzureAvailabilitySetArrayOutput struct{ *pulumi.OutputState }

func (AzureAvailabilitySetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureAvailabilitySet)(nil)).Elem()
}

func (o AzureAvailabilitySetArrayOutput) ToAzureAvailabilitySetArrayOutput() AzureAvailabilitySetArrayOutput {
	return o
}

func (o AzureAvailabilitySetArrayOutput) ToAzureAvailabilitySetArrayOutputWithContext(ctx context.Context) AzureAvailabilitySetArrayOutput {
	return o
}

func (o AzureAvailabilitySetArrayOutput) Index(i pulumi.IntInput) AzureAvailabilitySetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureAvailabilitySet {
		return vs[0].([]*AzureAvailabilitySet)[vs[1].(int)]
	}).(AzureAvailabilitySetOutput)
}

type AzureAvailabilitySetMapOutput struct{ *pulumi.OutputState }

func (AzureAvailabilitySetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureAvailabilitySet)(nil)).Elem()
}

func (o AzureAvailabilitySetMapOutput) ToAzureAvailabilitySetMapOutput() AzureAvailabilitySetMapOutput {
	return o
}

func (o AzureAvailabilitySetMapOutput) ToAzureAvailabilitySetMapOutputWithContext(ctx context.Context) AzureAvailabilitySetMapOutput {
	return o
}

func (o AzureAvailabilitySetMapOutput) MapIndex(k pulumi.StringInput) AzureAvailabilitySetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureAvailabilitySet {
		return vs[0].(map[string]*AzureAvailabilitySet)[vs[1].(string)]
	}).(AzureAvailabilitySetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureAvailabilitySetInput)(nil)).Elem(), &AzureAvailabilitySet{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureAvailabilitySetArrayInput)(nil)).Elem(), AzureAvailabilitySetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureAvailabilitySetMapInput)(nil)).Elem(), AzureAvailabilitySetMap{})
	pulumi.RegisterOutputType(AzureAvailabilitySetOutput{})
	pulumi.RegisterOutputType(AzureAvailabilitySetArrayOutput{})
	pulumi.RegisterOutputType(AzureAvailabilitySetMapOutput{})
}
