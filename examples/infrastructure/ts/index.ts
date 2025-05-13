import * as duplocloud from "@duplocloud/pulumi";

const infra = new duplocloud.Infrastructure("infra", {
    infraName: "nonprod",
    cloud: 0,
    region: "us-west-2",
    enableK8Cluster: true,
    addressPrefix: "10.11.0.0/16",
});

export const vpcId = infra.vpcId;