// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `duplocloud.getEcrRepository` get ecr repository in a Duplo tenant.
 */
export function getEcrRepository(args: GetEcrRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetEcrRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("duplocloud:index/getEcrRepository:getEcrRepository", {
        "name": args.name,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getEcrRepository.
 */
export interface GetEcrRepositoryArgs {
    /**
     * The name of the ECR Repository.
     */
    name: string;
    /**
     * The GUID of the tenant in which to list the hosts.
     */
    tenantId: string;
}

/**
 * A collection of values returned by getEcrRepository.
 */
export interface GetEcrRepositoryResult {
    readonly arn: string;
    readonly enableScanImageOnPush: boolean;
    readonly enableTagImmutability: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly kmsEncryptionKey: string;
    /**
     * The name of the ECR Repository.
     */
    readonly name: string;
    readonly registryId: string;
    readonly repositoryUrl: string;
    /**
     * The GUID of the tenant in which to list the hosts.
     */
    readonly tenantId: string;
}
/**
 * `duplocloud.getEcrRepository` get ecr repository in a Duplo tenant.
 */
export function getEcrRepositoryOutput(args: GetEcrRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEcrRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("duplocloud:index/getEcrRepository:getEcrRepository", {
        "name": args.name,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getEcrRepository.
 */
export interface GetEcrRepositoryOutputArgs {
    /**
     * The name of the ECR Repository.
     */
    name: pulumi.Input<string>;
    /**
     * The GUID of the tenant in which to list the hosts.
     */
    tenantId: pulumi.Input<string>;
}
