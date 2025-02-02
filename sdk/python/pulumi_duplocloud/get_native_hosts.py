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
    'GetNativeHostsResult',
    'AwaitableGetNativeHostsResult',
    'get_native_hosts',
    'get_native_hosts_output',
]

@pulumi.output_type
class GetNativeHostsResult:
    """
    A collection of values returned by getNativeHosts.
    """
    def __init__(__self__, hosts=None, id=None, tenant_id=None):
        if hosts and not isinstance(hosts, list):
            raise TypeError("Expected argument 'hosts' to be a list")
        pulumi.set(__self__, "hosts", hosts)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def hosts(self) -> Sequence['outputs.GetNativeHostsHostResult']:
        """
        The list of native hosts.
        """
        return pulumi.get(self, "hosts")

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


class AwaitableGetNativeHostsResult(GetNativeHostsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNativeHostsResult(
            hosts=self.hosts,
            id=self.id,
            tenant_id=self.tenant_id)


def get_native_hosts(tenant_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNativeHostsResult:
    """
    `get_native_hosts` lists native hosts in a Duplo tenant.


    :param str tenant_id: The GUID of the tenant in which to list the hosts.
    """
    __args__ = dict()
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('duplocloud:index/getNativeHosts:getNativeHosts', __args__, opts=opts, typ=GetNativeHostsResult).value

    return AwaitableGetNativeHostsResult(
        hosts=pulumi.get(__ret__, 'hosts'),
        id=pulumi.get(__ret__, 'id'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_native_hosts_output(tenant_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNativeHostsResult]:
    """
    `get_native_hosts` lists native hosts in a Duplo tenant.


    :param str tenant_id: The GUID of the tenant in which to list the hosts.
    """
    __args__ = dict()
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('duplocloud:index/getNativeHosts:getNativeHosts', __args__, opts=opts, typ=GetNativeHostsResult)
    return __ret__.apply(lambda __response__: GetNativeHostsResult(
        hosts=pulumi.get(__response__, 'hosts'),
        id=pulumi.get(__response__, 'id'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
