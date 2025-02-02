// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the duplocloud package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'duplocloud';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * This is the base URL to the Duplo REST API. It must be provided, but it can also be sourced from the `duploHost`
     * environment variable.
     */
    public readonly duploHost!: pulumi.Output<string | undefined>;
    /**
     * This is a bearer token used to authenticate to the Duplo REST API. It must be provided, but it can also be sourced from
     * the `duploToken` environment variable.
     */
    public readonly duploToken!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["duploHost"] = args ? args.duploHost : undefined;
            resourceInputs["duploToken"] = args?.duploToken ? pulumi.secret(args.duploToken) : undefined;
            resourceInputs["httpTimeout"] = pulumi.output(args ? args.httpTimeout : undefined).apply(JSON.stringify);
            resourceInputs["sslNoVerify"] = pulumi.output(args ? args.sslNoVerify : undefined).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["duploToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * This is the base URL to the Duplo REST API. It must be provided, but it can also be sourced from the `duploHost`
     * environment variable.
     */
    duploHost?: pulumi.Input<string>;
    /**
     * This is a bearer token used to authenticate to the Duplo REST API. It must be provided, but it can also be sourced from
     * the `duploToken` environment variable.
     */
    duploToken?: pulumi.Input<string>;
    /**
     * Timeout for HTTP requests in seconds.
     */
    httpTimeout?: pulumi.Input<number>;
    /**
     * Disable SSL certificate verification.
     */
    sslNoVerify?: pulumi.Input<boolean>;
}
