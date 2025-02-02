// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package duplocloud

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "duplocloud:index/adminSystemSetting:AdminSystemSetting":
		r = &AdminSystemSetting{}
	case "duplocloud:index/asgProfile:AsgProfile":
		r = &AsgProfile{}
	case "duplocloud:index/awsApiGatewayIntegration:AwsApiGatewayIntegration":
		r = &AwsApiGatewayIntegration{}
	case "duplocloud:index/awsApigatewayEvent:AwsApigatewayEvent":
		r = &AwsApigatewayEvent{}
	case "duplocloud:index/awsAppautoscalingPolicy:AwsAppautoscalingPolicy":
		r = &AwsAppautoscalingPolicy{}
	case "duplocloud:index/awsAppautoscalingTarget:AwsAppautoscalingTarget":
		r = &AwsAppautoscalingTarget{}
	case "duplocloud:index/awsBatchComputeEnvironment:AwsBatchComputeEnvironment":
		r = &AwsBatchComputeEnvironment{}
	case "duplocloud:index/awsBatchJobDefinition:AwsBatchJobDefinition":
		r = &AwsBatchJobDefinition{}
	case "duplocloud:index/awsBatchJobQueue:AwsBatchJobQueue":
		r = &AwsBatchJobQueue{}
	case "duplocloud:index/awsBatchSchedulingPolicy:AwsBatchSchedulingPolicy":
		r = &AwsBatchSchedulingPolicy{}
	case "duplocloud:index/awsCloudfrontDistribution:AwsCloudfrontDistribution":
		r = &AwsCloudfrontDistribution{}
	case "duplocloud:index/awsCloudwatchEventRule:AwsCloudwatchEventRule":
		r = &AwsCloudwatchEventRule{}
	case "duplocloud:index/awsCloudwatchEventTarget:AwsCloudwatchEventTarget":
		r = &AwsCloudwatchEventTarget{}
	case "duplocloud:index/awsCloudwatchMetricAlarm:AwsCloudwatchMetricAlarm":
		r = &AwsCloudwatchMetricAlarm{}
	case "duplocloud:index/awsDynamodbTable:AwsDynamodbTable":
		r = &AwsDynamodbTable{}
	case "duplocloud:index/awsDynamodbTableV2:AwsDynamodbTableV2":
		r = &AwsDynamodbTableV2{}
	case "duplocloud:index/awsEcrRepository:AwsEcrRepository":
		r = &AwsEcrRepository{}
	case "duplocloud:index/awsEfsFileSystem:AwsEfsFileSystem":
		r = &AwsEfsFileSystem{}
	case "duplocloud:index/awsEfsLifecyclePolicy:AwsEfsLifecyclePolicy":
		r = &AwsEfsLifecyclePolicy{}
	case "duplocloud:index/awsElasticsearch:AwsElasticsearch":
		r = &AwsElasticsearch{}
	case "duplocloud:index/awsHost:AwsHost":
		r = &AwsHost{}
	case "duplocloud:index/awsKafkaCluster:AwsKafkaCluster":
		r = &AwsKafkaCluster{}
	case "duplocloud:index/awsLambdaFunction:AwsLambdaFunction":
		r = &AwsLambdaFunction{}
	case "duplocloud:index/awsLambdaFunctionEventConfig:AwsLambdaFunctionEventConfig":
		r = &AwsLambdaFunctionEventConfig{}
	case "duplocloud:index/awsLambdaPermission:AwsLambdaPermission":
		r = &AwsLambdaPermission{}
	case "duplocloud:index/awsLaunchTemplate:AwsLaunchTemplate":
		r = &AwsLaunchTemplate{}
	case "duplocloud:index/awsLaunchTemplateDefaultVersion:AwsLaunchTemplateDefaultVersion":
		r = &AwsLaunchTemplateDefaultVersion{}
	case "duplocloud:index/awsLbListenerRule:AwsLbListenerRule":
		r = &AwsLbListenerRule{}
	case "duplocloud:index/awsLbTargetGroup:AwsLbTargetGroup":
		r = &AwsLbTargetGroup{}
	case "duplocloud:index/awsLoadBalancer:AwsLoadBalancer":
		r = &AwsLoadBalancer{}
	case "duplocloud:index/awsLoadBalancerListener:AwsLoadBalancerListener":
		r = &AwsLoadBalancerListener{}
	case "duplocloud:index/awsMwaaEnvironment:AwsMwaaEnvironment":
		r = &AwsMwaaEnvironment{}
	case "duplocloud:index/awsRdsTag:AwsRdsTag":
		r = &AwsRdsTag{}
	case "duplocloud:index/awsSnsTopic:AwsSnsTopic":
		r = &AwsSnsTopic{}
	case "duplocloud:index/awsSqsQueue:AwsSqsQueue":
		r = &AwsSqsQueue{}
	case "duplocloud:index/awsSsmParameter:AwsSsmParameter":
		r = &AwsSsmParameter{}
	case "duplocloud:index/awsTargetGroupAttributes:AwsTargetGroupAttributes":
		r = &AwsTargetGroupAttributes{}
	case "duplocloud:index/awsTimestreamwriteDatabase:AwsTimestreamwriteDatabase":
		r = &AwsTimestreamwriteDatabase{}
	case "duplocloud:index/awsTimestreamwriteTable:AwsTimestreamwriteTable":
		r = &AwsTimestreamwriteTable{}
	case "duplocloud:index/azureAvailabilitySet:AzureAvailabilitySet":
		r = &AzureAvailabilitySet{}
	case "duplocloud:index/azureDatafactory:AzureDatafactory":
		r = &AzureDatafactory{}
	case "duplocloud:index/azureK8NodePool:AzureK8NodePool":
		r = &AzureK8NodePool{}
	case "duplocloud:index/azureK8sCluster:AzureK8sCluster":
		r = &AzureK8sCluster{}
	case "duplocloud:index/azureKeyVaultSecret:AzureKeyVaultSecret":
		r = &AzureKeyVaultSecret{}
	case "duplocloud:index/azureLogAnalyticsWorkspace:AzureLogAnalyticsWorkspace":
		r = &AzureLogAnalyticsWorkspace{}
	case "duplocloud:index/azureMssqlDatabase:AzureMssqlDatabase":
		r = &AzureMssqlDatabase{}
	case "duplocloud:index/azureMssqlElasticpool:AzureMssqlElasticpool":
		r = &AzureMssqlElasticpool{}
	case "duplocloud:index/azureMssqlServer:AzureMssqlServer":
		r = &AzureMssqlServer{}
	case "duplocloud:index/azureMysqlDatabase:AzureMysqlDatabase":
		r = &AzureMysqlDatabase{}
	case "duplocloud:index/azureNetworkSecurityRule:AzureNetworkSecurityRule":
		r = &AzureNetworkSecurityRule{}
	case "duplocloud:index/azurePostgresqlDatabase:AzurePostgresqlDatabase":
		r = &AzurePostgresqlDatabase{}
	case "duplocloud:index/azurePostgresqlFlexibleDatabase:AzurePostgresqlFlexibleDatabase":
		r = &AzurePostgresqlFlexibleDatabase{}
	case "duplocloud:index/azurePrivateEndpoint:AzurePrivateEndpoint":
		r = &AzurePrivateEndpoint{}
	case "duplocloud:index/azureRecoveryServicesVault:AzureRecoveryServicesVault":
		r = &AzureRecoveryServicesVault{}
	case "duplocloud:index/azureRedisCache:AzureRedisCache":
		r = &AzureRedisCache{}
	case "duplocloud:index/azureSqlFirewallRule:AzureSqlFirewallRule":
		r = &AzureSqlFirewallRule{}
	case "duplocloud:index/azureSqlManagedDatabase:AzureSqlManagedDatabase":
		r = &AzureSqlManagedDatabase{}
	case "duplocloud:index/azureSqlVirtualNetworkRule:AzureSqlVirtualNetworkRule":
		r = &AzureSqlVirtualNetworkRule{}
	case "duplocloud:index/azureStorageAccount:AzureStorageAccount":
		r = &AzureStorageAccount{}
	case "duplocloud:index/azureStorageShareFile:AzureStorageShareFile":
		r = &AzureStorageShareFile{}
	case "duplocloud:index/azureStorageclassBlob:AzureStorageclassBlob":
		r = &AzureStorageclassBlob{}
	case "duplocloud:index/azureStorageclassQueue:AzureStorageclassQueue":
		r = &AzureStorageclassQueue{}
	case "duplocloud:index/azureStorageclassTable:AzureStorageclassTable":
		r = &AzureStorageclassTable{}
	case "duplocloud:index/azureTenantKeyVault:AzureTenantKeyVault":
		r = &AzureTenantKeyVault{}
	case "duplocloud:index/azureTenantKeyVaultSecret:AzureTenantKeyVaultSecret":
		r = &AzureTenantKeyVaultSecret{}
	case "duplocloud:index/azureVaultBackupPolicy:AzureVaultBackupPolicy":
		r = &AzureVaultBackupPolicy{}
	case "duplocloud:index/azureVirtualMachine:AzureVirtualMachine":
		r = &AzureVirtualMachine{}
	case "duplocloud:index/azureVirtualMachineScaleSet:AzureVirtualMachineScaleSet":
		r = &AzureVirtualMachineScaleSet{}
	case "duplocloud:index/azureVmFeature:AzureVmFeature":
		r = &AzureVmFeature{}
	case "duplocloud:index/azureVmMaintenanceConfiguration:AzureVmMaintenanceConfiguration":
		r = &AzureVmMaintenanceConfiguration{}
	case "duplocloud:index/byoh:Byoh":
		r = &Byoh{}
	case "duplocloud:index/dockerCredentials:DockerCredentials":
		r = &DockerCredentials{}
	case "duplocloud:index/duploService:DuploService":
		r = &DuploService{}
	case "duplocloud:index/duploServiceLbconfigs:DuploServiceLbconfigs":
		r = &DuploServiceLbconfigs{}
	case "duplocloud:index/duploServiceParams:DuploServiceParams":
		r = &DuploServiceParams{}
	case "duplocloud:index/ecacheInstance:EcacheInstance":
		r = &EcacheInstance{}
	case "duplocloud:index/ecsService:EcsService":
		r = &EcsService{}
	case "duplocloud:index/ecsTaskDefinition:EcsTaskDefinition":
		r = &EcsTaskDefinition{}
	case "duplocloud:index/emrCluster:EmrCluster":
		r = &EmrCluster{}
	case "duplocloud:index/gcpCloudFunction:GcpCloudFunction":
		r = &GcpCloudFunction{}
	case "duplocloud:index/gcpFirestore:GcpFirestore":
		r = &GcpFirestore{}
	case "duplocloud:index/gcpHost:GcpHost":
		r = &GcpHost{}
	case "duplocloud:index/gcpInfraMaintenanceWindow:GcpInfraMaintenanceWindow":
		r = &GcpInfraMaintenanceWindow{}
	case "duplocloud:index/gcpInfraSecurityRule:GcpInfraSecurityRule":
		r = &GcpInfraSecurityRule{}
	case "duplocloud:index/gcpNodePool:GcpNodePool":
		r = &GcpNodePool{}
	case "duplocloud:index/gcpPubsubTopic:GcpPubsubTopic":
		r = &GcpPubsubTopic{}
	case "duplocloud:index/gcpRedisInstance:GcpRedisInstance":
		r = &GcpRedisInstance{}
	case "duplocloud:index/gcpSchedulerJob:GcpSchedulerJob":
		r = &GcpSchedulerJob{}
	case "duplocloud:index/gcpSqlDatabaseInstance:GcpSqlDatabaseInstance":
		r = &GcpSqlDatabaseInstance{}
	case "duplocloud:index/gcpStorageBucket:GcpStorageBucket":
		r = &GcpStorageBucket{}
	case "duplocloud:index/gcpStorageBucketV2:GcpStorageBucketV2":
		r = &GcpStorageBucketV2{}
	case "duplocloud:index/gcpTenantSecurityRule:GcpTenantSecurityRule":
		r = &GcpTenantSecurityRule{}
	case "duplocloud:index/infrastructure:Infrastructure":
		r = &Infrastructure{}
	case "duplocloud:index/infrastructureOnprem:InfrastructureOnprem":
		r = &InfrastructureOnprem{}
	case "duplocloud:index/infrastructureSetting:InfrastructureSetting":
		r = &InfrastructureSetting{}
	case "duplocloud:index/infrastructureSubnet:InfrastructureSubnet":
		r = &InfrastructureSubnet{}
	case "duplocloud:index/k8ConfigMap:K8ConfigMap":
		r = &K8ConfigMap{}
	case "duplocloud:index/k8HelmRelease:K8HelmRelease":
		r = &K8HelmRelease{}
	case "duplocloud:index/k8HelmRepository:K8HelmRepository":
		r = &K8HelmRepository{}
	case "duplocloud:index/k8Ingress:K8Ingress":
		r = &K8Ingress{}
	case "duplocloud:index/k8PersistentVolumeClaim:K8PersistentVolumeClaim":
		r = &K8PersistentVolumeClaim{}
	case "duplocloud:index/k8Secret:K8Secret":
		r = &K8Secret{}
	case "duplocloud:index/k8SecretProviderClass:K8SecretProviderClass":
		r = &K8SecretProviderClass{}
	case "duplocloud:index/k8StorageClass:K8StorageClass":
		r = &K8StorageClass{}
	case "duplocloud:index/k8sCronJob:K8sCronJob":
		r = &K8sCronJob{}
	case "duplocloud:index/k8sJob:K8sJob":
		r = &K8sJob{}
	case "duplocloud:index/ociContainerengineNodePool:OciContainerengineNodePool":
		r = &OciContainerengineNodePool{}
	case "duplocloud:index/otherAgents:OtherAgents":
		r = &OtherAgents{}
	case "duplocloud:index/planCertificates:PlanCertificates":
		r = &PlanCertificates{}
	case "duplocloud:index/planConfigs:PlanConfigs":
		r = &PlanConfigs{}
	case "duplocloud:index/planImages:PlanImages":
		r = &PlanImages{}
	case "duplocloud:index/planKms:PlanKms":
		r = &PlanKms{}
	case "duplocloud:index/planKmsV2:PlanKmsV2":
		r = &PlanKmsV2{}
	case "duplocloud:index/planSettings:PlanSettings":
		r = &PlanSettings{}
	case "duplocloud:index/planWaf:PlanWaf":
		r = &PlanWaf{}
	case "duplocloud:index/planWafV2:PlanWafV2":
		r = &PlanWafV2{}
	case "duplocloud:index/rdsInstance:RdsInstance":
		r = &RdsInstance{}
	case "duplocloud:index/rdsReadReplica:RdsReadReplica":
		r = &RdsReadReplica{}
	case "duplocloud:index/s3Bucket:S3Bucket":
		r = &S3Bucket{}
	case "duplocloud:index/s3BucketReplication:S3BucketReplication":
		r = &S3BucketReplication{}
	case "duplocloud:index/tenant:Tenant":
		r = &Tenant{}
	case "duplocloud:index/tenantAccessGrant:TenantAccessGrant":
		r = &TenantAccessGrant{}
	case "duplocloud:index/tenantCleanupTimers:TenantCleanupTimers":
		r = &TenantCleanupTimers{}
	case "duplocloud:index/tenantConfig:TenantConfig":
		r = &TenantConfig{}
	case "duplocloud:index/tenantNetworkSecurityRule:TenantNetworkSecurityRule":
		r = &TenantNetworkSecurityRule{}
	case "duplocloud:index/tenantSecret:TenantSecret":
		r = &TenantSecret{}
	case "duplocloud:index/tenantTag:TenantTag":
		r = &TenantTag{}
	case "duplocloud:index/user:User":
		r = &User{}
	case "duplocloud:index/userTenantAccess:UserTenantAccess":
		r = &UserTenantAccess{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:duplocloud" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/adminSystemSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/asgProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsApiGatewayIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsApigatewayEvent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsAppautoscalingPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsAppautoscalingTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsBatchComputeEnvironment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsBatchJobDefinition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsBatchJobQueue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsBatchSchedulingPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsCloudfrontDistribution",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsCloudwatchEventRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsCloudwatchEventTarget",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsCloudwatchMetricAlarm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsDynamodbTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsDynamodbTableV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsEcrRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsEfsFileSystem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsEfsLifecyclePolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsElasticsearch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsHost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsKafkaCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsLambdaFunction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsLambdaFunctionEventConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsLambdaPermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsLaunchTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsLaunchTemplateDefaultVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsLbListenerRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsLbTargetGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsLoadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsLoadBalancerListener",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsMwaaEnvironment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsRdsTag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsSnsTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsSqsQueue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsSsmParameter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsTargetGroupAttributes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsTimestreamwriteDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/awsTimestreamwriteTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureAvailabilitySet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureDatafactory",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureK8NodePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureK8sCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureKeyVaultSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureLogAnalyticsWorkspace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureMssqlDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureMssqlElasticpool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureMssqlServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureMysqlDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureNetworkSecurityRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azurePostgresqlDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azurePostgresqlFlexibleDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azurePrivateEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureRecoveryServicesVault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureRedisCache",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureSqlFirewallRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureSqlManagedDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureSqlVirtualNetworkRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureStorageAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureStorageShareFile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureStorageclassBlob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureStorageclassQueue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureStorageclassTable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureTenantKeyVault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureTenantKeyVaultSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureVaultBackupPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureVirtualMachine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureVirtualMachineScaleSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureVmFeature",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/azureVmMaintenanceConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/byoh",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/dockerCredentials",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/duploService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/duploServiceLbconfigs",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/duploServiceParams",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/ecacheInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/ecsService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/ecsTaskDefinition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/emrCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpCloudFunction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpFirestore",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpHost",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpInfraMaintenanceWindow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpInfraSecurityRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpNodePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpPubsubTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpRedisInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpSchedulerJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpSqlDatabaseInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpStorageBucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpStorageBucketV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/gcpTenantSecurityRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/infrastructure",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/infrastructureOnprem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/infrastructureSetting",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/infrastructureSubnet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8ConfigMap",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8HelmRelease",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8HelmRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8Ingress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8PersistentVolumeClaim",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8Secret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8SecretProviderClass",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8StorageClass",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8sCronJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/k8sJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/ociContainerengineNodePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/otherAgents",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/planCertificates",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/planConfigs",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/planImages",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/planKms",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/planKmsV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/planSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/planWaf",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/planWafV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/rdsInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/rdsReadReplica",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/s3Bucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/s3BucketReplication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/tenant",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/tenantAccessGrant",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/tenantCleanupTimers",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/tenantConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/tenantNetworkSecurityRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/tenantSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/tenantTag",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"duplocloud",
		"index/userTenantAccess",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"duplocloud",
		&pkg{version},
	)
}
