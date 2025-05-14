package main

import (
	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Set up the infrastructure.

		infra, err := duplocloud.NewInfrastructure(ctx, "infra", &duplocloud.InfrastructureArgs{
			InfraName:       pulumi.String("pulumi-infra"),
			Cloud:           pulumi.Int(0),
			Region:          pulumi.String("us-east-2"),
			Azcount:         pulumi.Int(2),
			SubnetCidr:      pulumi.Int(24),
			EnableK8Cluster: pulumi.Bool(true),
			AddressPrefix:   pulumi.String("10.22.0.0/16"),
		})
		if err != nil {
			return err
		}
		ctx.Export("infraName", infra.InfraName)
		ctx.Export("vpcId", infra.VpcId)

		// Create DuploCloud Tenant
		tenant, err := duplocloud.NewTenant(ctx, "dev", &duplocloud.TenantArgs{
			AccountName:   pulumi.String("dev"),
			PlanId:        infra.InfraName,
			AllowDeletion: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{infra}))
		if err != nil {
			return err
		}
		ctx.Export("tenantId", tenant.ID())

		// Get native image
		image := duplocloud.GetNativeHostImageOutput(ctx, duplocloud.GetNativeHostImageOutputArgs{
			TenantId:     tenant.TenantId,
			IsKubernetes: pulumi.Bool(true),
		}, nil)

		// Create EKS Node
		node, err := duplocloud.NewAwsHost(ctx, "node", &duplocloud.AwsHostArgs{
			TenantId:      tenant.TenantId,
			FriendlyName:  pulumi.String("node01"),
			ImageId:       image.ImageId(),
			Capacity:      pulumi.String("t3a.medium"),
			AgentPlatform: pulumi.Int(7),
			Zone:          pulumi.Int(0),
			Metadatas: duplocloud.AwsHostMetadataArray{
				&duplocloud.AwsHostMetadataArgs{
					Key:   pulumi.String("OsDiskSize"),
					Value: pulumi.String("20"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{tenant, infra}))
		if err != nil {
			return err
		}

		// Create a service
		webapp, err := duplocloud.NewDuploService(ctx, "web-app", &duplocloud.DuploServiceArgs{
			TenantId:      tenant.TenantId,
			Name:          pulumi.String("web-app"),
			AgentPlatform: pulumi.Int(7),
			DockerImage:   pulumi.String("nginx:latest"),
			Replicas:      pulumi.Int(1),
		}, pulumi.DependsOn([]pulumi.Resource{node, tenant, infra}))
		if err != nil {
			return err
		}

		// Configure Load Balancer
		cert_arn := "arn:aws:acm:us-east-1:1234567890:certificate/d6c4138f-583e-4c75-a314-851142670b64"
		_, err = duplocloud.NewDuploServiceLbconfigs(ctx, "web-app-lb", &duplocloud.DuploServiceLbconfigsArgs{
			TenantId:                  webapp.TenantId,
			ReplicationControllerName: webapp.Name,
			Lbconfigs: duplocloud.DuploServiceLbconfigsLbconfigArray{
				&duplocloud.DuploServiceLbconfigsLbconfigArgs{
					ExternalPort:   pulumi.Int(443),
					HealthCheckUrl: pulumi.String("/"),
					IsNative:       pulumi.Bool(false),
					LbType:         pulumi.Int(1),
					Port:           pulumi.String("80"),
					Protocol:       pulumi.String("http"),
					CertificateArn: pulumi.String(cert_arn),
					HealthCheck: &duplocloud.DuploServiceLbconfigsLbconfigHealthCheckArgs{
						HealthyThreshold:   pulumi.Int(4),
						UnhealthyThreshold: pulumi.Int(4),
						Timeout:            pulumi.Int(50),
						Interval:           pulumi.Int(30),
						HttpSuccessCodes:   pulumi.String("200-399"),
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{webapp, node, tenant, infra}))
		if err != nil {
			return err
		}
		return nil
	})
}
