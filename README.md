# DuploCloud Resource Provider

The DuploCloud Resource Provider lets you manage [DuploCloud](https://duplocloud.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @duplocloud/pulumi
```

or `yarn`:

```bash
yarn add @duplocloud/pulumi
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-duplocloud
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/duplocloud/pulumi-duplocloud/sdk/go/...
```


## Configuration

The DuploCloud provider requires the following configuration parameters:

- `duplocloud:duploHost` - Base URL to the DuploCloud REST API
- `duplocloud:duploToken` - Bearer token for authentication

You can set these using environment variables:

```bash
export duplo_host="https://your-duplocloud-instance.com"
export duplo_token="<your_duplo_token>"
```