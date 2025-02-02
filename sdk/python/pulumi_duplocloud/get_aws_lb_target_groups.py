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
    'GetAwsLbTargetGroupsResult',
    'AwaitableGetAwsLbTargetGroupsResult',
    'get_aws_lb_target_groups',
    'get_aws_lb_target_groups_output',
]

@pulumi.output_type
class GetAwsLbTargetGroupsResult:
    """
    A collection of values returned by getAwsLbTargetGroups.
    """
    def __init__(__self__, id=None, target_groups=None, tenant_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if target_groups and not isinstance(target_groups, list):
            raise TypeError("Expected argument 'target_groups' to be a list")
        pulumi.set(__self__, "target_groups", target_groups)
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
    @pulumi.getter(name="targetGroups")
    def target_groups(self) -> Sequence['outputs.GetAwsLbTargetGroupsTargetGroupResult']:
        return pulumi.get(self, "target_groups")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        return pulumi.get(self, "tenant_id")


class AwaitableGetAwsLbTargetGroupsResult(GetAwsLbTargetGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAwsLbTargetGroupsResult(
            id=self.id,
            target_groups=self.target_groups,
            tenant_id=self.tenant_id)


def get_aws_lb_target_groups(tenant_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAwsLbTargetGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('duplocloud:index/getAwsLbTargetGroups:getAwsLbTargetGroups', __args__, opts=opts, typ=GetAwsLbTargetGroupsResult).value

    return AwaitableGetAwsLbTargetGroupsResult(
        id=pulumi.get(__ret__, 'id'),
        target_groups=pulumi.get(__ret__, 'target_groups'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_aws_lb_target_groups_output(tenant_id: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAwsLbTargetGroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('duplocloud:index/getAwsLbTargetGroups:getAwsLbTargetGroups', __args__, opts=opts, typ=GetAwsLbTargetGroupsResult)
    return __ret__.apply(lambda __response__: GetAwsLbTargetGroupsResult(
        id=pulumi.get(__response__, 'id'),
        target_groups=pulumi.get(__response__, 'target_groups'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
