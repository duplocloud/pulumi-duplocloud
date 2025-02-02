// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `duplocloud.AzureNetworkSecurityRule` manages an Azure security group rule in Duplo.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as duplocloud from "@pulumi/duplocloud";
 *
 * const securityRule = new duplocloud.AzureNetworkSecurityRule("security_rule", {
 *     infraName: "demo",
 *     networkSecurityGroupName: "duploinfra-sub01",
 *     name: "test-sg-rule",
 *     sourceRuleType: 0,
 *     destinationRuleType: 0,
 *     protocol: "tcp",
 *     sourcePortRange: "*",
 *     destinationPortRange: "*",
 *     sourceAddressPrefix: "49.207.178.47/32",
 *     destinationAddressPrefix: "*",
 *     access: "Allow",
 *     priority: 200,
 *     direction: "Inbound",
 * });
 * ```
 *
 * ## Import
 *
 * Example: Importing an existing Azure Security Group Rule
 *
 *  - *INFRA_NAME* is the Duplo Infra
 *
 *  - *SG_FULL_NAME* is the fullname of the Security Group, Example- "duploinfra-<SG_SHORT_NAME>"
 *
 *  - *RULE_FULL_NAME* is the fullname of the Security Group Rule
 *
 * # 
 *
 * ```sh
 * $ pulumi import duplocloud:index/azureNetworkSecurityRule:AzureNetworkSecurityRule security_rule *INFRA_NAME*&#47;*SG_FULL_NAME*&#47;*RULE_FULL_NAME*
 * ```
 */
export class AzureNetworkSecurityRule extends pulumi.CustomResource {
    /**
     * Get an existing AzureNetworkSecurityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureNetworkSecurityRuleState, opts?: pulumi.CustomResourceOptions): AzureNetworkSecurityRule {
        return new AzureNetworkSecurityRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'duplocloud:index/azureNetworkSecurityRule:AzureNetworkSecurityRule';

    /**
     * Returns true if the given object is an instance of AzureNetworkSecurityRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureNetworkSecurityRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureNetworkSecurityRule.__pulumiType;
    }

    /**
     * Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
     */
    public readonly access!: pulumi.Output<string>;
    /**
     * CIDR or destination IP range or * to match any IP.
     */
    public readonly destinationAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any.
     */
    public readonly destinationPortRange!: pulumi.Output<string | undefined>;
    /**
     * Type of the destination security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
     */
    public readonly destinationRuleType!: pulumi.Output<number>;
    /**
     * The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * The fullname of the security group rule.
     */
    public /*out*/ readonly fullname!: pulumi.Output<string>;
    /**
     * The name of the infrastructure.  Infrastructure names are globally unique and less than 13 characters.
     */
    public readonly infraName!: pulumi.Output<string>;
    /**
     * The name of the security group rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Network Security Group that we want to attach the rule to.
     */
    public readonly networkSecurityGroupName!: pulumi.Output<string>;
    /**
     * Specifies the priority of the rule.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * Network protocol this rule applies to. Possible values include `tcp`, `udp`, `icmp`, `esp`, `ah` or `*` (which matches all).
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * CIDR or source IP range or * to match any IP.
     */
    public readonly sourceAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * Source Port or Range. Integer or range between `0` and `65535` or `*` to match any.
     */
    public readonly sourcePortRange!: pulumi.Output<string | undefined>;
    /**
     * Type of the source security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
     */
    public readonly sourceRuleType!: pulumi.Output<number>;

    /**
     * Create a AzureNetworkSecurityRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureNetworkSecurityRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureNetworkSecurityRuleArgs | AzureNetworkSecurityRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureNetworkSecurityRuleState | undefined;
            resourceInputs["access"] = state ? state.access : undefined;
            resourceInputs["destinationAddressPrefix"] = state ? state.destinationAddressPrefix : undefined;
            resourceInputs["destinationPortRange"] = state ? state.destinationPortRange : undefined;
            resourceInputs["destinationRuleType"] = state ? state.destinationRuleType : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["fullname"] = state ? state.fullname : undefined;
            resourceInputs["infraName"] = state ? state.infraName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkSecurityGroupName"] = state ? state.networkSecurityGroupName : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["sourceAddressPrefix"] = state ? state.sourceAddressPrefix : undefined;
            resourceInputs["sourcePortRange"] = state ? state.sourcePortRange : undefined;
            resourceInputs["sourceRuleType"] = state ? state.sourceRuleType : undefined;
        } else {
            const args = argsOrState as AzureNetworkSecurityRuleArgs | undefined;
            if ((!args || args.access === undefined) && !opts.urn) {
                throw new Error("Missing required property 'access'");
            }
            if ((!args || args.destinationRuleType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationRuleType'");
            }
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.infraName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'infraName'");
            }
            if ((!args || args.networkSecurityGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkSecurityGroupName'");
            }
            if ((!args || args.priority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'priority'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.sourceRuleType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceRuleType'");
            }
            resourceInputs["access"] = args ? args.access : undefined;
            resourceInputs["destinationAddressPrefix"] = args ? args.destinationAddressPrefix : undefined;
            resourceInputs["destinationPortRange"] = args ? args.destinationPortRange : undefined;
            resourceInputs["destinationRuleType"] = args ? args.destinationRuleType : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["infraName"] = args ? args.infraName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkSecurityGroupName"] = args ? args.networkSecurityGroupName : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["sourceAddressPrefix"] = args ? args.sourceAddressPrefix : undefined;
            resourceInputs["sourcePortRange"] = args ? args.sourcePortRange : undefined;
            resourceInputs["sourceRuleType"] = args ? args.sourceRuleType : undefined;
            resourceInputs["fullname"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AzureNetworkSecurityRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureNetworkSecurityRule resources.
 */
export interface AzureNetworkSecurityRuleState {
    /**
     * Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
     */
    access?: pulumi.Input<string>;
    /**
     * CIDR or destination IP range or * to match any IP.
     */
    destinationAddressPrefix?: pulumi.Input<string>;
    /**
     * Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any.
     */
    destinationPortRange?: pulumi.Input<string>;
    /**
     * Type of the destination security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
     */
    destinationRuleType?: pulumi.Input<number>;
    /**
     * The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
     */
    direction?: pulumi.Input<string>;
    /**
     * The fullname of the security group rule.
     */
    fullname?: pulumi.Input<string>;
    /**
     * The name of the infrastructure.  Infrastructure names are globally unique and less than 13 characters.
     */
    infraName?: pulumi.Input<string>;
    /**
     * The name of the security group rule.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Network Security Group that we want to attach the rule to.
     */
    networkSecurityGroupName?: pulumi.Input<string>;
    /**
     * Specifies the priority of the rule.
     */
    priority?: pulumi.Input<number>;
    /**
     * Network protocol this rule applies to. Possible values include `tcp`, `udp`, `icmp`, `esp`, `ah` or `*` (which matches all).
     */
    protocol?: pulumi.Input<string>;
    /**
     * CIDR or source IP range or * to match any IP.
     */
    sourceAddressPrefix?: pulumi.Input<string>;
    /**
     * Source Port or Range. Integer or range between `0` and `65535` or `*` to match any.
     */
    sourcePortRange?: pulumi.Input<string>;
    /**
     * Type of the source security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
     */
    sourceRuleType?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AzureNetworkSecurityRule resource.
 */
export interface AzureNetworkSecurityRuleArgs {
    /**
     * Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
     */
    access: pulumi.Input<string>;
    /**
     * CIDR or destination IP range or * to match any IP.
     */
    destinationAddressPrefix?: pulumi.Input<string>;
    /**
     * Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any.
     */
    destinationPortRange?: pulumi.Input<string>;
    /**
     * Type of the destination security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
     */
    destinationRuleType: pulumi.Input<number>;
    /**
     * The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
     */
    direction: pulumi.Input<string>;
    /**
     * The name of the infrastructure.  Infrastructure names are globally unique and less than 13 characters.
     */
    infraName: pulumi.Input<string>;
    /**
     * The name of the security group rule.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Network Security Group that we want to attach the rule to.
     */
    networkSecurityGroupName: pulumi.Input<string>;
    /**
     * Specifies the priority of the rule.
     */
    priority: pulumi.Input<number>;
    /**
     * Network protocol this rule applies to. Possible values include `tcp`, `udp`, `icmp`, `esp`, `ah` or `*` (which matches all).
     */
    protocol: pulumi.Input<string>;
    /**
     * CIDR or source IP range or * to match any IP.
     */
    sourceAddressPrefix?: pulumi.Input<string>;
    /**
     * Source Port or Range. Integer or range between `0` and `65535` or `*` to match any.
     */
    sourcePortRange?: pulumi.Input<string>;
    /**
     * Type of the source security rule. Possible values include `0(IP Address)`, `1(Service Tag)`, `2(Application Security Group)`.
     */
    sourceRuleType: pulumi.Input<number>;
}
