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

__all__ = ['AzureKeyVaultSecretArgs', 'AzureKeyVaultSecret']

@pulumi.input_type
class AzureKeyVaultSecretArgs:
    def __init__(__self__, *,
                 tenant_id: pulumi.Input[str],
                 value: pulumi.Input[str],
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AzureKeyVaultSecret resource.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the key vault secret will be created in.
        :param pulumi.Input[str] value: Specifies the value of the Key vault secret.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Secret should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type: Specifies the content type for the Key Vault Secret. Defaults to `duplo_container_env`.
        """
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "value", value)
        if key_vault_id is not None:
            pulumi.set(__self__, "key_vault_id", key_vault_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The GUID of the tenant that the key vault secret will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Specifies the value of the Key vault secret.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Key Vault where the Secret should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @key_vault_id.setter
    def key_vault_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the content type for the Key Vault Secret. Defaults to `duplo_container_env`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _AzureKeyVaultSecretState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 fullname: Optional[pulumi.Input[str]] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recovery_level: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 vault_base_url: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AzureKeyVaultSecret resources.
        :param pulumi.Input[bool] enabled: Determines whether the object is enabled.
        :param pulumi.Input[str] fullname: Duplo will generate name of the Key Vault Secret.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Secret should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_level: Reflects the deletion recovery level currently in effect for secrets in the current vault.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the key vault secret will be created in.
        :param pulumi.Input[str] type: Specifies the content type for the Key Vault Secret. Defaults to `duplo_container_env`.
        :param pulumi.Input[str] value: Specifies the value of the Key vault secret.
        :param pulumi.Input[str] vault_base_url: Base URL of the Azure Key Vault
        :param pulumi.Input[str] version: The current version of the Key Vault Secret.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if fullname is not None:
            pulumi.set(__self__, "fullname", fullname)
        if key_vault_id is not None:
            pulumi.set(__self__, "key_vault_id", key_vault_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if recovery_level is not None:
            pulumi.set(__self__, "recovery_level", recovery_level)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if vault_base_url is not None:
            pulumi.set(__self__, "vault_base_url", vault_base_url)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether the object is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def fullname(self) -> Optional[pulumi.Input[str]]:
        """
        Duplo will generate name of the Key Vault Secret.
        """
        return pulumi.get(self, "fullname")

    @fullname.setter
    def fullname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fullname", value)

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Key Vault where the Secret should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @key_vault_id.setter
    def key_vault_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="recoveryLevel")
    def recovery_level(self) -> Optional[pulumi.Input[str]]:
        """
        Reflects the deletion recovery level currently in effect for secrets in the current vault.
        """
        return pulumi.get(self, "recovery_level")

    @recovery_level.setter
    def recovery_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recovery_level", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID of the tenant that the key vault secret will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the content type for the Key Vault Secret. Defaults to `duplo_container_env`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the value of the Key vault secret.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="vaultBaseUrl")
    def vault_base_url(self) -> Optional[pulumi.Input[str]]:
        """
        Base URL of the Azure Key Vault
        """
        return pulumi.get(self, "vault_base_url")

    @vault_base_url.setter
    def vault_base_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault_base_url", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The current version of the Key Vault Secret.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class AzureKeyVaultSecret(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `AzureKeyVaultSecret` manages a Key Vault Secret in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        myapp_azure_key_vault_secret = duplocloud.AzureKeyVaultSecret("myapp",
            tenant_id=myapp.tenant_id,
            name="base01-test",
            value="tst",
            type="duplo_container_env")
        ```

        ## Import

        Example: Importing an existing Azure Key Vault Secret

         - *TENANT_ID* is the tenant GUID

         - *SHORT_NAME* is the short name of the Azure Key Vault Secret

        # 

        ```sh
        $ pulumi import duplocloud:index/azureKeyVaultSecret:AzureKeyVaultSecret mykvsecret *TENANT_ID*/*SHORT_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Secret should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the key vault secret will be created in.
        :param pulumi.Input[str] type: Specifies the content type for the Key Vault Secret. Defaults to `duplo_container_env`.
        :param pulumi.Input[str] value: Specifies the value of the Key vault secret.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AzureKeyVaultSecretArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `AzureKeyVaultSecret` manages a Key Vault Secret in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        myapp_azure_key_vault_secret = duplocloud.AzureKeyVaultSecret("myapp",
            tenant_id=myapp.tenant_id,
            name="base01-test",
            value="tst",
            type="duplo_container_env")
        ```

        ## Import

        Example: Importing an existing Azure Key Vault Secret

         - *TENANT_ID* is the tenant GUID

         - *SHORT_NAME* is the short name of the Azure Key Vault Secret

        # 

        ```sh
        $ pulumi import duplocloud:index/azureKeyVaultSecret:AzureKeyVaultSecret mykvsecret *TENANT_ID*/*SHORT_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param AzureKeyVaultSecretArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AzureKeyVaultSecretArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_vault_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AzureKeyVaultSecretArgs.__new__(AzureKeyVaultSecretArgs)

            __props__.__dict__["key_vault_id"] = key_vault_id
            __props__.__dict__["name"] = name
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["type"] = type
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = None if value is None else pulumi.Output.secret(value)
            __props__.__dict__["enabled"] = None
            __props__.__dict__["fullname"] = None
            __props__.__dict__["recovery_level"] = None
            __props__.__dict__["vault_base_url"] = None
            __props__.__dict__["version"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["value"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AzureKeyVaultSecret, __self__).__init__(
            'duplocloud:index/azureKeyVaultSecret:AzureKeyVaultSecret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            fullname: Optional[pulumi.Input[str]] = None,
            key_vault_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            recovery_level: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None,
            vault_base_url: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'AzureKeyVaultSecret':
        """
        Get an existing AzureKeyVaultSecret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Determines whether the object is enabled.
        :param pulumi.Input[str] fullname: Duplo will generate name of the Key Vault Secret.
        :param pulumi.Input[str] key_vault_id: The ID of the Key Vault where the Secret should be created.
        :param pulumi.Input[str] name: Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_level: Reflects the deletion recovery level currently in effect for secrets in the current vault.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the key vault secret will be created in.
        :param pulumi.Input[str] type: Specifies the content type for the Key Vault Secret. Defaults to `duplo_container_env`.
        :param pulumi.Input[str] value: Specifies the value of the Key vault secret.
        :param pulumi.Input[str] vault_base_url: Base URL of the Azure Key Vault
        :param pulumi.Input[str] version: The current version of the Key Vault Secret.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AzureKeyVaultSecretState.__new__(_AzureKeyVaultSecretState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["fullname"] = fullname
        __props__.__dict__["key_vault_id"] = key_vault_id
        __props__.__dict__["name"] = name
        __props__.__dict__["recovery_level"] = recovery_level
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["type"] = type
        __props__.__dict__["value"] = value
        __props__.__dict__["vault_base_url"] = vault_base_url
        __props__.__dict__["version"] = version
        return AzureKeyVaultSecret(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Determines whether the object is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def fullname(self) -> pulumi.Output[str]:
        """
        Duplo will generate name of the Key Vault Secret.
        """
        return pulumi.get(self, "fullname")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Key Vault where the Secret should be created.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Key Vault Secret. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="recoveryLevel")
    def recovery_level(self) -> pulumi.Output[str]:
        """
        Reflects the deletion recovery level currently in effect for secrets in the current vault.
        """
        return pulumi.get(self, "recovery_level")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The GUID of the tenant that the key vault secret will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the content type for the Key Vault Secret. Defaults to `duplo_container_env`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        Specifies the value of the Key vault secret.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="vaultBaseUrl")
    def vault_base_url(self) -> pulumi.Output[str]:
        """
        Base URL of the Azure Key Vault
        """
        return pulumi.get(self, "vault_base_url")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The current version of the Key Vault Secret.
        """
        return pulumi.get(self, "version")

