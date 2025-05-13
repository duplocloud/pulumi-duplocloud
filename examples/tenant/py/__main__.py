"""
DuploCloud Tenant Example
This example demonstrates how to use the pulumi_duplocloud SDK to create DuploCloud tenant.
"""

import pulumi
import pulumi_duplocloud as duplo

tenant_name = "pulumi01"
plan_id = "default"
tenant = duplo.tenant.Tenant(resource_name="dev",
    account_name=tenant_name,
    plan_id=plan_id,
)

# Export the outputs
pulumi.export("tenantId", tenant.tenant_id)