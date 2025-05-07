"""
DuploCloud Infrastructure Example
This example demonstrates how to use the pulumi_duplocloud SDK to create various
DuploCloud infrastructure.
"""

import pulumi
import pulumi_duplocloud as duplo

# export duplo_host="https://qa-aws.duplocloud.net"
# export duplo_token="AQAAANCMnd8BFdERjHoAwE_Cl-sBAAAA0rFNCGCvrEOyv29SBXAqQgAAAAACAAAAAAAQZgAAAAEAACAAAADpXh1DBl0K4FqPGUc_JGlIzJvRLmCMu3a98Bl-YGi0XwAAAAAOgAAAAAIAACAAAACK81o1C2sKUZVN9iooY1ZSr-nwLd43xoudSmaYOlGUucAAAAB4nYxsiApbHRxf4_zxn-muFxzU1bpW4Va2P-1XpY7HV1inURBjoWrd01AYKFiC7bksdEw6aIVCLvT9BjZlcAFNeOlpYg1ZJ7I8C-myNiuIqYZ7E8LslePvknJjbNe-HGRwC8Y1XMe4HHNuFE2yOMz8CVLw0aaX4yKPHjIn3F7CQuPNTh09tbrYVw9-2HM30YzWTciWYey-Vq0whRB_Ng05bQxIqNes52dd7AQvL9ubj1fYHr0aaK2_llR3M4iJMh1AAAAAXjd4mya6g7N7ewLhKqyXxamffR2bMZdYLCiWVGO6mA4_VOVp6xb5TmysIPPQgNhzyPFwcs-dZEgJRoQ036cU1A"
# Create nonprod Infrastructure resource
infrastructure = duplo.infrastructure.Infrastructure(resource_name="nonprod",
    infra_name="nonprod",
    cloud=0,
    region="us-west-2",
    azcount=2,
    subnet_cidr=24,
    address_prefix="10.22.0.0/16",
    enable_k8_cluster=True,  # Enable Kubernetes
)

# Export the outputs
pulumi.export("vpc_id", infrastructure.vpc_id)