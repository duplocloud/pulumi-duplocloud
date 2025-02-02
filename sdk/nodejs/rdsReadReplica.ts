// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * `duplocloud.RdsReadReplica` manages an AWS RDS read replica in Duplo.
 *
 * ## Import
 *
 * Example: Importing an existing RDS read replica
 *
 *  - *TENANT_ID* is the tenant GUID
 *
 *  - *SHORTNAME* is the short name of the database read replica (without the duplo prefix)
 *
 * # 
 *
 * ```sh
 * $ pulumi import duplocloud:index/rdsReadReplica:RdsReadReplica read_replica v2/subscriptions/*TENANT_ID*&#47;RDSDBInstance/*SHORTNAME*
 * ```
 */
export class RdsReadReplica extends pulumi.CustomResource {
    /**
     * Get an existing RdsReadReplica resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RdsReadReplicaState, opts?: pulumi.CustomResourceOptions): RdsReadReplica {
        return new RdsReadReplica(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'duplocloud:index/rdsReadReplica:RdsReadReplica';

    /**
     * Returns true if the given object is an instance of RdsReadReplica.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RdsReadReplica {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RdsReadReplica.__pulumiType;
    }

    /**
     * The ARN of the RDS read replica.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The AZ for the RDS instance.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The full name of the RDS Cluster.
     */
    public readonly clusterIdentifier!: pulumi.Output<string>;
    /**
     * Parameter group associated with this instance's DB Cluster.
     */
    public /*out*/ readonly clusterParameterGroupName!: pulumi.Output<string>;
    /**
     * Whether or not to enable the RDS instance logging. This setting is not applicable for document db cluster instance.
     */
    public /*out*/ readonly enableLogging!: pulumi.Output<boolean>;
    /**
     * Whether or not to encrypt the RDS instance storage.
     */
    public /*out*/ readonly encryptStorage!: pulumi.Output<boolean>;
    /**
     * The endpoint of the RDS read replica.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The numerical index of database engine to be used the for the RDS read replica.
     */
    public /*out*/ readonly engine!: pulumi.Output<number>;
    /**
     * Engine type required to validate applicable parameter group setting for different instance. Should be referred from writer
     */
    public readonly engineType!: pulumi.Output<number>;
    /**
     * The database engine version to be used the for the RDS read replica.
     */
    public /*out*/ readonly engineVersion!: pulumi.Output<string>;
    /**
     * Interval to capture metrics in real time for the operating system (OS) that your Amazon RDS DB instance runs on.
     */
    public readonly enhancedMonitoring!: pulumi.Output<number>;
    /**
     * The DNS hostname of the RDS read replica.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * The full name of the RDS read replica.
     */
    public /*out*/ readonly identifier!: pulumi.Output<string>;
    /**
     * The globally unique identifier for the key.
     */
    public /*out*/ readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Specifies if the RDS instance is multi-AZ.
     */
    public /*out*/ readonly multiAz!: pulumi.Output<boolean>;
    /**
     * The short name of the RDS read replica.  Duplo will add a prefix to the name.  You can retrieve the full name from the `identifier` attribute.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A RDS parameter group name to apply to the RDS instance.
     */
    public readonly parameterGroupName!: pulumi.Output<string | undefined>;
    /**
     * Amazon RDS Performance Insights is a database performance tuning and monitoring feature that helps you quickly assess the load on your database, and determine when and where to take action. Perfomance Insights get apply when enable is set to true.
     */
    public readonly performanceInsights!: pulumi.Output<outputs.RdsReadReplicaPerformanceInsights>;
    /**
     * The listening port of the RDS read replica.
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * The current status of the RDS read replica.
     */
    public /*out*/ readonly replicaStatus!: pulumi.Output<string>;
    /**
     * The type of the RDS read replica.
     * See AWS documentation for the [available instance types](https://aws.amazon.com/rds/instance-types/).
     */
    public readonly size!: pulumi.Output<string>;
    /**
     * The GUID of the tenant that the RDS read replica will be created in.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a RdsReadReplica resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RdsReadReplicaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RdsReadReplicaArgs | RdsReadReplicaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RdsReadReplicaState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["clusterIdentifier"] = state ? state.clusterIdentifier : undefined;
            resourceInputs["clusterParameterGroupName"] = state ? state.clusterParameterGroupName : undefined;
            resourceInputs["enableLogging"] = state ? state.enableLogging : undefined;
            resourceInputs["encryptStorage"] = state ? state.encryptStorage : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["engineType"] = state ? state.engineType : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["enhancedMonitoring"] = state ? state.enhancedMonitoring : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["identifier"] = state ? state.identifier : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["multiAz"] = state ? state.multiAz : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parameterGroupName"] = state ? state.parameterGroupName : undefined;
            resourceInputs["performanceInsights"] = state ? state.performanceInsights : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["replicaStatus"] = state ? state.replicaStatus : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as RdsReadReplicaArgs | undefined;
            if ((!args || args.clusterIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterIdentifier'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["clusterIdentifier"] = args ? args.clusterIdentifier : undefined;
            resourceInputs["engineType"] = args ? args.engineType : undefined;
            resourceInputs["enhancedMonitoring"] = args ? args.enhancedMonitoring : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameterGroupName"] = args ? args.parameterGroupName : undefined;
            resourceInputs["performanceInsights"] = args ? args.performanceInsights : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["clusterParameterGroupName"] = undefined /*out*/;
            resourceInputs["enableLogging"] = undefined /*out*/;
            resourceInputs["encryptStorage"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["engine"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["identifier"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["multiAz"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["replicaStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RdsReadReplica.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RdsReadReplica resources.
 */
export interface RdsReadReplicaState {
    /**
     * The ARN of the RDS read replica.
     */
    arn?: pulumi.Input<string>;
    /**
     * The AZ for the RDS instance.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The full name of the RDS Cluster.
     */
    clusterIdentifier?: pulumi.Input<string>;
    /**
     * Parameter group associated with this instance's DB Cluster.
     */
    clusterParameterGroupName?: pulumi.Input<string>;
    /**
     * Whether or not to enable the RDS instance logging. This setting is not applicable for document db cluster instance.
     */
    enableLogging?: pulumi.Input<boolean>;
    /**
     * Whether or not to encrypt the RDS instance storage.
     */
    encryptStorage?: pulumi.Input<boolean>;
    /**
     * The endpoint of the RDS read replica.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The numerical index of database engine to be used the for the RDS read replica.
     */
    engine?: pulumi.Input<number>;
    /**
     * Engine type required to validate applicable parameter group setting for different instance. Should be referred from writer
     */
    engineType?: pulumi.Input<number>;
    /**
     * The database engine version to be used the for the RDS read replica.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Interval to capture metrics in real time for the operating system (OS) that your Amazon RDS DB instance runs on.
     */
    enhancedMonitoring?: pulumi.Input<number>;
    /**
     * The DNS hostname of the RDS read replica.
     */
    host?: pulumi.Input<string>;
    /**
     * The full name of the RDS read replica.
     */
    identifier?: pulumi.Input<string>;
    /**
     * The globally unique identifier for the key.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies if the RDS instance is multi-AZ.
     */
    multiAz?: pulumi.Input<boolean>;
    /**
     * The short name of the RDS read replica.  Duplo will add a prefix to the name.  You can retrieve the full name from the `identifier` attribute.
     */
    name?: pulumi.Input<string>;
    /**
     * A RDS parameter group name to apply to the RDS instance.
     */
    parameterGroupName?: pulumi.Input<string>;
    /**
     * Amazon RDS Performance Insights is a database performance tuning and monitoring feature that helps you quickly assess the load on your database, and determine when and where to take action. Perfomance Insights get apply when enable is set to true.
     */
    performanceInsights?: pulumi.Input<inputs.RdsReadReplicaPerformanceInsights>;
    /**
     * The listening port of the RDS read replica.
     */
    port?: pulumi.Input<number>;
    /**
     * The current status of the RDS read replica.
     */
    replicaStatus?: pulumi.Input<string>;
    /**
     * The type of the RDS read replica.
     * See AWS documentation for the [available instance types](https://aws.amazon.com/rds/instance-types/).
     */
    size?: pulumi.Input<string>;
    /**
     * The GUID of the tenant that the RDS read replica will be created in.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RdsReadReplica resource.
 */
export interface RdsReadReplicaArgs {
    /**
     * The AZ for the RDS instance.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The full name of the RDS Cluster.
     */
    clusterIdentifier: pulumi.Input<string>;
    /**
     * Engine type required to validate applicable parameter group setting for different instance. Should be referred from writer
     */
    engineType?: pulumi.Input<number>;
    /**
     * Interval to capture metrics in real time for the operating system (OS) that your Amazon RDS DB instance runs on.
     */
    enhancedMonitoring?: pulumi.Input<number>;
    /**
     * The short name of the RDS read replica.  Duplo will add a prefix to the name.  You can retrieve the full name from the `identifier` attribute.
     */
    name?: pulumi.Input<string>;
    /**
     * A RDS parameter group name to apply to the RDS instance.
     */
    parameterGroupName?: pulumi.Input<string>;
    /**
     * Amazon RDS Performance Insights is a database performance tuning and monitoring feature that helps you quickly assess the load on your database, and determine when and where to take action. Perfomance Insights get apply when enable is set to true.
     */
    performanceInsights?: pulumi.Input<inputs.RdsReadReplicaPerformanceInsights>;
    /**
     * The type of the RDS read replica.
     * See AWS documentation for the [available instance types](https://aws.amazon.com/rds/instance-types/).
     */
    size: pulumi.Input<string>;
    /**
     * The GUID of the tenant that the RDS read replica will be created in.
     */
    tenantId: pulumi.Input<string>;
}
