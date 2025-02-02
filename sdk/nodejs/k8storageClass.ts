// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * `duplocloud.K8StorageClass` manages a kubernetes storage class in a Duplo tenant.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as duplocloud from "@pulumi/duplocloud";
 *
 * const tenantId = "3a0b2ea5-7403-4765-ad6e-8771ca8fa0fd";
 * const sc = new duplocloud.K8StorageClass("sc", {
 *     tenantId: tenantId,
 *     name: "sc",
 *     storageProvisioner: "efs.csi.aws.com",
 *     reclaimPolicy: "Delete",
 *     volumeBindingMode: "Immediate",
 *     allowVolumeExpansion: false,
 *     parameters: {
 *         fileSystemId: "fs-0d2f79aca4712c6e8",
 *         basePath: "/dynamic_provisioning",
 *         directoryPerms: "700",
 *         gidRangeStart: "1000",
 *         gidRangeEnd: "2000",
 *         provisioningMode: "efs-ap",
 *     },
 *     annotations: {
 *         a1: "v1",
 *         a2: "v2",
 *     },
 *     labels: {
 *         l1: "v1",
 *         l2: "v2",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Example: Importing an existing kubernetes storage class
 *
 *  - *TENANT_ID* is the tenant GUID
 *
 *  - *FULL_NAME* is the Duplo provided storage class name
 *
 * # 
 *
 * ```sh
 * $ pulumi import duplocloud:index/k8StorageClass:K8StorageClass sc v3/subscriptions/*TENANT_ID*&#47;k8s/storageclass/*FULL_NAME*
 * ```
 */
export class K8StorageClass extends pulumi.CustomResource {
    /**
     * Get an existing K8StorageClass resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: K8StorageClassState, opts?: pulumi.CustomResourceOptions): K8StorageClass {
        return new K8StorageClass(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'duplocloud:index/k8StorageClass:K8StorageClass';

    /**
     * Returns true if the given object is an instance of K8StorageClass.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is K8StorageClass {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === K8StorageClass.__pulumiType;
    }

    /**
     * Indicates whether the storage class allow volume expand Defaults to `false`.
     */
    public readonly allowVolumeExpansion!: pulumi.Output<boolean | undefined>;
    /**
     * Restrict the node topologies where volumes can be dynamically provisioned.
     */
    public readonly allowedTopologies!: pulumi.Output<outputs.K8StorageClassAllowedTopologies | undefined>;
    /**
     * An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    /**
     * The full name of the storage class.
     */
    public /*out*/ readonly fullname!: pulumi.Output<string>;
    /**
     * Map of string keys and values that can be used to organize and categorize (scope and select) the service.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The name of the storage class.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parameters for the provisioner that should create volumes of this storage class
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Indicates the type of the reclaim policy Defaults to `Delete`.
     */
    public readonly reclaimPolicy!: pulumi.Output<string | undefined>;
    /**
     * Indicates the type of the provisioner
     */
    public readonly storageProvisioner!: pulumi.Output<string>;
    /**
     * The GUID of the tenant that the storage class will be created in.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Indicates when volume binding and dynamic provisioning should occur Defaults to `Immediate`.
     */
    public readonly volumeBindingMode!: pulumi.Output<string | undefined>;

    /**
     * Create a K8StorageClass resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: K8StorageClassArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: K8StorageClassArgs | K8StorageClassState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as K8StorageClassState | undefined;
            resourceInputs["allowVolumeExpansion"] = state ? state.allowVolumeExpansion : undefined;
            resourceInputs["allowedTopologies"] = state ? state.allowedTopologies : undefined;
            resourceInputs["annotations"] = state ? state.annotations : undefined;
            resourceInputs["fullname"] = state ? state.fullname : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["reclaimPolicy"] = state ? state.reclaimPolicy : undefined;
            resourceInputs["storageProvisioner"] = state ? state.storageProvisioner : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["volumeBindingMode"] = state ? state.volumeBindingMode : undefined;
        } else {
            const args = argsOrState as K8StorageClassArgs | undefined;
            if ((!args || args.storageProvisioner === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageProvisioner'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["allowVolumeExpansion"] = args ? args.allowVolumeExpansion : undefined;
            resourceInputs["allowedTopologies"] = args ? args.allowedTopologies : undefined;
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["reclaimPolicy"] = args ? args.reclaimPolicy : undefined;
            resourceInputs["storageProvisioner"] = args ? args.storageProvisioner : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["volumeBindingMode"] = args ? args.volumeBindingMode : undefined;
            resourceInputs["fullname"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(K8StorageClass.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering K8StorageClass resources.
 */
export interface K8StorageClassState {
    /**
     * Indicates whether the storage class allow volume expand Defaults to `false`.
     */
    allowVolumeExpansion?: pulumi.Input<boolean>;
    /**
     * Restrict the node topologies where volumes can be dynamically provisioned.
     */
    allowedTopologies?: pulumi.Input<inputs.K8StorageClassAllowedTopologies>;
    /**
     * An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The full name of the storage class.
     */
    fullname?: pulumi.Input<string>;
    /**
     * Map of string keys and values that can be used to organize and categorize (scope and select) the service.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the storage class.
     */
    name?: pulumi.Input<string>;
    /**
     * The parameters for the provisioner that should create volumes of this storage class
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the type of the reclaim policy Defaults to `Delete`.
     */
    reclaimPolicy?: pulumi.Input<string>;
    /**
     * Indicates the type of the provisioner
     */
    storageProvisioner?: pulumi.Input<string>;
    /**
     * The GUID of the tenant that the storage class will be created in.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Indicates when volume binding and dynamic provisioning should occur Defaults to `Immediate`.
     */
    volumeBindingMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a K8StorageClass resource.
 */
export interface K8StorageClassArgs {
    /**
     * Indicates whether the storage class allow volume expand Defaults to `false`.
     */
    allowVolumeExpansion?: pulumi.Input<boolean>;
    /**
     * Restrict the node topologies where volumes can be dynamically provisioned.
     */
    allowedTopologies?: pulumi.Input<inputs.K8StorageClassAllowedTopologies>;
    /**
     * An unstructured key value map stored with the storage class that may be used to store arbitrary metadata.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of string keys and values that can be used to organize and categorize (scope and select) the service.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the storage class.
     */
    name?: pulumi.Input<string>;
    /**
     * The parameters for the provisioner that should create volumes of this storage class
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates the type of the reclaim policy Defaults to `Delete`.
     */
    reclaimPolicy?: pulumi.Input<string>;
    /**
     * Indicates the type of the provisioner
     */
    storageProvisioner: pulumi.Input<string>;
    /**
     * The GUID of the tenant that the storage class will be created in.
     */
    tenantId: pulumi.Input<string>;
    /**
     * Indicates when volume binding and dynamic provisioning should occur Defaults to `Immediate`.
     */
    volumeBindingMode?: pulumi.Input<string>;
}
