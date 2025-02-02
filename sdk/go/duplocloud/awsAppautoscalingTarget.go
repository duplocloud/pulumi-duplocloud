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

// `AwsAppautoscalingTarget` manages an aws autoscaling target in Duplo.
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
//			_, err := duplocloud.NewTenant(ctx, "duplo-app", &duplocloud.TenantArgs{
//				AccountName: pulumi.String("duplo-app"),
//				PlanId:      pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			// ECS Service Autoscaling
//			_, err = duplocloud.NewAwsAppautoscalingTarget(ctx, "asg-target", &duplocloud.AwsAppautoscalingTargetArgs{
//				TenantId:          duplo_app.TenantId,
//				MaxCapacity:       pulumi.Int(4),
//				MinCapacity:       pulumi.Int(2),
//				ResourceId:        pulumi.String("duploservices-duplo-app-ecs-service"),
//				ScalableDimension: pulumi.String("ecs:service:DesiredCount"),
//				ServiceNamespace:  pulumi.String("ecs"),
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
// Example: Importing an existing Autoscaling Target
//
//   - *TENANT_ID* is the tenant GUID
//
//   - *SERVICE_NAMESPACE* The AWS service namespace of the scalable target
//
//   - *SCALABLE_DIMENSION*  The scalable dimension of the scalable target.
//
//   - *RESOURCE_ID* is the duploservices-<account_name>-<resource_name>
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/awsAppautoscalingTarget:AwsAppautoscalingTarget asgTarget *TENANT_ID*/*SERVICE_NAMESPACE*/*SCALABLE_DIMENSION*/*RESOURCE_ID*
// ```
type AwsAppautoscalingTarget struct {
	pulumi.CustomResourceState

	// The resource type and unique identifier string for the resource associated with the scaling policy.
	FullResourceId pulumi.StringOutput `pulumi:"fullResourceId"`
	// The max capacity of the scalable target.
	MaxCapacity pulumi.IntOutput `pulumi:"maxCapacity"`
	// The min capacity of the scalable target.
	MinCapacity pulumi.IntOutput `pulumi:"minCapacity"`
	// Resource name associated with the scaling policy.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The scalable dimension of the scalable target.
	ScalableDimension pulumi.StringOutput `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target.
	ServiceNamespace pulumi.StringOutput `pulumi:"serviceNamespace"`
	// The GUID of the tenant that the aws autoscaling target will be created in.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewAwsAppautoscalingTarget registers a new resource with the given unique name, arguments, and options.
func NewAwsAppautoscalingTarget(ctx *pulumi.Context,
	name string, args *AwsAppautoscalingTargetArgs, opts ...pulumi.ResourceOption) (*AwsAppautoscalingTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxCapacity == nil {
		return nil, errors.New("invalid value for required argument 'MaxCapacity'")
	}
	if args.MinCapacity == nil {
		return nil, errors.New("invalid value for required argument 'MinCapacity'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ScalableDimension == nil {
		return nil, errors.New("invalid value for required argument 'ScalableDimension'")
	}
	if args.ServiceNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ServiceNamespace'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AwsAppautoscalingTarget
	err := ctx.RegisterResource("duplocloud:index/awsAppautoscalingTarget:AwsAppautoscalingTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsAppautoscalingTarget gets an existing AwsAppautoscalingTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsAppautoscalingTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsAppautoscalingTargetState, opts ...pulumi.ResourceOption) (*AwsAppautoscalingTarget, error) {
	var resource AwsAppautoscalingTarget
	err := ctx.ReadResource("duplocloud:index/awsAppautoscalingTarget:AwsAppautoscalingTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsAppautoscalingTarget resources.
type awsAppautoscalingTargetState struct {
	// The resource type and unique identifier string for the resource associated with the scaling policy.
	FullResourceId *string `pulumi:"fullResourceId"`
	// The max capacity of the scalable target.
	MaxCapacity *int `pulumi:"maxCapacity"`
	// The min capacity of the scalable target.
	MinCapacity *int `pulumi:"minCapacity"`
	// Resource name associated with the scaling policy.
	ResourceId *string `pulumi:"resourceId"`
	// The ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf.
	RoleArn *string `pulumi:"roleArn"`
	// The scalable dimension of the scalable target.
	ScalableDimension *string `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target.
	ServiceNamespace *string `pulumi:"serviceNamespace"`
	// The GUID of the tenant that the aws autoscaling target will be created in.
	TenantId *string `pulumi:"tenantId"`
}

type AwsAppautoscalingTargetState struct {
	// The resource type and unique identifier string for the resource associated with the scaling policy.
	FullResourceId pulumi.StringPtrInput
	// The max capacity of the scalable target.
	MaxCapacity pulumi.IntPtrInput
	// The min capacity of the scalable target.
	MinCapacity pulumi.IntPtrInput
	// Resource name associated with the scaling policy.
	ResourceId pulumi.StringPtrInput
	// The ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf.
	RoleArn pulumi.StringPtrInput
	// The scalable dimension of the scalable target.
	ScalableDimension pulumi.StringPtrInput
	// The AWS service namespace of the scalable target.
	ServiceNamespace pulumi.StringPtrInput
	// The GUID of the tenant that the aws autoscaling target will be created in.
	TenantId pulumi.StringPtrInput
}

func (AwsAppautoscalingTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsAppautoscalingTargetState)(nil)).Elem()
}

type awsAppautoscalingTargetArgs struct {
	// The max capacity of the scalable target.
	MaxCapacity int `pulumi:"maxCapacity"`
	// The min capacity of the scalable target.
	MinCapacity int `pulumi:"minCapacity"`
	// Resource name associated with the scaling policy.
	ResourceId string `pulumi:"resourceId"`
	// The ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf.
	RoleArn *string `pulumi:"roleArn"`
	// The scalable dimension of the scalable target.
	ScalableDimension string `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target.
	ServiceNamespace string `pulumi:"serviceNamespace"`
	// The GUID of the tenant that the aws autoscaling target will be created in.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a AwsAppautoscalingTarget resource.
type AwsAppautoscalingTargetArgs struct {
	// The max capacity of the scalable target.
	MaxCapacity pulumi.IntInput
	// The min capacity of the scalable target.
	MinCapacity pulumi.IntInput
	// Resource name associated with the scaling policy.
	ResourceId pulumi.StringInput
	// The ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf.
	RoleArn pulumi.StringPtrInput
	// The scalable dimension of the scalable target.
	ScalableDimension pulumi.StringInput
	// The AWS service namespace of the scalable target.
	ServiceNamespace pulumi.StringInput
	// The GUID of the tenant that the aws autoscaling target will be created in.
	TenantId pulumi.StringInput
}

func (AwsAppautoscalingTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsAppautoscalingTargetArgs)(nil)).Elem()
}

type AwsAppautoscalingTargetInput interface {
	pulumi.Input

	ToAwsAppautoscalingTargetOutput() AwsAppautoscalingTargetOutput
	ToAwsAppautoscalingTargetOutputWithContext(ctx context.Context) AwsAppautoscalingTargetOutput
}

func (*AwsAppautoscalingTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsAppautoscalingTarget)(nil)).Elem()
}

func (i *AwsAppautoscalingTarget) ToAwsAppautoscalingTargetOutput() AwsAppautoscalingTargetOutput {
	return i.ToAwsAppautoscalingTargetOutputWithContext(context.Background())
}

func (i *AwsAppautoscalingTarget) ToAwsAppautoscalingTargetOutputWithContext(ctx context.Context) AwsAppautoscalingTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsAppautoscalingTargetOutput)
}

// AwsAppautoscalingTargetArrayInput is an input type that accepts AwsAppautoscalingTargetArray and AwsAppautoscalingTargetArrayOutput values.
// You can construct a concrete instance of `AwsAppautoscalingTargetArrayInput` via:
//
//	AwsAppautoscalingTargetArray{ AwsAppautoscalingTargetArgs{...} }
type AwsAppautoscalingTargetArrayInput interface {
	pulumi.Input

	ToAwsAppautoscalingTargetArrayOutput() AwsAppautoscalingTargetArrayOutput
	ToAwsAppautoscalingTargetArrayOutputWithContext(context.Context) AwsAppautoscalingTargetArrayOutput
}

type AwsAppautoscalingTargetArray []AwsAppautoscalingTargetInput

func (AwsAppautoscalingTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsAppautoscalingTarget)(nil)).Elem()
}

func (i AwsAppautoscalingTargetArray) ToAwsAppautoscalingTargetArrayOutput() AwsAppautoscalingTargetArrayOutput {
	return i.ToAwsAppautoscalingTargetArrayOutputWithContext(context.Background())
}

func (i AwsAppautoscalingTargetArray) ToAwsAppautoscalingTargetArrayOutputWithContext(ctx context.Context) AwsAppautoscalingTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsAppautoscalingTargetArrayOutput)
}

// AwsAppautoscalingTargetMapInput is an input type that accepts AwsAppautoscalingTargetMap and AwsAppautoscalingTargetMapOutput values.
// You can construct a concrete instance of `AwsAppautoscalingTargetMapInput` via:
//
//	AwsAppautoscalingTargetMap{ "key": AwsAppautoscalingTargetArgs{...} }
type AwsAppautoscalingTargetMapInput interface {
	pulumi.Input

	ToAwsAppautoscalingTargetMapOutput() AwsAppautoscalingTargetMapOutput
	ToAwsAppautoscalingTargetMapOutputWithContext(context.Context) AwsAppautoscalingTargetMapOutput
}

type AwsAppautoscalingTargetMap map[string]AwsAppautoscalingTargetInput

func (AwsAppautoscalingTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsAppautoscalingTarget)(nil)).Elem()
}

func (i AwsAppautoscalingTargetMap) ToAwsAppautoscalingTargetMapOutput() AwsAppautoscalingTargetMapOutput {
	return i.ToAwsAppautoscalingTargetMapOutputWithContext(context.Background())
}

func (i AwsAppautoscalingTargetMap) ToAwsAppautoscalingTargetMapOutputWithContext(ctx context.Context) AwsAppautoscalingTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsAppautoscalingTargetMapOutput)
}

type AwsAppautoscalingTargetOutput struct{ *pulumi.OutputState }

func (AwsAppautoscalingTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsAppautoscalingTarget)(nil)).Elem()
}

func (o AwsAppautoscalingTargetOutput) ToAwsAppautoscalingTargetOutput() AwsAppautoscalingTargetOutput {
	return o
}

func (o AwsAppautoscalingTargetOutput) ToAwsAppautoscalingTargetOutputWithContext(ctx context.Context) AwsAppautoscalingTargetOutput {
	return o
}

// The resource type and unique identifier string for the resource associated with the scaling policy.
func (o AwsAppautoscalingTargetOutput) FullResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAppautoscalingTarget) pulumi.StringOutput { return v.FullResourceId }).(pulumi.StringOutput)
}

// The max capacity of the scalable target.
func (o AwsAppautoscalingTargetOutput) MaxCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *AwsAppautoscalingTarget) pulumi.IntOutput { return v.MaxCapacity }).(pulumi.IntOutput)
}

// The min capacity of the scalable target.
func (o AwsAppautoscalingTargetOutput) MinCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *AwsAppautoscalingTarget) pulumi.IntOutput { return v.MinCapacity }).(pulumi.IntOutput)
}

// Resource name associated with the scaling policy.
func (o AwsAppautoscalingTargetOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAppautoscalingTarget) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf.
func (o AwsAppautoscalingTargetOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAppautoscalingTarget) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// The scalable dimension of the scalable target.
func (o AwsAppautoscalingTargetOutput) ScalableDimension() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAppautoscalingTarget) pulumi.StringOutput { return v.ScalableDimension }).(pulumi.StringOutput)
}

// The AWS service namespace of the scalable target.
func (o AwsAppautoscalingTargetOutput) ServiceNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAppautoscalingTarget) pulumi.StringOutput { return v.ServiceNamespace }).(pulumi.StringOutput)
}

// The GUID of the tenant that the aws autoscaling target will be created in.
func (o AwsAppautoscalingTargetOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *AwsAppautoscalingTarget) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type AwsAppautoscalingTargetArrayOutput struct{ *pulumi.OutputState }

func (AwsAppautoscalingTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsAppautoscalingTarget)(nil)).Elem()
}

func (o AwsAppautoscalingTargetArrayOutput) ToAwsAppautoscalingTargetArrayOutput() AwsAppautoscalingTargetArrayOutput {
	return o
}

func (o AwsAppautoscalingTargetArrayOutput) ToAwsAppautoscalingTargetArrayOutputWithContext(ctx context.Context) AwsAppautoscalingTargetArrayOutput {
	return o
}

func (o AwsAppautoscalingTargetArrayOutput) Index(i pulumi.IntInput) AwsAppautoscalingTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AwsAppautoscalingTarget {
		return vs[0].([]*AwsAppautoscalingTarget)[vs[1].(int)]
	}).(AwsAppautoscalingTargetOutput)
}

type AwsAppautoscalingTargetMapOutput struct{ *pulumi.OutputState }

func (AwsAppautoscalingTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsAppautoscalingTarget)(nil)).Elem()
}

func (o AwsAppautoscalingTargetMapOutput) ToAwsAppautoscalingTargetMapOutput() AwsAppautoscalingTargetMapOutput {
	return o
}

func (o AwsAppautoscalingTargetMapOutput) ToAwsAppautoscalingTargetMapOutputWithContext(ctx context.Context) AwsAppautoscalingTargetMapOutput {
	return o
}

func (o AwsAppautoscalingTargetMapOutput) MapIndex(k pulumi.StringInput) AwsAppautoscalingTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AwsAppautoscalingTarget {
		return vs[0].(map[string]*AwsAppautoscalingTarget)[vs[1].(string)]
	}).(AwsAppautoscalingTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsAppautoscalingTargetInput)(nil)).Elem(), &AwsAppautoscalingTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsAppautoscalingTargetArrayInput)(nil)).Elem(), AwsAppautoscalingTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsAppautoscalingTargetMapInput)(nil)).Elem(), AwsAppautoscalingTargetMap{})
	pulumi.RegisterOutputType(AwsAppautoscalingTargetOutput{})
	pulumi.RegisterOutputType(AwsAppautoscalingTargetArrayOutput{})
	pulumi.RegisterOutputType(AwsAppautoscalingTargetMapOutput{})
}
