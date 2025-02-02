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

__all__ = ['AwsLaunchTemplateDefaultVersionArgs', 'AwsLaunchTemplateDefaultVersion']

@pulumi.input_type
class AwsLaunchTemplateDefaultVersionArgs:
    def __init__(__self__, *,
                 default_version: pulumi.Input[str],
                 tenant_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AwsLaunchTemplateDefaultVersion resource.
        :param pulumi.Input[str] default_version: The default version of the launch template to be set.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the launch template will be created in.
        :param pulumi.Input[str] name: The fullname of the asg group
        """
        pulumi.set(__self__, "default_version", default_version)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="defaultVersion")
    def default_version(self) -> pulumi.Input[str]:
        """
        The default version of the launch template to be set.
        """
        return pulumi.get(self, "default_version")

    @default_version.setter
    def default_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_version", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The GUID of the tenant that the launch template will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The fullname of the asg group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AwsLaunchTemplateDefaultVersionState:
    def __init__(__self__, *,
                 default_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AwsLaunchTemplateDefaultVersion resources.
        :param pulumi.Input[str] default_version: The default version of the launch template to be set.
        :param pulumi.Input[str] name: The fullname of the asg group
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the launch template will be created in.
        """
        if default_version is not None:
            pulumi.set(__self__, "default_version", default_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="defaultVersion")
    def default_version(self) -> Optional[pulumi.Input[str]]:
        """
        The default version of the launch template to be set.
        """
        return pulumi.get(self, "default_version")

    @default_version.setter
    def default_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The fullname of the asg group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID of the tenant that the launch template will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


class AwsLaunchTemplateDefaultVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        duplocloud_aws_launch_template_default_version helps to set or update default version of launch template

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        op = duplocloud.get_aws_launch_template_output(tenant_id=myapp.tenant_id,
            name="launch_template_name")
        name = duplocloud.AwsLaunchTemplateDefaultVersion("name",
            tenant_id=op.tenant_id,
            name=op.name,
            default_version=op.latest_version)
        ```

        ## Import

        Example: Importing an existing AWS launch template default version

         - *TENANT_ID* is the tenant GUID

         - *NAME* is the name of the AWS launch template

        ```sh
        $ pulumi import duplocloud:index/awsLaunchTemplateDefaultVersion:AwsLaunchTemplateDefaultVersion dlt *TENANT_ID*/launch-template/*NAME*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_version: The default version of the launch template to be set.
        :param pulumi.Input[str] name: The fullname of the asg group
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the launch template will be created in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsLaunchTemplateDefaultVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        duplocloud_aws_launch_template_default_version helps to set or update default version of launch template

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        op = duplocloud.get_aws_launch_template_output(tenant_id=myapp.tenant_id,
            name="launch_template_name")
        name = duplocloud.AwsLaunchTemplateDefaultVersion("name",
            tenant_id=op.tenant_id,
            name=op.name,
            default_version=op.latest_version)
        ```

        ## Import

        Example: Importing an existing AWS launch template default version

         - *TENANT_ID* is the tenant GUID

         - *NAME* is the name of the AWS launch template

        ```sh
        $ pulumi import duplocloud:index/awsLaunchTemplateDefaultVersion:AwsLaunchTemplateDefaultVersion dlt *TENANT_ID*/launch-template/*NAME*
        ```

        :param str resource_name: The name of the resource.
        :param AwsLaunchTemplateDefaultVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsLaunchTemplateDefaultVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsLaunchTemplateDefaultVersionArgs.__new__(AwsLaunchTemplateDefaultVersionArgs)

            if default_version is None and not opts.urn:
                raise TypeError("Missing required property 'default_version'")
            __props__.__dict__["default_version"] = default_version
            __props__.__dict__["name"] = name
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
        super(AwsLaunchTemplateDefaultVersion, __self__).__init__(
            'duplocloud:index/awsLaunchTemplateDefaultVersion:AwsLaunchTemplateDefaultVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'AwsLaunchTemplateDefaultVersion':
        """
        Get an existing AwsLaunchTemplateDefaultVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_version: The default version of the launch template to be set.
        :param pulumi.Input[str] name: The fullname of the asg group
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the launch template will be created in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsLaunchTemplateDefaultVersionState.__new__(_AwsLaunchTemplateDefaultVersionState)

        __props__.__dict__["default_version"] = default_version
        __props__.__dict__["name"] = name
        __props__.__dict__["tenant_id"] = tenant_id
        return AwsLaunchTemplateDefaultVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultVersion")
    def default_version(self) -> pulumi.Output[str]:
        """
        The default version of the launch template to be set.
        """
        return pulumi.get(self, "default_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The fullname of the asg group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The GUID of the tenant that the launch template will be created in.
        """
        return pulumi.get(self, "tenant_id")

