// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAwsSsmParameters(ctx *pulumi.Context, args *GetAwsSsmParametersArgs, opts ...pulumi.InvokeOption) (*GetAwsSsmParametersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAwsSsmParametersResult
	err := ctx.Invoke("duplocloud:index/getAwsSsmParameters:getAwsSsmParameters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAwsSsmParameters.
type GetAwsSsmParametersArgs struct {
	TenantId string `pulumi:"tenantId"`
}

// A collection of values returned by getAwsSsmParameters.
type GetAwsSsmParametersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                         `pulumi:"id"`
	Parameters []GetAwsSsmParametersParameter `pulumi:"parameters"`
	TenantId   string                         `pulumi:"tenantId"`
}

func GetAwsSsmParametersOutput(ctx *pulumi.Context, args GetAwsSsmParametersOutputArgs, opts ...pulumi.InvokeOption) GetAwsSsmParametersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetAwsSsmParametersResultOutput, error) {
			args := v.(GetAwsSsmParametersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("duplocloud:index/getAwsSsmParameters:getAwsSsmParameters", args, GetAwsSsmParametersResultOutput{}, options).(GetAwsSsmParametersResultOutput), nil
		}).(GetAwsSsmParametersResultOutput)
}

// A collection of arguments for invoking getAwsSsmParameters.
type GetAwsSsmParametersOutputArgs struct {
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

func (GetAwsSsmParametersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAwsSsmParametersArgs)(nil)).Elem()
}

// A collection of values returned by getAwsSsmParameters.
type GetAwsSsmParametersResultOutput struct{ *pulumi.OutputState }

func (GetAwsSsmParametersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAwsSsmParametersResult)(nil)).Elem()
}

func (o GetAwsSsmParametersResultOutput) ToGetAwsSsmParametersResultOutput() GetAwsSsmParametersResultOutput {
	return o
}

func (o GetAwsSsmParametersResultOutput) ToGetAwsSsmParametersResultOutputWithContext(ctx context.Context) GetAwsSsmParametersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAwsSsmParametersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsSsmParametersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAwsSsmParametersResultOutput) Parameters() GetAwsSsmParametersParameterArrayOutput {
	return o.ApplyT(func(v GetAwsSsmParametersResult) []GetAwsSsmParametersParameter { return v.Parameters }).(GetAwsSsmParametersParameterArrayOutput)
}

func (o GetAwsSsmParametersResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsSsmParametersResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAwsSsmParametersResultOutput{})
}
