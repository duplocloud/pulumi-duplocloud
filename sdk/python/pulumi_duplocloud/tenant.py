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

__all__ = ['TenantArgs', 'Tenant']

@pulumi.input_type
class TenantArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 plan_id: pulumi.Input[str],
                 allow_deletion: Optional[pulumi.Input[bool]] = None,
                 existing_k8s_namespace: Optional[pulumi.Input[str]] = None,
                 wait_until_created: Optional[pulumi.Input[bool]] = None,
                 wait_until_deleted: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Tenant resource.
        :param pulumi.Input[str] account_name: The name of the tenant. Tenant names are globally unique, and cannot be a prefix of any other tenant name.
        :param pulumi.Input[str] plan_id: The name of the plan under which the tenant will be created.
        :param pulumi.Input[bool] allow_deletion: Whether or not to even try and delete the tenant. *NOTE: This only works if you have disabled deletion protection for the tenant.* Defaults to `false`.
        :param pulumi.Input[str] existing_k8s_namespace: Existing kubernetes namespace to use by the tenant. *NOTE: This is an advanced feature, please contact your DuploCloud administrator for help if you want to use this field.*
        :param pulumi.Input[bool] wait_until_created: Whether or not to wait until Duplo has created the tenant. Defaults to `true`.
        :param pulumi.Input[bool] wait_until_deleted: Whether or not to wait until Duplo has destroyed the tenant. Defaults to `false`.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "plan_id", plan_id)
        if allow_deletion is not None:
            pulumi.set(__self__, "allow_deletion", allow_deletion)
        if existing_k8s_namespace is not None:
            pulumi.set(__self__, "existing_k8s_namespace", existing_k8s_namespace)
        if wait_until_created is not None:
            pulumi.set(__self__, "wait_until_created", wait_until_created)
        if wait_until_deleted is not None:
            pulumi.set(__self__, "wait_until_deleted", wait_until_deleted)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        The name of the tenant. Tenant names are globally unique, and cannot be a prefix of any other tenant name.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> pulumi.Input[str]:
        """
        The name of the plan under which the tenant will be created.
        """
        return pulumi.get(self, "plan_id")

    @plan_id.setter
    def plan_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "plan_id", value)

    @property
    @pulumi.getter(name="allowDeletion")
    def allow_deletion(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to even try and delete the tenant. *NOTE: This only works if you have disabled deletion protection for the tenant.* Defaults to `false`.
        """
        return pulumi.get(self, "allow_deletion")

    @allow_deletion.setter
    def allow_deletion(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_deletion", value)

    @property
    @pulumi.getter(name="existingK8sNamespace")
    def existing_k8s_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Existing kubernetes namespace to use by the tenant. *NOTE: This is an advanced feature, please contact your DuploCloud administrator for help if you want to use this field.*
        """
        return pulumi.get(self, "existing_k8s_namespace")

    @existing_k8s_namespace.setter
    def existing_k8s_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "existing_k8s_namespace", value)

    @property
    @pulumi.getter(name="waitUntilCreated")
    def wait_until_created(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to wait until Duplo has created the tenant. Defaults to `true`.
        """
        return pulumi.get(self, "wait_until_created")

    @wait_until_created.setter
    def wait_until_created(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_until_created", value)

    @property
    @pulumi.getter(name="waitUntilDeleted")
    def wait_until_deleted(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to wait until Duplo has destroyed the tenant. Defaults to `false`.
        """
        return pulumi.get(self, "wait_until_deleted")

    @wait_until_deleted.setter
    def wait_until_deleted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_until_deleted", value)


@pulumi.input_type
class _TenantState:
    def __init__(__self__, *,
                 account_name: Optional[pulumi.Input[str]] = None,
                 allow_deletion: Optional[pulumi.Input[bool]] = None,
                 existing_k8s_namespace: Optional[pulumi.Input[str]] = None,
                 infra_owner: Optional[pulumi.Input[str]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input['TenantPolicyArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TenantTagArgs']]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 wait_until_created: Optional[pulumi.Input[bool]] = None,
                 wait_until_deleted: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Tenant resources.
        :param pulumi.Input[str] account_name: The name of the tenant. Tenant names are globally unique, and cannot be a prefix of any other tenant name.
        :param pulumi.Input[bool] allow_deletion: Whether or not to even try and delete the tenant. *NOTE: This only works if you have disabled deletion protection for the tenant.* Defaults to `false`.
        :param pulumi.Input[str] existing_k8s_namespace: Existing kubernetes namespace to use by the tenant. *NOTE: This is an advanced feature, please contact your DuploCloud administrator for help if you want to use this field.*
        :param pulumi.Input[str] plan_id: The name of the plan under which the tenant will be created.
        :param pulumi.Input[str] tenant_id: A GUID identifying the tenant. This is automatically generated by Duplo.
        :param pulumi.Input[bool] wait_until_created: Whether or not to wait until Duplo has created the tenant. Defaults to `true`.
        :param pulumi.Input[bool] wait_until_deleted: Whether or not to wait until Duplo has destroyed the tenant. Defaults to `false`.
        """
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if allow_deletion is not None:
            pulumi.set(__self__, "allow_deletion", allow_deletion)
        if existing_k8s_namespace is not None:
            pulumi.set(__self__, "existing_k8s_namespace", existing_k8s_namespace)
        if infra_owner is not None:
            pulumi.set(__self__, "infra_owner", infra_owner)
        if plan_id is not None:
            pulumi.set(__self__, "plan_id", plan_id)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if wait_until_created is not None:
            pulumi.set(__self__, "wait_until_created", wait_until_created)
        if wait_until_deleted is not None:
            pulumi.set(__self__, "wait_until_deleted", wait_until_deleted)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the tenant. Tenant names are globally unique, and cannot be a prefix of any other tenant name.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="allowDeletion")
    def allow_deletion(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to even try and delete the tenant. *NOTE: This only works if you have disabled deletion protection for the tenant.* Defaults to `false`.
        """
        return pulumi.get(self, "allow_deletion")

    @allow_deletion.setter
    def allow_deletion(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_deletion", value)

    @property
    @pulumi.getter(name="existingK8sNamespace")
    def existing_k8s_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Existing kubernetes namespace to use by the tenant. *NOTE: This is an advanced feature, please contact your DuploCloud administrator for help if you want to use this field.*
        """
        return pulumi.get(self, "existing_k8s_namespace")

    @existing_k8s_namespace.setter
    def existing_k8s_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "existing_k8s_namespace", value)

    @property
    @pulumi.getter(name="infraOwner")
    def infra_owner(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "infra_owner")

    @infra_owner.setter
    def infra_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "infra_owner", value)

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the plan under which the tenant will be created.
        """
        return pulumi.get(self, "plan_id")

    @plan_id.setter
    def plan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plan_id", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TenantPolicyArgs']]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TenantPolicyArgs']]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TenantTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TenantTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        A GUID identifying the tenant. This is automatically generated by Duplo.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="waitUntilCreated")
    def wait_until_created(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to wait until Duplo has created the tenant. Defaults to `true`.
        """
        return pulumi.get(self, "wait_until_created")

    @wait_until_created.setter
    def wait_until_created(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_until_created", value)

    @property
    @pulumi.getter(name="waitUntilDeleted")
    def wait_until_deleted(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to wait until Duplo has destroyed the tenant. Defaults to `false`.
        """
        return pulumi.get(self, "wait_until_deleted")

    @wait_until_deleted.setter
    def wait_until_deleted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_until_deleted", value)


class Tenant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 allow_deletion: Optional[pulumi.Input[bool]] = None,
                 existing_k8s_namespace: Optional[pulumi.Input[str]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 wait_until_created: Optional[pulumi.Input[bool]] = None,
                 wait_until_deleted: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### Create a DuploCloud tenant named 'prod'.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        # Before creating a tenant, you must first set up the infrastructure. Below is the resource for creating the infrastructure.
        infra = duplocloud.Infrastructure("infra",
            infra_name="prod",
            cloud=0,
            region="us-west-2",
            enable_k8_cluster=False,
            address_prefix="10.11.0.0/16")
        # Use the infrastructure name as the 'plan_id' from the 'duplocloud_infrastructure' resource while creating tenant.
        tenant = duplocloud.Tenant("tenant",
            account_name="prod",
            plan_id=infra.infra_name)
        ```

        ### Create a DuploCloud tenant named 'prod' inside the following prod infra.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.Infrastructure("infra",
            infra_name="prod",
            cloud=0,
            region="us-west-2",
            enable_k8_cluster=False,
            address_prefix="10.11.0.0/16")
        ```

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        # Use the infrastructure name as the 'plan_id' from the 'duplocloud_infrastructure' resource.
        tenant = duplocloud.Tenant("tenant",
            account_name="prod",
            plan_id=infra["infraName"])
        ```

        ### Create a DuploCloud tenant named 'dev' within the 'nonprod' infrastructure.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        # Ensure the 'nonprod' infrastructure is already created before setting up the tenant.
        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        ```

        ### Create a DuploCloud tenant named 'dev' with infra name variable and tenant id as output.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        config = pulumi.Config()
        infra_name = config.get("infraName")
        if infra_name is None:
            infra_name = "nonprod"
        # Ensure the 'nonprod' infrastructure is already created before setting up the tenant.
        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        pulumi.export("tenantId", tenant.tenant_id)
        ```

        ### Create a duplocloud tenant named dev with AWS Cognito power user access in the nonprod infrastructure.

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_duplocloud as duplocloud

        # A prerequisite for creating a tenant is having an existing infrastructure. Here’s how you can reference an existing infrastructure.
        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        # Here’s how to create a tenant by providing the infrastructure name for the plan_id field.
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        # Attaches a managed IAM policy to an IAM role.
        amazon_cognito_power_user = aws.index.IamRolePolicyAttachment("AmazonCognitoPowerUser",
            role=fduploservices-{tenant.account_name},
            policy_arn=arn:aws:iam::aws:policy/AmazonCognitoPowerUser)
        ```

        ### Create a DuploCloud tenant named 'qa' with full access to invoke AWS API Gateway in the nonprod infrastructure.

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_duplocloud as duplocloud

        # A prerequisite for creating a tenant is having an existing infrastructure. Here’s how you can reference an existing infrastructure.
        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        # Here’s how to create a tenant by providing the infrastructure name for the plan_id field.
        tenant = duplocloud.Tenant("tenant",
            account_name="qa",
            plan_id=infra.infra_name)
        # Attaches a managed IAM policy to an IAM role.
        amazon_api_gateway_invoke_full_access = aws.index.IamRolePolicyAttachment("AmazonAPIGatewayInvokeFullAccess",
            role=fduploservices-{tenant.account_name},
            policy_arn=arn:aws:iam::aws:policy/AmazonAPIGatewayInvokeFullAccess)
        ```

        ### Create duplocloud tenant named dev with security group rule to allow access from 10.220.0.0/16 on port 5432 in nonprod infra’

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        # Allow communication on port 5432 for the PostgreSQL database from the 10.220.0.0/16 subnet
        allow_from_vpn = duplocloud.TenantNetworkSecurityRule("allow_from_vpn",
            tenant_id=tenant.tenant_id,
            source_address="10.220.0.0/16",
            protocol="tcp",
            from_port=5432,
            to_port=5432,
            description="Allow communication from 10.220.0.0/16 on port 5432.")
        ```

        ### Setup duplocloud tenant named dev with security group rule to allow access from 10.220.0.0/16 on port 22 in nonprod infra’

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        # Allow communication on port 22 from the 10.220.0.0/16 subnet.
        allow_from_vpn = duplocloud.TenantNetworkSecurityRule("allow_from_vpn",
            tenant_id=tenant.tenant_id,
            source_address="10.220.0.0/16",
            protocol="tcp",
            from_port=22,
            to_port=22,
            description="Allow communication from 10.220.0.0/16 on port 22.")
        ```

        ### Provision a tenant named 'myapp' within the infrastructure 'myinfra' and disable deletion protection.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.get_infrastructure(infra_name="myinfra")
        tenant = duplocloud.Tenant("tenant",
            account_name="myapp",
            plan_id=infra.infra_name,
            allow_deletion=True)
        # Reference the tenant_id field from the duplocloud_tenant resource.
        tenant_config = duplocloud.TenantConfig("tenant_config",
            tenant_id=tenant.tenant_id,
            settings=[{
                "key": "delete_protection",
                "value": "false",
            }])
        ```

        ### Provision a tenant named 'myapp' within the infrastructure 'myinfra', and ensure that the S3 bucket has public access blocked and SSL enforcement enabled in the S3 bucket policy.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.get_infrastructure(infra_name="myinfra")
        tenant = duplocloud.Tenant("tenant",
            account_name="myapp",
            plan_id=infra.infra_name,
            allow_deletion=True)
        # Reference the tenant_id field from the duplocloud_tenant resource.
        tenant_config = duplocloud.TenantConfig("tenant_config",
            tenant_id=tenant.tenant_id,
            settings=[
                {
                    "key": "block_public_access_to_s3",
                    "value": "true",
                },
                {
                    "key": "enforce_ssl_for_s3",
                    "value": "true",
                },
            ])
        ```

        ## Import

        ```sh
        $ pulumi import duplocloud:index/tenant:Tenant myapp v2/admin/TenantV2/*TENANT_ID*
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the tenant. Tenant names are globally unique, and cannot be a prefix of any other tenant name.
        :param pulumi.Input[bool] allow_deletion: Whether or not to even try and delete the tenant. *NOTE: This only works if you have disabled deletion protection for the tenant.* Defaults to `false`.
        :param pulumi.Input[str] existing_k8s_namespace: Existing kubernetes namespace to use by the tenant. *NOTE: This is an advanced feature, please contact your DuploCloud administrator for help if you want to use this field.*
        :param pulumi.Input[str] plan_id: The name of the plan under which the tenant will be created.
        :param pulumi.Input[bool] wait_until_created: Whether or not to wait until Duplo has created the tenant. Defaults to `true`.
        :param pulumi.Input[bool] wait_until_deleted: Whether or not to wait until Duplo has destroyed the tenant. Defaults to `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TenantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### Create a DuploCloud tenant named 'prod'.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        # Before creating a tenant, you must first set up the infrastructure. Below is the resource for creating the infrastructure.
        infra = duplocloud.Infrastructure("infra",
            infra_name="prod",
            cloud=0,
            region="us-west-2",
            enable_k8_cluster=False,
            address_prefix="10.11.0.0/16")
        # Use the infrastructure name as the 'plan_id' from the 'duplocloud_infrastructure' resource while creating tenant.
        tenant = duplocloud.Tenant("tenant",
            account_name="prod",
            plan_id=infra.infra_name)
        ```

        ### Create a DuploCloud tenant named 'prod' inside the following prod infra.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.Infrastructure("infra",
            infra_name="prod",
            cloud=0,
            region="us-west-2",
            enable_k8_cluster=False,
            address_prefix="10.11.0.0/16")
        ```

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        # Use the infrastructure name as the 'plan_id' from the 'duplocloud_infrastructure' resource.
        tenant = duplocloud.Tenant("tenant",
            account_name="prod",
            plan_id=infra["infraName"])
        ```

        ### Create a DuploCloud tenant named 'dev' within the 'nonprod' infrastructure.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        # Ensure the 'nonprod' infrastructure is already created before setting up the tenant.
        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        ```

        ### Create a DuploCloud tenant named 'dev' with infra name variable and tenant id as output.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        config = pulumi.Config()
        infra_name = config.get("infraName")
        if infra_name is None:
            infra_name = "nonprod"
        # Ensure the 'nonprod' infrastructure is already created before setting up the tenant.
        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        pulumi.export("tenantId", tenant.tenant_id)
        ```

        ### Create a duplocloud tenant named dev with AWS Cognito power user access in the nonprod infrastructure.

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_duplocloud as duplocloud

        # A prerequisite for creating a tenant is having an existing infrastructure. Here’s how you can reference an existing infrastructure.
        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        # Here’s how to create a tenant by providing the infrastructure name for the plan_id field.
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        # Attaches a managed IAM policy to an IAM role.
        amazon_cognito_power_user = aws.index.IamRolePolicyAttachment("AmazonCognitoPowerUser",
            role=fduploservices-{tenant.account_name},
            policy_arn=arn:aws:iam::aws:policy/AmazonCognitoPowerUser)
        ```

        ### Create a DuploCloud tenant named 'qa' with full access to invoke AWS API Gateway in the nonprod infrastructure.

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_duplocloud as duplocloud

        # A prerequisite for creating a tenant is having an existing infrastructure. Here’s how you can reference an existing infrastructure.
        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        # Here’s how to create a tenant by providing the infrastructure name for the plan_id field.
        tenant = duplocloud.Tenant("tenant",
            account_name="qa",
            plan_id=infra.infra_name)
        # Attaches a managed IAM policy to an IAM role.
        amazon_api_gateway_invoke_full_access = aws.index.IamRolePolicyAttachment("AmazonAPIGatewayInvokeFullAccess",
            role=fduploservices-{tenant.account_name},
            policy_arn=arn:aws:iam::aws:policy/AmazonAPIGatewayInvokeFullAccess)
        ```

        ### Create duplocloud tenant named dev with security group rule to allow access from 10.220.0.0/16 on port 5432 in nonprod infra’

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        # Allow communication on port 5432 for the PostgreSQL database from the 10.220.0.0/16 subnet
        allow_from_vpn = duplocloud.TenantNetworkSecurityRule("allow_from_vpn",
            tenant_id=tenant.tenant_id,
            source_address="10.220.0.0/16",
            protocol="tcp",
            from_port=5432,
            to_port=5432,
            description="Allow communication from 10.220.0.0/16 on port 5432.")
        ```

        ### Setup duplocloud tenant named dev with security group rule to allow access from 10.220.0.0/16 on port 22 in nonprod infra’

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.get_infrastructure(infra_name="nonprod")
        tenant = duplocloud.Tenant("tenant",
            account_name="dev",
            plan_id=infra.infra_name)
        # Allow communication on port 22 from the 10.220.0.0/16 subnet.
        allow_from_vpn = duplocloud.TenantNetworkSecurityRule("allow_from_vpn",
            tenant_id=tenant.tenant_id,
            source_address="10.220.0.0/16",
            protocol="tcp",
            from_port=22,
            to_port=22,
            description="Allow communication from 10.220.0.0/16 on port 22.")
        ```

        ### Provision a tenant named 'myapp' within the infrastructure 'myinfra' and disable deletion protection.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.get_infrastructure(infra_name="myinfra")
        tenant = duplocloud.Tenant("tenant",
            account_name="myapp",
            plan_id=infra.infra_name,
            allow_deletion=True)
        # Reference the tenant_id field from the duplocloud_tenant resource.
        tenant_config = duplocloud.TenantConfig("tenant_config",
            tenant_id=tenant.tenant_id,
            settings=[{
                "key": "delete_protection",
                "value": "false",
            }])
        ```

        ### Provision a tenant named 'myapp' within the infrastructure 'myinfra', and ensure that the S3 bucket has public access blocked and SSL enforcement enabled in the S3 bucket policy.

        ```python
        import pulumi
        import pulumi_duplocloud as duplocloud

        infra = duplocloud.get_infrastructure(infra_name="myinfra")
        tenant = duplocloud.Tenant("tenant",
            account_name="myapp",
            plan_id=infra.infra_name,
            allow_deletion=True)
        # Reference the tenant_id field from the duplocloud_tenant resource.
        tenant_config = duplocloud.TenantConfig("tenant_config",
            tenant_id=tenant.tenant_id,
            settings=[
                {
                    "key": "block_public_access_to_s3",
                    "value": "true",
                },
                {
                    "key": "enforce_ssl_for_s3",
                    "value": "true",
                },
            ])
        ```

        ## Import

        ```sh
        $ pulumi import duplocloud:index/tenant:Tenant myapp v2/admin/TenantV2/*TENANT_ID*
        ```

        :param str resource_name: The name of the resource.
        :param TenantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TenantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 allow_deletion: Optional[pulumi.Input[bool]] = None,
                 existing_k8s_namespace: Optional[pulumi.Input[str]] = None,
                 plan_id: Optional[pulumi.Input[str]] = None,
                 wait_until_created: Optional[pulumi.Input[bool]] = None,
                 wait_until_deleted: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TenantArgs.__new__(TenantArgs)

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            __props__.__dict__["allow_deletion"] = allow_deletion
            __props__.__dict__["existing_k8s_namespace"] = existing_k8s_namespace
            if plan_id is None and not opts.urn:
                raise TypeError("Missing required property 'plan_id'")
            __props__.__dict__["plan_id"] = plan_id
            __props__.__dict__["wait_until_created"] = wait_until_created
            __props__.__dict__["wait_until_deleted"] = wait_until_deleted
            __props__.__dict__["infra_owner"] = None
            __props__.__dict__["policies"] = None
            __props__.__dict__["tags"] = None
            __props__.__dict__["tenant_id"] = None
        super(Tenant, __self__).__init__(
            'duplocloud:index/tenant:Tenant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            allow_deletion: Optional[pulumi.Input[bool]] = None,
            existing_k8s_namespace: Optional[pulumi.Input[str]] = None,
            infra_owner: Optional[pulumi.Input[str]] = None,
            plan_id: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TenantPolicyArgs', 'TenantPolicyArgsDict']]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TenantTagArgs', 'TenantTagArgsDict']]]]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            wait_until_created: Optional[pulumi.Input[bool]] = None,
            wait_until_deleted: Optional[pulumi.Input[bool]] = None) -> 'Tenant':
        """
        Get an existing Tenant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the tenant. Tenant names are globally unique, and cannot be a prefix of any other tenant name.
        :param pulumi.Input[bool] allow_deletion: Whether or not to even try and delete the tenant. *NOTE: This only works if you have disabled deletion protection for the tenant.* Defaults to `false`.
        :param pulumi.Input[str] existing_k8s_namespace: Existing kubernetes namespace to use by the tenant. *NOTE: This is an advanced feature, please contact your DuploCloud administrator for help if you want to use this field.*
        :param pulumi.Input[str] plan_id: The name of the plan under which the tenant will be created.
        :param pulumi.Input[str] tenant_id: A GUID identifying the tenant. This is automatically generated by Duplo.
        :param pulumi.Input[bool] wait_until_created: Whether or not to wait until Duplo has created the tenant. Defaults to `true`.
        :param pulumi.Input[bool] wait_until_deleted: Whether or not to wait until Duplo has destroyed the tenant. Defaults to `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TenantState.__new__(_TenantState)

        __props__.__dict__["account_name"] = account_name
        __props__.__dict__["allow_deletion"] = allow_deletion
        __props__.__dict__["existing_k8s_namespace"] = existing_k8s_namespace
        __props__.__dict__["infra_owner"] = infra_owner
        __props__.__dict__["plan_id"] = plan_id
        __props__.__dict__["policies"] = policies
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["wait_until_created"] = wait_until_created
        __props__.__dict__["wait_until_deleted"] = wait_until_deleted
        return Tenant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        The name of the tenant. Tenant names are globally unique, and cannot be a prefix of any other tenant name.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="allowDeletion")
    def allow_deletion(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to even try and delete the tenant. *NOTE: This only works if you have disabled deletion protection for the tenant.* Defaults to `false`.
        """
        return pulumi.get(self, "allow_deletion")

    @property
    @pulumi.getter(name="existingK8sNamespace")
    def existing_k8s_namespace(self) -> pulumi.Output[str]:
        """
        Existing kubernetes namespace to use by the tenant. *NOTE: This is an advanced feature, please contact your DuploCloud administrator for help if you want to use this field.*
        """
        return pulumi.get(self, "existing_k8s_namespace")

    @property
    @pulumi.getter(name="infraOwner")
    def infra_owner(self) -> pulumi.Output[str]:
        return pulumi.get(self, "infra_owner")

    @property
    @pulumi.getter(name="planId")
    def plan_id(self) -> pulumi.Output[str]:
        """
        The name of the plan under which the tenant will be created.
        """
        return pulumi.get(self, "plan_id")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence['outputs.TenantPolicy']]:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence['outputs.TenantTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        A GUID identifying the tenant. This is automatically generated by Duplo.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="waitUntilCreated")
    def wait_until_created(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to wait until Duplo has created the tenant. Defaults to `true`.
        """
        return pulumi.get(self, "wait_until_created")

    @property
    @pulumi.getter(name="waitUntilDeleted")
    def wait_until_deleted(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to wait until Duplo has destroyed the tenant. Defaults to `false`.
        """
        return pulumi.get(self, "wait_until_deleted")

