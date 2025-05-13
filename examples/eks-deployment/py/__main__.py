"""
DuploCloud EKS Deployment Example
This example demonstrates how to use the pulumi_duplocloud SDK to create DuploCloud EKS deployment.
"""

import pulumi
import pulumi_duplocloud as duplo

# Create DuploCloud Infrastructure
infra = duplo.infrastructure.Infrastructure(resource_name="pulumi-infra",
    infra_name="pulumi-infra",
    cloud=0,
    region="us-east-2",
    azcount=2,
    subnet_cidr=24,
    address_prefix="10.22.0.0/16",
    enable_k8_cluster=True,  # Enable Kubernetes
)

# Export the outputs
pulumi.export("vpc_id", infra.vpc_id)

# Create Tenant
tenant = duplo.tenant.Tenant(resource_name="dev",
    account_name="dev",
    plan_id=infra.infra_name,
)

# Export the outputs
pulumi.export("tenantId", tenant.tenant_id)

image = duplo.get_native_host_image_output(tenant_id=tenant.tenant_id,
            is_kubernetes=True)

# Create EKS Node
node = duplo.AwsHost(resource_name="Node01",
    tenant_id=tenant.tenant_id,
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

# Deploy nginx service on EKS

cert_arn="arn:aws:acm:us-east-1:1234567890:certificate/d6c4138f-583e-4c75-a314-851142670b64"
app_service = duplo.duplo_service.DuploService(resource_name="web-application",
    opts=pulumi.ResourceOptions(depends_on=[node]),
    tenant_id=tenant.tenant_id,
    name="web-app",
    docker_image="nginx:latest",
    replicas=1,
    agent_platform=7,
)

# Configure Load Balancer
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
        "certificate_arn":  cert_arn,
        "health_check": {
            "healthy_threshold": 4,
            "unhealthy_threshold": 4,
            "timeout": 10,
            "interval": 30,
            "http_success_codes": "200-399",
        },
    }]
)
