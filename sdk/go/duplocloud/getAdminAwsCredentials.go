// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"context"
	"reflect"

	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAdminAwsCredentials(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetAdminAwsCredentialsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAdminAwsCredentialsResult
	err := ctx.Invoke("duplocloud:index/getAdminAwsCredentials:getAdminAwsCredentials", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getAdminAwsCredentials.
type GetAdminAwsCredentialsResult struct {
	AccessKeyId string `pulumi:"accessKeyId"`
	ConsoleUrl  string `pulumi:"consoleUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	Region          string `pulumi:"region"`
	SecretAccessKey string `pulumi:"secretAccessKey"`
	SessionToken    string `pulumi:"sessionToken"`
}

func GetAdminAwsCredentialsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetAdminAwsCredentialsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetAdminAwsCredentialsResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("duplocloud:index/getAdminAwsCredentials:getAdminAwsCredentials", nil, GetAdminAwsCredentialsResultOutput{}, options).(GetAdminAwsCredentialsResultOutput), nil
	}).(GetAdminAwsCredentialsResultOutput)
}

// A collection of values returned by getAdminAwsCredentials.
type GetAdminAwsCredentialsResultOutput struct{ *pulumi.OutputState }

func (GetAdminAwsCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAdminAwsCredentialsResult)(nil)).Elem()
}

func (o GetAdminAwsCredentialsResultOutput) ToGetAdminAwsCredentialsResultOutput() GetAdminAwsCredentialsResultOutput {
	return o
}

func (o GetAdminAwsCredentialsResultOutput) ToGetAdminAwsCredentialsResultOutputWithContext(ctx context.Context) GetAdminAwsCredentialsResultOutput {
	return o
}

func (o GetAdminAwsCredentialsResultOutput) AccessKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdminAwsCredentialsResult) string { return v.AccessKeyId }).(pulumi.StringOutput)
}

func (o GetAdminAwsCredentialsResultOutput) ConsoleUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdminAwsCredentialsResult) string { return v.ConsoleUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAdminAwsCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdminAwsCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAdminAwsCredentialsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdminAwsCredentialsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetAdminAwsCredentialsResultOutput) SecretAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdminAwsCredentialsResult) string { return v.SecretAccessKey }).(pulumi.StringOutput)
}

func (o GetAdminAwsCredentialsResultOutput) SessionToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetAdminAwsCredentialsResult) string { return v.SessionToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAdminAwsCredentialsResultOutput{})
}
