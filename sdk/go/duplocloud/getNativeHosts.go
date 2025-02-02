// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getNativeHosts` lists native hosts in a Duplo tenant.
func GetNativeHosts(ctx *pulumi.Context, args *GetNativeHostsArgs, opts ...pulumi.InvokeOption) (*GetNativeHostsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNativeHostsResult
	err := ctx.Invoke("duplocloud:index/getNativeHosts:getNativeHosts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNativeHosts.
type GetNativeHostsArgs struct {
	// The GUID of the tenant in which to list the hosts.
	TenantId string `pulumi:"tenantId"`
}

// A collection of values returned by getNativeHosts.
type GetNativeHostsResult struct {
	// The list of native hosts.
	Hosts []GetNativeHostsHost `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The GUID of the tenant in which to list the hosts.
	TenantId string `pulumi:"tenantId"`
}

func GetNativeHostsOutput(ctx *pulumi.Context, args GetNativeHostsOutputArgs, opts ...pulumi.InvokeOption) GetNativeHostsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNativeHostsResultOutput, error) {
			args := v.(GetNativeHostsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("duplocloud:index/getNativeHosts:getNativeHosts", args, GetNativeHostsResultOutput{}, options).(GetNativeHostsResultOutput), nil
		}).(GetNativeHostsResultOutput)
}

// A collection of arguments for invoking getNativeHosts.
type GetNativeHostsOutputArgs struct {
	// The GUID of the tenant in which to list the hosts.
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (GetNativeHostsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNativeHostsArgs)(nil)).Elem()
}

// A collection of values returned by getNativeHosts.
type GetNativeHostsResultOutput struct{ *pulumi.OutputState }

func (GetNativeHostsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNativeHostsResult)(nil)).Elem()
}

func (o GetNativeHostsResultOutput) ToGetNativeHostsResultOutput() GetNativeHostsResultOutput {
	return o
}

func (o GetNativeHostsResultOutput) ToGetNativeHostsResultOutputWithContext(ctx context.Context) GetNativeHostsResultOutput {
	return o
}

// The list of native hosts.
func (o GetNativeHostsResultOutput) Hosts() GetNativeHostsHostArrayOutput {
	return o.ApplyT(func(v GetNativeHostsResult) []GetNativeHostsHost { return v.Hosts }).(GetNativeHostsHostArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNativeHostsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNativeHostsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The GUID of the tenant in which to list the hosts.
func (o GetNativeHostsResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNativeHostsResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNativeHostsResultOutput{})
}
