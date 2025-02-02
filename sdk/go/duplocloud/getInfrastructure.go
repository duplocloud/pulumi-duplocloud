// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Infrastructure` retrieves details of an infrastructure in Duplo.
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
//			// Example 1 - look up an infrastructure by tenant ID.
//			_, err := duplocloud.LookupInfrastructure(ctx, &duplocloud.LookupInfrastructureArgs{
//				TenantId: pulumi.StringRef(tenantId),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Example 2 - look up an infrastructure by name.
//			_, err = duplocloud.LookupInfrastructure(ctx, &duplocloud.LookupInfrastructureArgs{
//				InfraName: pulumi.StringRef("myinfra"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Example 3 - look up list of certificates by plan ID.
//			_, err = duplocloud.LookupPlanCertificates(ctx, &duplocloud.LookupPlanCertificatesArgs{
//				PlanId: "default",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Example 3 - look up plan certificates by plan ID and certificate name.
//			_, err = duplocloud.LookupPlanCertificate(ctx, &duplocloud.LookupPlanCertificateArgs{
//				PlanId: "default",
//				Name:   "poc.duplocloud.net",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInfrastructure(ctx *pulumi.Context, args *LookupInfrastructureArgs, opts ...pulumi.InvokeOption) (*LookupInfrastructureResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInfrastructureResult
	err := ctx.Invoke("duplocloud:index/getInfrastructure:getInfrastructure", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInfrastructure.
type LookupInfrastructureArgs struct {
	InfraName *string `pulumi:"infraName"`
	TenantId  *string `pulumi:"tenantId"`
}

// A collection of values returned by getInfrastructure.
type LookupInfrastructureResult struct {
	AccountId       string `pulumi:"accountId"`
	AddressPrefix   string `pulumi:"addressPrefix"`
	Azcount         int    `pulumi:"azcount"`
	Cloud           int    `pulumi:"cloud"`
	EnableK8Cluster bool   `pulumi:"enableK8Cluster"`
	// The provider-assigned unique ID for this managed resource.
	Id             string                           `pulumi:"id"`
	InfraName      string                           `pulumi:"infraName"`
	PrivateSubnets []GetInfrastructurePrivateSubnet `pulumi:"privateSubnets"`
	PublicSubnets  []GetInfrastructurePublicSubnet  `pulumi:"publicSubnets"`
	Region         string                           `pulumi:"region"`
	SecurityGroups []GetInfrastructureSecurityGroup `pulumi:"securityGroups"`
	Status         string                           `pulumi:"status"`
	SubnetCidr     int                              `pulumi:"subnetCidr"`
	TenantId       *string                          `pulumi:"tenantId"`
	VpcId          string                           `pulumi:"vpcId"`
	VpcName        string                           `pulumi:"vpcName"`
}

func LookupInfrastructureOutput(ctx *pulumi.Context, args LookupInfrastructureOutputArgs, opts ...pulumi.InvokeOption) LookupInfrastructureResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupInfrastructureResultOutput, error) {
			args := v.(LookupInfrastructureArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("duplocloud:index/getInfrastructure:getInfrastructure", args, LookupInfrastructureResultOutput{}, options).(LookupInfrastructureResultOutput), nil
		}).(LookupInfrastructureResultOutput)
}

// A collection of arguments for invoking getInfrastructure.
type LookupInfrastructureOutputArgs struct {
	InfraName pulumi.StringPtrInput `pulumi:"infraName"`
	TenantId  pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (LookupInfrastructureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInfrastructureArgs)(nil)).Elem()
}

// A collection of values returned by getInfrastructure.
type LookupInfrastructureResultOutput struct{ *pulumi.OutputState }

func (LookupInfrastructureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInfrastructureResult)(nil)).Elem()
}

func (o LookupInfrastructureResultOutput) ToLookupInfrastructureResultOutput() LookupInfrastructureResultOutput {
	return o
}

func (o LookupInfrastructureResultOutput) ToLookupInfrastructureResultOutputWithContext(ctx context.Context) LookupInfrastructureResultOutput {
	return o
}

func (o LookupInfrastructureResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupInfrastructureResultOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) string { return v.AddressPrefix }).(pulumi.StringOutput)
}

func (o LookupInfrastructureResultOutput) Azcount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) int { return v.Azcount }).(pulumi.IntOutput)
}

func (o LookupInfrastructureResultOutput) Cloud() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) int { return v.Cloud }).(pulumi.IntOutput)
}

func (o LookupInfrastructureResultOutput) EnableK8Cluster() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) bool { return v.EnableK8Cluster }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInfrastructureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInfrastructureResultOutput) InfraName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) string { return v.InfraName }).(pulumi.StringOutput)
}

func (o LookupInfrastructureResultOutput) PrivateSubnets() GetInfrastructurePrivateSubnetArrayOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) []GetInfrastructurePrivateSubnet { return v.PrivateSubnets }).(GetInfrastructurePrivateSubnetArrayOutput)
}

func (o LookupInfrastructureResultOutput) PublicSubnets() GetInfrastructurePublicSubnetArrayOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) []GetInfrastructurePublicSubnet { return v.PublicSubnets }).(GetInfrastructurePublicSubnetArrayOutput)
}

func (o LookupInfrastructureResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupInfrastructureResultOutput) SecurityGroups() GetInfrastructureSecurityGroupArrayOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) []GetInfrastructureSecurityGroup { return v.SecurityGroups }).(GetInfrastructureSecurityGroupArrayOutput)
}

func (o LookupInfrastructureResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupInfrastructureResultOutput) SubnetCidr() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) int { return v.SubnetCidr }).(pulumi.IntOutput)
}

func (o LookupInfrastructureResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o LookupInfrastructureResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func (o LookupInfrastructureResultOutput) VpcName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureResult) string { return v.VpcName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInfrastructureResultOutput{})
}
