// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDuploServiceParams(ctx *pulumi.Context, args *LookupDuploServiceParamsArgs, opts ...pulumi.InvokeOption) (*LookupDuploServiceParamsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDuploServiceParamsResult
	err := ctx.Invoke("duplocloud:index/getDuploServiceParams:getDuploServiceParams", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDuploServiceParams.
type LookupDuploServiceParamsArgs struct {
	DnsPrfx                   *string `pulumi:"dnsPrfx"`
	ReplicationControllerName *string `pulumi:"replicationControllerName"`
	TenantId                  string  `pulumi:"tenantId"`
	Webaclid                  *string `pulumi:"webaclid"`
}

// A collection of values returned by getDuploServiceParams.
type LookupDuploServiceParamsResult struct {
	DnsPrfx *string `pulumi:"dnsPrfx"`
	// The provider-assigned unique ID for this managed resource.
	Id                        string                        `pulumi:"id"`
	ReplicationControllerName *string                       `pulumi:"replicationControllerName"`
	Results                   []GetDuploServiceParamsResult `pulumi:"results"`
	TenantId                  string                        `pulumi:"tenantId"`
	Webaclid                  *string                       `pulumi:"webaclid"`
}

func LookupDuploServiceParamsOutput(ctx *pulumi.Context, args LookupDuploServiceParamsOutputArgs, opts ...pulumi.InvokeOption) LookupDuploServiceParamsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDuploServiceParamsResultOutput, error) {
			args := v.(LookupDuploServiceParamsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("duplocloud:index/getDuploServiceParams:getDuploServiceParams", args, LookupDuploServiceParamsResultOutput{}, options).(LookupDuploServiceParamsResultOutput), nil
		}).(LookupDuploServiceParamsResultOutput)
}

// A collection of arguments for invoking getDuploServiceParams.
type LookupDuploServiceParamsOutputArgs struct {
	DnsPrfx                   pulumi.StringPtrInput `pulumi:"dnsPrfx"`
	ReplicationControllerName pulumi.StringPtrInput `pulumi:"replicationControllerName"`
	TenantId                  pulumi.StringInput    `pulumi:"tenantId"`
	Webaclid                  pulumi.StringPtrInput `pulumi:"webaclid"`
}

func (LookupDuploServiceParamsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDuploServiceParamsArgs)(nil)).Elem()
}

// A collection of values returned by getDuploServiceParams.
type LookupDuploServiceParamsResultOutput struct{ *pulumi.OutputState }

func (LookupDuploServiceParamsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDuploServiceParamsResult)(nil)).Elem()
}

func (o LookupDuploServiceParamsResultOutput) ToLookupDuploServiceParamsResultOutput() LookupDuploServiceParamsResultOutput {
	return o
}

func (o LookupDuploServiceParamsResultOutput) ToLookupDuploServiceParamsResultOutputWithContext(ctx context.Context) LookupDuploServiceParamsResultOutput {
	return o
}

func (o LookupDuploServiceParamsResultOutput) DnsPrfx() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDuploServiceParamsResult) *string { return v.DnsPrfx }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDuploServiceParamsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDuploServiceParamsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDuploServiceParamsResultOutput) ReplicationControllerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDuploServiceParamsResult) *string { return v.ReplicationControllerName }).(pulumi.StringPtrOutput)
}

func (o LookupDuploServiceParamsResultOutput) Results() GetDuploServiceParamsResultArrayOutput {
	return o.ApplyT(func(v LookupDuploServiceParamsResult) []GetDuploServiceParamsResult { return v.Results }).(GetDuploServiceParamsResultArrayOutput)
}

func (o LookupDuploServiceParamsResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDuploServiceParamsResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupDuploServiceParamsResultOutput) Webaclid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDuploServiceParamsResult) *string { return v.Webaclid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDuploServiceParamsResultOutput{})
}
