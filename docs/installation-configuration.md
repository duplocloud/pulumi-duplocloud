---
title: DuploCloud provider installation & configuration
meta_desc: Detailed instructions for installing and configuring the Pulumi DuploCloud provider.
layout: installation
---

# Installation & Configuration

This guide will help you set up the DuploCloud provider for use with Pulumi.

## Installation

The DuploCloud provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@duplocloud/pulumi`](https://www.npmjs.com/package/@duplocloud/pulumi)
* Python: [`pulumi-duplocloud`](https://pypi.org/project/pulumi-duplocloud/)
* Go: [`github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud`](https://github.com/duplocloud/pulumi-duplocloud)

### Language-specific Installation

{{< chooser language "javascript,python,go" >}}
{{% choosable language javascript %}}
```bash
npm install @duplocloud/pulumi
```
{{% /choosable %}}

{{% choosable language python %}}
```bash
pip install pulumi-duplocloud
```
{{% /choosable %}}

{{% choosable language go %}}
```bash
go get github.com/duplocloud/pulumi-duplocloud/sdk/go/duplocloud
```
{{% /choosable %}}

### Provider Binary

For local development or specific version requirements, you can install the provider plugin directly:

```bash
pulumi plugin install resource duplocloud --server github://api.github.com/duplocloud/pulumi-duplocloud
```

To install a specific version:

```bash
pulumi plugin install resource duplocloud v0.0.1 --server github://api.github.com/duplocloud/pulumi-duplocloud
```

## Configuration

The DuploCloud provider requires the following configuration parameters:

- duplocloud:duploHost - Base URL to the DuploCloud REST API
- duplocloud:duploToken - Bearer token for authentication

You can set these using environment variables:

```bash
export duplo_host="https://your-duplocloud-instance.com"
export duplo_token="<your_duplo_token>"
```

Or using the Pulumi configuration system:

```bash
pulumi config set duplocloud:duploHost "https://your-duplocloud-instance.com"
pulumi config set duplocloud:duploToken "<your_duplo_token>"
```
