---
title: DuploCloud Provider
meta_desc: Provides an overview of the Pulumi DuploCloud provider for managing DuploCloud resources.
layout: package
---


## Overview

Use the DuploCloud Pulumi provider to interact with almost all of [DuploCloud](https://duplocloud.com/) resources.

## Example usage

<!-- These are copied from examples/templates -->

{{< chooser language "go,typescript,python" >}}
{{% choosable language go %}}

```golang
package main

import (
	"github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Set up the infrastructure.
		_, err := duplocloud.NewInfrastructure(ctx, "infra", &duplocloud.InfrastructureArgs{
			InfraName:       pulumi.String("pulumi"),
			Cloud:           pulumi.Int(0),
			Region:          pulumi.String("us-east-2"),
			EnableK8Cluster: pulumi.Bool(true),
			AddressPrefix:   pulumi.String("10.5.0.0/16"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}
```

{{% /choosable %}}

{{% choosable language python %}}

```python
import pulumi
import pulumi_duplocloud as duplo

# Create nonprod Infrastructure resource
infrastructure = duplo.infrastructure.Infrastructure(resource_name="nonprod",
    infra_name="nonprod",
    cloud=0,
    region="us-west-2",
    azcount=2,
    subnet_cidr=24,
    address_prefix="10.22.0.0/16",
    enable_k8_cluster=True,  # Enable Kubernetes
)

# Export the outputs
pulumi.export("vpc_id", infrastructure.vpc_id)
```

{{% /choosable %}}

{{% choosable language typescript %}}

```typescript
import * as duplocloud from "@duplocloud/pulumi";

const infra = new duplocloud.Infrastructure("infra", {
    infraName: "nonprod",
    cloud: 0,
    region: "us-west-2",
    enableK8Cluster: true,
    addressPrefix: "10.11.0.0/16",
});

export const vpcId = infra.vpcId;
```

{{% /choosable %}}

{{< /chooser >}}