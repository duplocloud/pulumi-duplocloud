import * as duplocloud from "@duplocloud/pulumi";

const tenantName = "pulumi02"
const certArn = "arn:aws:acm:us-east-1:1234567890:certificate/d6c4138f-583e-4c75-a314-851142670b64"
const tenant = duplocloud.getTenantOutput({
    name: tenantName,
});

const myservice = new duplocloud.DuploService("web-application", {
    tenantId: tenant.id,
    name: "web-app",
    agentPlatform: 0,
    dockerImage: "nginx:latest",
    replicas: 1,
});

const appLbconfigs = new duplocloud.DuploServiceLbconfigs("web-application-lbconfigs", {
    tenantId: myservice.tenantId,
    replicationControllerName: myservice.name,
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