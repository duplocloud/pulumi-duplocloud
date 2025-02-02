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

// `AzureNetworkSecurityRule` manages an Azure security group rule in Duplo.
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
//			_, err := duplocloud.NewAzureNetworkSecurityRule(ctx, "security_rule", &duplocloud.AzureNetworkSecurityRuleArgs{
//				InfraName:                pulumi.String("demo"),
//				NetworkSecurityGroupName: pulumi.String("duploinfra-sub01"),
//				Name:                     pulumi.String("test-sg-rule"),
//				SourceRuleType:           pulumi.Int(0),
//				DestinationRuleType:      pulumi.Int(0),
//				Protocol:                 pulumi.String("tcp"),
//				SourcePortRange:          pulumi.String("*"),
//				DestinationPortRange:     pulumi.String("*"),
//				SourceAddressPrefix:      pulumi.String("49.207.178.47/32"),
//				DestinationAddressPrefix: pulumi.String("*"),
//				Access:                   pulumi.String("Allow"),
//				Priority:                 pulumi.Int(200),
//				Direction:                pulumi.String("Inbound"),
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
// Example: Importing an existing Azure Security Group Rule
//
//   - *INFRA_NAME* is the Duplo Infra
//
//   - *SG_FULL_NAME* is the fullname of the Security Group, Example- "duploinfra-<SG_SHORT_NAME>"
//
//   - *RULE_FULL_NAME* is the fullname of the Security Group Rule
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/azureNetworkSecurityRule:AzureNetworkSecurityRule security_rule *INFRA_NAME*/*SG_FULL_NAME*/*RULE_FULL_NAME*
// ```
type AzureNetworkSecurityRule struct {
	pulumi.CustomResourceState

	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access pulumi.StringOutput `pulumi:"access"`
	// CIDR or destination IP range or * to match any IP.
	DestinationAddressPrefix pulumi.StringPtrOutput `pulumi:"destinationAddressPrefix"`
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	DestinationPortRange pulumi.StringPtrOutput `pulumi:"destinationPortRange"`
	// Type of the destination security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	DestinationRuleType pulumi.IntOutput `pulumi:"destinationRuleType"`
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// The fullname of the security group rule.
	Fullname pulumi.StringOutput `pulumi:"fullname"`
	// The name of the infrastructure.  Infrastructure names are globally unique and less than 13 characters.
	InfraName pulumi.StringOutput `pulumi:"infraName"`
	// The name of the security group rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Network Security Group that we want to attach the rule to.
	NetworkSecurityGroupName pulumi.StringOutput `pulumi:"networkSecurityGroupName"`
	// Specifies the priority of the rule.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Network protocol this rule applies to. Possible values include `tcp`, `udp`, `icmp`, `esp`, `ah` or `*` (which matches all).
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// CIDR or source IP range or * to match any IP.
	SourceAddressPrefix pulumi.StringPtrOutput `pulumi:"sourceAddressPrefix"`
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	SourcePortRange pulumi.StringPtrOutput `pulumi:"sourcePortRange"`
	// Type of the source security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	SourceRuleType pulumi.IntOutput `pulumi:"sourceRuleType"`
}

// NewAzureNetworkSecurityRule registers a new resource with the given unique name, arguments, and options.
func NewAzureNetworkSecurityRule(ctx *pulumi.Context,
	name string, args *AzureNetworkSecurityRuleArgs, opts ...pulumi.ResourceOption) (*AzureNetworkSecurityRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Access == nil {
		return nil, errors.New("invalid value for required argument 'Access'")
	}
	if args.DestinationRuleType == nil {
		return nil, errors.New("invalid value for required argument 'DestinationRuleType'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.InfraName == nil {
		return nil, errors.New("invalid value for required argument 'InfraName'")
	}
	if args.NetworkSecurityGroupName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityGroupName'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.SourceRuleType == nil {
		return nil, errors.New("invalid value for required argument 'SourceRuleType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AzureNetworkSecurityRule
	err := ctx.RegisterResource("duplocloud:index/azureNetworkSecurityRule:AzureNetworkSecurityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureNetworkSecurityRule gets an existing AzureNetworkSecurityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureNetworkSecurityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureNetworkSecurityRuleState, opts ...pulumi.ResourceOption) (*AzureNetworkSecurityRule, error) {
	var resource AzureNetworkSecurityRule
	err := ctx.ReadResource("duplocloud:index/azureNetworkSecurityRule:AzureNetworkSecurityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureNetworkSecurityRule resources.
type azureNetworkSecurityRuleState struct {
	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access *string `pulumi:"access"`
	// CIDR or destination IP range or * to match any IP.
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// Type of the destination security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	DestinationRuleType *int `pulumi:"destinationRuleType"`
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction *string `pulumi:"direction"`
	// The fullname of the security group rule.
	Fullname *string `pulumi:"fullname"`
	// The name of the infrastructure.  Infrastructure names are globally unique and less than 13 characters.
	InfraName *string `pulumi:"infraName"`
	// The name of the security group rule.
	Name *string `pulumi:"name"`
	// The name of the Network Security Group that we want to attach the rule to.
	NetworkSecurityGroupName *string `pulumi:"networkSecurityGroupName"`
	// Specifies the priority of the rule.
	Priority *int `pulumi:"priority"`
	// Network protocol this rule applies to. Possible values include `tcp`, `udp`, `icmp`, `esp`, `ah` or `*` (which matches all).
	Protocol *string `pulumi:"protocol"`
	// CIDR or source IP range or * to match any IP.
	SourceAddressPrefix *string `pulumi:"sourceAddressPrefix"`
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// Type of the source security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	SourceRuleType *int `pulumi:"sourceRuleType"`
}

type AzureNetworkSecurityRuleState struct {
	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access pulumi.StringPtrInput
	// CIDR or destination IP range or * to match any IP.
	DestinationAddressPrefix pulumi.StringPtrInput
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	DestinationPortRange pulumi.StringPtrInput
	// Type of the destination security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	DestinationRuleType pulumi.IntPtrInput
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction pulumi.StringPtrInput
	// The fullname of the security group rule.
	Fullname pulumi.StringPtrInput
	// The name of the infrastructure.  Infrastructure names are globally unique and less than 13 characters.
	InfraName pulumi.StringPtrInput
	// The name of the security group rule.
	Name pulumi.StringPtrInput
	// The name of the Network Security Group that we want to attach the rule to.
	NetworkSecurityGroupName pulumi.StringPtrInput
	// Specifies the priority of the rule.
	Priority pulumi.IntPtrInput
	// Network protocol this rule applies to. Possible values include `tcp`, `udp`, `icmp`, `esp`, `ah` or `*` (which matches all).
	Protocol pulumi.StringPtrInput
	// CIDR or source IP range or * to match any IP.
	SourceAddressPrefix pulumi.StringPtrInput
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	SourcePortRange pulumi.StringPtrInput
	// Type of the source security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	SourceRuleType pulumi.IntPtrInput
}

func (AzureNetworkSecurityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureNetworkSecurityRuleState)(nil)).Elem()
}

type azureNetworkSecurityRuleArgs struct {
	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access string `pulumi:"access"`
	// CIDR or destination IP range or * to match any IP.
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// Type of the destination security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	DestinationRuleType int `pulumi:"destinationRuleType"`
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction string `pulumi:"direction"`
	// The name of the infrastructure.  Infrastructure names are globally unique and less than 13 characters.
	InfraName string `pulumi:"infraName"`
	// The name of the security group rule.
	Name *string `pulumi:"name"`
	// The name of the Network Security Group that we want to attach the rule to.
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	// Specifies the priority of the rule.
	Priority int `pulumi:"priority"`
	// Network protocol this rule applies to. Possible values include `tcp`, `udp`, `icmp`, `esp`, `ah` or `*` (which matches all).
	Protocol string `pulumi:"protocol"`
	// CIDR or source IP range or * to match any IP.
	SourceAddressPrefix *string `pulumi:"sourceAddressPrefix"`
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// Type of the source security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	SourceRuleType int `pulumi:"sourceRuleType"`
}

// The set of arguments for constructing a AzureNetworkSecurityRule resource.
type AzureNetworkSecurityRuleArgs struct {
	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access pulumi.StringInput
	// CIDR or destination IP range or * to match any IP.
	DestinationAddressPrefix pulumi.StringPtrInput
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	DestinationPortRange pulumi.StringPtrInput
	// Type of the destination security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	DestinationRuleType pulumi.IntInput
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction pulumi.StringInput
	// The name of the infrastructure.  Infrastructure names are globally unique and less than 13 characters.
	InfraName pulumi.StringInput
	// The name of the security group rule.
	Name pulumi.StringPtrInput
	// The name of the Network Security Group that we want to attach the rule to.
	NetworkSecurityGroupName pulumi.StringInput
	// Specifies the priority of the rule.
	Priority pulumi.IntInput
	// Network protocol this rule applies to. Possible values include `tcp`, `udp`, `icmp`, `esp`, `ah` or `*` (which matches all).
	Protocol pulumi.StringInput
	// CIDR or source IP range or * to match any IP.
	SourceAddressPrefix pulumi.StringPtrInput
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any.
	SourcePortRange pulumi.StringPtrInput
	// Type of the source security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
	SourceRuleType pulumi.IntInput
}

func (AzureNetworkSecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureNetworkSecurityRuleArgs)(nil)).Elem()
}

type AzureNetworkSecurityRuleInput interface {
	pulumi.Input

	ToAzureNetworkSecurityRuleOutput() AzureNetworkSecurityRuleOutput
	ToAzureNetworkSecurityRuleOutputWithContext(ctx context.Context) AzureNetworkSecurityRuleOutput
}

func (*AzureNetworkSecurityRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureNetworkSecurityRule)(nil)).Elem()
}

func (i *AzureNetworkSecurityRule) ToAzureNetworkSecurityRuleOutput() AzureNetworkSecurityRuleOutput {
	return i.ToAzureNetworkSecurityRuleOutputWithContext(context.Background())
}

func (i *AzureNetworkSecurityRule) ToAzureNetworkSecurityRuleOutputWithContext(ctx context.Context) AzureNetworkSecurityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureNetworkSecurityRuleOutput)
}

// AzureNetworkSecurityRuleArrayInput is an input type that accepts AzureNetworkSecurityRuleArray and AzureNetworkSecurityRuleArrayOutput values.
// You can construct a concrete instance of `AzureNetworkSecurityRuleArrayInput` via:
//
//	AzureNetworkSecurityRuleArray{ AzureNetworkSecurityRuleArgs{...} }
type AzureNetworkSecurityRuleArrayInput interface {
	pulumi.Input

	ToAzureNetworkSecurityRuleArrayOutput() AzureNetworkSecurityRuleArrayOutput
	ToAzureNetworkSecurityRuleArrayOutputWithContext(context.Context) AzureNetworkSecurityRuleArrayOutput
}

type AzureNetworkSecurityRuleArray []AzureNetworkSecurityRuleInput

func (AzureNetworkSecurityRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureNetworkSecurityRule)(nil)).Elem()
}

func (i AzureNetworkSecurityRuleArray) ToAzureNetworkSecurityRuleArrayOutput() AzureNetworkSecurityRuleArrayOutput {
	return i.ToAzureNetworkSecurityRuleArrayOutputWithContext(context.Background())
}

func (i AzureNetworkSecurityRuleArray) ToAzureNetworkSecurityRuleArrayOutputWithContext(ctx context.Context) AzureNetworkSecurityRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureNetworkSecurityRuleArrayOutput)
}

// AzureNetworkSecurityRuleMapInput is an input type that accepts AzureNetworkSecurityRuleMap and AzureNetworkSecurityRuleMapOutput values.
// You can construct a concrete instance of `AzureNetworkSecurityRuleMapInput` via:
//
//	AzureNetworkSecurityRuleMap{ "key": AzureNetworkSecurityRuleArgs{...} }
type AzureNetworkSecurityRuleMapInput interface {
	pulumi.Input

	ToAzureNetworkSecurityRuleMapOutput() AzureNetworkSecurityRuleMapOutput
	ToAzureNetworkSecurityRuleMapOutputWithContext(context.Context) AzureNetworkSecurityRuleMapOutput
}

type AzureNetworkSecurityRuleMap map[string]AzureNetworkSecurityRuleInput

func (AzureNetworkSecurityRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureNetworkSecurityRule)(nil)).Elem()
}

func (i AzureNetworkSecurityRuleMap) ToAzureNetworkSecurityRuleMapOutput() AzureNetworkSecurityRuleMapOutput {
	return i.ToAzureNetworkSecurityRuleMapOutputWithContext(context.Background())
}

func (i AzureNetworkSecurityRuleMap) ToAzureNetworkSecurityRuleMapOutputWithContext(ctx context.Context) AzureNetworkSecurityRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureNetworkSecurityRuleMapOutput)
}

type AzureNetworkSecurityRuleOutput struct{ *pulumi.OutputState }

func (AzureNetworkSecurityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureNetworkSecurityRule)(nil)).Elem()
}

func (o AzureNetworkSecurityRuleOutput) ToAzureNetworkSecurityRuleOutput() AzureNetworkSecurityRuleOutput {
	return o
}

func (o AzureNetworkSecurityRuleOutput) ToAzureNetworkSecurityRuleOutputWithContext(ctx context.Context) AzureNetworkSecurityRuleOutput {
	return o
}

// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
func (o AzureNetworkSecurityRuleOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringOutput { return v.Access }).(pulumi.StringOutput)
}

// CIDR or destination IP range or * to match any IP.
func (o AzureNetworkSecurityRuleOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringPtrOutput { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any.
func (o AzureNetworkSecurityRuleOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringPtrOutput { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

// Type of the destination security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
func (o AzureNetworkSecurityRuleOutput) DestinationRuleType() pulumi.IntOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.IntOutput { return v.DestinationRuleType }).(pulumi.IntOutput)
}

// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
func (o AzureNetworkSecurityRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// The fullname of the security group rule.
func (o AzureNetworkSecurityRuleOutput) Fullname() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringOutput { return v.Fullname }).(pulumi.StringOutput)
}

// The name of the infrastructure.  Infrastructure names are globally unique and less than 13 characters.
func (o AzureNetworkSecurityRuleOutput) InfraName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringOutput { return v.InfraName }).(pulumi.StringOutput)
}

// The name of the security group rule.
func (o AzureNetworkSecurityRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the Network Security Group that we want to attach the rule to.
func (o AzureNetworkSecurityRuleOutput) NetworkSecurityGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringOutput { return v.NetworkSecurityGroupName }).(pulumi.StringOutput)
}

// Specifies the priority of the rule.
func (o AzureNetworkSecurityRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// Network protocol this rule applies to. Possible values include `tcp`, `udp`, `icmp`, `esp`, `ah` or `*` (which matches all).
func (o AzureNetworkSecurityRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// CIDR or source IP range or * to match any IP.
func (o AzureNetworkSecurityRuleOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringPtrOutput { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any.
func (o AzureNetworkSecurityRuleOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.StringPtrOutput { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

// Type of the source security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
func (o AzureNetworkSecurityRuleOutput) SourceRuleType() pulumi.IntOutput {
	return o.ApplyT(func(v *AzureNetworkSecurityRule) pulumi.IntOutput { return v.SourceRuleType }).(pulumi.IntOutput)
}

type AzureNetworkSecurityRuleArrayOutput struct{ *pulumi.OutputState }

func (AzureNetworkSecurityRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureNetworkSecurityRule)(nil)).Elem()
}

func (o AzureNetworkSecurityRuleArrayOutput) ToAzureNetworkSecurityRuleArrayOutput() AzureNetworkSecurityRuleArrayOutput {
	return o
}

func (o AzureNetworkSecurityRuleArrayOutput) ToAzureNetworkSecurityRuleArrayOutputWithContext(ctx context.Context) AzureNetworkSecurityRuleArrayOutput {
	return o
}

func (o AzureNetworkSecurityRuleArrayOutput) Index(i pulumi.IntInput) AzureNetworkSecurityRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureNetworkSecurityRule {
		return vs[0].([]*AzureNetworkSecurityRule)[vs[1].(int)]
	}).(AzureNetworkSecurityRuleOutput)
}

type AzureNetworkSecurityRuleMapOutput struct{ *pulumi.OutputState }

func (AzureNetworkSecurityRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureNetworkSecurityRule)(nil)).Elem()
}

func (o AzureNetworkSecurityRuleMapOutput) ToAzureNetworkSecurityRuleMapOutput() AzureNetworkSecurityRuleMapOutput {
	return o
}

func (o AzureNetworkSecurityRuleMapOutput) ToAzureNetworkSecurityRuleMapOutputWithContext(ctx context.Context) AzureNetworkSecurityRuleMapOutput {
	return o
}

func (o AzureNetworkSecurityRuleMapOutput) MapIndex(k pulumi.StringInput) AzureNetworkSecurityRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureNetworkSecurityRule {
		return vs[0].(map[string]*AzureNetworkSecurityRule)[vs[1].(string)]
	}).(AzureNetworkSecurityRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureNetworkSecurityRuleInput)(nil)).Elem(), &AzureNetworkSecurityRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureNetworkSecurityRuleArrayInput)(nil)).Elem(), AzureNetworkSecurityRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureNetworkSecurityRuleMapInput)(nil)).Elem(), AzureNetworkSecurityRuleMap{})
	pulumi.RegisterOutputType(AzureNetworkSecurityRuleOutput{})
	pulumi.RegisterOutputType(AzureNetworkSecurityRuleArrayOutput{})
	pulumi.RegisterOutputType(AzureNetworkSecurityRuleMapOutput{})
}
