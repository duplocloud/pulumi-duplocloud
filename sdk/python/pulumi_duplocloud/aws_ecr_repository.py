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

__all__ = ['AwsEcrRepositoryArgs', 'AwsEcrRepository']

@pulumi.input_type
class AwsEcrRepositoryArgs:
    def __init__(__self__, *,
                 tenant_id: pulumi.Input[str],
                 enable_scan_image_on_push: Optional[pulumi.Input[bool]] = None,
                 enable_tag_immutability: Optional[pulumi.Input[bool]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 kms_encryption_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AwsEcrRepository resource.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the aws ecr repository will be created in.
        :param pulumi.Input[bool] enable_scan_image_on_push: Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
        :param pulumi.Input[bool] enable_tag_immutability: The tag mutability setting for the repository.
        :param pulumi.Input[bool] force_delete: Whether to force delete the repository on destroy operations Defaults to `false`.
        :param pulumi.Input[str] kms_encryption_key: The ARN of the KMS key to use.
        :param pulumi.Input[str] name: The name of the ECR Repository.
        """
        pulumi.set(__self__, "tenant_id", tenant_id)
        if enable_scan_image_on_push is not None:
            pulumi.set(__self__, "enable_scan_image_on_push", enable_scan_image_on_push)
        if enable_tag_immutability is not None:
            pulumi.set(__self__, "enable_tag_immutability", enable_tag_immutability)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if kms_encryption_key is not None:
            pulumi.set(__self__, "kms_encryption_key", kms_encryption_key)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The GUID of the tenant that the aws ecr repository will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="enableScanImageOnPush")
    def enable_scan_image_on_push(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
        """
        return pulumi.get(self, "enable_scan_image_on_push")

    @enable_scan_image_on_push.setter
    def enable_scan_image_on_push(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_scan_image_on_push", value)

    @property
    @pulumi.getter(name="enableTagImmutability")
    def enable_tag_immutability(self) -> Optional[pulumi.Input[bool]]:
        """
        The tag mutability setting for the repository.
        """
        return pulumi.get(self, "enable_tag_immutability")

    @enable_tag_immutability.setter
    def enable_tag_immutability(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_tag_immutability", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to force delete the repository on destroy operations Defaults to `false`.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter(name="kmsEncryptionKey")
    def kms_encryption_key(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the KMS key to use.
        """
        return pulumi.get(self, "kms_encryption_key")

    @kms_encryption_key.setter
    def kms_encryption_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_encryption_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ECR Repository.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AwsEcrRepositoryState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 enable_scan_image_on_push: Optional[pulumi.Input[bool]] = None,
                 enable_tag_immutability: Optional[pulumi.Input[bool]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 kms_encryption_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 repository_url: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AwsEcrRepository resources.
        :param pulumi.Input[str] arn: Full ARN of the ECR repository.
        :param pulumi.Input[bool] enable_scan_image_on_push: Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
        :param pulumi.Input[bool] enable_tag_immutability: The tag mutability setting for the repository.
        :param pulumi.Input[bool] force_delete: Whether to force delete the repository on destroy operations Defaults to `false`.
        :param pulumi.Input[str] kms_encryption_key: The ARN of the KMS key to use.
        :param pulumi.Input[str] name: The name of the ECR Repository.
        :param pulumi.Input[str] registry_id: The registry ID where the repository was created.
        :param pulumi.Input[str] repository_url: The URL of the repository.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the aws ecr repository will be created in.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if enable_scan_image_on_push is not None:
            pulumi.set(__self__, "enable_scan_image_on_push", enable_scan_image_on_push)
        if enable_tag_immutability is not None:
            pulumi.set(__self__, "enable_tag_immutability", enable_tag_immutability)
        if force_delete is not None:
            pulumi.set(__self__, "force_delete", force_delete)
        if kms_encryption_key is not None:
            pulumi.set(__self__, "kms_encryption_key", kms_encryption_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if registry_id is not None:
            pulumi.set(__self__, "registry_id", registry_id)
        if repository_url is not None:
            pulumi.set(__self__, "repository_url", repository_url)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Full ARN of the ECR repository.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="enableScanImageOnPush")
    def enable_scan_image_on_push(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
        """
        return pulumi.get(self, "enable_scan_image_on_push")

    @enable_scan_image_on_push.setter
    def enable_scan_image_on_push(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_scan_image_on_push", value)

    @property
    @pulumi.getter(name="enableTagImmutability")
    def enable_tag_immutability(self) -> Optional[pulumi.Input[bool]]:
        """
        The tag mutability setting for the repository.
        """
        return pulumi.get(self, "enable_tag_immutability")

    @enable_tag_immutability.setter
    def enable_tag_immutability(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_tag_immutability", value)

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to force delete the repository on destroy operations Defaults to `false`.
        """
        return pulumi.get(self, "force_delete")

    @force_delete.setter
    def force_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_delete", value)

    @property
    @pulumi.getter(name="kmsEncryptionKey")
    def kms_encryption_key(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the KMS key to use.
        """
        return pulumi.get(self, "kms_encryption_key")

    @kms_encryption_key.setter
    def kms_encryption_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_encryption_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ECR Repository.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> Optional[pulumi.Input[str]]:
        """
        The registry ID where the repository was created.
        """
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the repository.
        """
        return pulumi.get(self, "repository_url")

    @repository_url.setter
    def repository_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_url", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID of the tenant that the aws ecr repository will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


class AwsEcrRepository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_scan_image_on_push: Optional[pulumi.Input[bool]] = None,
                 enable_tag_immutability: Optional[pulumi.Input[bool]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 kms_encryption_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `AwsEcrRepository` manages an aws ecr repository in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        test_ecr = duplocloud.AwsEcrRepository("test-ecr",
            tenant_id=myapp.tenant_id,
            name="test-ecr",
            enable_scan_image_on_push=True,
            enable_tag_immutability=True,
            force_delete=False)
        ```

        ## Import

        Example: Importing an existing AWS ECR repository

         - *TENANT_ID* is the tenant GUID

         - *SHORT_NAME* is the short name of the AWS ECR repository

        # 

        ```sh
        $ pulumi import duplocloud:index/awsEcrRepository:AwsEcrRepository myecr *TENANT_ID*/*SHORT_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_scan_image_on_push: Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
        :param pulumi.Input[bool] enable_tag_immutability: The tag mutability setting for the repository.
        :param pulumi.Input[bool] force_delete: Whether to force delete the repository on destroy operations Defaults to `false`.
        :param pulumi.Input[str] kms_encryption_key: The ARN of the KMS key to use.
        :param pulumi.Input[str] name: The name of the ECR Repository.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the aws ecr repository will be created in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsEcrRepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `AwsEcrRepository` manages an aws ecr repository in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        test_ecr = duplocloud.AwsEcrRepository("test-ecr",
            tenant_id=myapp.tenant_id,
            name="test-ecr",
            enable_scan_image_on_push=True,
            enable_tag_immutability=True,
            force_delete=False)
        ```

        ## Import

        Example: Importing an existing AWS ECR repository

         - *TENANT_ID* is the tenant GUID

         - *SHORT_NAME* is the short name of the AWS ECR repository

        # 

        ```sh
        $ pulumi import duplocloud:index/awsEcrRepository:AwsEcrRepository myecr *TENANT_ID*/*SHORT_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param AwsEcrRepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsEcrRepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_scan_image_on_push: Optional[pulumi.Input[bool]] = None,
                 enable_tag_immutability: Optional[pulumi.Input[bool]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 kms_encryption_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsEcrRepositoryArgs.__new__(AwsEcrRepositoryArgs)

            __props__.__dict__["enable_scan_image_on_push"] = enable_scan_image_on_push
            __props__.__dict__["enable_tag_immutability"] = enable_tag_immutability
            __props__.__dict__["force_delete"] = force_delete
            __props__.__dict__["kms_encryption_key"] = kms_encryption_key
            __props__.__dict__["name"] = name
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["registry_id"] = None
            __props__.__dict__["repository_url"] = None
        super(AwsEcrRepository, __self__).__init__(
            'duplocloud:index/awsEcrRepository:AwsEcrRepository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            enable_scan_image_on_push: Optional[pulumi.Input[bool]] = None,
            enable_tag_immutability: Optional[pulumi.Input[bool]] = None,
            force_delete: Optional[pulumi.Input[bool]] = None,
            kms_encryption_key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            registry_id: Optional[pulumi.Input[str]] = None,
            repository_url: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'AwsEcrRepository':
        """
        Get an existing AwsEcrRepository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Full ARN of the ECR repository.
        :param pulumi.Input[bool] enable_scan_image_on_push: Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
        :param pulumi.Input[bool] enable_tag_immutability: The tag mutability setting for the repository.
        :param pulumi.Input[bool] force_delete: Whether to force delete the repository on destroy operations Defaults to `false`.
        :param pulumi.Input[str] kms_encryption_key: The ARN of the KMS key to use.
        :param pulumi.Input[str] name: The name of the ECR Repository.
        :param pulumi.Input[str] registry_id: The registry ID where the repository was created.
        :param pulumi.Input[str] repository_url: The URL of the repository.
        :param pulumi.Input[str] tenant_id: The GUID of the tenant that the aws ecr repository will be created in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsEcrRepositoryState.__new__(_AwsEcrRepositoryState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["enable_scan_image_on_push"] = enable_scan_image_on_push
        __props__.__dict__["enable_tag_immutability"] = enable_tag_immutability
        __props__.__dict__["force_delete"] = force_delete
        __props__.__dict__["kms_encryption_key"] = kms_encryption_key
        __props__.__dict__["name"] = name
        __props__.__dict__["registry_id"] = registry_id
        __props__.__dict__["repository_url"] = repository_url
        __props__.__dict__["tenant_id"] = tenant_id
        return AwsEcrRepository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Full ARN of the ECR repository.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="enableScanImageOnPush")
    def enable_scan_image_on_push(self) -> pulumi.Output[bool]:
        """
        Indicates whether images are scanned after being pushed to the repository (true) or not scanned (false).
        """
        return pulumi.get(self, "enable_scan_image_on_push")

    @property
    @pulumi.getter(name="enableTagImmutability")
    def enable_tag_immutability(self) -> pulumi.Output[bool]:
        """
        The tag mutability setting for the repository.
        """
        return pulumi.get(self, "enable_tag_immutability")

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to force delete the repository on destroy operations Defaults to `false`.
        """
        return pulumi.get(self, "force_delete")

    @property
    @pulumi.getter(name="kmsEncryptionKey")
    def kms_encryption_key(self) -> pulumi.Output[str]:
        """
        The ARN of the KMS key to use.
        """
        return pulumi.get(self, "kms_encryption_key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the ECR Repository.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[str]:
        """
        The registry ID where the repository was created.
        """
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter(name="repositoryUrl")
    def repository_url(self) -> pulumi.Output[str]:
        """
        The URL of the repository.
        """
        return pulumi.get(self, "repository_url")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The GUID of the tenant that the aws ecr repository will be created in.
        """
        return pulumi.get(self, "tenant_id")

