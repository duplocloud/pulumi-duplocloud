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

__all__ = ['GcpRedisInstanceArgs', 'GcpRedisInstance']

@pulumi.input_type
class GcpRedisInstanceArgs:
    def __init__(__self__, *,
                 memory_size_gb: pulumi.Input[int],
                 tenant_id: pulumi.Input[str],
                 tier: pulumi.Input[str],
                 auth_enabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_replicas_enabled: Optional[pulumi.Input[bool]] = None,
                 redis_configs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 redis_version: Optional[pulumi.Input[str]] = None,
                 replica_count: Optional[pulumi.Input[int]] = None,
                 transit_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 wait_until_ready: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GcpRedisInstance resource.
        :param pulumi.Input[int] memory_size_gb: Redis memory size in GiB.
        :param pulumi.Input[str] tenant_id: GUID of the tenant the Redis instance will be created in.
        :param pulumi.Input[str] tier: Service tier. Must be one of ['BASIC', 'STANDARD_HA'].
        :param pulumi.Input[bool] auth_enabled: Enable OSS Redis AUTH. Defaults to false (AUTH disabled).
        :param pulumi.Input[str] display_name: User-provided name for the instance.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels for user-provided metadata.
        :param pulumi.Input[str] name: Short name of the Redis instance. Duplo adds a prefix; retrieve the full name from `fullname`.
        :param pulumi.Input[bool] read_replicas_enabled: Enable read replica mode (can only be set during instance creation).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] redis_configs: Redis configuration parameters. See https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs for supported parameters.
        :param pulumi.Input[str] redis_version: Version of Redis software. Defaults to the latest supported version.
        :param pulumi.Input[int] replica_count: Number of replica nodes. Valid range for Standard Tier with read replicas enabled is [1-5], default is 2. For basic tier, valid value is 0, default is 0. Defaults to `0`.
        :param pulumi.Input[bool] transit_encryption_enabled: Enable TLS for the Redis instance. Defaults to disabled.
        :param pulumi.Input[bool] wait_until_ready: Whether or not to wait until redis instance to be ready, after creation. Defaults to `true`.
        """
        pulumi.set(__self__, "memory_size_gb", memory_size_gb)
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "tier", tier)
        if auth_enabled is not None:
            pulumi.set(__self__, "auth_enabled", auth_enabled)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if read_replicas_enabled is not None:
            pulumi.set(__self__, "read_replicas_enabled", read_replicas_enabled)
        if redis_configs is not None:
            pulumi.set(__self__, "redis_configs", redis_configs)
        if redis_version is not None:
            pulumi.set(__self__, "redis_version", redis_version)
        if replica_count is not None:
            pulumi.set(__self__, "replica_count", replica_count)
        if transit_encryption_enabled is not None:
            pulumi.set(__self__, "transit_encryption_enabled", transit_encryption_enabled)
        if wait_until_ready is not None:
            pulumi.set(__self__, "wait_until_ready", wait_until_ready)

    @property
    @pulumi.getter(name="memorySizeGb")
    def memory_size_gb(self) -> pulumi.Input[int]:
        """
        Redis memory size in GiB.
        """
        return pulumi.get(self, "memory_size_gb")

    @memory_size_gb.setter
    def memory_size_gb(self, value: pulumi.Input[int]):
        pulumi.set(self, "memory_size_gb", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        GUID of the tenant the Redis instance will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Input[str]:
        """
        Service tier. Must be one of ['BASIC', 'STANDARD_HA'].
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: pulumi.Input[str]):
        pulumi.set(self, "tier", value)

    @property
    @pulumi.getter(name="authEnabled")
    def auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable OSS Redis AUTH. Defaults to false (AUTH disabled).
        """
        return pulumi.get(self, "auth_enabled")

    @auth_enabled.setter
    def auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auth_enabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        User-provided name for the instance.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels for user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Short name of the Redis instance. Duplo adds a prefix; retrieve the full name from `fullname`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="readReplicasEnabled")
    def read_replicas_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable read replica mode (can only be set during instance creation).
        """
        return pulumi.get(self, "read_replicas_enabled")

    @read_replicas_enabled.setter
    def read_replicas_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "read_replicas_enabled", value)

    @property
    @pulumi.getter(name="redisConfigs")
    def redis_configs(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Redis configuration parameters. See https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs for supported parameters.
        """
        return pulumi.get(self, "redis_configs")

    @redis_configs.setter
    def redis_configs(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "redis_configs", value)

    @property
    @pulumi.getter(name="redisVersion")
    def redis_version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of Redis software. Defaults to the latest supported version.
        """
        return pulumi.get(self, "redis_version")

    @redis_version.setter
    def redis_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redis_version", value)

    @property
    @pulumi.getter(name="replicaCount")
    def replica_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number of replica nodes. Valid range for Standard Tier with read replicas enabled is [1-5], default is 2. For basic tier, valid value is 0, default is 0. Defaults to `0`.
        """
        return pulumi.get(self, "replica_count")

    @replica_count.setter
    def replica_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "replica_count", value)

    @property
    @pulumi.getter(name="transitEncryptionEnabled")
    def transit_encryption_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable TLS for the Redis instance. Defaults to disabled.
        """
        return pulumi.get(self, "transit_encryption_enabled")

    @transit_encryption_enabled.setter
    def transit_encryption_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "transit_encryption_enabled", value)

    @property
    @pulumi.getter(name="waitUntilReady")
    def wait_until_ready(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to wait until redis instance to be ready, after creation. Defaults to `true`.
        """
        return pulumi.get(self, "wait_until_ready")

    @wait_until_ready.setter
    def wait_until_ready(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_until_ready", value)


@pulumi.input_type
class _GcpRedisInstanceState:
    def __init__(__self__, *,
                 auth_enabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 fullname: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 memory_size_gb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_replicas_enabled: Optional[pulumi.Input[bool]] = None,
                 redis_configs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 redis_version: Optional[pulumi.Input[str]] = None,
                 replica_count: Optional[pulumi.Input[int]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 transit_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 wait_until_ready: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering GcpRedisInstance resources.
        :param pulumi.Input[bool] auth_enabled: Enable OSS Redis AUTH. Defaults to false (AUTH disabled).
        :param pulumi.Input[str] display_name: User-provided name for the instance.
        :param pulumi.Input[str] fullname: The full name of the of the Redis instance.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels for user-provided metadata.
        :param pulumi.Input[int] memory_size_gb: Redis memory size in GiB.
        :param pulumi.Input[str] name: Short name of the Redis instance. Duplo adds a prefix; retrieve the full name from `fullname`.
        :param pulumi.Input[bool] read_replicas_enabled: Enable read replica mode (can only be set during instance creation).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] redis_configs: Redis configuration parameters. See https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs for supported parameters.
        :param pulumi.Input[str] redis_version: Version of Redis software. Defaults to the latest supported version.
        :param pulumi.Input[int] replica_count: Number of replica nodes. Valid range for Standard Tier with read replicas enabled is [1-5], default is 2. For basic tier, valid value is 0, default is 0. Defaults to `0`.
        :param pulumi.Input[str] tenant_id: GUID of the tenant the Redis instance will be created in.
        :param pulumi.Input[str] tier: Service tier. Must be one of ['BASIC', 'STANDARD_HA'].
        :param pulumi.Input[bool] transit_encryption_enabled: Enable TLS for the Redis instance. Defaults to disabled.
        :param pulumi.Input[bool] wait_until_ready: Whether or not to wait until redis instance to be ready, after creation. Defaults to `true`.
        """
        if auth_enabled is not None:
            pulumi.set(__self__, "auth_enabled", auth_enabled)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if fullname is not None:
            pulumi.set(__self__, "fullname", fullname)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if memory_size_gb is not None:
            pulumi.set(__self__, "memory_size_gb", memory_size_gb)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if read_replicas_enabled is not None:
            pulumi.set(__self__, "read_replicas_enabled", read_replicas_enabled)
        if redis_configs is not None:
            pulumi.set(__self__, "redis_configs", redis_configs)
        if redis_version is not None:
            pulumi.set(__self__, "redis_version", redis_version)
        if replica_count is not None:
            pulumi.set(__self__, "replica_count", replica_count)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)
        if transit_encryption_enabled is not None:
            pulumi.set(__self__, "transit_encryption_enabled", transit_encryption_enabled)
        if wait_until_ready is not None:
            pulumi.set(__self__, "wait_until_ready", wait_until_ready)

    @property
    @pulumi.getter(name="authEnabled")
    def auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable OSS Redis AUTH. Defaults to false (AUTH disabled).
        """
        return pulumi.get(self, "auth_enabled")

    @auth_enabled.setter
    def auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auth_enabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        User-provided name for the instance.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def fullname(self) -> Optional[pulumi.Input[str]]:
        """
        The full name of the of the Redis instance.
        """
        return pulumi.get(self, "fullname")

    @fullname.setter
    def fullname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fullname", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels for user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="memorySizeGb")
    def memory_size_gb(self) -> Optional[pulumi.Input[int]]:
        """
        Redis memory size in GiB.
        """
        return pulumi.get(self, "memory_size_gb")

    @memory_size_gb.setter
    def memory_size_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_size_gb", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Short name of the Redis instance. Duplo adds a prefix; retrieve the full name from `fullname`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="readReplicasEnabled")
    def read_replicas_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable read replica mode (can only be set during instance creation).
        """
        return pulumi.get(self, "read_replicas_enabled")

    @read_replicas_enabled.setter
    def read_replicas_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "read_replicas_enabled", value)

    @property
    @pulumi.getter(name="redisConfigs")
    def redis_configs(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Redis configuration parameters. See https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs for supported parameters.
        """
        return pulumi.get(self, "redis_configs")

    @redis_configs.setter
    def redis_configs(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "redis_configs", value)

    @property
    @pulumi.getter(name="redisVersion")
    def redis_version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of Redis software. Defaults to the latest supported version.
        """
        return pulumi.get(self, "redis_version")

    @redis_version.setter
    def redis_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redis_version", value)

    @property
    @pulumi.getter(name="replicaCount")
    def replica_count(self) -> Optional[pulumi.Input[int]]:
        """
        Number of replica nodes. Valid range for Standard Tier with read replicas enabled is [1-5], default is 2. For basic tier, valid value is 0, default is 0. Defaults to `0`.
        """
        return pulumi.get(self, "replica_count")

    @replica_count.setter
    def replica_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "replica_count", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        GUID of the tenant the Redis instance will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        Service tier. Must be one of ['BASIC', 'STANDARD_HA'].
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)

    @property
    @pulumi.getter(name="transitEncryptionEnabled")
    def transit_encryption_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable TLS for the Redis instance. Defaults to disabled.
        """
        return pulumi.get(self, "transit_encryption_enabled")

    @transit_encryption_enabled.setter
    def transit_encryption_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "transit_encryption_enabled", value)

    @property
    @pulumi.getter(name="waitUntilReady")
    def wait_until_ready(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to wait until redis instance to be ready, after creation. Defaults to `true`.
        """
        return pulumi.get(self, "wait_until_ready")

    @wait_until_ready.setter
    def wait_until_ready(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_until_ready", value)


class GcpRedisInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_enabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 memory_size_gb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_replicas_enabled: Optional[pulumi.Input[bool]] = None,
                 redis_configs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 redis_version: Optional[pulumi.Input[str]] = None,
                 replica_count: Optional[pulumi.Input[int]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 transit_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 wait_until_ready: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        `GcpRedisInstance` manages a GCP redis instance in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        redis_demo = duplocloud.GcpRedisInstance("redis-demo",
            tenant_id=myapp.tenant_id,
            name="redis-demo",
            display_name="redis-demo",
            tier="BASIC",
            redis_version="REDIS_4_0",
            memory_size_gb=1)
        ```

        ## Import

        Example: Importing an existing GCP RedisInstance

         - *TENANT_ID* is the tenant GUID

         - *NAME* is the  name of the RedisInstance

        # 

        ```sh
        $ pulumi import duplocloud:index/gcpRedisInstance:GcpRedisInstance memory-cache *TENANT_ID*/*NAME*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auth_enabled: Enable OSS Redis AUTH. Defaults to false (AUTH disabled).
        :param pulumi.Input[str] display_name: User-provided name for the instance.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels for user-provided metadata.
        :param pulumi.Input[int] memory_size_gb: Redis memory size in GiB.
        :param pulumi.Input[str] name: Short name of the Redis instance. Duplo adds a prefix; retrieve the full name from `fullname`.
        :param pulumi.Input[bool] read_replicas_enabled: Enable read replica mode (can only be set during instance creation).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] redis_configs: Redis configuration parameters. See https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs for supported parameters.
        :param pulumi.Input[str] redis_version: Version of Redis software. Defaults to the latest supported version.
        :param pulumi.Input[int] replica_count: Number of replica nodes. Valid range for Standard Tier with read replicas enabled is [1-5], default is 2. For basic tier, valid value is 0, default is 0. Defaults to `0`.
        :param pulumi.Input[str] tenant_id: GUID of the tenant the Redis instance will be created in.
        :param pulumi.Input[str] tier: Service tier. Must be one of ['BASIC', 'STANDARD_HA'].
        :param pulumi.Input[bool] transit_encryption_enabled: Enable TLS for the Redis instance. Defaults to disabled.
        :param pulumi.Input[bool] wait_until_ready: Whether or not to wait until redis instance to be ready, after creation. Defaults to `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GcpRedisInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `GcpRedisInstance` manages a GCP redis instance in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        myapp = duplocloud.Tenant("myapp",
            account_name="myapp",
            plan_id="default")
        redis_demo = duplocloud.GcpRedisInstance("redis-demo",
            tenant_id=myapp.tenant_id,
            name="redis-demo",
            display_name="redis-demo",
            tier="BASIC",
            redis_version="REDIS_4_0",
            memory_size_gb=1)
        ```

        ## Import

        Example: Importing an existing GCP RedisInstance

         - *TENANT_ID* is the tenant GUID

         - *NAME* is the  name of the RedisInstance

        # 

        ```sh
        $ pulumi import duplocloud:index/gcpRedisInstance:GcpRedisInstance memory-cache *TENANT_ID*/*NAME*
        ```

        :param str resource_name: The name of the resource.
        :param GcpRedisInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GcpRedisInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_enabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 memory_size_gb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 read_replicas_enabled: Optional[pulumi.Input[bool]] = None,
                 redis_configs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 redis_version: Optional[pulumi.Input[str]] = None,
                 replica_count: Optional[pulumi.Input[int]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 transit_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 wait_until_ready: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GcpRedisInstanceArgs.__new__(GcpRedisInstanceArgs)

            __props__.__dict__["auth_enabled"] = auth_enabled
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["labels"] = labels
            if memory_size_gb is None and not opts.urn:
                raise TypeError("Missing required property 'memory_size_gb'")
            __props__.__dict__["memory_size_gb"] = memory_size_gb
            __props__.__dict__["name"] = name
            __props__.__dict__["read_replicas_enabled"] = read_replicas_enabled
            __props__.__dict__["redis_configs"] = redis_configs
            __props__.__dict__["redis_version"] = redis_version
            __props__.__dict__["replica_count"] = replica_count
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__.__dict__["tenant_id"] = tenant_id
            if tier is None and not opts.urn:
                raise TypeError("Missing required property 'tier'")
            __props__.__dict__["tier"] = tier
            __props__.__dict__["transit_encryption_enabled"] = transit_encryption_enabled
            __props__.__dict__["wait_until_ready"] = wait_until_ready
            __props__.__dict__["fullname"] = None
        super(GcpRedisInstance, __self__).__init__(
            'duplocloud:index/gcpRedisInstance:GcpRedisInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_enabled: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            fullname: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            memory_size_gb: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            read_replicas_enabled: Optional[pulumi.Input[bool]] = None,
            redis_configs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            redis_version: Optional[pulumi.Input[str]] = None,
            replica_count: Optional[pulumi.Input[int]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            tier: Optional[pulumi.Input[str]] = None,
            transit_encryption_enabled: Optional[pulumi.Input[bool]] = None,
            wait_until_ready: Optional[pulumi.Input[bool]] = None) -> 'GcpRedisInstance':
        """
        Get an existing GcpRedisInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auth_enabled: Enable OSS Redis AUTH. Defaults to false (AUTH disabled).
        :param pulumi.Input[str] display_name: User-provided name for the instance.
        :param pulumi.Input[str] fullname: The full name of the of the Redis instance.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels for user-provided metadata.
        :param pulumi.Input[int] memory_size_gb: Redis memory size in GiB.
        :param pulumi.Input[str] name: Short name of the Redis instance. Duplo adds a prefix; retrieve the full name from `fullname`.
        :param pulumi.Input[bool] read_replicas_enabled: Enable read replica mode (can only be set during instance creation).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] redis_configs: Redis configuration parameters. See https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs for supported parameters.
        :param pulumi.Input[str] redis_version: Version of Redis software. Defaults to the latest supported version.
        :param pulumi.Input[int] replica_count: Number of replica nodes. Valid range for Standard Tier with read replicas enabled is [1-5], default is 2. For basic tier, valid value is 0, default is 0. Defaults to `0`.
        :param pulumi.Input[str] tenant_id: GUID of the tenant the Redis instance will be created in.
        :param pulumi.Input[str] tier: Service tier. Must be one of ['BASIC', 'STANDARD_HA'].
        :param pulumi.Input[bool] transit_encryption_enabled: Enable TLS for the Redis instance. Defaults to disabled.
        :param pulumi.Input[bool] wait_until_ready: Whether or not to wait until redis instance to be ready, after creation. Defaults to `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GcpRedisInstanceState.__new__(_GcpRedisInstanceState)

        __props__.__dict__["auth_enabled"] = auth_enabled
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["fullname"] = fullname
        __props__.__dict__["labels"] = labels
        __props__.__dict__["memory_size_gb"] = memory_size_gb
        __props__.__dict__["name"] = name
        __props__.__dict__["read_replicas_enabled"] = read_replicas_enabled
        __props__.__dict__["redis_configs"] = redis_configs
        __props__.__dict__["redis_version"] = redis_version
        __props__.__dict__["replica_count"] = replica_count
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["tier"] = tier
        __props__.__dict__["transit_encryption_enabled"] = transit_encryption_enabled
        __props__.__dict__["wait_until_ready"] = wait_until_ready
        return GcpRedisInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authEnabled")
    def auth_enabled(self) -> pulumi.Output[bool]:
        """
        Enable OSS Redis AUTH. Defaults to false (AUTH disabled).
        """
        return pulumi.get(self, "auth_enabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        User-provided name for the instance.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def fullname(self) -> pulumi.Output[str]:
        """
        The full name of the of the Redis instance.
        """
        return pulumi.get(self, "fullname")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Resource labels for user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="memorySizeGb")
    def memory_size_gb(self) -> pulumi.Output[int]:
        """
        Redis memory size in GiB.
        """
        return pulumi.get(self, "memory_size_gb")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Short name of the Redis instance. Duplo adds a prefix; retrieve the full name from `fullname`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="readReplicasEnabled")
    def read_replicas_enabled(self) -> pulumi.Output[bool]:
        """
        Enable read replica mode (can only be set during instance creation).
        """
        return pulumi.get(self, "read_replicas_enabled")

    @property
    @pulumi.getter(name="redisConfigs")
    def redis_configs(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Redis configuration parameters. See https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs for supported parameters.
        """
        return pulumi.get(self, "redis_configs")

    @property
    @pulumi.getter(name="redisVersion")
    def redis_version(self) -> pulumi.Output[str]:
        """
        Version of Redis software. Defaults to the latest supported version.
        """
        return pulumi.get(self, "redis_version")

    @property
    @pulumi.getter(name="replicaCount")
    def replica_count(self) -> pulumi.Output[Optional[int]]:
        """
        Number of replica nodes. Valid range for Standard Tier with read replicas enabled is [1-5], default is 2. For basic tier, valid value is 0, default is 0. Defaults to `0`.
        """
        return pulumi.get(self, "replica_count")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        GUID of the tenant the Redis instance will be created in.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Output[str]:
        """
        Service tier. Must be one of ['BASIC', 'STANDARD_HA'].
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter(name="transitEncryptionEnabled")
    def transit_encryption_enabled(self) -> pulumi.Output[bool]:
        """
        Enable TLS for the Redis instance. Defaults to disabled.
        """
        return pulumi.get(self, "transit_encryption_enabled")

    @property
    @pulumi.getter(name="waitUntilReady")
    def wait_until_ready(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to wait until redis instance to be ready, after creation. Defaults to `true`.
        """
        return pulumi.get(self, "wait_until_ready")

