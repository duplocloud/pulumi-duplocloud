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

// `K8StorageClass` manages a kubernetes storage class in a Duplo tenant.
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
//			tenantId := "3a0b2ea5-7403-4765-ad6e-8771ca8fa0fd"
//			_, err := duplocloud.NewK8StorageClass(ctx, "sc", &duplocloud.K8StorageClassArgs{
//				TenantId:             pulumi.String(tenantId),
//				Name:                 pulumi.String("sc"),
//				StorageProvisioner:   pulumi.String("efs.csi.aws.com"),
//				ReclaimPolicy:        pulumi.String("Delete"),
//				VolumeBindingMode:    pulumi.String("Immediate"),
//				AllowVolumeExpansion: pulumi.Bool(false),
//				Parameters: pulumi.StringMap{
//					"fileSystemId":     pulumi.String("fs-0d2f79aca4712c6e8"),
//					"basePath":         pulumi.String("/dynamic_provisioning"),
//					"directoryPerms":   pulumi.String("700"),
//					"gidRangeStart":    pulumi.String("1000"),
//					"gidRangeEnd":      pulumi.String("2000"),
//					"provisioningMode": pulumi.String("efs-ap"),
//				},
//				Annotations: pulumi.StringMap{
//					"a1": pulumi.String("v1"),
//					"a2": pulumi.String("v2"),
//				},
//				Labels: pulumi.StringMap{
//					"l1": pulumi.String("v1"),
//					"l2": pulumi.String("v2"),
//				},
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
// Example: Importing an existing kubernetes storage class
//
//   - *TENANT_ID* is the tenant GUID
//
//   - *FULL_NAME* is the Duplo provided storage class name
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/k8StorageClass:K8StorageClass sc v3/subscriptions/*TENANT_ID*/k8s/storageclass/*FULL_NAME*
// ```
type K8StorageClass struct {
	pulumi.CustomResourceState

	// Indicates whether the storage class allow volume expand Defaults to `false`.
	AllowVolumeExpansion pulumi.BoolPtrOutput `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned.
	AllowedTopologies K8StorageClassAllowedTopologiesPtrOutput `pulumi:"allowedTopologies"`
	// An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// The full name of the storage class.
	Fullname pulumi.StringOutput `pulumi:"fullname"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the service.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the storage class.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters for the provisioner that should create volumes of this storage class
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Indicates the type of the reclaim policy Defaults to `Delete`.
	ReclaimPolicy pulumi.StringPtrOutput `pulumi:"reclaimPolicy"`
	// Indicates the type of the provisioner
	StorageProvisioner pulumi.StringOutput `pulumi:"storageProvisioner"`
	// The GUID of the tenant that the storage class will be created in.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Indicates when volume binding and dynamic provisioning should occur Defaults to `Immediate`.
	VolumeBindingMode pulumi.StringPtrOutput `pulumi:"volumeBindingMode"`
}

// NewK8StorageClass registers a new resource with the given unique name, arguments, and options.
func NewK8StorageClass(ctx *pulumi.Context,
	name string, args *K8StorageClassArgs, opts ...pulumi.ResourceOption) (*K8StorageClass, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageProvisioner == nil {
		return nil, errors.New("invalid value for required argument 'StorageProvisioner'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource K8StorageClass
	err := ctx.RegisterResource("duplocloud:index/k8StorageClass:K8StorageClass", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetK8StorageClass gets an existing K8StorageClass resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetK8StorageClass(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *K8StorageClassState, opts ...pulumi.ResourceOption) (*K8StorageClass, error) {
	var resource K8StorageClass
	err := ctx.ReadResource("duplocloud:index/k8StorageClass:K8StorageClass", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering K8StorageClass resources.
type k8storageClassState struct {
	// Indicates whether the storage class allow volume expand Defaults to `false`.
	AllowVolumeExpansion *bool `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned.
	AllowedTopologies *K8StorageClassAllowedTopologies `pulumi:"allowedTopologies"`
	// An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
	Annotations map[string]string `pulumi:"annotations"`
	// The full name of the storage class.
	Fullname *string `pulumi:"fullname"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the service.
	Labels map[string]string `pulumi:"labels"`
	// The name of the storage class.
	Name *string `pulumi:"name"`
	// The parameters for the provisioner that should create volumes of this storage class
	Parameters map[string]string `pulumi:"parameters"`
	// Indicates the type of the reclaim policy Defaults to `Delete`.
	ReclaimPolicy *string `pulumi:"reclaimPolicy"`
	// Indicates the type of the provisioner
	StorageProvisioner *string `pulumi:"storageProvisioner"`
	// The GUID of the tenant that the storage class will be created in.
	TenantId *string `pulumi:"tenantId"`
	// Indicates when volume binding and dynamic provisioning should occur Defaults to `Immediate`.
	VolumeBindingMode *string `pulumi:"volumeBindingMode"`
}

type K8StorageClassState struct {
	// Indicates whether the storage class allow volume expand Defaults to `false`.
	AllowVolumeExpansion pulumi.BoolPtrInput
	// Restrict the node topologies where volumes can be dynamically provisioned.
	AllowedTopologies K8StorageClassAllowedTopologiesPtrInput
	// An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
	Annotations pulumi.StringMapInput
	// The full name of the storage class.
	Fullname pulumi.StringPtrInput
	// Map of string keys and values that can be used to organize and categorize (scope and select) the service.
	Labels pulumi.StringMapInput
	// The name of the storage class.
	Name pulumi.StringPtrInput
	// The parameters for the provisioner that should create volumes of this storage class
	Parameters pulumi.StringMapInput
	// Indicates the type of the reclaim policy Defaults to `Delete`.
	ReclaimPolicy pulumi.StringPtrInput
	// Indicates the type of the provisioner
	StorageProvisioner pulumi.StringPtrInput
	// The GUID of the tenant that the storage class will be created in.
	TenantId pulumi.StringPtrInput
	// Indicates when volume binding and dynamic provisioning should occur Defaults to `Immediate`.
	VolumeBindingMode pulumi.StringPtrInput
}

func (K8StorageClassState) ElementType() reflect.Type {
	return reflect.TypeOf((*k8storageClassState)(nil)).Elem()
}

type k8storageClassArgs struct {
	// Indicates whether the storage class allow volume expand Defaults to `false`.
	AllowVolumeExpansion *bool `pulumi:"allowVolumeExpansion"`
	// Restrict the node topologies where volumes can be dynamically provisioned.
	AllowedTopologies *K8StorageClassAllowedTopologies `pulumi:"allowedTopologies"`
	// An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
	Annotations map[string]string `pulumi:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) the service.
	Labels map[string]string `pulumi:"labels"`
	// The name of the storage class.
	Name *string `pulumi:"name"`
	// The parameters for the provisioner that should create volumes of this storage class
	Parameters map[string]string `pulumi:"parameters"`
	// Indicates the type of the reclaim policy Defaults to `Delete`.
	ReclaimPolicy *string `pulumi:"reclaimPolicy"`
	// Indicates the type of the provisioner
	StorageProvisioner string `pulumi:"storageProvisioner"`
	// The GUID of the tenant that the storage class will be created in.
	TenantId string `pulumi:"tenantId"`
	// Indicates when volume binding and dynamic provisioning should occur Defaults to `Immediate`.
	VolumeBindingMode *string `pulumi:"volumeBindingMode"`
}

// The set of arguments for constructing a K8StorageClass resource.
type K8StorageClassArgs struct {
	// Indicates whether the storage class allow volume expand Defaults to `false`.
	AllowVolumeExpansion pulumi.BoolPtrInput
	// Restrict the node topologies where volumes can be dynamically provisioned.
	AllowedTopologies K8StorageClassAllowedTopologiesPtrInput
	// An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
	Annotations pulumi.StringMapInput
	// Map of string keys and values that can be used to organize and categorize (scope and select) the service.
	Labels pulumi.StringMapInput
	// The name of the storage class.
	Name pulumi.StringPtrInput
	// The parameters for the provisioner that should create volumes of this storage class
	Parameters pulumi.StringMapInput
	// Indicates the type of the reclaim policy Defaults to `Delete`.
	ReclaimPolicy pulumi.StringPtrInput
	// Indicates the type of the provisioner
	StorageProvisioner pulumi.StringInput
	// The GUID of the tenant that the storage class will be created in.
	TenantId pulumi.StringInput
	// Indicates when volume binding and dynamic provisioning should occur Defaults to `Immediate`.
	VolumeBindingMode pulumi.StringPtrInput
}

func (K8StorageClassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*k8storageClassArgs)(nil)).Elem()
}

type K8StorageClassInput interface {
	pulumi.Input

	ToK8StorageClassOutput() K8StorageClassOutput
	ToK8StorageClassOutputWithContext(ctx context.Context) K8StorageClassOutput
}

func (*K8StorageClass) ElementType() reflect.Type {
	return reflect.TypeOf((**K8StorageClass)(nil)).Elem()
}

func (i *K8StorageClass) ToK8StorageClassOutput() K8StorageClassOutput {
	return i.ToK8StorageClassOutputWithContext(context.Background())
}

func (i *K8StorageClass) ToK8StorageClassOutputWithContext(ctx context.Context) K8StorageClassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8StorageClassOutput)
}

// K8StorageClassArrayInput is an input type that accepts K8StorageClassArray and K8StorageClassArrayOutput values.
// You can construct a concrete instance of `K8StorageClassArrayInput` via:
//
//	K8StorageClassArray{ K8StorageClassArgs{...} }
type K8StorageClassArrayInput interface {
	pulumi.Input

	ToK8StorageClassArrayOutput() K8StorageClassArrayOutput
	ToK8StorageClassArrayOutputWithContext(context.Context) K8StorageClassArrayOutput
}

type K8StorageClassArray []K8StorageClassInput

func (K8StorageClassArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*K8StorageClass)(nil)).Elem()
}

func (i K8StorageClassArray) ToK8StorageClassArrayOutput() K8StorageClassArrayOutput {
	return i.ToK8StorageClassArrayOutputWithContext(context.Background())
}

func (i K8StorageClassArray) ToK8StorageClassArrayOutputWithContext(ctx context.Context) K8StorageClassArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8StorageClassArrayOutput)
}

// K8StorageClassMapInput is an input type that accepts K8StorageClassMap and K8StorageClassMapOutput values.
// You can construct a concrete instance of `K8StorageClassMapInput` via:
//
//	K8StorageClassMap{ "key": K8StorageClassArgs{...} }
type K8StorageClassMapInput interface {
	pulumi.Input

	ToK8StorageClassMapOutput() K8StorageClassMapOutput
	ToK8StorageClassMapOutputWithContext(context.Context) K8StorageClassMapOutput
}

type K8StorageClassMap map[string]K8StorageClassInput

func (K8StorageClassMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*K8StorageClass)(nil)).Elem()
}

func (i K8StorageClassMap) ToK8StorageClassMapOutput() K8StorageClassMapOutput {
	return i.ToK8StorageClassMapOutputWithContext(context.Background())
}

func (i K8StorageClassMap) ToK8StorageClassMapOutputWithContext(ctx context.Context) K8StorageClassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(K8StorageClassMapOutput)
}

type K8StorageClassOutput struct{ *pulumi.OutputState }

func (K8StorageClassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**K8StorageClass)(nil)).Elem()
}

func (o K8StorageClassOutput) ToK8StorageClassOutput() K8StorageClassOutput {
	return o
}

func (o K8StorageClassOutput) ToK8StorageClassOutputWithContext(ctx context.Context) K8StorageClassOutput {
	return o
}

// Indicates whether the storage class allow volume expand Defaults to `false`.
func (o K8StorageClassOutput) AllowVolumeExpansion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.BoolPtrOutput { return v.AllowVolumeExpansion }).(pulumi.BoolPtrOutput)
}

// Restrict the node topologies where volumes can be dynamically provisioned.
func (o K8StorageClassOutput) AllowedTopologies() K8StorageClassAllowedTopologiesPtrOutput {
	return o.ApplyT(func(v *K8StorageClass) K8StorageClassAllowedTopologiesPtrOutput { return v.AllowedTopologies }).(K8StorageClassAllowedTopologiesPtrOutput)
}

// An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
func (o K8StorageClassOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// The full name of the storage class.
func (o K8StorageClassOutput) Fullname() pulumi.StringOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.StringOutput { return v.Fullname }).(pulumi.StringOutput)
}

// Map of string keys and values that can be used to organize and categorize (scope and select) the service.
func (o K8StorageClassOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the storage class.
func (o K8StorageClassOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parameters for the provisioner that should create volumes of this storage class
func (o K8StorageClassOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

// Indicates the type of the reclaim policy Defaults to `Delete`.
func (o K8StorageClassOutput) ReclaimPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.StringPtrOutput { return v.ReclaimPolicy }).(pulumi.StringPtrOutput)
}

// Indicates the type of the provisioner
func (o K8StorageClassOutput) StorageProvisioner() pulumi.StringOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.StringOutput { return v.StorageProvisioner }).(pulumi.StringOutput)
}

// The GUID of the tenant that the storage class will be created in.
func (o K8StorageClassOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Indicates when volume binding and dynamic provisioning should occur Defaults to `Immediate`.
func (o K8StorageClassOutput) VolumeBindingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *K8StorageClass) pulumi.StringPtrOutput { return v.VolumeBindingMode }).(pulumi.StringPtrOutput)
}

type K8StorageClassArrayOutput struct{ *pulumi.OutputState }

func (K8StorageClassArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*K8StorageClass)(nil)).Elem()
}

func (o K8StorageClassArrayOutput) ToK8StorageClassArrayOutput() K8StorageClassArrayOutput {
	return o
}

func (o K8StorageClassArrayOutput) ToK8StorageClassArrayOutputWithContext(ctx context.Context) K8StorageClassArrayOutput {
	return o
}

func (o K8StorageClassArrayOutput) Index(i pulumi.IntInput) K8StorageClassOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *K8StorageClass {
		return vs[0].([]*K8StorageClass)[vs[1].(int)]
	}).(K8StorageClassOutput)
}

type K8StorageClassMapOutput struct{ *pulumi.OutputState }

func (K8StorageClassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*K8StorageClass)(nil)).Elem()
}

func (o K8StorageClassMapOutput) ToK8StorageClassMapOutput() K8StorageClassMapOutput {
	return o
}

func (o K8StorageClassMapOutput) ToK8StorageClassMapOutputWithContext(ctx context.Context) K8StorageClassMapOutput {
	return o
}

func (o K8StorageClassMapOutput) MapIndex(k pulumi.StringInput) K8StorageClassOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *K8StorageClass {
		return vs[0].(map[string]*K8StorageClass)[vs[1].(string)]
	}).(K8StorageClassOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*K8StorageClassInput)(nil)).Elem(), &K8StorageClass{})
	pulumi.RegisterInputType(reflect.TypeOf((*K8StorageClassArrayInput)(nil)).Elem(), K8StorageClassArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*K8StorageClassMapInput)(nil)).Elem(), K8StorageClassMap{})
	pulumi.RegisterOutputType(K8StorageClassOutput{})
	pulumi.RegisterOutputType(K8StorageClassArrayOutput{})
	pulumi.RegisterOutputType(K8StorageClassMapOutput{})
}
