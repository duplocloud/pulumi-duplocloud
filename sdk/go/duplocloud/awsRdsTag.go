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

// `AwsRdsTag` manages an AWS RDS tag in Duplo.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud"
//	"github.com/pulumi/pulumi-random/sdk/go/random"
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
//			// Generate a random password.
//			mypassword, err := random.NewPassword(ctx, "mypassword", &random.PasswordArgs{
//				Length:  16,
//				Special: false,
//			})
//			if err != nil {
//				return err
//			}
//			// Create an RDS instance.
//			mydb, err := duplocloud.NewRdsInstance(ctx, "mydb", &duplocloud.RdsInstanceArgs{
//				TenantId:       myapp.TenantId,
//				Name:           pulumi.String("mydb"),
//				Engine:         pulumi.Int(1),
//				EngineVersion:  pulumi.String("12.5"),
//				Size:           pulumi.String("db.t3.medium"),
//				MasterUsername: pulumi.String("myuser"),
//				MasterPassword: mypassword.Result,
//				EncryptStorage: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			// Create RDS Tag for type "instance".
//			_, err = duplocloud.NewAwsRdsTag(ctx, "tag", &duplocloud.AwsRdsTagArgs{
//				TenantId:     myapp.TenantId,
//				ResourceType: pulumi.String("instance"),
//				ResourceId:   mydb.Identifier,
//				Key:          pulumi.String("CreatedBy"),
//				Value:        pulumi.String("DuploCloud"),
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
// Example: Importing an existing RDS Tag.
//
//   - *TENANT_ID* is the tenant GUID.
//
//   - *RESOURCE_TYPE* The type of the RDS resource, Valid vaues are - "cluster" and "instance".
//
//   - *RESOURCE_ID* The RDS identifier.
//
//   - *TAG_KEY* The tag name.
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/awsRdsTag:AwsRdsTag tag1 *TENANT_ID*/*RESOURCE_TYPE*/*RESOURCE_ID*/*TAG_KEY*
// ```
type AwsRdsTag struct {
	pulumi.CustomResourceState

	// The tag name.
	Key pulumi.StringOutput `pulumi:"key"`
	// The ID of the RDS resource to manage the tag for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The type of the RDS resource to manage the tag for. Valid values are `cluster` and `instance`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The GUID of the tenant that the RDS tag will be created in.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The value of the tag.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewAwsRdsTag registers a new resource with the given unique name, arguments, and options.
func NewAwsRdsTag(ctx *pulumi.Context,
	name string, args *AwsRdsTagArgs, opts ...pulumi.ResourceOption) (*AwsRdsTag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AwsRdsTag
	err := ctx.RegisterResource("duplocloud:index/awsRdsTag:AwsRdsTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsRdsTag gets an existing AwsRdsTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsRdsTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsRdsTagState, opts ...pulumi.ResourceOption) (*AwsRdsTag, error) {
	var resource AwsRdsTag
	err := ctx.ReadResource("duplocloud:index/awsRdsTag:AwsRdsTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsRdsTag resources.
type awsRdsTagState struct {
	// The tag name.
	Key *string `pulumi:"key"`
	// The ID of the RDS resource to manage the tag for.
	ResourceId *string `pulumi:"resourceId"`
	// The type of the RDS resource to manage the tag for. Valid values are `cluster` and `instance`.
	ResourceType *string `pulumi:"resourceType"`
	// The GUID of the tenant that the RDS tag will be created in.
	TenantId *string `pulumi:"tenantId"`
	// The value of the tag.
	Value *string `pulumi:"value"`
}

type AwsRdsTagState struct {
	// The tag name.
	Key pulumi.StringPtrInput
	// The ID of the RDS resource to manage the tag for.
	ResourceId pulumi.StringPtrInput
	// The type of the RDS resource to manage the tag for. Valid values are `cluster` and `instance`.
	ResourceType pulumi.StringPtrInput
	// The GUID of the tenant that the RDS tag will be created in.
	TenantId pulumi.StringPtrInput
	// The value of the tag.
	Value pulumi.StringPtrInput
}

func (AwsRdsTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsRdsTagState)(nil)).Elem()
}

type awsRdsTagArgs struct {
	// The tag name.
	Key string `pulumi:"key"`
	// The ID of the RDS resource to manage the tag for.
	ResourceId string `pulumi:"resourceId"`
	// The type of the RDS resource to manage the tag for. Valid values are `cluster` and `instance`.
	ResourceType string `pulumi:"resourceType"`
	// The GUID of the tenant that the RDS tag will be created in.
	TenantId string `pulumi:"tenantId"`
	// The value of the tag.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a AwsRdsTag resource.
type AwsRdsTagArgs struct {
	// The tag name.
	Key pulumi.StringInput
	// The ID of the RDS resource to manage the tag for.
	ResourceId pulumi.StringInput
	// The type of the RDS resource to manage the tag for. Valid values are `cluster` and `instance`.
	ResourceType pulumi.StringInput
	// The GUID of the tenant that the RDS tag will be created in.
	TenantId pulumi.StringInput
	// The value of the tag.
	Value pulumi.StringInput
}

func (AwsRdsTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsRdsTagArgs)(nil)).Elem()
}

type AwsRdsTagInput interface {
	pulumi.Input

	ToAwsRdsTagOutput() AwsRdsTagOutput
	ToAwsRdsTagOutputWithContext(ctx context.Context) AwsRdsTagOutput
}

func (*AwsRdsTag) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsRdsTag)(nil)).Elem()
}

func (i *AwsRdsTag) ToAwsRdsTagOutput() AwsRdsTagOutput {
	return i.ToAwsRdsTagOutputWithContext(context.Background())
}

func (i *AwsRdsTag) ToAwsRdsTagOutputWithContext(ctx context.Context) AwsRdsTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsRdsTagOutput)
}

// AwsRdsTagArrayInput is an input type that accepts AwsRdsTagArray and AwsRdsTagArrayOutput values.
// You can construct a concrete instance of `AwsRdsTagArrayInput` via:
//
//	AwsRdsTagArray{ AwsRdsTagArgs{...} }
type AwsRdsTagArrayInput interface {
	pulumi.Input

	ToAwsRdsTagArrayOutput() AwsRdsTagArrayOutput
	ToAwsRdsTagArrayOutputWithContext(context.Context) AwsRdsTagArrayOutput
}

type AwsRdsTagArray []AwsRdsTagInput

func (AwsRdsTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsRdsTag)(nil)).Elem()
}

func (i AwsRdsTagArray) ToAwsRdsTagArrayOutput() AwsRdsTagArrayOutput {
	return i.ToAwsRdsTagArrayOutputWithContext(context.Background())
}

func (i AwsRdsTagArray) ToAwsRdsTagArrayOutputWithContext(ctx context.Context) AwsRdsTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsRdsTagArrayOutput)
}

// AwsRdsTagMapInput is an input type that accepts AwsRdsTagMap and AwsRdsTagMapOutput values.
// You can construct a concrete instance of `AwsRdsTagMapInput` via:
//
//	AwsRdsTagMap{ "key": AwsRdsTagArgs{...} }
type AwsRdsTagMapInput interface {
	pulumi.Input

	ToAwsRdsTagMapOutput() AwsRdsTagMapOutput
	ToAwsRdsTagMapOutputWithContext(context.Context) AwsRdsTagMapOutput
}

type AwsRdsTagMap map[string]AwsRdsTagInput

func (AwsRdsTagMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsRdsTag)(nil)).Elem()
}

func (i AwsRdsTagMap) ToAwsRdsTagMapOutput() AwsRdsTagMapOutput {
	return i.ToAwsRdsTagMapOutputWithContext(context.Background())
}

func (i AwsRdsTagMap) ToAwsRdsTagMapOutputWithContext(ctx context.Context) AwsRdsTagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsRdsTagMapOutput)
}

type AwsRdsTagOutput struct{ *pulumi.OutputState }

func (AwsRdsTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsRdsTag)(nil)).Elem()
}

func (o AwsRdsTagOutput) ToAwsRdsTagOutput() AwsRdsTagOutput {
	return o
}

func (o AwsRdsTagOutput) ToAwsRdsTagOutputWithContext(ctx context.Context) AwsRdsTagOutput {
	return o
}

// The tag name.
func (o AwsRdsTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsRdsTag) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The ID of the RDS resource to manage the tag for.
func (o AwsRdsTagOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsRdsTag) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The type of the RDS resource to manage the tag for. Valid values are `cluster` and `instance`.
func (o AwsRdsTagOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsRdsTag) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// The GUID of the tenant that the RDS tag will be created in.
func (o AwsRdsTagOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsRdsTag) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The value of the tag.
func (o AwsRdsTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsRdsTag) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type AwsRdsTagArrayOutput struct{ *pulumi.OutputState }

func (AwsRdsTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsRdsTag)(nil)).Elem()
}

func (o AwsRdsTagArrayOutput) ToAwsRdsTagArrayOutput() AwsRdsTagArrayOutput {
	return o
}

func (o AwsRdsTagArrayOutput) ToAwsRdsTagArrayOutputWithContext(ctx context.Context) AwsRdsTagArrayOutput {
	return o
}

func (o AwsRdsTagArrayOutput) Index(i pulumi.IntInput) AwsRdsTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AwsRdsTag {
		return vs[0].([]*AwsRdsTag)[vs[1].(int)]
	}).(AwsRdsTagOutput)
}

type AwsRdsTagMapOutput struct{ *pulumi.OutputState }

func (AwsRdsTagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsRdsTag)(nil)).Elem()
}

func (o AwsRdsTagMapOutput) ToAwsRdsTagMapOutput() AwsRdsTagMapOutput {
	return o
}

func (o AwsRdsTagMapOutput) ToAwsRdsTagMapOutputWithContext(ctx context.Context) AwsRdsTagMapOutput {
	return o
}

func (o AwsRdsTagMapOutput) MapIndex(k pulumi.StringInput) AwsRdsTagOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AwsRdsTag {
		return vs[0].(map[string]*AwsRdsTag)[vs[1].(string)]
	}).(AwsRdsTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsRdsTagInput)(nil)).Elem(), &AwsRdsTag{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsRdsTagArrayInput)(nil)).Elem(), AwsRdsTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsRdsTagMapInput)(nil)).Elem(), AwsRdsTagMap{})
	pulumi.RegisterOutputType(AwsRdsTagOutput{})
	pulumi.RegisterOutputType(AwsRdsTagArrayOutput{})
	pulumi.RegisterOutputType(AwsRdsTagMapOutput{})
}
