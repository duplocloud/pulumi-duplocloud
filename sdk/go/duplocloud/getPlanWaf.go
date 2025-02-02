// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `PlanWaf` retrieves details of a web acl in Duplo.
func LookupPlanWaf(ctx *pulumi.Context, args *LookupPlanWafArgs, opts ...pulumi.InvokeOption) (*LookupPlanWafResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPlanWafResult
	err := ctx.Invoke("duplocloud:index/getPlanWaf:getPlanWaf", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlanWaf.
type LookupPlanWafArgs struct {
	// The ID of the plan for waf.
	PlanId  string `pulumi:"planId"`
	WafName string `pulumi:"wafName"`
}

// A collection of values returned by getPlanWaf.
type LookupPlanWafResult struct {
	DashboardUrl string `pulumi:"dashboardUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the plan for waf.
	PlanId  string `pulumi:"planId"`
	WafArn  string `pulumi:"wafArn"`
	WafName string `pulumi:"wafName"`
}

func LookupPlanWafOutput(ctx *pulumi.Context, args LookupPlanWafOutputArgs, opts ...pulumi.InvokeOption) LookupPlanWafResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPlanWafResultOutput, error) {
			args := v.(LookupPlanWafArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("duplocloud:index/getPlanWaf:getPlanWaf", args, LookupPlanWafResultOutput{}, options).(LookupPlanWafResultOutput), nil
		}).(LookupPlanWafResultOutput)
}

// A collection of arguments for invoking getPlanWaf.
type LookupPlanWafOutputArgs struct {
	// The ID of the plan for waf.
	PlanId  pulumi.StringInput `pulumi:"planId"`
	WafName pulumi.StringInput `pulumi:"wafName"`
}

func (LookupPlanWafOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlanWafArgs)(nil)).Elem()
}

// A collection of values returned by getPlanWaf.
type LookupPlanWafResultOutput struct{ *pulumi.OutputState }

func (LookupPlanWafResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlanWafResult)(nil)).Elem()
}

func (o LookupPlanWafResultOutput) ToLookupPlanWafResultOutput() LookupPlanWafResultOutput {
	return o
}

func (o LookupPlanWafResultOutput) ToLookupPlanWafResultOutputWithContext(ctx context.Context) LookupPlanWafResultOutput {
	return o
}

func (o LookupPlanWafResultOutput) DashboardUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanWafResult) string { return v.DashboardUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPlanWafResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanWafResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the plan for waf.
func (o LookupPlanWafResultOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanWafResult) string { return v.PlanId }).(pulumi.StringOutput)
}

func (o LookupPlanWafResultOutput) WafArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanWafResult) string { return v.WafArn }).(pulumi.StringOutput)
}

func (o LookupPlanWafResultOutput) WafName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanWafResult) string { return v.WafName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlanWafResultOutput{})
}
