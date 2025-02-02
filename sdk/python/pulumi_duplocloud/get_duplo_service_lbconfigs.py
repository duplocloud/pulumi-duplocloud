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
    'GetDuploServiceLbconfigsResult',
    'AwaitableGetDuploServiceLbconfigsResult',
    'get_duplo_service_lbconfigs',
    'get_duplo_service_lbconfigs_output',
]

@pulumi.output_type
class GetDuploServiceLbconfigsResult:
    """
    A collection of values returned by getDuploServiceLbconfigs.
    """
    def __init__(__self__, id=None, name=None, services=None, tenant_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)
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
    def name(self) -> Optional[str]:
        """
        The name of the duplo service.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetDuploServiceLbconfigsServiceResult']:
        return pulumi.get(self, "services")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The GUID of the tenant that hosts the duplo service.
        """
        return pulumi.get(self, "tenant_id")


class AwaitableGetDuploServiceLbconfigsResult(GetDuploServiceLbconfigsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDuploServiceLbconfigsResult(
            id=self.id,
            name=self.name,
            services=self.services,
            tenant_id=self.tenant_id)


def get_duplo_service_lbconfigs(name: Optional[str] = None,
                                tenant_id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDuploServiceLbconfigsResult:
    """
    `DuploServiceLbconfigs` retrieves load balancer configuration(s) for container-based service(s) in Duplo.

    NOTE: For Amazon ECS services, see the `get_ecs_services` data source.


    :param str name: The name of the duplo service.
    :param str tenant_id: The GUID of the tenant that hosts the duplo service.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('duplocloud:index/getDuploServiceLbconfigs:getDuploServiceLbconfigs', __args__, opts=opts, typ=GetDuploServiceLbconfigsResult).value

    return AwaitableGetDuploServiceLbconfigsResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        services=pulumi.get(__ret__, 'services'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_duplo_service_lbconfigs_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                       tenant_id: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDuploServiceLbconfigsResult]:
    """
    `DuploServiceLbconfigs` retrieves load balancer configuration(s) for container-based service(s) in Duplo.

    NOTE: For Amazon ECS services, see the `get_ecs_services` data source.


    :param str name: The name of the duplo service.
    :param str tenant_id: The GUID of the tenant that hosts the duplo service.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('duplocloud:index/getDuploServiceLbconfigs:getDuploServiceLbconfigs', __args__, opts=opts, typ=GetDuploServiceLbconfigsResult)
    return __ret__.apply(lambda __response__: GetDuploServiceLbconfigsResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        services=pulumi.get(__response__, 'services'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
