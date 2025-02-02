# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['AzureSqlManagedDatabaseArgs', 'AzureSqlManagedDatabase']

@pulumi.input_type
class AzureSqlManagedDatabaseArgs:
    def __init__(__self__, *,
                 administrator_login: pulumi.Input[str],
                 administrator_login_password: pulumi.Input[str],
                 sku_name: pulumi.Input[str],
                 storage_size_in_gb: pulumi.Input[int],
                 subnet_id: pulumi.Input[str],
                 tenant_id: pulumi.Input[str],
                 vcores: pulumi.Input[int],
                 minimum_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_data_endpoint_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AzureSqlManagedDatabase resource.
        :param pulumi.Input[str] administrator_login: The administrator login name for the new server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] administrator_login_password: The password associated with the `administrator_login` user. Needs to comply with Azure's Password Policy
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] storage_size_in_gb: Maximum storage space for your instance. It should be a multiple of 32GB.
        :param pulumi.Input[str] subnet_id: The subnet resource id that the SQL Managed Instance will be associated with.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the azure sql managed database will be created in.
        :param pulumi.Input[int] vcores: Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `sku_name` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `sku_name` is `GP_Gen5`
        :param pulumi.Input[str] minimum_tls_version: The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        :param pulumi.Input[str] name: The name of the SQL Managed Instance. This needs to be globally unique within Azure.
        :param pulumi.Input[bool] public_data_endpoint_enabled: Is the public data endpoint enabled? Default value is `false`.
        """
        pulumi.set(__self__, "administrator_login", administrator_login)
        pulumi.set(__self__, "administrator_login_password", administrator_login_password)
        pulumi.set(__self__, "sku_name", sku_name)
        pulumi.set(__self__, "storage_size_in_gb", storage_size_in_gb)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "vcores", vcores)
        if minimum_tls_version is not None:
            pulumi.set(__self__, "minimum_tls_version", minimum_tls_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_data_endpoint_enabled is not None:
            pulumi.set(__self__, "public_data_endpoint_enabled", public_data_endpoint_enabled)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> pulumi.Input[str]:
        """
        The administrator login name for the new server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "administrator_login")

    @administrator_login.setter
    def administrator_login(self, value: pulumi.Input[str]):
        pulumi.set(self, "administrator_login", value)

    @property
    @pulumi.getter(name="administratorLoginPassword")
    def administrator_login_password(self) -> pulumi.Input[str]:
        """
        The password associated with the `administrator_login` user. Needs to comply with Azure's Password Policy
        """
        return pulumi.get(self, "administrator_login_password")

    @administrator_login_password.setter
    def administrator_login_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "administrator_login_password", value)

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Input[str]:
        """
        Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku_name")

    @sku_name.setter
    def sku_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "sku_name", value)

    @property
    @pulumi.getter(name="storageSizeInGb")
    def storage_size_in_gb(self) -> pulumi.Input[int]:
        """
        Maximum storage space for your instance. It should be a multiple of 32GB.
        """
        return pulumi.get(self, "storage_size_in_gb")

    @storage_size_in_gb.setter
    def storage_size_in_gb(self, value: pulumi.Input[int]):
        pulumi.set(self, "storage_size_in_gb", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The subnet resource id that the SQL Managed Instance will be associated with.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The GUID of the tenant that the azure sql managed database will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def vcores(self) -> pulumi.Input[int]:
        """
        Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `sku_name` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `sku_name` is `GP_Gen5`
        """
        return pulumi.get(self, "vcores")

    @vcores.setter
    def vcores(self, value: pulumi.Input[int]):
        pulumi.set(self, "vcores", value)

    @property
    @pulumi.getter(name="minimumTlsVersion")
    def minimum_tls_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        """
        return pulumi.get(self, "minimum_tls_version")

    @minimum_tls_version.setter
    def minimum_tls_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minimum_tls_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SQL Managed Instance. This needs to be globally unique within Azure.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicDataEndpointEnabled")
    def public_data_endpoint_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the public data endpoint enabled? Default value is `false`.
        """
        return pulumi.get(self, "public_data_endpoint_enabled")

    @public_data_endpoint_enabled.setter
    def public_data_endpoint_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public_data_endpoint_enabled", value)


@pulumi.input_type
class _AzureSqlManagedDatabaseState:
    def __init__(__self__, *,
                 administrator_login: Optional[pulumi.Input[str]] = None,
                 administrator_login_password: Optional[pulumi.Input[str]] = None,
                 collation: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 minimum_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_data_endpoint_enabled: Optional[pulumi.Input[bool]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 storage_size_in_gb: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 vcores: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AzureSqlManagedDatabase resources.
        :param pulumi.Input[str] administrator_login: The administrator login name for the new server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] administrator_login_password: The password associated with the `administrator_login` user. Needs to comply with Azure's Password Policy
        :param pulumi.Input[str] fqdn: The fully qualified domain name of the sql managed instance.
        :param pulumi.Input[str] minimum_tls_version: The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        :param pulumi.Input[str] name: The name of the SQL Managed Instance. This needs to be globally unique within Azure.
        :param pulumi.Input[bool] public_data_endpoint_enabled: Is the public data endpoint enabled? Default value is `false`.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] storage_size_in_gb: Maximum storage space for your instance. It should be a multiple of 32GB.
        :param pulumi.Input[str] subnet_id: The subnet resource id that the SQL Managed Instance will be associated with.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the azure sql managed database will be created in.
        :param pulumi.Input[int] vcores: Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `sku_name` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `sku_name` is `GP_Gen5`
        """
        if administrator_login is not None:
            pulumi.set(__self__, "administrator_login", administrator_login)
        if administrator_login_password is not None:
            pulumi.set(__self__, "administrator_login_password", administrator_login_password)
        if collation is not None:
            pulumi.set(__self__, "collation", collation)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if minimum_tls_version is not None:
            pulumi.set(__self__, "minimum_tls_version", minimum_tls_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_data_endpoint_enabled is not None:
            pulumi.set(__self__, "public_data_endpoint_enabled", public_data_endpoint_enabled)
        if sku_name is not None:
            pulumi.set(__self__, "sku_name", sku_name)
        if storage_size_in_gb is not None:
            pulumi.set(__self__, "storage_size_in_gb", storage_size_in_gb)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if vcores is not None:
            pulumi.set(__self__, "vcores", vcores)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> Optional[pulumi.Input[str]]:
        """
        The administrator login name for the new server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "administrator_login")

    @administrator_login.setter
    def administrator_login(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "administrator_login", value)

    @property
    @pulumi.getter(name="administratorLoginPassword")
    def administrator_login_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password associated with the `administrator_login` user. Needs to comply with Azure's Password Policy
        """
        return pulumi.get(self, "administrator_login_password")

    @administrator_login_password.setter
    def administrator_login_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "administrator_login_password", value)

    @property
    @pulumi.getter
    def collation(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "collation")

    @collation.setter
    def collation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collation", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified domain name of the sql managed instance.
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter(name="minimumTlsVersion")
    def minimum_tls_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        """
        return pulumi.get(self, "minimum_tls_version")

    @minimum_tls_version.setter
    def minimum_tls_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minimum_tls_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SQL Managed Instance. This needs to be globally unique within Azure.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicDataEndpointEnabled")
    def public_data_endpoint_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the public data endpoint enabled? Default value is `false`.
        """
        return pulumi.get(self, "public_data_endpoint_enabled")

    @public_data_endpoint_enabled.setter
    def public_data_endpoint_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public_data_endpoint_enabled", value)

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku_name")

    @sku_name.setter
    def sku_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sku_name", value)

    @property
    @pulumi.getter(name="storageSizeInGb")
    def storage_size_in_gb(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum storage space for your instance. It should be a multiple of 32GB.
        """
        return pulumi.get(self, "storage_size_in_gb")

    @storage_size_in_gb.setter
    def storage_size_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage_size_in_gb", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet resource id that the SQL Managed Instance will be associated with.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID of the tenant that the azure sql managed database will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def vcores(self) -> Optional[pulumi.Input[int]]:
        """
        Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `sku_name` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `sku_name` is `GP_Gen5`
        """
        return pulumi.get(self, "vcores")

    @vcores.setter
    def vcores(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vcores", value)


class AzureSqlManagedDatabase(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 administrator_login: Optional[pulumi.Input[str]] = None,
                 administrator_login_password: Optional[pulumi.Input[str]] = None,
                 minimum_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_data_endpoint_enabled: Optional[pulumi.Input[bool]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 storage_size_in_gb: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 vcores: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        `AzureSqlManagedDatabase` manages an azure sql managed database in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        mydb = duplocloud.AzureSqlManagedDatabase("mydb",
            tenant_id=myapp.tenant_id,
            name="db-test",
            administrator_login="testroot",
            administrator_login_password="P@ssword12345",
            minimum_tls_version="1.2",
            sku_name="GP_Gen5",
            vcores=4,
            storage_size_in_gb=32,
            subnet_id="/subscriptions/0c84b91e-95f5-409e-9cff-6c2e60affbb3/resourceGroups/duploinfra-demo/providers/Microsoft.Network/virtualNetworks/demo/subnets/duploinfra-default")
        ```

        ## Import

        Example: Importing an existing Azure Sql Managed Database

         - *TENANT_ID* is the tenant GUID

         - *SHORT_NAME* is the short name of the Azure Sql Managed Database

        # 

        ```sh
        $ pulumi import duplocloud:index/azureSqlManagedDatabase:AzureSqlManagedDatabase myManagedSQLDatabase *TENANT_ID*/*SHORT_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administrator_login: The administrator login name for the new server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] administrator_login_password: The password associated with the `administrator_login` user. Needs to comply with Azure's Password Policy
        :param pulumi.Input[str] minimum_tls_version: The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        :param pulumi.Input[str] name: The name of the SQL Managed Instance. This needs to be globally unique within Azure.
        :param pulumi.Input[bool] public_data_endpoint_enabled: Is the public data endpoint enabled? Default value is `false`.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] storage_size_in_gb: Maximum storage space for your instance. It should be a multiple of 32GB.
        :param pulumi.Input[str] subnet_id: The subnet resource id that the SQL Managed Instance will be associated with.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the azure sql managed database will be created in.
        :param pulumi.Input[int] vcores: Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `sku_name` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `sku_name` is `GP_Gen5`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AzureSqlManagedDatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `AzureSqlManagedDatabase` manages an azure sql managed database in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        mydb = duplocloud.AzureSqlManagedDatabase("mydb",
            tenant_id=myapp.tenant_id,
            name="db-test",
            administrator_login="testroot",
            administrator_login_password="P@ssword12345",
            minimum_tls_version="1.2",
            sku_name="GP_Gen5",
            vcores=4,
            storage_size_in_gb=32,
            subnet_id="/subscriptions/0c84b91e-95f5-409e-9cff-6c2e60affbb3/resourceGroups/duploinfra-demo/providers/Microsoft.Network/virtualNetworks/demo/subnets/duploinfra-default")
        ```

        ## Import

        Example: Importing an existing Azure Sql Managed Database

         - *TENANT_ID* is the tenant GUID

         - *SHORT_NAME* is the short name of the Azure Sql Managed Database

        # 

        ```sh
        $ pulumi import duplocloud:index/azureSqlManagedDatabase:AzureSqlManagedDatabase myManagedSQLDatabase *TENANT_ID*/*SHORT_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param AzureSqlManagedDatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AzureSqlManagedDatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 administrator_login: Optional[pulumi.Input[str]] = None,
                 administrator_login_password: Optional[pulumi.Input[str]] = None,
                 minimum_tls_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_data_endpoint_enabled: Optional[pulumi.Input[bool]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 storage_size_in_gb: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 vcores: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AzureSqlManagedDatabaseArgs.__new__(AzureSqlManagedDatabaseArgs)

            if administrator_login is None and not opts.urn:
                raise TypeError("Missing required property 'administrator_login'")
            __props__.__dict__["administrator_login"] = administrator_login
            if administrator_login_password is None and not opts.urn:
                raise TypeError("Missing required property 'administrator_login_password'")
            __props__.__dict__["administrator_login_password"] = None if administrator_login_password is None else pulumi.Output.secret(administrator_login_password)
            __props__.__dict__["minimum_tls_version"] = minimum_tls_version
            __props__.__dict__["name"] = name
            __props__.__dict__["public_data_endpoint_enabled"] = public_data_endpoint_enabled
            if sku_name is None and not opts.urn:
                raise TypeError("Missing required property 'sku_name'")
            __props__.__dict__["sku_name"] = sku_name
            if storage_size_in_gb is None and not opts.urn:
                raise TypeError("Missing required property 'storage_size_in_gb'")
            __props__.__dict__["storage_size_in_gb"] = storage_size_in_gb
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
            if vcores is None and not opts.urn:
                raise TypeError("Missing required property 'vcores'")
            __props__.__dict__["vcores"] = vcores
            __props__.__dict__["collation"] = None
            __props__.__dict__["fqdn"] = None
            __props__.__dict__["tags"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["administratorLoginPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AzureSqlManagedDatabase, __self__).__init__(
            'duplocloud:index/azureSqlManagedDatabase:AzureSqlManagedDatabase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            administrator_login: Optional[pulumi.Input[str]] = None,
            administrator_login_password: Optional[pulumi.Input[str]] = None,
            collation: Optional[pulumi.Input[str]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            minimum_tls_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            public_data_endpoint_enabled: Optional[pulumi.Input[bool]] = None,
            sku_name: Optional[pulumi.Input[str]] = None,
            storage_size_in_gb: Optional[pulumi.Input[int]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            vcores: Optional[pulumi.Input[int]] = None) -> 'AzureSqlManagedDatabase':
        """
        Get an existing AzureSqlManagedDatabase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administrator_login: The administrator login name for the new server. Changing this forces a new resource to be created.
        :param pulumi.Input[str] administrator_login_password: The password associated with the `administrator_login` user. Needs to comply with Azure's Password Policy
        :param pulumi.Input[str] fqdn: The fully qualified domain name of the sql managed instance.
        :param pulumi.Input[str] minimum_tls_version: The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        :param pulumi.Input[str] name: The name of the SQL Managed Instance. This needs to be globally unique within Azure.
        :param pulumi.Input[bool] public_data_endpoint_enabled: Is the public data endpoint enabled? Default value is `false`.
        :param pulumi.Input[str] sku_name: Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
        :param pulumi.Input[int] storage_size_in_gb: Maximum storage space for your instance. It should be a multiple of 32GB.
        :param pulumi.Input[str] subnet_id: The subnet resource id that the SQL Managed Instance will be associated with.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the azure sql managed database will be created in.
        :param pulumi.Input[int] vcores: Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `sku_name` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `sku_name` is `GP_Gen5`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AzureSqlManagedDatabaseState.__new__(_AzureSqlManagedDatabaseState)

        __props__.__dict__["administrator_login"] = administrator_login
        __props__.__dict__["administrator_login_password"] = administrator_login_password
        __props__.__dict__["collation"] = collation
        __props__.__dict__["fqdn"] = fqdn
        __props__.__dict__["minimum_tls_version"] = minimum_tls_version
        __props__.__dict__["name"] = name
        __props__.__dict__["public_data_endpoint_enabled"] = public_data_endpoint_enabled
        __props__.__dict__["sku_name"] = sku_name
        __props__.__dict__["storage_size_in_gb"] = storage_size_in_gb
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["vcores"] = vcores
        return AzureSqlManagedDatabase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> pulumi.Output[str]:
        """
        The administrator login name for the new server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "administrator_login")

    @property
    @pulumi.getter(name="administratorLoginPassword")
    def administrator_login_password(self) -> pulumi.Output[str]:
        """
        The password associated with the `administrator_login` user. Needs to comply with Azure's Password Policy
        """
        return pulumi.get(self, "administrator_login_password")

    @property
    @pulumi.getter
    def collation(self) -> pulumi.Output[str]:
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        The fully qualified domain name of the sql managed instance.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="minimumTlsVersion")
    def minimum_tls_version(self) -> pulumi.Output[str]:
        """
        The Minimum TLS Version for all SQL managed Database and SQL Data Warehouse databases associated with the server. Valid values are: `1.0`, `1.1` and `1.2`.
        """
        return pulumi.get(self, "minimum_tls_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the SQL Managed Instance. This needs to be globally unique within Azure.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicDataEndpointEnabled")
    def public_data_endpoint_enabled(self) -> pulumi.Output[bool]:
        """
        Is the public data endpoint enabled? Default value is `false`.
        """
        return pulumi.get(self, "public_data_endpoint_enabled")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Output[str]:
        """
        Specifies the SKU Name for the SQL Managed Instance. Valid values include `GP_Gen4`, `GP_Gen5`, `BC_Gen4`, `BC_Gen5`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter(name="storageSizeInGb")
    def storage_size_in_gb(self) -> pulumi.Output[int]:
        """
        Maximum storage space for your instance. It should be a multiple of 32GB.
        """
        return pulumi.get(self, "storage_size_in_gb")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The subnet resource id that the SQL Managed Instance will be associated with.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The GUID of the tenant that the azure sql managed database will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def vcores(self) -> pulumi.Output[int]:
        """
        Number of cores that should be assigned to your instance. Values can be `8`, `16`, or `24` if `sku_name` is `GP_Gen4`, or `8`, `16`, `24`, `32`, or `40` if `sku_name` is `GP_Gen5`
        """
        return pulumi.get(self, "vcores")

