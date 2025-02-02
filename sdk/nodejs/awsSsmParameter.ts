// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `duplocloud.AwsSsmParameter` manages an AWS SSM parameter in Duplo.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as duplocloud from "@pulumi/duplocloud";
 *
 * const myapp = new duplocloud.Tenant("myapp", {
 *     accountName: "myapp",
 *     planId: "default",
 * });
 * const ssmParam = new duplocloud.AwsSsmParameter("ssm_param", {
 *     tenantId: myapp.tenantId,
 *     name: "ssm_param",
 *     type: "String",
 *     value: "ssm_param_value",
 * });
 * ```
 *
 * ## Import
 *
 * Example: Importing an existing AWS SSM Parameter
 *
 *  - *TENANT_ID* is the tenant GUID
 *
 *  - *NAME* The name for the created Amazon SSM Parameter.
 *
 * # 
 *
 * ```sh
 * $ pulumi import duplocloud:index/awsSsmParameter:AwsSsmParameter ssm_param *TENANT_ID*&#47;*NAME*
 * ```
 */
export class AwsSsmParameter extends pulumi.CustomResource {
    /**
     * Get an existing AwsSsmParameter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AwsSsmParameterState, opts?: pulumi.CustomResourceOptions): AwsSsmParameter {
        return new AwsSsmParameter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'duplocloud:index/awsSsmParameter:AwsSsmParameter';

    /**
     * Returns true if the given object is an instance of AwsSsmParameter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AwsSsmParameter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AwsSsmParameter.__pulumiType;
    }

    public readonly allowedPattern!: pulumi.Output<string>;
    /**
     * The description of the SSM parameter.
     */
    public readonly description!: pulumi.Output<string>;
    public readonly keyId!: pulumi.Output<string>;
    public /*out*/ readonly lastModifiedDate!: pulumi.Output<string>;
    public /*out*/ readonly lastModifiedUser!: pulumi.Output<string>;
    /**
     * The name of the SSM parameter.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The GUID of the tenant that the SSM parameter will be created in.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The type of the SSM parameter. Valid values are `String`, `StringList`, and `SecureString`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The value of the SSM parameter.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a AwsSsmParameter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AwsSsmParameterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AwsSsmParameterArgs | AwsSsmParameterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AwsSsmParameterState | undefined;
            resourceInputs["allowedPattern"] = state ? state.allowedPattern : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["lastModifiedDate"] = state ? state.lastModifiedDate : undefined;
            resourceInputs["lastModifiedUser"] = state ? state.lastModifiedUser : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as AwsSsmParameterArgs | undefined;
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["allowedPattern"] = args ? args.allowedPattern : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args?.value ? pulumi.secret(args.value) : undefined;
            resourceInputs["lastModifiedDate"] = undefined /*out*/;
            resourceInputs["lastModifiedUser"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["value"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AwsSsmParameter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AwsSsmParameter resources.
 */
export interface AwsSsmParameterState {
    allowedPattern?: pulumi.Input<string>;
    /**
     * The description of the SSM parameter.
     */
    description?: pulumi.Input<string>;
    keyId?: pulumi.Input<string>;
    lastModifiedDate?: pulumi.Input<string>;
    lastModifiedUser?: pulumi.Input<string>;
    /**
     * The name of the SSM parameter.
     */
    name?: pulumi.Input<string>;
    /**
     * The GUID of the tenant that the SSM parameter will be created in.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The type of the SSM parameter. Valid values are `String`, `StringList`, and `SecureString`.
     */
    type?: pulumi.Input<string>;
    /**
     * The value of the SSM parameter.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AwsSsmParameter resource.
 */
export interface AwsSsmParameterArgs {
    allowedPattern?: pulumi.Input<string>;
    /**
     * The description of the SSM parameter.
     */
    description?: pulumi.Input<string>;
    keyId?: pulumi.Input<string>;
    /**
     * The name of the SSM parameter.
     */
    name?: pulumi.Input<string>;
    /**
     * The GUID of the tenant that the SSM parameter will be created in.
     */
    tenantId: pulumi.Input<string>;
    /**
     * The type of the SSM parameter. Valid values are `String`, `StringList`, and `SecureString`.
     */
    type: pulumi.Input<string>;
    /**
     * The value of the SSM parameter.
     */
    value: pulumi.Input<string>;
}
