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

// `K8ConfigMap` manages a kubernetes configmap in a Duplo tenant.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
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
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"foo": "bar2",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = duplocloud.NewK8ConfigMap(ctx, "myapp", &duplocloud.K8ConfigMapArgs{
//				TenantId: myapp.TenantId,
//				Name:     pulumi.String("myconfigmap"),
//				Data:     pulumi.String(json0),
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
// Example: Importing an existing kubernetes config map
//
//   - *TENANT_ID* is the tenant GUID
//
//   - *NAME* is the config map name
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/k8ConfigMap:K8ConfigMap myapp v2/subscriptions/*TENANT_ID*/K8ConfigMapApiV2/*NAME*
// ```
type K8ConfigMap struct {
	pulumi.CustomResourceState

	// A JSON encoded string representing the configmap data. You can use the `jsonencode()` function to build this from JSON.
	Data pulumi.StringOutput `pulumi:"data"`
	// A JSON encoded string representing the configmap metadata. You can use the `jsondecode()` function to parse this, if needed.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The name of the configmap.
	Name pulumi.StringOutput `pulumi:"name"`
	// The GUID of the tenant that the configmap will be created in.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewK8ConfigMap registers a new resource with the given unique name, arguments, and options.
func NewK8ConfigMap(ctx *pulumi.Context,
	name string, args *K8ConfigMapArgs, opts ...pulumi.ResourceOption) (*K8ConfigMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource K8ConfigMap
	err := ctx.RegisterResource("duplocloud:index/k8ConfigMap:K8ConfigMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetK8ConfigMap gets an existing K8ConfigMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetK8ConfigMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *K8ConfigMapState, opts ...pulumi.ResourceOption) (*K8ConfigMap, error) {
	var resource K8ConfigMap
	err := ctx.ReadResource("duplocloud:index/k8ConfigMap:K8ConfigMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering K8ConfigMap resources.
type k8configMapState struct {
	// A JSON encoded string representing the configmap data. You can use the `jsonencode()` function to build this from JSON.
	Data *string `pulumi:"data"`
	// A JSON encoded string representing the configmap metadata. You can use the `jsondecode()` function to parse this, if needed.
	Metadata *string `pulumi:"metadata"`
	// The name of the configmap.
	Name *string `pulumi:"name"`
	// The GUID of the tenant that the configmap will be created in.
	TenantId *string `pulumi:"tenantId"`
}

type K8ConfigMapState struct {
	// A JSON encoded string representing the configmap data. You can use the `jsonencode()` function to build this from JSON.
	Data pulumi.StringPtrInput
	// A JSON encoded string representing the configmap metadata. You can use the `jsondecode()` function to parse this, if needed.
	Metadata pulumi.StringPtrInput
	// The name of the configmap.
	Name pulumi.StringPtrInput
	// The GUID of the tenant that the configmap will be created in.
	TenantId pulumi.StringPtrInput
}

func (K8ConfigMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*k8configMapState)(nil)).Elem()
}

type k8configMapArgs struct {
	// A JSON encoded string representing the configmap data. You can use the `jsonencode()` function to build this from JSON.
	Data string `pulumi:"data"`
	// The name of the configmap.
	Name *string `pulumi:"name"`
	// The GUID of the tenant that the configmap will be created in.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a K8ConfigMap resource.
type K8ConfigMapArgs struct {
	// A JSON encoded string representing the configmap data. You can use the `jsonencode()` function to build this from JSON.
	Data pulumi.StringInput
	// The name of the configmap.
	Name pulumi.StringPtrInput
	// The GUID of the tenant that the configmap will be created in.
	TenantId pulumi.StringInput
}

func (K8ConfigMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*k8configMapArgs)(nil)).Elem()
}

type K8ConfigMapInput interface {
	pulumi.Input

	ToK8ConfigMapOutput() K8ConfigMapOutput
	ToK8ConfigMapOutputWithContext(ctx context.Context) K8ConfigMapOutput
}

func (*K8ConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((**K8ConfigMap)(nil)).Elem()
}

func (i *K8ConfigMap) ToK8ConfigMapOutput() K8ConfigMapOutput {
	return i.ToK8ConfigMapOutputWithContext(context.Background())
}

func (i *K8ConfigMap) ToK8ConfigMapOutputWithContext(ctx context.Context) K8ConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8ConfigMapOutput)
}

// K8ConfigMapArrayInput is an input type that accepts K8ConfigMapArray and K8ConfigMapArrayOutput values.
// You can construct a concrete instance of `K8ConfigMapArrayInput` via:
//
//	K8ConfigMapArray{ K8ConfigMapArgs{...} }
type K8ConfigMapArrayInput interface {
	pulumi.Input

	ToK8ConfigMapArrayOutput() K8ConfigMapArrayOutput
	ToK8ConfigMapArrayOutputWithContext(context.Context) K8ConfigMapArrayOutput
}

type K8ConfigMapArray []K8ConfigMapInput

func (K8ConfigMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*K8ConfigMap)(nil)).Elem()
}

func (i K8ConfigMapArray) ToK8ConfigMapArrayOutput() K8ConfigMapArrayOutput {
	return i.ToK8ConfigMapArrayOutputWithContext(context.Background())
}

func (i K8ConfigMapArray) ToK8ConfigMapArrayOutputWithContext(ctx context.Context) K8ConfigMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8ConfigMapArrayOutput)
}

// K8ConfigMapMapInput is an input type that accepts K8ConfigMapMap and K8ConfigMapMapOutput values.
// You can construct a concrete instance of `K8ConfigMapMapInput` via:
//
//	K8ConfigMapMap{ "key": K8ConfigMapArgs{...} }
type K8ConfigMapMapInput interface {
	pulumi.Input

	ToK8ConfigMapMapOutput() K8ConfigMapMapOutput
	ToK8ConfigMapMapOutputWithContext(context.Context) K8ConfigMapMapOutput
}

type K8ConfigMapMap map[string]K8ConfigMapInput

func (K8ConfigMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*K8ConfigMap)(nil)).Elem()
}

func (i K8ConfigMapMap) ToK8ConfigMapMapOutput() K8ConfigMapMapOutput {
	return i.ToK8ConfigMapMapOutputWithContext(context.Background())
}

func (i K8ConfigMapMap) ToK8ConfigMapMapOutputWithContext(ctx context.Context) K8ConfigMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8ConfigMapMapOutput)
}

type K8ConfigMapOutput struct{ *pulumi.OutputState }

func (K8ConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K8ConfigMap)(nil)).Elem()
}

func (o K8ConfigMapOutput) ToK8ConfigMapOutput() K8ConfigMapOutput {
	return o
}

func (o K8ConfigMapOutput) ToK8ConfigMapOutputWithContext(ctx context.Context) K8ConfigMapOutput {
	return o
}

// A JSON encoded string representing the configmap data. You can use the `jsonencode()` function to build this from JSON.
func (o K8ConfigMapOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v *K8ConfigMap) pulumi.StringOutput { return v.Data }).(pulumi.StringOutput)
}

// A JSON encoded string representing the configmap metadata. You can use the `jsondecode()` function to parse this, if needed.
func (o K8ConfigMapOutput) Metadata() pulumi.StringOutput {
	return o.ApplyT(func(v *K8ConfigMap) pulumi.StringOutput { return v.Metadata }).(pulumi.StringOutput)
}

// The name of the configmap.
func (o K8ConfigMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *K8ConfigMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The GUID of the tenant that the configmap will be created in.
func (o K8ConfigMapOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *K8ConfigMap) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type K8ConfigMapArrayOutput struct{ *pulumi.OutputState }

func (K8ConfigMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*K8ConfigMap)(nil)).Elem()
}

func (o K8ConfigMapArrayOutput) ToK8ConfigMapArrayOutput() K8ConfigMapArrayOutput {
	return o
}

func (o K8ConfigMapArrayOutput) ToK8ConfigMapArrayOutputWithContext(ctx context.Context) K8ConfigMapArrayOutput {
	return o
}

func (o K8ConfigMapArrayOutput) Index(i pulumi.IntInput) K8ConfigMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *K8ConfigMap {
		return vs[0].([]*K8ConfigMap)[vs[1].(int)]
	}).(K8ConfigMapOutput)
}

type K8ConfigMapMapOutput struct{ *pulumi.OutputState }

func (K8ConfigMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*K8ConfigMap)(nil)).Elem()
}

func (o K8ConfigMapMapOutput) ToK8ConfigMapMapOutput() K8ConfigMapMapOutput {
	return o
}

func (o K8ConfigMapMapOutput) ToK8ConfigMapMapOutputWithContext(ctx context.Context) K8ConfigMapMapOutput {
	return o
}

func (o K8ConfigMapMapOutput) MapIndex(k pulumi.StringInput) K8ConfigMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *K8ConfigMap {
		return vs[0].(map[string]*K8ConfigMap)[vs[1].(string)]
	}).(K8ConfigMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*K8ConfigMapInput)(nil)).Elem(), &K8ConfigMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*K8ConfigMapArrayInput)(nil)).Elem(), K8ConfigMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*K8ConfigMapMapInput)(nil)).Elem(), K8ConfigMapMap{})
	pulumi.RegisterOutputType(K8ConfigMapOutput{})
	pulumi.RegisterOutputType(K8ConfigMapArrayOutput{})
	pulumi.RegisterOutputType(K8ConfigMapMapOutput{})
}
