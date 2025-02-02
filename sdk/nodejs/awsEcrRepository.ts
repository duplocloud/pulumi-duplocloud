// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `duplocloud.AwsEcrRepository` manages an aws ecr repository in Duplo.
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
 * const test_ecr = new duplocloud.AwsEcrRepository("test-ecr", {
 *     tenantId: myapp.tenantId,
 *     name: "test-ecr",
 *     enableScanImageOnPush: true,
 *     enableTagImmutability: true,
 *     forceDelete: false,
 * });
 * ```
 *
 * ## Import
 *
 * Example: Importing an existing AWS ECR repository
 *
 *  - *TENANT_ID* is the tenant GUID
 *
 *  - *SHORT_NAME* is the short name of the AWS ECR repository
 *
 * # 
 *
 * ```sh
 * $ pulumi import duplocloud:index/awsEcrRepository:AwsEcrRepository myecr *TENANT_ID*&#47;*SHORT_NAME*
 * ```
 */
export class AwsEcrRepository extends pulumi.CustomResource {
    /**
     * Get an existing AwsEcrRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AwsEcrRepositoryState, opts?: pulumi.CustomResourceOptions): AwsEcrRepository {
        return new AwsEcrRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'duplocloud:index/awsEcrRepository:AwsEcrRepository';

    /**
     * Returns true if the given object is an instance of AwsEcrRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AwsEcrRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AwsEcrRepository.__pulumiType;
    }

    /**
     * Full ARN of the ECR repository.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
     */
    public readonly enableScanImageOnPush!: pulumi.Output<boolean>;
    /**
     * The tag mutability setting for the repository.
     */
    public readonly enableTagImmutability!: pulumi.Output<boolean>;
    /**
     * Whether to force delete the repository on destroy operations Defaults to `false`.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * The ARN of the KMS key to use.
     */
    public readonly kmsEncryptionKey!: pulumi.Output<string>;
    /**
     * The name of the ECR Repository.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The registry ID where the repository was created.
     */
    public /*out*/ readonly registryId!: pulumi.Output<string>;
    /**
     * The URL of the repository.
     */
    public /*out*/ readonly repositoryUrl!: pulumi.Output<string>;
    /**
     * The GUID of the tenant that the aws ecr repository will be created in.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a AwsEcrRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AwsEcrRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AwsEcrRepositoryArgs | AwsEcrRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AwsEcrRepositoryState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["enableScanImageOnPush"] = state ? state.enableScanImageOnPush : undefined;
            resourceInputs["enableTagImmutability"] = state ? state.enableTagImmutability : undefined;
            resourceInputs["forceDelete"] = state ? state.forceDelete : undefined;
            resourceInputs["kmsEncryptionKey"] = state ? state.kmsEncryptionKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["registryId"] = state ? state.registryId : undefined;
            resourceInputs["repositoryUrl"] = state ? state.repositoryUrl : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as AwsEcrRepositoryArgs | undefined;
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["enableScanImageOnPush"] = args ? args.enableScanImageOnPush : undefined;
            resourceInputs["enableTagImmutability"] = args ? args.enableTagImmutability : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["kmsEncryptionKey"] = args ? args.kmsEncryptionKey : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["registryId"] = undefined /*out*/;
            resourceInputs["repositoryUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AwsEcrRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AwsEcrRepository resources.
 */
export interface AwsEcrRepositoryState {
    /**
     * Full ARN of the ECR repository.
     */
    arn?: pulumi.Input<string>;
    /**
     * Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
     */
    enableScanImageOnPush?: pulumi.Input<boolean>;
    /**
     * The tag mutability setting for the repository.
     */
    enableTagImmutability?: pulumi.Input<boolean>;
    /**
     * Whether to force delete the repository on destroy operations Defaults to `false`.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * The ARN of the KMS key to use.
     */
    kmsEncryptionKey?: pulumi.Input<string>;
    /**
     * The name of the ECR Repository.
     */
    name?: pulumi.Input<string>;
    /**
     * The registry ID where the repository was created.
     */
    registryId?: pulumi.Input<string>;
    /**
     * The URL of the repository.
     */
    repositoryUrl?: pulumi.Input<string>;
    /**
     * The GUID of the tenant that the aws ecr repository will be created in.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AwsEcrRepository resource.
 */
export interface AwsEcrRepositoryArgs {
    /**
     * Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
     */
    enableScanImageOnPush?: pulumi.Input<boolean>;
    /**
     * The tag mutability setting for the repository.
     */
    enableTagImmutability?: pulumi.Input<boolean>;
    /**
     * Whether to force delete the repository on destroy operations Defaults to `false`.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * The ARN of the KMS key to use.
     */
    kmsEncryptionKey?: pulumi.Input<string>;
    /**
     * The name of the ECR Repository.
     */
    name?: pulumi.Input<string>;
    /**
     * The GUID of the tenant that the aws ecr repository will be created in.
     */
    tenantId: pulumi.Input<string>;
}
