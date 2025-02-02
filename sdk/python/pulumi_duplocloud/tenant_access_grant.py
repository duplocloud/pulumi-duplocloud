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

__all__ = ['TenantAccessGrantArgs', 'TenantAccessGrant']

@pulumi.input_type
class TenantAccessGrantArgs:
    def __init__(__self__, *,
                 grant_area: pulumi.Input[str],
                 grantee_tenant_id: pulumi.Input[str],
                 grantor_tenant_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a TenantAccessGrant resource.
        :param pulumi.Input[str] grant_area: The area the grant allows access to. Currently supported: ['s3', 'dynamodb', 'kms']
        :param pulumi.Input[str] grantee_tenant_id: The GUID of the tenant that will receive the granted access.
        :param pulumi.Input[str] grantor_tenant_id: The GUID of the tenant that will grant the access.
        """
        pulumi.set(__self__, "grant_area", grant_area)
        pulumi.set(__self__, "grantee_tenant_id", grantee_tenant_id)
        pulumi.set(__self__, "grantor_tenant_id", grantor_tenant_id)

    @property
    @pulumi.getter(name="grantArea")
    def grant_area(self) -> pulumi.Input[str]:
        """
        The area the grant allows access to. Currently supported: ['s3', 'dynamodb', 'kms']
        """
        return pulumi.get(self, "grant_area")

    @grant_area.setter
    def grant_area(self, value: pulumi.Input[str]):
        pulumi.set(self, "grant_area", value)

    @property
    @pulumi.getter(name="granteeTenantId")
    def grantee_tenant_id(self) -> pulumi.Input[str]:
        """
        The GUID of the tenant that will receive the granted access.
        """
        return pulumi.get(self, "grantee_tenant_id")

    @grantee_tenant_id.setter
    def grantee_tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "grantee_tenant_id", value)

    @property
    @pulumi.getter(name="grantorTenantId")
    def grantor_tenant_id(self) -> pulumi.Input[str]:
        """
        The GUID of the tenant that will grant the access.
        """
        return pulumi.get(self, "grantor_tenant_id")

    @grantor_tenant_id.setter
    def grantor_tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "grantor_tenant_id", value)


@pulumi.input_type
class _TenantAccessGrantState:
    def __init__(__self__, *,
                 grant_area: Optional[pulumi.Input[str]] = None,
                 grantee_tenant_id: Optional[pulumi.Input[str]] = None,
                 grantor_tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TenantAccessGrant resources.
        :param pulumi.Input[str] grant_area: The area the grant allows access to. Currently supported: ['s3', 'dynamodb', 'kms']
        :param pulumi.Input[str] grantee_tenant_id: The GUID of the tenant that will receive the granted access.
        :param pulumi.Input[str] grantor_tenant_id: The GUID of the tenant that will grant the access.
        """
        if grant_area is not None:
            pulumi.set(__self__, "grant_area", grant_area)
        if grantee_tenant_id is not None:
            pulumi.set(__self__, "grantee_tenant_id", grantee_tenant_id)
        if grantor_tenant_id is not None:
            pulumi.set(__self__, "grantor_tenant_id", grantor_tenant_id)

    @property
    @pulumi.getter(name="grantArea")
    def grant_area(self) -> Optional[pulumi.Input[str]]:
        """
        The area the grant allows access to. Currently supported: ['s3', 'dynamodb', 'kms']
        """
        return pulumi.get(self, "grant_area")

    @grant_area.setter
    def grant_area(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant_area", value)

    @property
    @pulumi.getter(name="granteeTenantId")
    def grantee_tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID of the tenant that will receive the granted access.
        """
        return pulumi.get(self, "grantee_tenant_id")

    @grantee_tenant_id.setter
    def grantee_tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grantee_tenant_id", value)

    @property
    @pulumi.getter(name="grantorTenantId")
    def grantor_tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID of the tenant that will grant the access.
        """
        return pulumi.get(self, "grantor_tenant_id")

    @grantor_tenant_id.setter
    def grantor_tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grantor_tenant_id", value)


class TenantAccessGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grant_area: Optional[pulumi.Input[str]] = None,
                 grantee_tenant_id: Optional[pulumi.Input[str]] = None,
                 grantor_tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `TenantAccessGrant` manages a tenant access grant in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        grantor = duplocloud.get_tenant(name="tenant1")
        grantee = duplocloud.get_tenant(name="tenant2")
        dynamodb_grant = duplocloud.TenantAccessGrant("dynamodbGrant",
            grantee_tenant_id=grantee.id,
            grantor_tenant_id=grantor.id,
            grant_area="dynamodb")
        ```

        ## Import

        ```sh
        $ pulumi import duplocloud:index/tenantAccessGrant:TenantAccessGrant dynamodbGrant *GRANEE_TENANT_ID*/*GRANTOR_TENANT_ID*/*GRANTED_AREA*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grant_area: The area the grant allows access to. Currently supported: ['s3', 'dynamodb', 'kms']
        :param pulumi.Input[str] grantee_tenant_id: The GUID of the tenant that will receive the granted access.
        :param pulumi.Input[str] grantor_tenant_id: The GUID of the tenant that will grant the access.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TenantAccessGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `TenantAccessGrant` manages a tenant access grant in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        grantor = duplocloud.get_tenant(name="tenant1")
        grantee = duplocloud.get_tenant(name="tenant2")
        dynamodb_grant = duplocloud.TenantAccessGrant("dynamodbGrant",
            grantee_tenant_id=grantee.id,
            grantor_tenant_id=grantor.id,
            grant_area="dynamodb")
        ```

        ## Import

        ```sh
        $ pulumi import duplocloud:index/tenantAccessGrant:TenantAccessGrant dynamodbGrant *GRANEE_TENANT_ID*/*GRANTOR_TENANT_ID*/*GRANTED_AREA*
        ```

        :param str resource_name: The name of the resource.
        :param TenantAccessGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TenantAccessGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grant_area: Optional[pulumi.Input[str]] = None,
                 grantee_tenant_id: Optional[pulumi.Input[str]] = None,
                 grantor_tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TenantAccessGrantArgs.__new__(TenantAccessGrantArgs)

            if grant_area is None and not opts.urn:
                raise TypeError("Missing required property 'grant_area'")
            __props__.__dict__["grant_area"] = grant_area
            if grantee_tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'grantee_tenant_id'")
            __props__.__dict__["grantee_tenant_id"] = grantee_tenant_id
            if grantor_tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'grantor_tenant_id'")
            __props__.__dict__["grantor_tenant_id"] = grantor_tenant_id
        super(TenantAccessGrant, __self__).__init__(
            'duplocloud:index/tenantAccessGrant:TenantAccessGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            grant_area: Optional[pulumi.Input[str]] = None,
            grantee_tenant_id: Optional[pulumi.Input[str]] = None,
            grantor_tenant_id: Optional[pulumi.Input[str]] = None) -> 'TenantAccessGrant':
        """
        Get an existing TenantAccessGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grant_area: The area the grant allows access to. Currently supported: ['s3', 'dynamodb', 'kms']
        :param pulumi.Input[str] grantee_tenant_id: The GUID of the tenant that will receive the granted access.
        :param pulumi.Input[str] grantor_tenant_id: The GUID of the tenant that will grant the access.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TenantAccessGrantState.__new__(_TenantAccessGrantState)

        __props__.__dict__["grant_area"] = grant_area
        __props__.__dict__["grantee_tenant_id"] = grantee_tenant_id
        __props__.__dict__["grantor_tenant_id"] = grantor_tenant_id
        return TenantAccessGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="grantArea")
    def grant_area(self) -> pulumi.Output[str]:
        """
        The area the grant allows access to. Currently supported: ['s3', 'dynamodb', 'kms']
        """
        return pulumi.get(self, "grant_area")

    @property
    @pulumi.getter(name="granteeTenantId")
    def grantee_tenant_id(self) -> pulumi.Output[str]:
        """
        The GUID of the tenant that will receive the granted access.
        """
        return pulumi.get(self, "grantee_tenant_id")

    @property
    @pulumi.getter(name="grantorTenantId")
    def grantor_tenant_id(self) -> pulumi.Output[str]:
        """
        The GUID of the tenant that will grant the access.
        """
        return pulumi.get(self, "grantor_tenant_id")

