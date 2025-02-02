// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `duplocloud.AzureSqlManagedDatabase` manages an azure sql managed database in Duplo.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as duplocloud from "@pulumi/duplocloud";
 *
 * const myapp = new duplocloud.Tenant("myapp", {
 *     accountName: "myapp",
 *     planId: "default",
 * });
 * const mydb = new duplocloud.AzureSqlManagedDatabase("mydb", {
 *     tenantId: myapp.tenantId,
 *     name: "db-test",
 *     administratorLogin: "testroot",
 *     administratorLoginPassword: "P@ssword12345",
 *     minimumTlsVersion: "1.2",
 *     skuName: "GP_Gen5",
 *     vcores: 4,
 *     storageSizeInGb: 32,
 *     subnetId: "/subscriptions/0c84b91e-95f5-409e-9cff-6c2e60affbb3/resourceGroups/duploinfra-demo/providers/Microsoft.Network/virtualNetworks/demo/subnets/duploinfra-default",
 * });
 * ```
 *
 * ## Import
 *
 * Example: Importing an existing Azure Sql Managed Database
 *
 *  - *TENANT_ID* is the tenant GUID
 *
 *  - *SHORT_NAME* is the short name of the Azure Sql Managed Database
 *
 * # 
 *
 * ```sh
 * $ pulumi import duplocloud:index/azureSqlManagedDatabase:AzureSqlManagedDatabase myManagedSQLDatabase *TENANT_ID*&#47;*SHORT_NAME*
 * ```
 */
export class AzureSqlManagedDatabase extends pulumi.CustomResource {
    /**
     * Get an existing AzureSqlManagedDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureSqlManagedDatabaseState, opts?: pulumi.CustomResourceOptions): AzureSqlManagedDatabase {
        return new AzureSqlManagedDatabase(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'duplocloud:index/azureSqlManagedDatabase:AzureSqlManagedDatabase';

    /**
     * Returns true if the given object is an instance of AzureSqlManagedDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureSqlManagedDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureSqlManagedDatabase.__pulumiType;
    }

    /**
     * The administrator login name for the new server. Changing this forces a new resource to be created.
     */
    public readonly administratorLogin!: pulumi.Output<string>;
    /**
     * The password associated with the `administratorLogin` user. Needs to comply with Azure's Password Policy
     */
    public readonly administratorLoginPassword!: pulumi.Output<string>;
    public /*out*/ readonly collation!: pulumi.Output<string>;
    /**
     * The fully qualified domain name of the sql managed instance.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
     */
    public readonly minimumTlsVersion!: pulumi.Output<string>;
    /**
     * The name of the SQL Managed Instance. This needs to be globally unique within Azure.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Is the public data endpoint enabled? Default value is `false`.
     */
    public readonly publicDataEndpointEnabled!: pulumi.Output<boolean>;
    /**
     * Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * Maximum storage space for your instance. It should be a multiple of 32GB.
     */
    public readonly storageSizeInGb!: pulumi.Output<number>;
    /**
     * The subnet resource id that the SQL Managed Instance will be associated with.
     */
    public readonly subnetId!: pulumi.Output<string>;
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * The GUID of the tenant that the azure sql managed database will be created in.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `skuName` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `skuName` is `GP_Gen5`
     */
    public readonly vcores!: pulumi.Output<number>;

    /**
     * Create a AzureSqlManagedDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureSqlManagedDatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureSqlManagedDatabaseArgs | AzureSqlManagedDatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureSqlManagedDatabaseState | undefined;
            resourceInputs["administratorLogin"] = state ? state.administratorLogin : undefined;
            resourceInputs["administratorLoginPassword"] = state ? state.administratorLoginPassword : undefined;
            resourceInputs["collation"] = state ? state.collation : undefined;
            resourceInputs["fqdn"] = state ? state.fqdn : undefined;
            resourceInputs["minimumTlsVersion"] = state ? state.minimumTlsVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publicDataEndpointEnabled"] = state ? state.publicDataEndpointEnabled : undefined;
            resourceInputs["skuName"] = state ? state.skuName : undefined;
            resourceInputs["storageSizeInGb"] = state ? state.storageSizeInGb : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["vcores"] = state ? state.vcores : undefined;
        } else {
            const args = argsOrState as AzureSqlManagedDatabaseArgs | undefined;
            if ((!args || args.administratorLogin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'administratorLogin'");
            }
            if ((!args || args.administratorLoginPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'administratorLoginPassword'");
            }
            if ((!args || args.skuName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'skuName'");
            }
            if ((!args || args.storageSizeInGb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageSizeInGb'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            if ((!args || args.vcores === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vcores'");
            }
            resourceInputs["administratorLogin"] = args ? args.administratorLogin : undefined;
            resourceInputs["administratorLoginPassword"] = args?.administratorLoginPassword ? pulumi.secret(args.administratorLoginPassword) : undefined;
            resourceInputs["minimumTlsVersion"] = args ? args.minimumTlsVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publicDataEndpointEnabled"] = args ? args.publicDataEndpointEnabled : undefined;
            resourceInputs["skuName"] = args ? args.skuName : undefined;
            resourceInputs["storageSizeInGb"] = args ? args.storageSizeInGb : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["vcores"] = args ? args.vcores : undefined;
            resourceInputs["collation"] = undefined /*out*/;
            resourceInputs["fqdn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["administratorLoginPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AzureSqlManagedDatabase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureSqlManagedDatabase resources.
 */
export interface AzureSqlManagedDatabaseState {
    /**
     * The administrator login name for the new server. Changing this forces a new resource to be created.
     */
    administratorLogin?: pulumi.Input<string>;
    /**
     * The password associated with the `administratorLogin` user. Needs to comply with Azure's Password Policy
     */
    administratorLoginPassword?: pulumi.Input<string>;
    collation?: pulumi.Input<string>;
    /**
     * The fully qualified domain name of the sql managed instance.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
     */
    minimumTlsVersion?: pulumi.Input<string>;
    /**
     * The name of the SQL Managed Instance. This needs to be globally unique within Azure.
     */
    name?: pulumi.Input<string>;
    /**
     * Is the public data endpoint enabled? Default value is `false`.
     */
    publicDataEndpointEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
     */
    skuName?: pulumi.Input<string>;
    /**
     * Maximum storage space for your instance. It should be a multiple of 32GB.
     */
    storageSizeInGb?: pulumi.Input<number>;
    /**
     * The subnet resource id that the SQL Managed Instance will be associated with.
     */
    subnetId?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The GUID of the tenant that the azure sql managed database will be created in.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `skuName` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `skuName` is `GP_Gen5`
     */
    vcores?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AzureSqlManagedDatabase resource.
 */
export interface AzureSqlManagedDatabaseArgs {
    /**
     * The administrator login name for the new server. Changing this forces a new resource to be created.
     */
    administratorLogin: pulumi.Input<string>;
    /**
     * The password associated with the `administratorLogin` user. Needs to comply with Azure's Password Policy
     */
    administratorLoginPassword: pulumi.Input<string>;
    /**
     * The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
     */
    minimumTlsVersion?: pulumi.Input<string>;
    /**
     * The name of the SQL Managed Instance. This needs to be globally unique within Azure.
     */
    name?: pulumi.Input<string>;
    /**
     * Is the public data endpoint enabled? Default value is `false`.
     */
    publicDataEndpointEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
     */
    skuName: pulumi.Input<string>;
    /**
     * Maximum storage space for your instance. It should be a multiple of 32GB.
     */
    storageSizeInGb: pulumi.Input<number>;
    /**
     * The subnet resource id that the SQL Managed Instance will be associated with.
     */
    subnetId: pulumi.Input<string>;
    /**
     * The GUID of the tenant that the azure sql managed database will be created in.
     */
    tenantId: pulumi.Input<string>;
    /**
     * Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `skuName` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `skuName` is `GP_Gen5`
     */
    vcores: pulumi.Input<number>;
}
