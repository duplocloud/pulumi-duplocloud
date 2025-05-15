using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Duplocloud = DuploCloud.Pulumi;

return await Deployment.RunAsync(() =>
{

    // Create a new infrastructure in AWS
    var infra = new Duplocloud.Infrastructure("pulumi-infra", new Duplocloud.InfrastructureArgs
    {
        InfraName = "pulumi-infra",
        Cloud = 0,
        Region = "us-east-2",
        Azcount = 2,
        EnableK8Cluster = true,
        AddressPrefix = "10.22.0.0/16",
        SubnetCidr = 24,
    });
    infra.VpcId.Apply(id =>
    {
        Pulumi.Log.Info($"VPC ID : {id}");
        return id;
    });

    // Create tenant
    var tenant = new Duplocloud.Tenant("dev", new()
    {
        AccountName = "dev",
        PlanId = infra.InfraName,
    },
    new CustomResourceOptions { DependsOn = { infra } });

    tenant.TenantId.Apply(id =>
    {
        Pulumi.Log.Info($"Tenant ID : {id}");
        return id;
    });
    
    var imageOutput = Duplocloud.GetNativeHostImage.Invoke(new()
    {
        TenantId = tenant.TenantId,
        IsKubernetes = true,
    });
    
     imageOutput.Apply(image =>
    {
        Pulumi.Log.Info($"AMI ID : {image.ImageId}");
        return image.ImageId;
    });
    
    // Create EKS Node
    var host = new Duplocloud.AwsHost("Node01", new()
    {
        TenantId = tenant.TenantId,
        FriendlyName = "Node01",
        ImageId = imageOutput.Apply(image => image.ImageId),
        Capacity = "t3a.small",
        AgentPlatform = 7,
        Zone = 0,
        Metadatas = new[]
        {
            new Duplocloud.Inputs.AwsHostMetadataArgs
            {
                Key = "OsDiskSize",
                Value = "20",
            },
        },
    },
    new CustomResourceOptions { DependsOn = { tenant } });

    // Create EKS Deployment
    var certArn="arn:aws:acm:us-east-1:1234567890:certificate/d6c4138f-583e-4c75-a314-851142670b64";
    var appService = new Duplocloud.DuploService("web-application", new()
    {
        TenantId = tenant.TenantId,
        Name = "web-application",
        AgentPlatform = 7,
        DockerImage = "nginx:latest",
        Replicas = 1,
    },
    new CustomResourceOptions { DependsOn = { host } });

    // Create EKS Service
    var appLBConfigs = new Duplocloud.DuploServiceLbconfigs("web-application-lb", new()
    {
        TenantId = appService.TenantId,
        ReplicationControllerName = appService.Name,
        Lbconfigs = new[]
        {
            new Duplocloud.Inputs.DuploServiceLbconfigsLbconfigArgs
            {
                ExternalPort = 443,
                HealthCheckUrl = "/",
                IsNative = false,
                LbType = 1,
                Port = "80",
                Protocol = "http",
                CertificateArn = certArn,
                HealthCheck = new Duplocloud.Inputs.DuploServiceLbconfigsLbconfigHealthCheckArgs
                {
                    HealthyThreshold = 4,
                    UnhealthyThreshold = 4,
                    Timeout = 10,
                    Interval = 30,
                    HttpSuccessCodes = "200-399",
                },
            },
        },
    },
    new CustomResourceOptions { DependsOn = { appService } });

    return new Dictionary<string, object>
    {
        { "VpcId",  infra.VpcId },
        { "tenantId",  tenant.TenantId },
    };
});