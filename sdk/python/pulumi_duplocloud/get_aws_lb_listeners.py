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
    'GetAwsLbListenersResult',
    'AwaitableGetAwsLbListenersResult',
    'get_aws_lb_listeners',
    'get_aws_lb_listeners_output',
]

@pulumi.output_type
class GetAwsLbListenersResult:
    """
    A collection of values returned by getAwsLbListeners.
    """
    def __init__(__self__, id=None, listeners=None, name=None, tenant_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listeners and not isinstance(listeners, list):
            raise TypeError("Expected argument 'listeners' to be a list")
        pulumi.set(__self__, "listeners", listeners)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def listeners(self) -> Sequence['outputs.GetAwsLbListenersListenerResult']:
        return pulumi.get(self, "listeners")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        return pulumi.get(self, "tenant_id")


class AwaitableGetAwsLbListenersResult(GetAwsLbListenersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAwsLbListenersResult(
            id=self.id,
            listeners=self.listeners,
            name=self.name,
            tenant_id=self.tenant_id)


def get_aws_lb_listeners(name: Optional[str] = None,
                         tenant_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAwsLbListenersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('duplocloud:index/getAwsLbListeners:getAwsLbListeners', __args__, opts=opts, typ=GetAwsLbListenersResult).value

    return AwaitableGetAwsLbListenersResult(
        id=pulumi.get(__ret__, 'id'),
        listeners=pulumi.get(__ret__, 'listeners'),
        name=pulumi.get(__ret__, 'name'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_aws_lb_listeners_output(name: Optional[pulumi.Input[str]] = None,
                                tenant_id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAwsLbListenersResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('duplocloud:index/getAwsLbListeners:getAwsLbListeners', __args__, opts=opts, typ=GetAwsLbListenersResult)
    return __ret__.apply(lambda __response__: GetAwsLbListenersResult(
        id=pulumi.get(__response__, 'id'),
        listeners=pulumi.get(__response__, 'listeners'),
        name=pulumi.get(__response__, 'name'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
