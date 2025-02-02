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

// `Byoh` manages BYOH in Duplo.
//
// ## Import
//
// Example: Importing an existing BYOH Instance.
//
//   - *TENANT_ID* is the tenant GUID
//
//   - *NAME* is the name of BYOH Instance.
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/byoh:Byoh byoh *TENANT_ID*/*NAME*
// ```
type Byoh struct {
	pulumi.CustomResourceState

	// The numeric ID of the container agent pool that this instance is added to. Defaults to `0`.
	AgentPlatform pulumi.IntPtrOutput `pulumi:"agentPlatform"`
	// Allocation tag for BYOH instance.
	AllocationTag pulumi.StringOutput `pulumi:"allocationTag"`
	ConnectionUrl pulumi.StringOutput `pulumi:"connectionUrl"`
	// IP address of the BYOH instance.
	DirectAddress pulumi.StringOutput `pulumi:"directAddress"`
	// The name of the BYOH instance. Changing this forces a new resource to be created.
	Name            pulumi.StringOutput `pulumi:"name"`
	NetworkAgentUrl pulumi.StringOutput `pulumi:"networkAgentUrl"`
	// Password of the BYOH instance.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Private key for BYOH instance.
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	Tags       ByohTagArrayOutput     `pulumi:"tags"`
	// The GUID of the tenant that the BYHO will be created in.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Username of the BYOH instance.
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// Whether or not to wait until BYOH instance to be connected to the fleet, after creation. Defaults to `false`.
	WaitUntilReady pulumi.BoolPtrOutput `pulumi:"waitUntilReady"`
}

// NewByoh registers a new resource with the given unique name, arguments, and options.
func NewByoh(ctx *pulumi.Context,
	name string, args *ByohArgs, opts ...pulumi.ResourceOption) (*Byoh, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectAddress == nil {
		return nil, errors.New("invalid value for required argument 'DirectAddress'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Byoh
	err := ctx.RegisterResource("duplocloud:index/byoh:Byoh", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetByoh gets an existing Byoh resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetByoh(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ByohState, opts ...pulumi.ResourceOption) (*Byoh, error) {
	var resource Byoh
	err := ctx.ReadResource("duplocloud:index/byoh:Byoh", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Byoh resources.
type byohState struct {
	// The numeric ID of the container agent pool that this instance is added to. Defaults to `0`.
	AgentPlatform *int `pulumi:"agentPlatform"`
	// Allocation tag for BYOH instance.
	AllocationTag *string `pulumi:"allocationTag"`
	ConnectionUrl *string `pulumi:"connectionUrl"`
	// IP address of the BYOH instance.
	DirectAddress *string `pulumi:"directAddress"`
	// The name of the BYOH instance. Changing this forces a new resource to be created.
	Name            *string `pulumi:"name"`
	NetworkAgentUrl *string `pulumi:"networkAgentUrl"`
	// Password of the BYOH instance.
	Password *string `pulumi:"password"`
	// Private key for BYOH instance.
	PrivateKey *string   `pulumi:"privateKey"`
	Tags       []ByohTag `pulumi:"tags"`
	// The GUID of the tenant that the BYHO will be created in.
	TenantId *string `pulumi:"tenantId"`
	// Username of the BYOH instance.
	Username *string `pulumi:"username"`
	// Whether or not to wait until BYOH instance to be connected to the fleet, after creation. Defaults to `false`.
	WaitUntilReady *bool `pulumi:"waitUntilReady"`
}

type ByohState struct {
	// The numeric ID of the container agent pool that this instance is added to. Defaults to `0`.
	AgentPlatform pulumi.IntPtrInput
	// Allocation tag for BYOH instance.
	AllocationTag pulumi.StringPtrInput
	ConnectionUrl pulumi.StringPtrInput
	// IP address of the BYOH instance.
	DirectAddress pulumi.StringPtrInput
	// The name of the BYOH instance. Changing this forces a new resource to be created.
	Name            pulumi.StringPtrInput
	NetworkAgentUrl pulumi.StringPtrInput
	// Password of the BYOH instance.
	Password pulumi.StringPtrInput
	// Private key for BYOH instance.
	PrivateKey pulumi.StringPtrInput
	Tags       ByohTagArrayInput
	// The GUID of the tenant that the BYHO will be created in.
	TenantId pulumi.StringPtrInput
	// Username of the BYOH instance.
	Username pulumi.StringPtrInput
	// Whether or not to wait until BYOH instance to be connected to the fleet, after creation. Defaults to `false`.
	WaitUntilReady pulumi.BoolPtrInput
}

func (ByohState) ElementType() reflect.Type {
	return reflect.TypeOf((*byohState)(nil)).Elem()
}

type byohArgs struct {
	// The numeric ID of the container agent pool that this instance is added to. Defaults to `0`.
	AgentPlatform *int `pulumi:"agentPlatform"`
	// Allocation tag for BYOH instance.
	AllocationTag *string `pulumi:"allocationTag"`
	// IP address of the BYOH instance.
	DirectAddress string `pulumi:"directAddress"`
	// The name of the BYOH instance. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Password of the BYOH instance.
	Password *string `pulumi:"password"`
	// Private key for BYOH instance.
	PrivateKey *string `pulumi:"privateKey"`
	// The GUID of the tenant that the BYHO will be created in.
	TenantId string `pulumi:"tenantId"`
	// Username of the BYOH instance.
	Username *string `pulumi:"username"`
	// Whether or not to wait until BYOH instance to be connected to the fleet, after creation. Defaults to `false`.
	WaitUntilReady *bool `pulumi:"waitUntilReady"`
}

// The set of arguments for constructing a Byoh resource.
type ByohArgs struct {
	// The numeric ID of the container agent pool that this instance is added to. Defaults to `0`.
	AgentPlatform pulumi.IntPtrInput
	// Allocation tag for BYOH instance.
	AllocationTag pulumi.StringPtrInput
	// IP address of the BYOH instance.
	DirectAddress pulumi.StringInput
	// The name of the BYOH instance. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Password of the BYOH instance.
	Password pulumi.StringPtrInput
	// Private key for BYOH instance.
	PrivateKey pulumi.StringPtrInput
	// The GUID of the tenant that the BYHO will be created in.
	TenantId pulumi.StringInput
	// Username of the BYOH instance.
	Username pulumi.StringPtrInput
	// Whether or not to wait until BYOH instance to be connected to the fleet, after creation. Defaults to `false`.
	WaitUntilReady pulumi.BoolPtrInput
}

func (ByohArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*byohArgs)(nil)).Elem()
}

type ByohInput interface {
	pulumi.Input

	ToByohOutput() ByohOutput
	ToByohOutputWithContext(ctx context.Context) ByohOutput
}

func (*Byoh) ElementType() reflect.Type {
	return reflect.TypeOf((**Byoh)(nil)).Elem()
}

func (i *Byoh) ToByohOutput() ByohOutput {
	return i.ToByohOutputWithContext(context.Background())
}

func (i *Byoh) ToByohOutputWithContext(ctx context.Context) ByohOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByohOutput)
}

// ByohArrayInput is an input type that accepts ByohArray and ByohArrayOutput values.
// You can construct a concrete instance of `ByohArrayInput` via:
//
//	ByohArray{ ByohArgs{...} }
type ByohArrayInput interface {
	pulumi.Input

	ToByohArrayOutput() ByohArrayOutput
	ToByohArrayOutputWithContext(context.Context) ByohArrayOutput
}

type ByohArray []ByohInput

func (ByohArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Byoh)(nil)).Elem()
}

func (i ByohArray) ToByohArrayOutput() ByohArrayOutput {
	return i.ToByohArrayOutputWithContext(context.Background())
}

func (i ByohArray) ToByohArrayOutputWithContext(ctx context.Context) ByohArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByohArrayOutput)
}

// ByohMapInput is an input type that accepts ByohMap and ByohMapOutput values.
// You can construct a concrete instance of `ByohMapInput` via:
//
//	ByohMap{ "key": ByohArgs{...} }
type ByohMapInput interface {
	pulumi.Input

	ToByohMapOutput() ByohMapOutput
	ToByohMapOutputWithContext(context.Context) ByohMapOutput
}

type ByohMap map[string]ByohInput

func (ByohMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Byoh)(nil)).Elem()
}

func (i ByohMap) ToByohMapOutput() ByohMapOutput {
	return i.ToByohMapOutputWithContext(context.Background())
}

func (i ByohMap) ToByohMapOutputWithContext(ctx context.Context) ByohMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByohMapOutput)
}

type ByohOutput struct{ *pulumi.OutputState }

func (ByohOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Byoh)(nil)).Elem()
}

func (o ByohOutput) ToByohOutput() ByohOutput {
	return o
}

func (o ByohOutput) ToByohOutputWithContext(ctx context.Context) ByohOutput {
	return o
}

// The numeric ID of the container agent pool that this instance is added to. Defaults to `0`.
func (o ByohOutput) AgentPlatform() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Byoh) pulumi.IntPtrOutput { return v.AgentPlatform }).(pulumi.IntPtrOutput)
}

// Allocation tag for BYOH instance.
func (o ByohOutput) AllocationTag() pulumi.StringOutput {
	return o.ApplyT(func(v *Byoh) pulumi.StringOutput { return v.AllocationTag }).(pulumi.StringOutput)
}

func (o ByohOutput) ConnectionUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Byoh) pulumi.StringOutput { return v.ConnectionUrl }).(pulumi.StringOutput)
}

// IP address of the BYOH instance.
func (o ByohOutput) DirectAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Byoh) pulumi.StringOutput { return v.DirectAddress }).(pulumi.StringOutput)
}

// The name of the BYOH instance. Changing this forces a new resource to be created.
func (o ByohOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Byoh) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ByohOutput) NetworkAgentUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Byoh) pulumi.StringOutput { return v.NetworkAgentUrl }).(pulumi.StringOutput)
}

// Password of the BYOH instance.
func (o ByohOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Byoh) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Private key for BYOH instance.
func (o ByohOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Byoh) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o ByohOutput) Tags() ByohTagArrayOutput {
	return o.ApplyT(func(v *Byoh) ByohTagArrayOutput { return v.Tags }).(ByohTagArrayOutput)
}

// The GUID of the tenant that the BYHO will be created in.
func (o ByohOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Byoh) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Username of the BYOH instance.
func (o ByohOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Byoh) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

// Whether or not to wait until BYOH instance to be connected to the fleet, after creation. Defaults to `false`.
func (o ByohOutput) WaitUntilReady() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Byoh) pulumi.BoolPtrOutput { return v.WaitUntilReady }).(pulumi.BoolPtrOutput)
}

type ByohArrayOutput struct{ *pulumi.OutputState }

func (ByohArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Byoh)(nil)).Elem()
}

func (o ByohArrayOutput) ToByohArrayOutput() ByohArrayOutput {
	return o
}

func (o ByohArrayOutput) ToByohArrayOutputWithContext(ctx context.Context) ByohArrayOutput {
	return o
}

func (o ByohArrayOutput) Index(i pulumi.IntInput) ByohOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Byoh {
		return vs[0].([]*Byoh)[vs[1].(int)]
	}).(ByohOutput)
}

type ByohMapOutput struct{ *pulumi.OutputState }

func (ByohMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Byoh)(nil)).Elem()
}

func (o ByohMapOutput) ToByohMapOutput() ByohMapOutput {
	return o
}

func (o ByohMapOutput) ToByohMapOutputWithContext(ctx context.Context) ByohMapOutput {
	return o
}

func (o ByohMapOutput) MapIndex(k pulumi.StringInput) ByohOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Byoh {
		return vs[0].(map[string]*Byoh)[vs[1].(string)]
	}).(ByohOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ByohInput)(nil)).Elem(), &Byoh{})
	pulumi.RegisterInputType(reflect.TypeOf((*ByohArrayInput)(nil)).Elem(), ByohArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ByohMapInput)(nil)).Elem(), ByohMap{})
	pulumi.RegisterOutputType(ByohOutput{})
	pulumi.RegisterOutputType(ByohArrayOutput{})
	pulumi.RegisterOutputType(ByohMapOutput{})
}
