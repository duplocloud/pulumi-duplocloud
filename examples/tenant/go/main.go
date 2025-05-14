package main

import (
	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Set up the infrastructure.
		tenant, err := duplocloud.NewTenant(ctx, "tenant", &duplocloud.TenantArgs{
			AccountName:   pulumi.String("pulumi03"),
			PlanId:        pulumi.String("nonprod"),
			AllowDeletion: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		ctx.Export("tenantId", tenant.TenantId)
		return nil
	})

}
