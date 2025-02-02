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
from ._inputs import *

__all__ = ['OtherAgentsArgs', 'OtherAgents']

@pulumi.input_type
class OtherAgentsArgs:
    def __init__(__self__, *,
                 agents: pulumi.Input[Sequence[pulumi.Input['OtherAgentsAgentArgs']]],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OtherAgents resource.
        :param pulumi.Input[str] name: Resource name to create other agents in duplo.
        """
        pulumi.set(__self__, "agents", agents)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def agents(self) -> pulumi.Input[Sequence[pulumi.Input['OtherAgentsAgentArgs']]]:
        return pulumi.get(self, "agents")

    @agents.setter
    def agents(self, value: pulumi.Input[Sequence[pulumi.Input['OtherAgentsAgentArgs']]]):
        pulumi.set(self, "agents", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name to create other agents in duplo.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _OtherAgentsState:
    def __init__(__self__, *,
                 agents: Optional[pulumi.Input[Sequence[pulumi.Input['OtherAgentsAgentArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OtherAgents resources.
        :param pulumi.Input[str] name: Resource name to create other agents in duplo.
        """
        if agents is not None:
            pulumi.set(__self__, "agents", agents)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def agents(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['OtherAgentsAgentArgs']]]]:
        return pulumi.get(self, "agents")

    @agents.setter
    def agents(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['OtherAgentsAgentArgs']]]]):
        pulumi.set(self, "agents", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name to create other agents in duplo.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class OtherAgents(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents: Optional[pulumi.Input[Sequence[pulumi.Input[Union['OtherAgentsAgentArgs', 'OtherAgentsAgentArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `OtherAgents` manages an other agents in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        agents = duplocloud.OtherAgents("agents",
            name="duplo-agents",
            agents=[{
                "agent_name": "CloudWatchAgent_0",
                "agent_linux_package_path": "https://s3.amazonaws.com/amazoncloudwatch-agent/ubuntu/amd64/latest/amazon-cloudwatch-agent.deb",
                "linux_agent_install_status_cmd": "sudo service amazon-cloudwatch-agent status | grep -wc 'running'",
                "linux_agent_service_name": "amazon-cloudwatch-agent",
                "linux_agent_uninstall_status_cmd": "OS_FAMILY=$(cat /etc/os-release | grep PRETTY_NAME); if [[ $OS_FAMILY == *'Ubuntu'* ]]; then sudo apt-get purge --yes --force-yes amazon-cloudwatch-agent; else sudo yum remove -y AwsAgent; fi",
                "linux_install_cmd": "OS_FAMILY=$(cat /etc/os-release | grep PRETTY_NAME); if [[ $OS_FAMILY == *'Ubuntu'* ]]; then wget https://s3.amazonaws.com/amazoncloudwatch-agent/ubuntu/amd64/latest/amazon-cloudwatch-agent.deb; sudo dpkg -i -E ./amazon-cloudwatch-agent.deb; sudo wget -O /opt/aws/amazon-cloudwatch-agent/etc/amazon-cloudwatch-agent.json https://cf-templates-3qf987fmmv5g-us-east-2.s3.us-east-2.amazonaws.com/amazon-cloudwatch-agent.json; sudo service amazon-cloudwatch-agent restart; else wget https://s3.amazonaws.com/amazoncloudwatch-agent/amazon_linux/amd64/latest/amazon-cloudwatch-agent.rpm; sudo rpm -U ./amazon-cloudwatch-agent.rpm; sudo wget -O /opt/aws/amazon-cloudwatch-agent/etc/amazon-cloudwatch-agent.json https://cf-templates-3qf987fmmv5g-us-east-2.s3.us-east-2.amazonaws.com/amazon-cloudwatch-agent.json && sudo service amazon-cloudwatch-agent restart; fi",
                "windows_agent_service_name": "amazon-cloudwatch-agent",
            }])
        ```

        ## Import

        Example: Importing an existing other agents

         - *SHORTNAME* is the short name of the other agents

        # 

        ```sh
        $ pulumi import duplocloud:index/otherAgents:OtherAgents agents *SHORTNAME*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Resource name to create other agents in duplo.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OtherAgentsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `OtherAgents` manages an other agents in Duplo.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        agents = duplocloud.OtherAgents("agents",
            name="duplo-agents",
            agents=[{
                "agent_name": "CloudWatchAgent_0",
                "agent_linux_package_path": "https://s3.amazonaws.com/amazoncloudwatch-agent/ubuntu/amd64/latest/amazon-cloudwatch-agent.deb",
                "linux_agent_install_status_cmd": "sudo service amazon-cloudwatch-agent status | grep -wc 'running'",
                "linux_agent_service_name": "amazon-cloudwatch-agent",
                "linux_agent_uninstall_status_cmd": "OS_FAMILY=$(cat /etc/os-release | grep PRETTY_NAME); if [[ $OS_FAMILY == *'Ubuntu'* ]]; then sudo apt-get purge --yes --force-yes amazon-cloudwatch-agent; else sudo yum remove -y AwsAgent; fi",
                "linux_install_cmd": "OS_FAMILY=$(cat /etc/os-release | grep PRETTY_NAME); if [[ $OS_FAMILY == *'Ubuntu'* ]]; then wget https://s3.amazonaws.com/amazoncloudwatch-agent/ubuntu/amd64/latest/amazon-cloudwatch-agent.deb; sudo dpkg -i -E ./amazon-cloudwatch-agent.deb; sudo wget -O /opt/aws/amazon-cloudwatch-agent/etc/amazon-cloudwatch-agent.json https://cf-templates-3qf987fmmv5g-us-east-2.s3.us-east-2.amazonaws.com/amazon-cloudwatch-agent.json; sudo service amazon-cloudwatch-agent restart; else wget https://s3.amazonaws.com/amazoncloudwatch-agent/amazon_linux/amd64/latest/amazon-cloudwatch-agent.rpm; sudo rpm -U ./amazon-cloudwatch-agent.rpm; sudo wget -O /opt/aws/amazon-cloudwatch-agent/etc/amazon-cloudwatch-agent.json https://cf-templates-3qf987fmmv5g-us-east-2.s3.us-east-2.amazonaws.com/amazon-cloudwatch-agent.json && sudo service amazon-cloudwatch-agent restart; fi",
                "windows_agent_service_name": "amazon-cloudwatch-agent",
            }])
        ```

        ## Import

        Example: Importing an existing other agents

         - *SHORTNAME* is the short name of the other agents

        # 

        ```sh
        $ pulumi import duplocloud:index/otherAgents:OtherAgents agents *SHORTNAME*
        ```

        :param str resource_name: The name of the resource.
        :param OtherAgentsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OtherAgentsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents: Optional[pulumi.Input[Sequence[pulumi.Input[Union['OtherAgentsAgentArgs', 'OtherAgentsAgentArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OtherAgentsArgs.__new__(OtherAgentsArgs)

            if agents is None and not opts.urn:
                raise TypeError("Missing required property 'agents'")
            __props__.__dict__["agents"] = agents
            __props__.__dict__["name"] = name
        super(OtherAgents, __self__).__init__(
            'duplocloud:index/otherAgents:OtherAgents',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agents: Optional[pulumi.Input[Sequence[pulumi.Input[Union['OtherAgentsAgentArgs', 'OtherAgentsAgentArgsDict']]]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'OtherAgents':
        """
        Get an existing OtherAgents resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Resource name to create other agents in duplo.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OtherAgentsState.__new__(_OtherAgentsState)

        __props__.__dict__["agents"] = agents
        __props__.__dict__["name"] = name
        return OtherAgents(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def agents(self) -> pulumi.Output[Sequence['outputs.OtherAgentsAgent']]:
        return pulumi.get(self, "agents")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name to create other agents in duplo.
        """
        return pulumi.get(self, "name")

