// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * `duplocloud.AzureMssqlDatabase` manages an azure mssql database in Duplo.
 *
 * ## Import
 *
 * Example: Importing an existing Azure MS SQL databse
 *
 *  - *TENANT_ID* is the tenant GUID
 *
 *  - *SERVER_NAME* is the short name of the Azure MS SQL Server
 *
 *  - *DB_NAME* is the short name of the Azure MS SQL Database
 *
 * ```sh
 * $ pulumi import duplocloud:index/azureMssqlDatabase:AzureMssqlDatabase myMsSqlDb *TENANT_ID*&#47;*SERVER_NAME*&#47;*DB_NAME*
 * ```
 */
export class AzureMssqlDatabase extends pulumi.CustomResource {
    /**
     * Get an existing AzureMssqlDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureMssqlDatabaseState, opts?: pulumi.CustomResourceOptions): AzureMssqlDatabase {
        return new AzureMssqlDatabase(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'duplocloud:index/azureMssqlDatabase:AzureMssqlDatabase';

    /**
     * Returns true if the given object is an instance of AzureMssqlDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureMssqlDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureMssqlDatabase.__pulumiType;
    }

    /**
     * Specifies the collation of the database.
     */
    public readonly collation!: pulumi.Output<string>;
    /**
     * Specifies the id of the elastic pool containing this database.
     */
    public readonly elasticPoolId!: pulumi.Output<string | undefined>;
    /**
     * The name of the MS SQL Database.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the MS SQL Server on which to create the database.
     */
    public readonly serverName!: pulumi.Output<string>;
    public readonly sku!: pulumi.Output<outputs.AzureMssqlDatabaseSku | undefined>;
    /**
     * The GUID of the tenant that the azure mssql database will be created in.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a AzureMssqlDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureMssqlDatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureMssqlDatabaseArgs | AzureMssqlDatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureMssqlDatabaseState | undefined;
            resourceInputs["collation"] = state ? state.collation : undefined;
            resourceInputs["elasticPoolId"] = state ? state.elasticPoolId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serverName"] = state ? state.serverName : undefined;
            resourceInputs["sku"] = state ? state.sku : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as AzureMssqlDatabaseArgs | undefined;
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["collation"] = args ? args.collation : undefined;
            resourceInputs["elasticPoolId"] = args ? args.elasticPoolId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["sku"] = args ? args.sku : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AzureMssqlDatabase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureMssqlDatabase resources.
 */
export interface AzureMssqlDatabaseState {
    /**
     * Specifies the collation of the database.
     */
    collation?: pulumi.Input<string>;
    /**
     * Specifies the id of the elastic pool containing this database.
     */
    elasticPoolId?: pulumi.Input<string>;
    /**
     * The name of the MS SQL Database.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the MS SQL Server on which to create the database.
     */
    serverName?: pulumi.Input<string>;
    sku?: pulumi.Input<inputs.AzureMssqlDatabaseSku>;
    /**
     * The GUID of the tenant that the azure mssql database will be created in.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AzureMssqlDatabase resource.
 */
export interface AzureMssqlDatabaseArgs {
    /**
     * Specifies the collation of the database.
     */
    collation?: pulumi.Input<string>;
    /**
     * Specifies the id of the elastic pool containing this database.
     */
    elasticPoolId?: pulumi.Input<string>;
    /**
     * The name of the MS SQL Database.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the MS SQL Server on which to create the database.
     */
    serverName: pulumi.Input<string>;
    sku?: pulumi.Input<inputs.AzureMssqlDatabaseSku>;
    /**
     * The GUID of the tenant that the azure mssql database will be created in.
     */
    tenantId: pulumi.Input<string>;
}
