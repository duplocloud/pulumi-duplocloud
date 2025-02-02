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
from . import outputs

__all__ = [
    'GetEmrClusterResult',
    'AwaitableGetEmrClusterResult',
    'get_emr_cluster',
    'get_emr_cluster_output',
]

@pulumi.output_type
class GetEmrClusterResult:
    """
    A collection of values returned by getEmrCluster.
    """
    def __init__(__self__, datas=None, id=None, tenant_id=None):
        if datas and not isinstance(datas, list):
            raise TypeError("Expected argument 'datas' to be a list")
        pulumi.set(__self__, "datas", datas)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def datas(self) -> Sequence['outputs.GetEmrClusterDataResult']:
        """
        The list of native hosts.
        """
        return pulumi.get(self, "datas")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The GUID of the tenant in which to list the hosts.
        """
        return pulumi.get(self, "tenant_id")


class AwaitableGetEmrClusterResult(GetEmrClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEmrClusterResult(
            datas=self.datas,
            id=self.id,
            tenant_id=self.tenant_id)


def get_emr_cluster(tenant_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEmrClusterResult:
    """
    `EmrCluster` lists EmrClusters in a Duplo tenant.


    :param str tenant_id: The GUID of the tenant in which to list the hosts.
    """
    __args__ = dict()
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('duplocloud:index/getEmrCluster:getEmrCluster', __args__, opts=opts, typ=GetEmrClusterResult).value

    return AwaitableGetEmrClusterResult(
        datas=pulumi.get(__ret__, 'datas'),
        id=pulumi.get(__ret__, 'id'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_emr_cluster_output(tenant_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEmrClusterResult]:
    """
    `EmrCluster` lists EmrClusters in a Duplo tenant.


    :param str tenant_id: The GUID of the tenant in which to list the hosts.
    """
    __args__ = dict()
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('duplocloud:index/getEmrCluster:getEmrCluster', __args__, opts=opts, typ=GetEmrClusterResult)
    return __ret__.apply(lambda __response__: GetEmrClusterResult(
        datas=pulumi.get(__response__, 'datas'),
        id=pulumi.get(__response__, 'id'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
