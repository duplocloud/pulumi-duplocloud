import * as duplocloud from "@duplocloud/pulumi";

const certArn = "arn:aws:acm:us-east-1:1234567890:certificate/d6c4138f-583e-4c75-a314-851142670b64"

// Create DuploCloud infrastructure
const infra = new duplocloud.Infrastructure("infra", {
    infraName: "pulumi-infra",
    cloud: 0,
    region: "us-east-2",
    azcount: 2,
    subnetCidr: 24,
    enableK8Cluster: true,
    addressPrefix: "10.22.0.0/16",
});

export const vpcId = infra.vpcId;

// Create DuploCloud tenant
const tenant = new duplocloud.Tenant("dev", {
    accountName: "dev",
    planId: infra.infraName,
}, { dependsOn: [infra] });

export const tenantId = tenant.tenantId;

const image = duplocloud.getNativeHostImageOutput({
    tenantId: tenant.tenantId,
    isKubernetes: true,
});

// Create DuploCloud node
const node = new duplocloud.AwsHost("Node01", {
    tenantId: tenant.tenantId,
    friendlyName: "Node01",
    imageId: image.imageId,
    capacity: "t3a.medium",
    agentPlatform: 7,
    zone: 0,
    metadatas: [{
        key: "OsDiskSize",
        value: "20",
    }],
}, { dependsOn: [tenant] });

// Create DuploCloud service
const app_service = new duplocloud.DuploService("web-application", {
    tenantId: tenant.tenantId,
    name: "web-app",
    agentPlatform: 7,
    dockerImage: "nginx:latest",
    replicas: 1,
}, { dependsOn: [node] });

// Create DuploCloud load balancer
const appLbconfigs = new duplocloud.DuploServiceLbconfigs("web-application-lbconfigs", {
    tenantId: app_service.tenantId,
    replicationControllerName: app_service.name,
    lbconfigs: [{
        externalPort: 80,
        healthCheckUrl: "/",
        isNative: false,
        lbType: 1,
        port: "80",
        protocol: "http",
        certificateArn: certArn,
        healthCheck: {
            healthyThreshold: 4,
            unhealthyThreshold: 4,
            timeout: 50,
            interval: 30,
            httpSuccessCodes: "200-399",
        },
    }],
});