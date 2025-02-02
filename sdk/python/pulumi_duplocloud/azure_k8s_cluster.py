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

__all__ = ['AzureK8sClusterArgs', 'AzureK8sCluster']

@pulumi.input_type
class AzureK8sClusterArgs:
    def __init__(__self__, *,
                 infra_name: pulumi.Input[str],
                 kubernetes_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_plugin: Optional[pulumi.Input[str]] = None,
                 outbound_type: Optional[pulumi.Input[str]] = None,
                 private_cluster_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 vm_size: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AzureK8sCluster resource.
        :param pulumi.Input[str] infra_name: The name of the infrastructure.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the AKS managed cluster.
        :param pulumi.Input[str] name: The name of the aks. If not specified default name would be infra name
        :param pulumi.Input[str] network_plugin: Network plugin to use for networking. Valid values are: `azure` and `kubenet`.
        :param pulumi.Input[str] outbound_type: The outbound (egress) routing method which should be used for this Kubernetes Cluster. Valid values are: `loadBalancer` and `userDefinedRouting`.
        :param pulumi.Input[bool] private_cluster_enabled: Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: The name of the aks resource group.
        :param pulumi.Input[str] vm_size: The size of the Virtual Machine.
        """
        pulumi.set(__self__, "infra_name", infra_name)
        if kubernetes_version is not None:
            pulumi.set(__self__, "kubernetes_version", kubernetes_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_plugin is not None:
            pulumi.set(__self__, "network_plugin", network_plugin)
        if outbound_type is not None:
            pulumi.set(__self__, "outbound_type", outbound_type)
        if private_cluster_enabled is not None:
            pulumi.set(__self__, "private_cluster_enabled", private_cluster_enabled)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if vm_size is not None:
            pulumi.set(__self__, "vm_size", vm_size)

    @property
    @pulumi.getter(name="infraName")
    def infra_name(self) -> pulumi.Input[str]:
        """
        The name of the infrastructure.
        """
        return pulumi.get(self, "infra_name")

    @infra_name.setter
    def infra_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "infra_name", value)

    @property
    @pulumi.getter(name="kubernetesVersion")
    def kubernetes_version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of Kubernetes specified when creating the AKS managed cluster.
        """
        return pulumi.get(self, "kubernetes_version")

    @kubernetes_version.setter
    def kubernetes_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the aks. If not specified default name would be infra name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkPlugin")
    def network_plugin(self) -> Optional[pulumi.Input[str]]:
        """
        Network plugin to use for networking. Valid values are: `azure` and `kubenet`.
        """
        return pulumi.get(self, "network_plugin")

    @network_plugin.setter
    def network_plugin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_plugin", value)

    @property
    @pulumi.getter(name="outboundType")
    def outbound_type(self) -> Optional[pulumi.Input[str]]:
        """
        The outbound (egress) routing method which should be used for this Kubernetes Cluster. Valid values are: `loadBalancer` and `userDefinedRouting`.
        """
        return pulumi.get(self, "outbound_type")

    @outbound_type.setter
    def outbound_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outbound_type", value)

    @property
    @pulumi.getter(name="privateClusterEnabled")
    def private_cluster_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`.
        """
        return pulumi.get(self, "private_cluster_enabled")

    @private_cluster_enabled.setter
    def private_cluster_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "private_cluster_enabled", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the aks resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> Optional[pulumi.Input[str]]:
        """
        The size of the Virtual Machine.
        """
        return pulumi.get(self, "vm_size")

    @vm_size.setter
    def vm_size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vm_size", value)


@pulumi.input_type
class _AzureK8sClusterState:
    def __init__(__self__, *,
                 infra_name: Optional[pulumi.Input[str]] = None,
                 kubernetes_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_plugin: Optional[pulumi.Input[str]] = None,
                 outbound_type: Optional[pulumi.Input[str]] = None,
                 private_cluster_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 vm_size: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AzureK8sCluster resources.
        :param pulumi.Input[str] infra_name: The name of the infrastructure.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the AKS managed cluster.
        :param pulumi.Input[str] name: The name of the aks. If not specified default name would be infra name
        :param pulumi.Input[str] network_plugin: Network plugin to use for networking. Valid values are: `azure` and `kubenet`.
        :param pulumi.Input[str] outbound_type: The outbound (egress) routing method which should be used for this Kubernetes Cluster. Valid values are: `loadBalancer` and `userDefinedRouting`.
        :param pulumi.Input[bool] private_cluster_enabled: Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: The name of the aks resource group.
        :param pulumi.Input[str] vm_size: The size of the Virtual Machine.
        """
        if infra_name is not None:
            pulumi.set(__self__, "infra_name", infra_name)
        if kubernetes_version is not None:
            pulumi.set(__self__, "kubernetes_version", kubernetes_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_plugin is not None:
            pulumi.set(__self__, "network_plugin", network_plugin)
        if outbound_type is not None:
            pulumi.set(__self__, "outbound_type", outbound_type)
        if private_cluster_enabled is not None:
            pulumi.set(__self__, "private_cluster_enabled", private_cluster_enabled)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if vm_size is not None:
            pulumi.set(__self__, "vm_size", vm_size)

    @property
    @pulumi.getter(name="infraName")
    def infra_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the infrastructure.
        """
        return pulumi.get(self, "infra_name")

    @infra_name.setter
    def infra_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "infra_name", value)

    @property
    @pulumi.getter(name="kubernetesVersion")
    def kubernetes_version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of Kubernetes specified when creating the AKS managed cluster.
        """
        return pulumi.get(self, "kubernetes_version")

    @kubernetes_version.setter
    def kubernetes_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the aks. If not specified default name would be infra name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkPlugin")
    def network_plugin(self) -> Optional[pulumi.Input[str]]:
        """
        Network plugin to use for networking. Valid values are: `azure` and `kubenet`.
        """
        return pulumi.get(self, "network_plugin")

    @network_plugin.setter
    def network_plugin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_plugin", value)

    @property
    @pulumi.getter(name="outboundType")
    def outbound_type(self) -> Optional[pulumi.Input[str]]:
        """
        The outbound (egress) routing method which should be used for this Kubernetes Cluster. Valid values are: `loadBalancer` and `userDefinedRouting`.
        """
        return pulumi.get(self, "outbound_type")

    @outbound_type.setter
    def outbound_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outbound_type", value)

    @property
    @pulumi.getter(name="privateClusterEnabled")
    def private_cluster_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`.
        """
        return pulumi.get(self, "private_cluster_enabled")

    @private_cluster_enabled.setter
    def private_cluster_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "private_cluster_enabled", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the aks resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> Optional[pulumi.Input[str]]:
        """
        The size of the Virtual Machine.
        """
        return pulumi.get(self, "vm_size")

    @vm_size.setter
    def vm_size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vm_size", value)


class AzureK8sCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 infra_name: Optional[pulumi.Input[str]] = None,
                 kubernetes_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_plugin: Optional[pulumi.Input[str]] = None,
                 outbound_type: Optional[pulumi.Input[str]] = None,
                 private_cluster_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 vm_size: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `AzureK8sCluster` manages an azure kubernetes cluster in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.Infrastructure("infra",
            infra_name="tst-0206",
            account_id="143ffc59-9394-4ec6-8f5a-c408a238be62",
            cloud=2,
            azcount=2,
            region="West US 2",
            enable_k8_cluster=False,
            address_prefix="10.50.0.0/16",
            subnet_cidr=0,
            subnet_name="sub01",
            subnet_address_prefix="10.50.1.0/24")
        cluster = duplocloud.AzureK8sCluster("cluster", infra_name=infra.infra_name)
        ```

        ## Import

        Example: Importing an existing Azure K8S Cluster

         - *INFRA_NAME* is the infrastructure name

        # 

        ```sh
        $ pulumi import duplocloud:index/azureK8sCluster:AzureK8sCluster cluster v2/admin/InfrastructureV2/*INFRA_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] infra_name: The name of the infrastructure.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the AKS managed cluster.
        :param pulumi.Input[str] name: The name of the aks. If not specified default name would be infra name
        :param pulumi.Input[str] network_plugin: Network plugin to use for networking. Valid values are: `azure` and `kubenet`.
        :param pulumi.Input[str] outbound_type: The outbound (egress) routing method which should be used for this Kubernetes Cluster. Valid values are: `loadBalancer` and `userDefinedRouting`.
        :param pulumi.Input[bool] private_cluster_enabled: Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: The name of the aks resource group.
        :param pulumi.Input[str] vm_size: The size of the Virtual Machine.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AzureK8sClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `AzureK8sCluster` manages an azure kubernetes cluster in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.Infrastructure("infra",
            infra_name="tst-0206",
            account_id="143ffc59-9394-4ec6-8f5a-c408a238be62",
            cloud=2,
            azcount=2,
            region="West US 2",
            enable_k8_cluster=False,
            address_prefix="10.50.0.0/16",
            subnet_cidr=0,
            subnet_name="sub01",
            subnet_address_prefix="10.50.1.0/24")
        cluster = duplocloud.AzureK8sCluster("cluster", infra_name=infra.infra_name)
        ```

        ## Import

        Example: Importing an existing Azure K8S Cluster

         - *INFRA_NAME* is the infrastructure name

        # 

        ```sh
        $ pulumi import duplocloud:index/azureK8sCluster:AzureK8sCluster cluster v2/admin/InfrastructureV2/*INFRA_NAME*
        ```

        :param str resource_name: The name of the resource.
        :param AzureK8sClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AzureK8sClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 infra_name: Optional[pulumi.Input[str]] = None,
                 kubernetes_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_plugin: Optional[pulumi.Input[str]] = None,
                 outbound_type: Optional[pulumi.Input[str]] = None,
                 private_cluster_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 vm_size: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AzureK8sClusterArgs.__new__(AzureK8sClusterArgs)

            if infra_name is None and not opts.urn:
                raise TypeError("Missing required property 'infra_name'")
            __props__.__dict__["infra_name"] = infra_name
            __props__.__dict__["kubernetes_version"] = kubernetes_version
            __props__.__dict__["name"] = name
            __props__.__dict__["network_plugin"] = network_plugin
            __props__.__dict__["outbound_type"] = outbound_type
            __props__.__dict__["private_cluster_enabled"] = private_cluster_enabled
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["vm_size"] = vm_size
        super(AzureK8sCluster, __self__).__init__(
            'duplocloud:index/azureK8sCluster:AzureK8sCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            infra_name: Optional[pulumi.Input[str]] = None,
            kubernetes_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_plugin: Optional[pulumi.Input[str]] = None,
            outbound_type: Optional[pulumi.Input[str]] = None,
            private_cluster_enabled: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            vm_size: Optional[pulumi.Input[str]] = None) -> 'AzureK8sCluster':
        """
        Get an existing AzureK8sCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] infra_name: The name of the infrastructure.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the AKS managed cluster.
        :param pulumi.Input[str] name: The name of the aks. If not specified default name would be infra name
        :param pulumi.Input[str] network_plugin: Network plugin to use for networking. Valid values are: `azure` and `kubenet`.
        :param pulumi.Input[str] outbound_type: The outbound (egress) routing method which should be used for this Kubernetes Cluster. Valid values are: `loadBalancer` and `userDefinedRouting`.
        :param pulumi.Input[bool] private_cluster_enabled: Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: The name of the aks resource group.
        :param pulumi.Input[str] vm_size: The size of the Virtual Machine.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AzureK8sClusterState.__new__(_AzureK8sClusterState)

        __props__.__dict__["infra_name"] = infra_name
        __props__.__dict__["kubernetes_version"] = kubernetes_version
        __props__.__dict__["name"] = name
        __props__.__dict__["network_plugin"] = network_plugin
        __props__.__dict__["outbound_type"] = outbound_type
        __props__.__dict__["private_cluster_enabled"] = private_cluster_enabled
        __props__.__dict__["resource_group_name"] = resource_group_name
        __props__.__dict__["vm_size"] = vm_size
        return AzureK8sCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="infraName")
    def infra_name(self) -> pulumi.Output[str]:
        """
        The name of the infrastructure.
        """
        return pulumi.get(self, "infra_name")

    @property
    @pulumi.getter(name="kubernetesVersion")
    def kubernetes_version(self) -> pulumi.Output[str]:
        """
        Version of Kubernetes specified when creating the AKS managed cluster.
        """
        return pulumi.get(self, "kubernetes_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the aks. If not specified default name would be infra name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkPlugin")
    def network_plugin(self) -> pulumi.Output[str]:
        """
        Network plugin to use for networking. Valid values are: `azure` and `kubenet`.
        """
        return pulumi.get(self, "network_plugin")

    @property
    @pulumi.getter(name="outboundType")
    def outbound_type(self) -> pulumi.Output[str]:
        """
        The outbound (egress) routing method which should be used for this Kubernetes Cluster. Valid values are: `loadBalancer` and `userDefinedRouting`.
        """
        return pulumi.get(self, "outbound_type")

    @property
    @pulumi.getter(name="privateClusterEnabled")
    def private_cluster_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`.
        """
        return pulumi.get(self, "private_cluster_enabled")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the aks resource group.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> pulumi.Output[str]:
        """
        The size of the Virtual Machine.
        """
        return pulumi.get(self, "vm_size")

