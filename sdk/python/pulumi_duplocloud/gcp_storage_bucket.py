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

__all__ = ['GcpStorageBucketArgs', 'GcpStorageBucket']

@pulumi.input_type
class GcpStorageBucketArgs:
    def __init__(__self__, *,
                 tenant_id: pulumi.Input[str],
                 allow_public_access: Optional[pulumi.Input[bool]] = None,
                 enable_versioning: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GcpStorageBucket resource.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the storage bucket will be created in.
        :param pulumi.Input[bool] allow_public_access: Whether or not public access might be allowed for the storage bucket. Defaults to `false`.
        :param pulumi.Input[bool] enable_versioning: Whether or not versioning is enabled for the storage bucket. Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels assigned to this storage bucket.
        :param pulumi.Input[str] name: The short name of the storage bucket.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
        """
        pulumi.set(__self__, "tenant_id", tenant_id)
        if allow_public_access is not None:
            pulumi.set(__self__, "allow_public_access", allow_public_access)
        if enable_versioning is not None:
            pulumi.set(__self__, "enable_versioning", enable_versioning)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The GUID of the tenant that the storage bucket will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="allowPublicAccess")
    def allow_public_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not public access might be allowed for the storage bucket. Defaults to `false`.
        """
        return pulumi.get(self, "allow_public_access")

    @allow_public_access.setter
    def allow_public_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_public_access", value)

    @property
    @pulumi.getter(name="enableVersioning")
    def enable_versioning(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not versioning is enabled for the storage bucket. Defaults to `false`.
        """
        return pulumi.get(self, "enable_versioning")

    @enable_versioning.setter
    def enable_versioning(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_versioning", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels assigned to this storage bucket.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The short name of the storage bucket.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GcpStorageBucketState:
    def __init__(__self__, *,
                 allow_public_access: Optional[pulumi.Input[bool]] = None,
                 enable_versioning: Optional[pulumi.Input[bool]] = None,
                 fullname: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GcpStorageBucket resources.
        :param pulumi.Input[bool] allow_public_access: Whether or not public access might be allowed for the storage bucket. Defaults to `false`.
        :param pulumi.Input[bool] enable_versioning: Whether or not versioning is enabled for the storage bucket. Defaults to `false`.
        :param pulumi.Input[str] fullname: The full name of the storage bucket.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels assigned to this storage bucket.
        :param pulumi.Input[str] name: The short name of the storage bucket.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
        :param pulumi.Input[str] self_link: The SelfLink of the storage bucket.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the storage bucket will be created in.
        """
        if allow_public_access is not None:
            pulumi.set(__self__, "allow_public_access", allow_public_access)
        if enable_versioning is not None:
            pulumi.set(__self__, "enable_versioning", enable_versioning)
        if fullname is not None:
            pulumi.set(__self__, "fullname", fullname)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="allowPublicAccess")
    def allow_public_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not public access might be allowed for the storage bucket. Defaults to `false`.
        """
        return pulumi.get(self, "allow_public_access")

    @allow_public_access.setter
    def allow_public_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_public_access", value)

    @property
    @pulumi.getter(name="enableVersioning")
    def enable_versioning(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not versioning is enabled for the storage bucket. Defaults to `false`.
        """
        return pulumi.get(self, "enable_versioning")

    @enable_versioning.setter
    def enable_versioning(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_versioning", value)

    @property
    @pulumi.getter
    def fullname(self) -> Optional[pulumi.Input[str]]:
        """
        The full name of the storage bucket.
        """
        return pulumi.get(self, "fullname")

    @fullname.setter
    def fullname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fullname", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels assigned to this storage bucket.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The short name of the storage bucket.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The SelfLink of the storage bucket.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID of the tenant that the storage bucket will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


class GcpStorageBucket(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_public_access: Optional[pulumi.Input[bool]] = None,
                 enable_versioning: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `GcpStorageBucket` manages a GCP storage bucket in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        mybucket = duplocloud.GcpStorageBucket("mybucket",
            tenant_id=myapp.tenant_id,
            name="mybucket",
            enable_versioning=False)
        ```

        ## Import

        Example: Importing an existing GCP storage bucket

         - *TENANT_ID* is the tenant GUID

         - *SHORT_NAME* is the short name of the GCP storage bucket

        # 

        ```sh
        $ pulumi import duplocloud:index/gcpStorageBucket:GcpStorageBucket mybucket *TENANT_ID*/*SHORT_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_public_access: Whether or not public access might be allowed for the storage bucket. Defaults to `false`.
        :param pulumi.Input[bool] enable_versioning: Whether or not versioning is enabled for the storage bucket. Defaults to `false`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels assigned to this storage bucket.
        :param pulumi.Input[str] name: The short name of the storage bucket.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the storage bucket will be created in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GcpStorageBucketArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `GcpStorageBucket` manages a GCP storage bucket in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        mybucket = duplocloud.GcpStorageBucket("mybucket",
            tenant_id=myapp.tenant_id,
            name="mybucket",
            enable_versioning=False)
        ```

        ## Import

        Example: Importing an existing GCP storage bucket

         - *TENANT_ID* is the tenant GUID

         - *SHORT_NAME* is the short name of the GCP storage bucket

        # 

        ```sh
        $ pulumi import duplocloud:index/gcpStorageBucket:GcpStorageBucket mybucket *TENANT_ID*/*SHORT_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param GcpStorageBucketArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GcpStorageBucketArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_public_access: Optional[pulumi.Input[bool]] = None,
                 enable_versioning: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GcpStorageBucketArgs.__new__(GcpStorageBucketArgs)

            __props__.__dict__["allow_public_access"] = allow_public_access
            __props__.__dict__["enable_versioning"] = enable_versioning
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["fullname"] = None
            __props__.__dict__["self_link"] = None
        super(GcpStorageBucket, __self__).__init__(
            'duplocloud:index/gcpStorageBucket:GcpStorageBucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_public_access: Optional[pulumi.Input[bool]] = None,
            enable_versioning: Optional[pulumi.Input[bool]] = None,
            fullname: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'GcpStorageBucket':
        """
        Get an existing GcpStorageBucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_public_access: Whether or not public access might be allowed for the storage bucket. Defaults to `false`.
        :param pulumi.Input[bool] enable_versioning: Whether or not versioning is enabled for the storage bucket. Defaults to `false`.
        :param pulumi.Input[str] fullname: The full name of the storage bucket.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels assigned to this storage bucket.
        :param pulumi.Input[str] name: The short name of the storage bucket.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
        :param pulumi.Input[str] self_link: The SelfLink of the storage bucket.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the storage bucket will be created in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GcpStorageBucketState.__new__(_GcpStorageBucketState)

        __props__.__dict__["allow_public_access"] = allow_public_access
        __props__.__dict__["enable_versioning"] = enable_versioning
        __props__.__dict__["fullname"] = fullname
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["tenant_id"] = tenant_id
        return GcpStorageBucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowPublicAccess")
    def allow_public_access(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not public access might be allowed for the storage bucket. Defaults to `false`.
        """
        return pulumi.get(self, "allow_public_access")

    @property
    @pulumi.getter(name="enableVersioning")
    def enable_versioning(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not versioning is enabled for the storage bucket. Defaults to `false`.
        """
        return pulumi.get(self, "enable_versioning")

    @property
    @pulumi.getter
    def fullname(self) -> pulumi.Output[str]:
        """
        The full name of the storage bucket.
        """
        return pulumi.get(self, "fullname")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The labels assigned to this storage bucket.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The short name of the storage bucket.  Duplo will add a prefix to the name.  You can retrieve the full name from the `fullname` attribute.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The SelfLink of the storage bucket.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The GUID of the tenant that the storage bucket will be created in.
        """
        return pulumi.get(self, "tenant_id")

