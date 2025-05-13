import * as duplocloud from "@duplocloud/pulumi";

const tenant = new duplocloud.Tenant("tenant", {
    accountName: "pulumi02",
    planId: "nonprod",
});

export const tenantId = tenant.tenantId;