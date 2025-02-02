// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `duplocloud.AwsKafkaCluster` manages an AWS MSK cluster in Duplo.
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
 * const mycluster = new duplocloud.AwsKafkaCluster("mycluster", {
 *     tenantId: _this.tenantId,
 *     name: "mycluster",
 *     kafkaVersion: "2.4.1.1",
 *     instanceType: "kafka.m5.large",
 *     storageSize: 20,
 * });
 * ```
 *
 * ## Import
 *
 * Example: Importing an existing AWS Kafka cluster
 *
 *  - *TENANT_ID* is the tenant GUID
 *
 *  - *SHORT_NAME* is the short name of the AWS Kafka cluster
 *
 * # 
 *
 * ```sh
 * $ pulumi import duplocloud:index/awsKafkaCluster:AwsKafkaCluster mycluster *TENANT_ID*&#47;*SHORT_NAME*
 * ```
 */
export class AwsKafkaCluster extends pulumi.CustomResource {
    /**
     * Get an existing AwsKafkaCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AwsKafkaClusterState, opts?: pulumi.CustomResourceOptions): AwsKafkaCluster {
        return new AwsKafkaCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'duplocloud:index/awsKafkaCluster:AwsKafkaCluster';

    /**
     * Returns true if the given object is an instance of AwsKafkaCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AwsKafkaCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AwsKafkaCluster.__pulumiType;
    }

    /**
     * The ARN of the Kafka cluster.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The availability zone distribution used by the cluster.
     */
    public /*out*/ readonly azDistribution!: pulumi.Output<string>;
    /**
     * An ARN of a Kafka configuration to apply to the cluster.
     */
    public readonly configurationArn!: pulumi.Output<string>;
    /**
     * An revision of a Kafka configuration to apply to the cluster.
     */
    public readonly configurationRevision!: pulumi.Output<number>;
    /**
     * Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`
     */
    public readonly encryptionInTransit!: pulumi.Output<string>;
    /**
     * The full name of the Kakfa cluster.
     */
    public /*out*/ readonly fullname!: pulumi.Output<string>;
    /**
     * The Kafka node instance type to use.
     * See the [AWS documentation](https://docs.aws.amazon.com/msk/latest/developerguide/msk-create-cluster.html) for more information.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The version of the Kafka cluster.
     */
    public readonly kafkaVersion!: pulumi.Output<string>;
    /**
     * The short name of the Kafka cluster.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The desired total number of broker nodes in the kafka cluster.
     */
    public /*out*/ readonly numberOfBrokerNodes!: pulumi.Output<number>;
    /**
     * The bootstrap broker connect string for plaintext (unencrypted) connections.
     */
    public /*out*/ readonly plaintextBootstrapBrokerString!: pulumi.Output<string>;
    /**
     * The Zookeeper connect string for plaintext (unencrypted) connections.
     */
    public /*out*/ readonly plaintextZookeeperConnectString!: pulumi.Output<string>;
    /**
     * The list of security groups used by the cluster.
     */
    public /*out*/ readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * The current state of the cluster.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The size of the Kafka storage, in gigabytes.
     */
    public readonly storageSize!: pulumi.Output<number>;
    /**
     * The list of subnets that the cluster will be launched in.
     */
    public readonly subnets!: pulumi.Output<string[]>;
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * The GUID of the tenant that the Kafka cluster will be created in.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The bootstrap broker connect string for TLS (encrypted) connections.
     */
    public /*out*/ readonly tlsBootstrapBrokerString!: pulumi.Output<string>;
    /**
     * The Zookeeper connect string for TLS (encrypted) connections.
     */
    public /*out*/ readonly tlsZookeeperConnectString!: pulumi.Output<string>;

    /**
     * Create a AwsKafkaCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AwsKafkaClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AwsKafkaClusterArgs | AwsKafkaClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AwsKafkaClusterState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["azDistribution"] = state ? state.azDistribution : undefined;
            resourceInputs["configurationArn"] = state ? state.configurationArn : undefined;
            resourceInputs["configurationRevision"] = state ? state.configurationRevision : undefined;
            resourceInputs["encryptionInTransit"] = state ? state.encryptionInTransit : undefined;
            resourceInputs["fullname"] = state ? state.fullname : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["kafkaVersion"] = state ? state.kafkaVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["numberOfBrokerNodes"] = state ? state.numberOfBrokerNodes : undefined;
            resourceInputs["plaintextBootstrapBrokerString"] = state ? state.plaintextBootstrapBrokerString : undefined;
            resourceInputs["plaintextZookeeperConnectString"] = state ? state.plaintextZookeeperConnectString : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["storageSize"] = state ? state.storageSize : undefined;
            resourceInputs["subnets"] = state ? state.subnets : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["tlsBootstrapBrokerString"] = state ? state.tlsBootstrapBrokerString : undefined;
            resourceInputs["tlsZookeeperConnectString"] = state ? state.tlsZookeeperConnectString : undefined;
        } else {
            const args = argsOrState as AwsKafkaClusterArgs | undefined;
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.kafkaVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kafkaVersion'");
            }
            if ((!args || args.storageSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageSize'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["configurationArn"] = args ? args.configurationArn : undefined;
            resourceInputs["configurationRevision"] = args ? args.configurationRevision : undefined;
            resourceInputs["encryptionInTransit"] = args ? args.encryptionInTransit : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["kafkaVersion"] = args ? args.kafkaVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["storageSize"] = args ? args.storageSize : undefined;
            resourceInputs["subnets"] = args ? args.subnets : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["azDistribution"] = undefined /*out*/;
            resourceInputs["fullname"] = undefined /*out*/;
            resourceInputs["numberOfBrokerNodes"] = undefined /*out*/;
            resourceInputs["plaintextBootstrapBrokerString"] = undefined /*out*/;
            resourceInputs["plaintextZookeeperConnectString"] = undefined /*out*/;
            resourceInputs["securityGroups"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["tlsBootstrapBrokerString"] = undefined /*out*/;
            resourceInputs["tlsZookeeperConnectString"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AwsKafkaCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AwsKafkaCluster resources.
 */
export interface AwsKafkaClusterState {
    /**
     * The ARN of the Kafka cluster.
     */
    arn?: pulumi.Input<string>;
    /**
     * The availability zone distribution used by the cluster.
     */
    azDistribution?: pulumi.Input<string>;
    /**
     * An ARN of a Kafka configuration to apply to the cluster.
     */
    configurationArn?: pulumi.Input<string>;
    /**
     * An revision of a Kafka configuration to apply to the cluster.
     */
    configurationRevision?: pulumi.Input<number>;
    /**
     * Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`
     */
    encryptionInTransit?: pulumi.Input<string>;
    /**
     * The full name of the Kakfa cluster.
     */
    fullname?: pulumi.Input<string>;
    /**
     * The Kafka node instance type to use.
     * See the [AWS documentation](https://docs.aws.amazon.com/msk/latest/developerguide/msk-create-cluster.html) for more information.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The version of the Kafka cluster.
     */
    kafkaVersion?: pulumi.Input<string>;
    /**
     * The short name of the Kafka cluster.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
     */
    name?: pulumi.Input<string>;
    /**
     * The desired total number of broker nodes in the kafka cluster.
     */
    numberOfBrokerNodes?: pulumi.Input<number>;
    /**
     * The bootstrap broker connect string for plaintext (unencrypted) connections.
     */
    plaintextBootstrapBrokerString?: pulumi.Input<string>;
    /**
     * The Zookeeper connect string for plaintext (unencrypted) connections.
     */
    plaintextZookeeperConnectString?: pulumi.Input<string>;
    /**
     * The list of security groups used by the cluster.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The current state of the cluster.
     */
    state?: pulumi.Input<string>;
    /**
     * The size of the Kafka storage, in gigabytes.
     */
    storageSize?: pulumi.Input<number>;
    /**
     * The list of subnets that the cluster will be launched in.
     */
    subnets?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The GUID of the tenant that the Kafka cluster will be created in.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The bootstrap broker connect string for TLS (encrypted) connections.
     */
    tlsBootstrapBrokerString?: pulumi.Input<string>;
    /**
     * The Zookeeper connect string for TLS (encrypted) connections.
     */
    tlsZookeeperConnectString?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AwsKafkaCluster resource.
 */
export interface AwsKafkaClusterArgs {
    /**
     * An ARN of a Kafka configuration to apply to the cluster.
     */
    configurationArn?: pulumi.Input<string>;
    /**
     * An revision of a Kafka configuration to apply to the cluster.
     */
    configurationRevision?: pulumi.Input<number>;
    /**
     * Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`
     */
    encryptionInTransit?: pulumi.Input<string>;
    /**
     * The Kafka node instance type to use.
     * See the [AWS documentation](https://docs.aws.amazon.com/msk/latest/developerguide/msk-create-cluster.html) for more information.
     */
    instanceType: pulumi.Input<string>;
    /**
     * The version of the Kafka cluster.
     */
    kafkaVersion: pulumi.Input<string>;
    /**
     * The short name of the Kafka cluster.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
     */
    name?: pulumi.Input<string>;
    /**
     * The size of the Kafka storage, in gigabytes.
     */
    storageSize: pulumi.Input<number>;
    /**
     * The list of subnets that the cluster will be launched in.
     */
    subnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The GUID of the tenant that the Kafka cluster will be created in.
     */
    tenantId: pulumi.Input<string>;
}
