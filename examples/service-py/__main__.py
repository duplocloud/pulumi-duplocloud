"""
DuploCloud Service Example
This example demonstrates how to use the pulumi_duplocloud SDK to create DuploCloud service.
"""

import pulumi
import pulumi_duplocloud as duplo

tenant_name = "pulumi01"
plan_id = "argocd01"
cert_arn="arn:aws:acm:us-east-1:884446924812:certificate/d6c4138f-583e-4c75-a314-851142670b64"
tenant = duplo.tenant.Tenant(resource_name="dev",
    account_name=tenant_name,
    plan_id=plan_id,
)

# Export the outputs
pulumi.export("tenantId", tenant.tenant_id)

app_service = duplo.duplo_service.DuploService(resource_name="web-application",
    tenant_id=tenant.tenant_id,
    name="web-app",
    docker_image="nginx:latest",
    replicas=1,
    agent_platform=7,
)

app_lbconfigs = duplo.DuploServiceLbconfigs(resource_name="web-application-lb",
            tenant_id=tenant.tenant_id,
            replication_controller_name=app_service.name,
            lbconfigs=[{
                "external_port": 443,
                "health_check_url": "/",
                "is_native": False,
                "lb_type": 1,
                "port": "80",
                "protocol": "http",
                "certificate_arn": cert_arn,
                "health_check": {
                    "healthy_threshold": 4,
                    "unhealthy_threshold": 4,
                    "timeout": 10,
                    "interval": 30,
                    "http_success_codes": "200-399",
                },
            }])