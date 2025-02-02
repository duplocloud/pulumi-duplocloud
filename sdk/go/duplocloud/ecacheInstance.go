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

// ## Example Usage
//
// ### Create an Amazon ElastiCache cluster of type Redis.
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
//			// Before creating a ElastiCache cluster, you must first set up the infrastructure and tenant. Below is the resource for creating the infrastructure.
//			infra, err := duplocloud.NewInfrastructure(ctx, "infra", &duplocloud.InfrastructureArgs{
//				InfraName:       pulumi.String("prod"),
//				Cloud:           pulumi.Int(0),
//				Region:          pulumi.String("us-west-2"),
//				EnableK8Cluster: pulumi.Bool(false),
//				AddressPrefix:   pulumi.String("10.11.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			// Use the infrastructure name as the 'plan_id' from the 'duplocloud_infrastructure' resource while creating tenant.
//			tenant, err := duplocloud.NewTenant(ctx, "tenant", &duplocloud.TenantArgs{
//				AccountName: pulumi.String("prod"),
//				PlanId:      infra.InfraName,
//			})
//			if err != nil {
//				return err
//			}
//			// Use the tenant_id from the duplocloud_tenant, which will be populated after the tenant resource is created, when setting up the Redis ElastiCache cluster.
//			_, err = duplocloud.NewEcacheInstance(ctx, "redis_cache", &duplocloud.EcacheInstanceArgs{
//				TenantId:          tenant.TenantId,
//				Name:              pulumi.String("mycache"),
//				CacheType:         pulumi.Int(0),
//				Replicas:          pulumi.Int(1),
//				Size:              pulumi.String("cache.t2.small"),
//				EnableClusterMode: pulumi.Bool(true),
//				NumberOfShards:    pulumi.Int(1),
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
// ### Create an Amazon ElastiCache cluster of type Redis with 2 replicas of type cache.t2.small in dev tenant.
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
//			// Assuming the 'dev' tenant is already created, use a data source to fetch the tenant ID.
//			tenant, err := duplocloud.LookupTenant(ctx, &duplocloud.LookupTenantArgs{
//				Name: pulumi.StringRef("dev"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Use the tenant_id from the duplocloud_tenant, which will be populated after the tenant resource is created, when setting up the Redis ElastiCache cluster.
//			_, err = duplocloud.NewEcacheInstance(ctx, "redis_cache", &duplocloud.EcacheInstanceArgs{
//				TenantId:          pulumi.String(tenant.Id),
//				Name:              pulumi.String("mycache"),
//				CacheType:         pulumi.Int(0),
//				Replicas:          pulumi.Int(2),
//				Size:              pulumi.String("cache.t2.small"),
//				EnableClusterMode: pulumi.Bool(true),
//				NumberOfShards:    pulumi.Int(1),
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
// ### Create an Amazon ElastiCache of type Redis with log delivery configuration and automatic failover enabled in dev tenant.
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
//			// Assuming the 'dev' tenant is already created, use a data source to fetch the tenant ID.
//			tenant, err := duplocloud.LookupTenant(ctx, &duplocloud.LookupTenantArgs{
//				Name: pulumi.StringRef("dev"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Use the tenant_id from the duplocloud_tenant, which will be populated after the tenant resource is created, when setting up the Redis ElastiCache.
//			_, err = duplocloud.NewEcacheInstance(ctx, "redis_cache", &duplocloud.EcacheInstanceArgs{
//				TenantId:                 pulumi.String(tenant.Id),
//				Name:                     pulumi.String("mycache"),
//				CacheType:                pulumi.Int(0),
//				Replicas:                 pulumi.Int(2),
//				Size:                     pulumi.String("cache.t2.small"),
//				AutomaticFailoverEnabled: pulumi.Bool(true),
//				LogDeliveryConfigurations: duplocloud.EcacheInstanceLogDeliveryConfigurationArray{
//					&duplocloud.EcacheInstanceLogDeliveryConfigurationArgs{
//						LogGroup:        pulumi.String("/elasticache/redis"),
//						DestinationType: pulumi.String("cloudwatch-logs"),
//						LogFormat:       pulumi.String("text"),
//						LogType:         pulumi.String("slow-log"),
//					},
//					&duplocloud.EcacheInstanceLogDeliveryConfigurationArgs{
//						LogGroup:        pulumi.String("/elasticache/redis"),
//						DestinationType: pulumi.String("cloudwatch-logs"),
//						LogFormat:       pulumi.String("json"),
//						LogType:         pulumi.String("engine-log"),
//					},
//				},
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
// ### Set up an ElastiCache Redis cluster with 2 shards and 2 cache.t2.small replicas in the dev tenant.
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
//			// Assuming the 'dev' tenant is already created, use a data source to fetch the tenant ID.
//			tenant, err := duplocloud.LookupTenant(ctx, &duplocloud.LookupTenantArgs{
//				Name: pulumi.StringRef("dev"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Use the tenant_id from the duplocloud_tenant, which will be populated after the tenant resource is created, when setting up the Redis ElastiCache cluster.
//			_, err = duplocloud.NewEcacheInstance(ctx, "redis_cache", &duplocloud.EcacheInstanceArgs{
//				TenantId:          pulumi.String(tenant.Id),
//				Name:              pulumi.String("mycache"),
//				CacheType:         pulumi.Int(0),
//				Replicas:          pulumi.Int(2),
//				Size:              pulumi.String("cache.t2.small"),
//				EnableClusterMode: pulumi.Bool(true),
//				NumberOfShards:    pulumi.Int(2),
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
// ### Create an Amazon ElastiCache cluster of type Memcached.
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
//			// Assuming the 'dev' tenant is already created, use a data source to fetch the tenant ID.
//			tenant, err := duplocloud.LookupTenant(ctx, &duplocloud.LookupTenantArgs{
//				Name: pulumi.StringRef("dev"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// Use the tenant_id from the duplocloud_tenant data source, which will be populated after the tenant data source is created, when setting up the Memcached ElastiCache cluster.
//			_, err = duplocloud.NewEcacheInstance(ctx, "mem_cache", &duplocloud.EcacheInstanceArgs{
//				TenantId:  pulumi.String(tenant.Id),
//				Name:      pulumi.String("mycache"),
//				CacheType: pulumi.Int(1),
//				Replicas:  pulumi.Int(1),
//				Size:      pulumi.String("cache.t2.small"),
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
// ### Create an Amazon ElastiCache with snapshot window
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
//			_, err := duplocloud.NewEcacheInstance(ctx, "mycaches", &duplocloud.EcacheInstanceArgs{
//				TenantId:       pulumi.Any(tenant.Id),
//				Name:           pulumi.String("mycache"),
//				CacheType:      pulumi.Int(0),
//				Replicas:       pulumi.Int(2),
//				Size:           pulumi.String("cache.t3.small"),
//				EngineVersion:  pulumi.String("7.1"),
//				SnapshotWindow: pulumi.String("04:00-13:00"),
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
// Example: Importing an existing AWS ElastiCache cluster
//
//   - *TENANT_ID* is the tenant GUID
//
//   - *SHORT_NAME* is the short name of the AWS ElastiCache cluster
//
// #
//
// ```sh
// $ pulumi import duplocloud:index/ecacheInstance:EcacheInstance mycluster v2/subscriptions/*TENANT_ID*/ECacheDBInstance/*SHORT_NAME*
// ```
type EcacheInstance struct {
	pulumi.CustomResourceState

	// The ARN of the elasticache instance.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Set a password for authenticating to the ElastiCache instance.  Only supported if `encryptionInTransit` is to to `true`.
	AuthToken pulumi.StringPtrOutput `pulumi:"authToken"`
	// Enables automatic failover.
	AutomaticFailoverEnabled pulumi.BoolPtrOutput `pulumi:"automaticFailoverEnabled"`
	// The numerical index of elasticache instance type. Should be one of: - `0` : Redis - `1` : Memcache
	CacheType pulumi.IntPtrOutput `pulumi:"cacheType"`
	// Flag to enable/disable redis cluster mode.
	EnableClusterMode pulumi.BoolOutput `pulumi:"enableClusterMode"`
	// Enables encryption-at-rest.
	EncryptionAtRest pulumi.BoolPtrOutput `pulumi:"encryptionAtRest"`
	// Enables encryption-in-transit.
	EncryptionInTransit pulumi.BoolPtrOutput `pulumi:"encryptionInTransit"`
	// The endpoint of the elasticache instance.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The engine version of the elastic instance. See AWS documentation for the [available Redis instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/supported-engine-versions.html) or the [available
	// Memcached instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/supported-engine-versions-mc.html).
	EngineVersion pulumi.StringPtrOutput `pulumi:"engineVersion"`
	// The DNS hostname of the elasticache instance.
	Host pulumi.StringOutput `pulumi:"host"`
	// The full name of the elasticache instance.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// The status of the elasticache instance.
	InstanceStatus pulumi.StringOutput `pulumi:"instanceStatus"`
	// The globally unique identifier for the key.
	KmsKeyId                  pulumi.StringOutput                               `pulumi:"kmsKeyId"`
	LogDeliveryConfigurations EcacheInstanceLogDeliveryConfigurationArrayOutput `pulumi:"logDeliveryConfigurations"`
	// The short name of the elasticache instance.  Duplo will add a prefix to the name.  You can retrieve the full name from the `identifier` attribute.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of shards to create. Applicable only if enableClusterMode is set to true
	NumberOfShards pulumi.IntOutput `pulumi:"numberOfShards"`
	// The REDIS parameter group to supply.
	ParameterGroupName pulumi.StringOutput `pulumi:"parameterGroupName"`
	// The listening port of the elasticache instance.
	Port pulumi.IntOutput `pulumi:"port"`
	// The number of replicas to create. Supported number of replicas is 1 to 6
	Replicas pulumi.IntPtrOutput `pulumi:"replicas"`
	// The instance type of the elasticache instance.
	// See AWS documentation for the [available instance types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html).
	Size pulumi.StringOutput `pulumi:"size"`
	// Specify the ARN of a Redis RDB snapshot file stored in Amazon S3. User should have the access to export snapshot to s3
	// bucket. One can find steps to provide access to export snapshot to s3 on following link
	// https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html
	SnapshotArns pulumi.StringArrayOutput `pulumi:"snapshotArns"`
	// Select the snapshot/backup you want to use for creating redis.
	SnapshotName pulumi.StringOutput `pulumi:"snapshotName"`
	// Specify retention limit in days. Accepted values - 1-35.
	SnapshotRetentionLimit pulumi.IntOutput `pulumi:"snapshotRetentionLimit"`
	// Specify snapshot window limit The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of
	// your node group (shard). Example: 05:00-09:00. If you do not specify this parameter, ElastiCache automatically chooses
	// an appropriate time range.
	SnapshotWindow pulumi.StringOutput `pulumi:"snapshotWindow"`
	// The GUID of the tenant that the elasticache instance will be created in.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewEcacheInstance registers a new resource with the given unique name, arguments, and options.
func NewEcacheInstance(ctx *pulumi.Context,
	name string, args *EcacheInstanceArgs, opts ...pulumi.ResourceOption) (*EcacheInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EcacheInstance
	err := ctx.RegisterResource("duplocloud:index/ecacheInstance:EcacheInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEcacheInstance gets an existing EcacheInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEcacheInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EcacheInstanceState, opts ...pulumi.ResourceOption) (*EcacheInstance, error) {
	var resource EcacheInstance
	err := ctx.ReadResource("duplocloud:index/ecacheInstance:EcacheInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EcacheInstance resources.
type ecacheInstanceState struct {
	// The ARN of the elasticache instance.
	Arn *string `pulumi:"arn"`
	// Set a password for authenticating to the ElastiCache instance.  Only supported if `encryptionInTransit` is to to `true`.
	AuthToken *string `pulumi:"authToken"`
	// Enables automatic failover.
	AutomaticFailoverEnabled *bool `pulumi:"automaticFailoverEnabled"`
	// The numerical index of elasticache instance type. Should be one of: - `0` : Redis - `1` : Memcache
	CacheType *int `pulumi:"cacheType"`
	// Flag to enable/disable redis cluster mode.
	EnableClusterMode *bool `pulumi:"enableClusterMode"`
	// Enables encryption-at-rest.
	EncryptionAtRest *bool `pulumi:"encryptionAtRest"`
	// Enables encryption-in-transit.
	EncryptionInTransit *bool `pulumi:"encryptionInTransit"`
	// The endpoint of the elasticache instance.
	Endpoint *string `pulumi:"endpoint"`
	// The engine version of the elastic instance. See AWS documentation for the [available Redis instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/supported-engine-versions.html) or the [available
	// Memcached instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/supported-engine-versions-mc.html).
	EngineVersion *string `pulumi:"engineVersion"`
	// The DNS hostname of the elasticache instance.
	Host *string `pulumi:"host"`
	// The full name of the elasticache instance.
	Identifier *string `pulumi:"identifier"`
	// The status of the elasticache instance.
	InstanceStatus *string `pulumi:"instanceStatus"`
	// The globally unique identifier for the key.
	KmsKeyId                  *string                                  `pulumi:"kmsKeyId"`
	LogDeliveryConfigurations []EcacheInstanceLogDeliveryConfiguration `pulumi:"logDeliveryConfigurations"`
	// The short name of the elasticache instance.  Duplo will add a prefix to the name.  You can retrieve the full name from the `identifier` attribute.
	Name *string `pulumi:"name"`
	// The number of shards to create. Applicable only if enableClusterMode is set to true
	NumberOfShards *int `pulumi:"numberOfShards"`
	// The REDIS parameter group to supply.
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// The listening port of the elasticache instance.
	Port *int `pulumi:"port"`
	// The number of replicas to create. Supported number of replicas is 1 to 6
	Replicas *int `pulumi:"replicas"`
	// The instance type of the elasticache instance.
	// See AWS documentation for the [available instance types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html).
	Size *string `pulumi:"size"`
	// Specify the ARN of a Redis RDB snapshot file stored in Amazon S3. User should have the access to export snapshot to s3
	// bucket. One can find steps to provide access to export snapshot to s3 on following link
	// https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html
	SnapshotArns []string `pulumi:"snapshotArns"`
	// Select the snapshot/backup you want to use for creating redis.
	SnapshotName *string `pulumi:"snapshotName"`
	// Specify retention limit in days. Accepted values - 1-35.
	SnapshotRetentionLimit *int `pulumi:"snapshotRetentionLimit"`
	// Specify snapshot window limit The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of
	// your node group (shard). Example: 05:00-09:00. If you do not specify this parameter, ElastiCache automatically chooses
	// an appropriate time range.
	SnapshotWindow *string `pulumi:"snapshotWindow"`
	// The GUID of the tenant that the elasticache instance will be created in.
	TenantId *string `pulumi:"tenantId"`
}

type EcacheInstanceState struct {
	// The ARN of the elasticache instance.
	Arn pulumi.StringPtrInput
	// Set a password for authenticating to the ElastiCache instance.  Only supported if `encryptionInTransit` is to to `true`.
	AuthToken pulumi.StringPtrInput
	// Enables automatic failover.
	AutomaticFailoverEnabled pulumi.BoolPtrInput
	// The numerical index of elasticache instance type. Should be one of: - `0` : Redis - `1` : Memcache
	CacheType pulumi.IntPtrInput
	// Flag to enable/disable redis cluster mode.
	EnableClusterMode pulumi.BoolPtrInput
	// Enables encryption-at-rest.
	EncryptionAtRest pulumi.BoolPtrInput
	// Enables encryption-in-transit.
	EncryptionInTransit pulumi.BoolPtrInput
	// The endpoint of the elasticache instance.
	Endpoint pulumi.StringPtrInput
	// The engine version of the elastic instance. See AWS documentation for the [available Redis instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/supported-engine-versions.html) or the [available
	// Memcached instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/supported-engine-versions-mc.html).
	EngineVersion pulumi.StringPtrInput
	// The DNS hostname of the elasticache instance.
	Host pulumi.StringPtrInput
	// The full name of the elasticache instance.
	Identifier pulumi.StringPtrInput
	// The status of the elasticache instance.
	InstanceStatus pulumi.StringPtrInput
	// The globally unique identifier for the key.
	KmsKeyId                  pulumi.StringPtrInput
	LogDeliveryConfigurations EcacheInstanceLogDeliveryConfigurationArrayInput
	// The short name of the elasticache instance.  Duplo will add a prefix to the name.  You can retrieve the full name from the `identifier` attribute.
	Name pulumi.StringPtrInput
	// The number of shards to create. Applicable only if enableClusterMode is set to true
	NumberOfShards pulumi.IntPtrInput
	// The REDIS parameter group to supply.
	ParameterGroupName pulumi.StringPtrInput
	// The listening port of the elasticache instance.
	Port pulumi.IntPtrInput
	// The number of replicas to create. Supported number of replicas is 1 to 6
	Replicas pulumi.IntPtrInput
	// The instance type of the elasticache instance.
	// See AWS documentation for the [available instance types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html).
	Size pulumi.StringPtrInput
	// Specify the ARN of a Redis RDB snapshot file stored in Amazon S3. User should have the access to export snapshot to s3
	// bucket. One can find steps to provide access to export snapshot to s3 on following link
	// https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html
	SnapshotArns pulumi.StringArrayInput
	// Select the snapshot/backup you want to use for creating redis.
	SnapshotName pulumi.StringPtrInput
	// Specify retention limit in days. Accepted values - 1-35.
	SnapshotRetentionLimit pulumi.IntPtrInput
	// Specify snapshot window limit The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of
	// your node group (shard). Example: 05:00-09:00. If you do not specify this parameter, ElastiCache automatically chooses
	// an appropriate time range.
	SnapshotWindow pulumi.StringPtrInput
	// The GUID of the tenant that the elasticache instance will be created in.
	TenantId pulumi.StringPtrInput
}

func (EcacheInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ecacheInstanceState)(nil)).Elem()
}

type ecacheInstanceArgs struct {
	// Set a password for authenticating to the ElastiCache instance.  Only supported if `encryptionInTransit` is to to `true`.
	AuthToken *string `pulumi:"authToken"`
	// Enables automatic failover.
	AutomaticFailoverEnabled *bool `pulumi:"automaticFailoverEnabled"`
	// The numerical index of elasticache instance type. Should be one of: - `0` : Redis - `1` : Memcache
	CacheType *int `pulumi:"cacheType"`
	// Flag to enable/disable redis cluster mode.
	EnableClusterMode *bool `pulumi:"enableClusterMode"`
	// Enables encryption-at-rest.
	EncryptionAtRest *bool `pulumi:"encryptionAtRest"`
	// Enables encryption-in-transit.
	EncryptionInTransit *bool `pulumi:"encryptionInTransit"`
	// The engine version of the elastic instance. See AWS documentation for the [available Redis instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/supported-engine-versions.html) or the [available
	// Memcached instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/supported-engine-versions-mc.html).
	EngineVersion *string `pulumi:"engineVersion"`
	// The globally unique identifier for the key.
	KmsKeyId                  *string                                  `pulumi:"kmsKeyId"`
	LogDeliveryConfigurations []EcacheInstanceLogDeliveryConfiguration `pulumi:"logDeliveryConfigurations"`
	// The short name of the elasticache instance.  Duplo will add a prefix to the name.  You can retrieve the full name from the `identifier` attribute.
	Name *string `pulumi:"name"`
	// The number of shards to create. Applicable only if enableClusterMode is set to true
	NumberOfShards *int `pulumi:"numberOfShards"`
	// The REDIS parameter group to supply.
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// The number of replicas to create. Supported number of replicas is 1 to 6
	Replicas *int `pulumi:"replicas"`
	// The instance type of the elasticache instance.
	// See AWS documentation for the [available instance types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html).
	Size string `pulumi:"size"`
	// Specify the ARN of a Redis RDB snapshot file stored in Amazon S3. User should have the access to export snapshot to s3
	// bucket. One can find steps to provide access to export snapshot to s3 on following link
	// https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html
	SnapshotArns []string `pulumi:"snapshotArns"`
	// Select the snapshot/backup you want to use for creating redis.
	SnapshotName *string `pulumi:"snapshotName"`
	// Specify retention limit in days. Accepted values - 1-35.
	SnapshotRetentionLimit *int `pulumi:"snapshotRetentionLimit"`
	// Specify snapshot window limit The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of
	// your node group (shard). Example: 05:00-09:00. If you do not specify this parameter, ElastiCache automatically chooses
	// an appropriate time range.
	SnapshotWindow *string `pulumi:"snapshotWindow"`
	// The GUID of the tenant that the elasticache instance will be created in.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a EcacheInstance resource.
type EcacheInstanceArgs struct {
	// Set a password for authenticating to the ElastiCache instance.  Only supported if `encryptionInTransit` is to to `true`.
	AuthToken pulumi.StringPtrInput
	// Enables automatic failover.
	AutomaticFailoverEnabled pulumi.BoolPtrInput
	// The numerical index of elasticache instance type. Should be one of: - `0` : Redis - `1` : Memcache
	CacheType pulumi.IntPtrInput
	// Flag to enable/disable redis cluster mode.
	EnableClusterMode pulumi.BoolPtrInput
	// Enables encryption-at-rest.
	EncryptionAtRest pulumi.BoolPtrInput
	// Enables encryption-in-transit.
	EncryptionInTransit pulumi.BoolPtrInput
	// The engine version of the elastic instance. See AWS documentation for the [available Redis instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/supported-engine-versions.html) or the [available
	// Memcached instance
	// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/supported-engine-versions-mc.html).
	EngineVersion pulumi.StringPtrInput
	// The globally unique identifier for the key.
	KmsKeyId                  pulumi.StringPtrInput
	LogDeliveryConfigurations EcacheInstanceLogDeliveryConfigurationArrayInput
	// The short name of the elasticache instance.  Duplo will add a prefix to the name.  You can retrieve the full name from the `identifier` attribute.
	Name pulumi.StringPtrInput
	// The number of shards to create. Applicable only if enableClusterMode is set to true
	NumberOfShards pulumi.IntPtrInput
	// The REDIS parameter group to supply.
	ParameterGroupName pulumi.StringPtrInput
	// The number of replicas to create. Supported number of replicas is 1 to 6
	Replicas pulumi.IntPtrInput
	// The instance type of the elasticache instance.
	// See AWS documentation for the [available instance types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html).
	Size pulumi.StringInput
	// Specify the ARN of a Redis RDB snapshot file stored in Amazon S3. User should have the access to export snapshot to s3
	// bucket. One can find steps to provide access to export snapshot to s3 on following link
	// https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html
	SnapshotArns pulumi.StringArrayInput
	// Select the snapshot/backup you want to use for creating redis.
	SnapshotName pulumi.StringPtrInput
	// Specify retention limit in days. Accepted values - 1-35.
	SnapshotRetentionLimit pulumi.IntPtrInput
	// Specify snapshot window limit The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of
	// your node group (shard). Example: 05:00-09:00. If you do not specify this parameter, ElastiCache automatically chooses
	// an appropriate time range.
	SnapshotWindow pulumi.StringPtrInput
	// The GUID of the tenant that the elasticache instance will be created in.
	TenantId pulumi.StringInput
}

func (EcacheInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ecacheInstanceArgs)(nil)).Elem()
}

type EcacheInstanceInput interface {
	pulumi.Input

	ToEcacheInstanceOutput() EcacheInstanceOutput
	ToEcacheInstanceOutputWithContext(ctx context.Context) EcacheInstanceOutput
}

func (*EcacheInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**EcacheInstance)(nil)).Elem()
}

func (i *EcacheInstance) ToEcacheInstanceOutput() EcacheInstanceOutput {
	return i.ToEcacheInstanceOutputWithContext(context.Background())
}

func (i *EcacheInstance) ToEcacheInstanceOutputWithContext(ctx context.Context) EcacheInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcacheInstanceOutput)
}

// EcacheInstanceArrayInput is an input type that accepts EcacheInstanceArray and EcacheInstanceArrayOutput values.
// You can construct a concrete instance of `EcacheInstanceArrayInput` via:
//
//	EcacheInstanceArray{ EcacheInstanceArgs{...} }
type EcacheInstanceArrayInput interface {
	pulumi.Input

	ToEcacheInstanceArrayOutput() EcacheInstanceArrayOutput
	ToEcacheInstanceArrayOutputWithContext(context.Context) EcacheInstanceArrayOutput
}

type EcacheInstanceArray []EcacheInstanceInput

func (EcacheInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcacheInstance)(nil)).Elem()
}

func (i EcacheInstanceArray) ToEcacheInstanceArrayOutput() EcacheInstanceArrayOutput {
	return i.ToEcacheInstanceArrayOutputWithContext(context.Background())
}

func (i EcacheInstanceArray) ToEcacheInstanceArrayOutputWithContext(ctx context.Context) EcacheInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcacheInstanceArrayOutput)
}

// EcacheInstanceMapInput is an input type that accepts EcacheInstanceMap and EcacheInstanceMapOutput values.
// You can construct a concrete instance of `EcacheInstanceMapInput` via:
//
//	EcacheInstanceMap{ "key": EcacheInstanceArgs{...} }
type EcacheInstanceMapInput interface {
	pulumi.Input

	ToEcacheInstanceMapOutput() EcacheInstanceMapOutput
	ToEcacheInstanceMapOutputWithContext(context.Context) EcacheInstanceMapOutput
}

type EcacheInstanceMap map[string]EcacheInstanceInput

func (EcacheInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcacheInstance)(nil)).Elem()
}

func (i EcacheInstanceMap) ToEcacheInstanceMapOutput() EcacheInstanceMapOutput {
	return i.ToEcacheInstanceMapOutputWithContext(context.Background())
}

func (i EcacheInstanceMap) ToEcacheInstanceMapOutputWithContext(ctx context.Context) EcacheInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EcacheInstanceMapOutput)
}

type EcacheInstanceOutput struct{ *pulumi.OutputState }

func (EcacheInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EcacheInstance)(nil)).Elem()
}

func (o EcacheInstanceOutput) ToEcacheInstanceOutput() EcacheInstanceOutput {
	return o
}

func (o EcacheInstanceOutput) ToEcacheInstanceOutputWithContext(ctx context.Context) EcacheInstanceOutput {
	return o
}

// The ARN of the elasticache instance.
func (o EcacheInstanceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Set a password for authenticating to the ElastiCache instance.  Only supported if `encryptionInTransit` is to to `true`.
func (o EcacheInstanceOutput) AuthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringPtrOutput { return v.AuthToken }).(pulumi.StringPtrOutput)
}

// Enables automatic failover.
func (o EcacheInstanceOutput) AutomaticFailoverEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.BoolPtrOutput { return v.AutomaticFailoverEnabled }).(pulumi.BoolPtrOutput)
}

// The numerical index of elasticache instance type. Should be one of: - `0` : Redis - `1` : Memcache
func (o EcacheInstanceOutput) CacheType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.IntPtrOutput { return v.CacheType }).(pulumi.IntPtrOutput)
}

// Flag to enable/disable redis cluster mode.
func (o EcacheInstanceOutput) EnableClusterMode() pulumi.BoolOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.BoolOutput { return v.EnableClusterMode }).(pulumi.BoolOutput)
}

// Enables encryption-at-rest.
func (o EcacheInstanceOutput) EncryptionAtRest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.BoolPtrOutput { return v.EncryptionAtRest }).(pulumi.BoolPtrOutput)
}

// Enables encryption-in-transit.
func (o EcacheInstanceOutput) EncryptionInTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.BoolPtrOutput { return v.EncryptionInTransit }).(pulumi.BoolPtrOutput)
}

// The endpoint of the elasticache instance.
func (o EcacheInstanceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The engine version of the elastic instance. See AWS documentation for the [available Redis instance
// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/supported-engine-versions.html) or the [available
// Memcached instance
// types](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/supported-engine-versions-mc.html).
func (o EcacheInstanceOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringPtrOutput { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

// The DNS hostname of the elasticache instance.
func (o EcacheInstanceOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The full name of the elasticache instance.
func (o EcacheInstanceOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// The status of the elasticache instance.
func (o EcacheInstanceOutput) InstanceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.InstanceStatus }).(pulumi.StringOutput)
}

// The globally unique identifier for the key.
func (o EcacheInstanceOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.KmsKeyId }).(pulumi.StringOutput)
}

func (o EcacheInstanceOutput) LogDeliveryConfigurations() EcacheInstanceLogDeliveryConfigurationArrayOutput {
	return o.ApplyT(func(v *EcacheInstance) EcacheInstanceLogDeliveryConfigurationArrayOutput {
		return v.LogDeliveryConfigurations
	}).(EcacheInstanceLogDeliveryConfigurationArrayOutput)
}

// The short name of the elasticache instance.  Duplo will add a prefix to the name.  You can retrieve the full name from the `identifier` attribute.
func (o EcacheInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of shards to create. Applicable only if enableClusterMode is set to true
func (o EcacheInstanceOutput) NumberOfShards() pulumi.IntOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.IntOutput { return v.NumberOfShards }).(pulumi.IntOutput)
}

// The REDIS parameter group to supply.
func (o EcacheInstanceOutput) ParameterGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.ParameterGroupName }).(pulumi.StringOutput)
}

// The listening port of the elasticache instance.
func (o EcacheInstanceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The number of replicas to create. Supported number of replicas is 1 to 6
func (o EcacheInstanceOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.IntPtrOutput { return v.Replicas }).(pulumi.IntPtrOutput)
}

// The instance type of the elasticache instance.
// See AWS documentation for the [available instance types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html).
func (o EcacheInstanceOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

// Specify the ARN of a Redis RDB snapshot file stored in Amazon S3. User should have the access to export snapshot to s3
// bucket. One can find steps to provide access to export snapshot to s3 on following link
// https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html
func (o EcacheInstanceOutput) SnapshotArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringArrayOutput { return v.SnapshotArns }).(pulumi.StringArrayOutput)
}

// Select the snapshot/backup you want to use for creating redis.
func (o EcacheInstanceOutput) SnapshotName() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.SnapshotName }).(pulumi.StringOutput)
}

// Specify retention limit in days. Accepted values - 1-35.
func (o EcacheInstanceOutput) SnapshotRetentionLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.IntOutput { return v.SnapshotRetentionLimit }).(pulumi.IntOutput)
}

// Specify snapshot window limit The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of
// your node group (shard). Example: 05:00-09:00. If you do not specify this parameter, ElastiCache automatically chooses
// an appropriate time range.
func (o EcacheInstanceOutput) SnapshotWindow() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.SnapshotWindow }).(pulumi.StringOutput)
}

// The GUID of the tenant that the elasticache instance will be created in.
func (o EcacheInstanceOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *EcacheInstance) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type EcacheInstanceArrayOutput struct{ *pulumi.OutputState }

func (EcacheInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EcacheInstance)(nil)).Elem()
}

func (o EcacheInstanceArrayOutput) ToEcacheInstanceArrayOutput() EcacheInstanceArrayOutput {
	return o
}

func (o EcacheInstanceArrayOutput) ToEcacheInstanceArrayOutputWithContext(ctx context.Context) EcacheInstanceArrayOutput {
	return o
}

func (o EcacheInstanceArrayOutput) Index(i pulumi.IntInput) EcacheInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EcacheInstance {
		return vs[0].([]*EcacheInstance)[vs[1].(int)]
	}).(EcacheInstanceOutput)
}

type EcacheInstanceMapOutput struct{ *pulumi.OutputState }

func (EcacheInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EcacheInstance)(nil)).Elem()
}

func (o EcacheInstanceMapOutput) ToEcacheInstanceMapOutput() EcacheInstanceMapOutput {
	return o
}

func (o EcacheInstanceMapOutput) ToEcacheInstanceMapOutputWithContext(ctx context.Context) EcacheInstanceMapOutput {
	return o
}

func (o EcacheInstanceMapOutput) MapIndex(k pulumi.StringInput) EcacheInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EcacheInstance {
		return vs[0].(map[string]*EcacheInstance)[vs[1].(string)]
	}).(EcacheInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EcacheInstanceInput)(nil)).Elem(), &EcacheInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcacheInstanceArrayInput)(nil)).Elem(), EcacheInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EcacheInstanceMapInput)(nil)).Elem(), EcacheInstanceMap{})
	pulumi.RegisterOutputType(EcacheInstanceOutput{})
	pulumi.RegisterOutputType(EcacheInstanceArrayOutput{})
	pulumi.RegisterOutputType(EcacheInstanceMapOutput{})
}
