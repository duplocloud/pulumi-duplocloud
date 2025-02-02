// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getInfrastructures` retrieves a list of infrastructures in Duplo.
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
//			// Retrieve a list of all infrastructures
//			_, err := duplocloud.GetInfrastructures(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInfrastructures(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetInfrastructuresResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInfrastructuresResult
	err := ctx.Invoke("duplocloud:index/getInfrastructures:getInfrastructures", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getInfrastructures.
type GetInfrastructuresResult struct {
	Datas []GetInfrastructuresData `pulumi:"datas"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetInfrastructuresOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetInfrastructuresResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetInfrastructuresResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("duplocloud:index/getInfrastructures:getInfrastructures", nil, GetInfrastructuresResultOutput{}, options).(GetInfrastructuresResultOutput), nil
	}).(GetInfrastructuresResultOutput)
}

// A collection of values returned by getInfrastructures.
type GetInfrastructuresResultOutput struct{ *pulumi.OutputState }

func (GetInfrastructuresResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInfrastructuresResult)(nil)).Elem()
}

func (o GetInfrastructuresResultOutput) ToGetInfrastructuresResultOutput() GetInfrastructuresResultOutput {
	return o
}

func (o GetInfrastructuresResultOutput) ToGetInfrastructuresResultOutputWithContext(ctx context.Context) GetInfrastructuresResultOutput {
	return o
}

func (o GetInfrastructuresResultOutput) Datas() GetInfrastructuresDataArrayOutput {
	return o.ApplyT(func(v GetInfrastructuresResult) []GetInfrastructuresData { return v.Datas }).(GetInfrastructuresDataArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInfrastructuresResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInfrastructuresResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInfrastructuresResultOutput{})
}
