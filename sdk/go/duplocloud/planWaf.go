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

// `PlanWaf` manages the list of waf's avaialble to a plan in Duplo.
//
// This resource allows you take control of individual waf's for a specific plan.
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
//			_, err := duplocloud.NewPlanWaf(ctx, "myplan", &duplocloud.PlanWafArgs{
//				PlanId:       pulumi.String("plan-name"),
//				WafName:      pulumi.String("WebAcl name"),
//				WafArn:       pulumi.String("WebAcl ARN"),
//				DashboardUrl: pulumi.String("dashboard url"),
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
// Example: Importing an existing WAF instance
//
//   - *PLAN_ID* is the plan name
//
//   - *WAF_NAME* is the name of the WAF
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/planWaf:PlanWaf myplan *PLAN_ID*/*WAF_NAME*
// ```
type PlanWaf struct {
	pulumi.CustomResourceState

	DashboardUrl pulumi.StringOutput `pulumi:"dashboardUrl"`
	// The ID of the plan for waf.
	PlanId  pulumi.StringOutput `pulumi:"planId"`
	WafArn  pulumi.StringOutput `pulumi:"wafArn"`
	WafName pulumi.StringOutput `pulumi:"wafName"`
}

// NewPlanWaf registers a new resource with the given unique name, arguments, and options.
func NewPlanWaf(ctx *pulumi.Context,
	name string, args *PlanWafArgs, opts ...pulumi.ResourceOption) (*PlanWaf, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PlanId == nil {
		return nil, errors.New("invalid value for required argument 'PlanId'")
	}
	if args.WafArn == nil {
		return nil, errors.New("invalid value for required argument 'WafArn'")
	}
	if args.WafName == nil {
		return nil, errors.New("invalid value for required argument 'WafName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PlanWaf
	err := ctx.RegisterResource("duplocloud:index/planWaf:PlanWaf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlanWaf gets an existing PlanWaf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlanWaf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlanWafState, opts ...pulumi.ResourceOption) (*PlanWaf, error) {
	var resource PlanWaf
	err := ctx.ReadResource("duplocloud:index/planWaf:PlanWaf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlanWaf resources.
type planWafState struct {
	DashboardUrl *string `pulumi:"dashboardUrl"`
	// The ID of the plan for waf.
	PlanId  *string `pulumi:"planId"`
	WafArn  *string `pulumi:"wafArn"`
	WafName *string `pulumi:"wafName"`
}

type PlanWafState struct {
	DashboardUrl pulumi.StringPtrInput
	// The ID of the plan for waf.
	PlanId  pulumi.StringPtrInput
	WafArn  pulumi.StringPtrInput
	WafName pulumi.StringPtrInput
}

func (PlanWafState) ElementType() reflect.Type {
	return reflect.TypeOf((*planWafState)(nil)).Elem()
}

type planWafArgs struct {
	DashboardUrl *string `pulumi:"dashboardUrl"`
	// The ID of the plan for waf.
	PlanId  string `pulumi:"planId"`
	WafArn  string `pulumi:"wafArn"`
	WafName string `pulumi:"wafName"`
}

// The set of arguments for constructing a PlanWaf resource.
type PlanWafArgs struct {
	DashboardUrl pulumi.StringPtrInput
	// The ID of the plan for waf.
	PlanId  pulumi.StringInput
	WafArn  pulumi.StringInput
	WafName pulumi.StringInput
}

func (PlanWafArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*planWafArgs)(nil)).Elem()
}

type PlanWafInput interface {
	pulumi.Input

	ToPlanWafOutput() PlanWafOutput
	ToPlanWafOutputWithContext(ctx context.Context) PlanWafOutput
}

func (*PlanWaf) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanWaf)(nil)).Elem()
}

func (i *PlanWaf) ToPlanWafOutput() PlanWafOutput {
	return i.ToPlanWafOutputWithContext(context.Background())
}

func (i *PlanWaf) ToPlanWafOutputWithContext(ctx context.Context) PlanWafOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanWafOutput)
}

// PlanWafArrayInput is an input type that accepts PlanWafArray and PlanWafArrayOutput values.
// You can construct a concrete instance of `PlanWafArrayInput` via:
//
//	PlanWafArray{ PlanWafArgs{...} }
type PlanWafArrayInput interface {
	pulumi.Input

	ToPlanWafArrayOutput() PlanWafArrayOutput
	ToPlanWafArrayOutputWithContext(context.Context) PlanWafArrayOutput
}

type PlanWafArray []PlanWafInput

func (PlanWafArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PlanWaf)(nil)).Elem()
}

func (i PlanWafArray) ToPlanWafArrayOutput() PlanWafArrayOutput {
	return i.ToPlanWafArrayOutputWithContext(context.Background())
}

func (i PlanWafArray) ToPlanWafArrayOutputWithContext(ctx context.Context) PlanWafArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanWafArrayOutput)
}

// PlanWafMapInput is an input type that accepts PlanWafMap and PlanWafMapOutput values.
// You can construct a concrete instance of `PlanWafMapInput` via:
//
//	PlanWafMap{ "key": PlanWafArgs{...} }
type PlanWafMapInput interface {
	pulumi.Input

	ToPlanWafMapOutput() PlanWafMapOutput
	ToPlanWafMapOutputWithContext(context.Context) PlanWafMapOutput
}

type PlanWafMap map[string]PlanWafInput

func (PlanWafMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PlanWaf)(nil)).Elem()
}

func (i PlanWafMap) ToPlanWafMapOutput() PlanWafMapOutput {
	return i.ToPlanWafMapOutputWithContext(context.Background())
}

func (i PlanWafMap) ToPlanWafMapOutputWithContext(ctx context.Context) PlanWafMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanWafMapOutput)
}

type PlanWafOutput struct{ *pulumi.OutputState }

func (PlanWafOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanWaf)(nil)).Elem()
}

func (o PlanWafOutput) ToPlanWafOutput() PlanWafOutput {
	return o
}

func (o PlanWafOutput) ToPlanWafOutputWithContext(ctx context.Context) PlanWafOutput {
	return o
}

func (o PlanWafOutput) DashboardUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *PlanWaf) pulumi.StringOutput { return v.DashboardUrl }).(pulumi.StringOutput)
}

// The ID of the plan for waf.
func (o PlanWafOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *PlanWaf) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

func (o PlanWafOutput) WafArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PlanWaf) pulumi.StringOutput { return v.WafArn }).(pulumi.StringOutput)
}

func (o PlanWafOutput) WafName() pulumi.StringOutput {
	return o.ApplyT(func(v *PlanWaf) pulumi.StringOutput { return v.WafName }).(pulumi.StringOutput)
}

type PlanWafArrayOutput struct{ *pulumi.OutputState }

func (PlanWafArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PlanWaf)(nil)).Elem()
}

func (o PlanWafArrayOutput) ToPlanWafArrayOutput() PlanWafArrayOutput {
	return o
}

func (o PlanWafArrayOutput) ToPlanWafArrayOutputWithContext(ctx context.Context) PlanWafArrayOutput {
	return o
}

func (o PlanWafArrayOutput) Index(i pulumi.IntInput) PlanWafOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PlanWaf {
		return vs[0].([]*PlanWaf)[vs[1].(int)]
	}).(PlanWafOutput)
}

type PlanWafMapOutput struct{ *pulumi.OutputState }

func (PlanWafMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PlanWaf)(nil)).Elem()
}

func (o PlanWafMapOutput) ToPlanWafMapOutput() PlanWafMapOutput {
	return o
}

func (o PlanWafMapOutput) ToPlanWafMapOutputWithContext(ctx context.Context) PlanWafMapOutput {
	return o
}

func (o PlanWafMapOutput) MapIndex(k pulumi.StringInput) PlanWafOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PlanWaf {
		return vs[0].(map[string]*PlanWaf)[vs[1].(string)]
	}).(PlanWafOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlanWafInput)(nil)).Elem(), &PlanWaf{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanWafArrayInput)(nil)).Elem(), PlanWafArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanWafMapInput)(nil)).Elem(), PlanWafMap{})
	pulumi.RegisterOutputType(PlanWafOutput{})
	pulumi.RegisterOutputType(PlanWafArrayOutput{})
	pulumi.RegisterOutputType(PlanWafMapOutput{})
}
