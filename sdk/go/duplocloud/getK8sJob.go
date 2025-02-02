// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			job, err := duplocloud.LookupK8sJob(ctx, &duplocloud.LookupK8sJobArgs{
//				TenantId: tenantId,
//				Metadata: duplocloud.GetK8sJobMetadata{
//					Name: "datajob",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("metadata", job.Metadata.Namespace)
//			return nil
//		})
//	}
//
// ```
func LookupK8sJob(ctx *pulumi.Context, args *LookupK8sJobArgs, opts ...pulumi.InvokeOption) (*LookupK8sJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupK8sJobResult
	err := ctx.Invoke("duplocloud:index/getK8sJob:getK8sJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getK8sJob.
type LookupK8sJobArgs struct {
	// Defaults to `false`.
	IsAnyHostAllowed *bool `pulumi:"isAnyHostAllowed"`
	// Standard job's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata GetK8sJobMetadata `pulumi:"metadata"`
	// The GUID of the tenant that the job will be created in.
	TenantId string `pulumi:"tenantId"`
}

// A collection of values returned by getK8sJob.
type LookupK8sJobResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Defaults to `false`.
	IsAnyHostAllowed *bool `pulumi:"isAnyHostAllowed"`
	// Standard job's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata GetK8sJobMetadata `pulumi:"metadata"`
	// Spec of the job owned by the cluster
	Specs []GetK8sJobSpec `pulumi:"specs"`
	// The GUID of the tenant that the job will be created in.
	TenantId          string `pulumi:"tenantId"`
	WaitForCompletion bool   `pulumi:"waitForCompletion"`
}

func LookupK8sJobOutput(ctx *pulumi.Context, args LookupK8sJobOutputArgs, opts ...pulumi.InvokeOption) LookupK8sJobResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupK8sJobResultOutput, error) {
			args := v.(LookupK8sJobArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("duplocloud:index/getK8sJob:getK8sJob", args, LookupK8sJobResultOutput{}, options).(LookupK8sJobResultOutput), nil
		}).(LookupK8sJobResultOutput)
}

// A collection of arguments for invoking getK8sJob.
type LookupK8sJobOutputArgs struct {
	// Defaults to `false`.
	IsAnyHostAllowed pulumi.BoolPtrInput `pulumi:"isAnyHostAllowed"`
	// Standard job's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata GetK8sJobMetadataInput `pulumi:"metadata"`
	// The GUID of the tenant that the job will be created in.
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (LookupK8sJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupK8sJobArgs)(nil)).Elem()
}

// A collection of values returned by getK8sJob.
type LookupK8sJobResultOutput struct{ *pulumi.OutputState }

func (LookupK8sJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupK8sJobResult)(nil)).Elem()
}

func (o LookupK8sJobResultOutput) ToLookupK8sJobResultOutput() LookupK8sJobResultOutput {
	return o
}

func (o LookupK8sJobResultOutput) ToLookupK8sJobResultOutputWithContext(ctx context.Context) LookupK8sJobResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupK8sJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sJobResult) string { return v.Id }).(pulumi.StringOutput)
}

// Defaults to `false`.
func (o LookupK8sJobResultOutput) IsAnyHostAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupK8sJobResult) *bool { return v.IsAnyHostAllowed }).(pulumi.BoolPtrOutput)
}

// Standard job's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o LookupK8sJobResultOutput) Metadata() GetK8sJobMetadataOutput {
	return o.ApplyT(func(v LookupK8sJobResult) GetK8sJobMetadata { return v.Metadata }).(GetK8sJobMetadataOutput)
}

// Spec of the job owned by the cluster
func (o LookupK8sJobResultOutput) Specs() GetK8sJobSpecArrayOutput {
	return o.ApplyT(func(v LookupK8sJobResult) []GetK8sJobSpec { return v.Specs }).(GetK8sJobSpecArrayOutput)
}

// The GUID of the tenant that the job will be created in.
func (o LookupK8sJobResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sJobResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupK8sJobResultOutput) WaitForCompletion() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupK8sJobResult) bool { return v.WaitForCompletion }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupK8sJobResultOutput{})
}
