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
    'GetDuploServiceParamsResult',
    'AwaitableGetDuploServiceParamsResult',
    'get_duplo_service_params',
    'get_duplo_service_params_output',
]

@pulumi.output_type
class GetDuploServiceParamsResult:
    """
    A collection of values returned by getDuploServiceParams.
    """
    def __init__(__self__, dns_prfx=None, id=None, replication_controller_name=None, results=None, tenant_id=None, webaclid=None):
        if dns_prfx and not isinstance(dns_prfx, str):
            raise TypeError("Expected argument 'dns_prfx' to be a str")
        pulumi.set(__self__, "dns_prfx", dns_prfx)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if replication_controller_name and not isinstance(replication_controller_name, str):
            raise TypeError("Expected argument 'replication_controller_name' to be a str")
        pulumi.set(__self__, "replication_controller_name", replication_controller_name)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)
        if webaclid and not isinstance(webaclid, str):
            raise TypeError("Expected argument 'webaclid' to be a str")
        pulumi.set(__self__, "webaclid", webaclid)

    @property
    @pulumi.getter(name="dnsPrfx")
    def dns_prfx(self) -> Optional[str]:
        return pulumi.get(self, "dns_prfx")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="replicationControllerName")
    def replication_controller_name(self) -> Optional[str]:
        return pulumi.get(self, "replication_controller_name")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetDuploServiceParamsResultResult']:
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def webaclid(self) -> Optional[str]:
        return pulumi.get(self, "webaclid")


class AwaitableGetDuploServiceParamsResult(GetDuploServiceParamsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDuploServiceParamsResult(
            dns_prfx=self.dns_prfx,
            id=self.id,
            replication_controller_name=self.replication_controller_name,
            results=self.results,
            tenant_id=self.tenant_id,
            webaclid=self.webaclid)


def get_duplo_service_params(dns_prfx: Optional[str] = None,
                             replication_controller_name: Optional[str] = None,
                             tenant_id: Optional[str] = None,
                             webaclid: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDuploServiceParamsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['dnsPrfx'] = dns_prfx
    __args__['replicationControllerName'] = replication_controller_name
    __args__['tenantId'] = tenant_id
    __args__['webaclid'] = webaclid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('duplocloud:index/getDuploServiceParams:getDuploServiceParams', __args__, opts=opts, typ=GetDuploServiceParamsResult).value

    return AwaitableGetDuploServiceParamsResult(
        dns_prfx=pulumi.get(__ret__, 'dns_prfx'),
        id=pulumi.get(__ret__, 'id'),
        replication_controller_name=pulumi.get(__ret__, 'replication_controller_name'),
        results=pulumi.get(__ret__, 'results'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'),
        webaclid=pulumi.get(__ret__, 'webaclid'))
def get_duplo_service_params_output(dns_prfx: Optional[pulumi.Input[Optional[str]]] = None,
                                    replication_controller_name: Optional[pulumi.Input[Optional[str]]] = None,
                                    tenant_id: Optional[pulumi.Input[str]] = None,
                                    webaclid: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDuploServiceParamsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['dnsPrfx'] = dns_prfx
    __args__['replicationControllerName'] = replication_controller_name
    __args__['tenantId'] = tenant_id
    __args__['webaclid'] = webaclid
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('duplocloud:index/getDuploServiceParams:getDuploServiceParams', __args__, opts=opts, typ=GetDuploServiceParamsResult)
    return __ret__.apply(lambda __response__: GetDuploServiceParamsResult(
        dns_prfx=pulumi.get(__response__, 'dns_prfx'),
        id=pulumi.get(__response__, 'id'),
        replication_controller_name=pulumi.get(__response__, 'replication_controller_name'),
        results=pulumi.get(__response__, 'results'),
        tenant_id=pulumi.get(__response__, 'tenant_id'),
        webaclid=pulumi.get(__response__, 'webaclid')))
