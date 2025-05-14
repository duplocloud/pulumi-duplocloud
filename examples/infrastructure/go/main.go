package main

import (
	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Set up the infrastructure.
		_, err := duplocloud.NewInfrastructure(ctx, "infra", &duplocloud.InfrastructureArgs{
			InfraName:       pulumi.String("pulumi"),
			Cloud:           pulumi.Int(0),
			Region:          pulumi.String("us-east-2"),
			Azcount:         pulumi.Int(2),
			SubnetCidr:      pulumi.Int(24),
			EnableK8Cluster: pulumi.Bool(true),
			AddressPrefix:   pulumi.String("10.5.0.0/16"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
