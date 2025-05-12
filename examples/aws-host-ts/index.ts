import * as duplocloud from "@duplocloud/pulumi";

const tenantName = "pulumi02"

const tenant = duplocloud.getTenantOutput({
    name: tenantName,
});

const image = duplocloud.getNativeHostImageOutput({
    tenantId: tenant.id,
    isKubernetes: true,
});

const host = new duplocloud.AwsHost("Node01", {
    tenantId: tenant.id,
    friendlyName: "Node01",
    imageId: image.imageId,
    capacity: "t3a.medium",
    agentPlatform: 7,
    zone: 0,
    metadatas: [{
        key: "OsDiskSize",
        value: "20",
    }],
});

export const hostId = host.instanceId;
