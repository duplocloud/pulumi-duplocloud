// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export function getTenantAwsKmsKeys(args: GetTenantAwsKmsKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetTenantAwsKmsKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("duplocloud:index/getTenantAwsKmsKeys:getTenantAwsKmsKeys", {
        "selectable": args.selectable,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTenantAwsKmsKeys.
 */
export interface GetTenantAwsKmsKeysArgs {
    /**
     * Defaults to `true`.
     */
    selectable?: boolean;
    tenantId: string;
}

/**
 * A collection of values returned by getTenantAwsKmsKeys.
 */
export interface GetTenantAwsKmsKeysResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keys: outputs.GetTenantAwsKmsKeysKey[];
    /**
     * Defaults to `true`.
     */
    readonly selectable?: boolean;
    readonly tenantId: string;
}
export function getTenantAwsKmsKeysOutput(args: GetTenantAwsKmsKeysOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTenantAwsKmsKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("duplocloud:index/getTenantAwsKmsKeys:getTenantAwsKmsKeys", {
        "selectable": args.selectable,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTenantAwsKmsKeys.
 */
export interface GetTenantAwsKmsKeysOutputArgs {
    /**
     * Defaults to `true`.
     */
    selectable?: pulumi.Input<boolean>;
    tenantId: pulumi.Input<string>;
}
