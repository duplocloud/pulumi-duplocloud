import * as pulumi from "@pulumi/pulumi";
import * as duplocloud from "@duplocloud/pulumi";

const resource = new duplocloud.Resource("Resource", { sampleAttribute: "attr" });

export const sampleAttribute = resource.sampleAttribute;
