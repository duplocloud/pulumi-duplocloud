"""
DuploCloud Host Example
This example demonstrates how to use the pulumi_duplocloud SDK to create DuploCloud host.
"""

import pulumi
import pulumi_duplocloud as duplo

tenant_name="pulumi01"
tenant = duplo.get_tenant(name=tenant_name)
image = duplo.get_native_host_image_output(tenant_id=tenant.id,
            is_kubernetes=True)

host = duplo.AwsHost(resource_name="Node01",
    tenant_id=tenant.id,
    friendly_name="Node01",
    image_id=image.image_id,
    capacity="t3a.medium",
    agent_platform=7,
    zone=0,
    metadatas=[{
        "key": "OsDiskSize",
        "value": "20",
    }]
)