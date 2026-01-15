# Upgrading the Pulumi DuploCloud Provider

This document describes how to manually upgrade the Pulumi DuploCloud provider to use a newer version of the upstream Terraform provider (`terraform-provider-duplocloud`).

## Table of Contents

- [Automated Upgrades](#automated-upgrades)
- [Manual Upgrade Process](#manual-upgrade-process)
- [Verification Steps](#verification-steps)
- [Troubleshooting](#troubleshooting)

## Automated Upgrades

The repository includes automated workflows that handle provider upgrades:

### Daily Automatic Checks

The `.github/workflows/upgrade-provider.yml` workflow runs daily at 3 AM UTC to:
- Check for new versions of `terraform-provider-duplocloud`
- Create GitHub issues when new versions are detected
- Automatically open Pull Requests with the upgrade

### Manual Trigger via GitHub UI

1. Go to: https://github.com/duplocloud/pulumi-duplocloud/actions/workflows/upgrade-provider.yml
2. Click "Run workflow"
3. Enter the target version (e.g., `0.11.31`) or leave blank to auto-detect latest
4. Click "Run workflow"

### Manual Trigger via GitHub CLI

```bash
# Upgrade to a specific version
gh workflow run upgrade-provider.yml -f version=0.11.31

# Auto-detect and upgrade to latest version
gh workflow run upgrade-provider.yml
```

## Manual Upgrade Process

If you need to upgrade the provider manually (e.g., for testing or when automation is unavailable), follow these steps:

### Prerequisites

Ensure you have the following tools installed:
- Go 1.22.7 or later
- Node.js (for SDK generation)
- Python 3.x (for SDK generation)
- .NET SDK (for SDK generation)
- Make

### Step 1: Update the Terraform Provider Version

Edit `provider/go.mod` and update the version of the Terraform provider:

```bash
# Open the file
vim provider/go.mod

# Find the line:
# github.com/duplocloud/terraform-provider-duplocloud v0.11.0

# Change it to the new version, e.g.:
# github.com/duplocloud/terraform-provider-duplocloud v0.11.31
```

**Example change:**
```diff
require (
-	github.com/duplocloud/terraform-provider-duplocloud v0.11.0
+	github.com/duplocloud/terraform-provider-duplocloud v0.11.31
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.102.0
)
```

### Step 2: Update Go Dependencies

Run `go mod tidy` to download the new version and update `go.sum`:

```bash
cd provider
go mod tidy
cd ..
```

This will:
- Download the new Terraform provider version
- Update the `go.sum` file with new checksums
- Resolve any dependency conflicts

### Step 3: Regenerate the Provider Schema

Generate the Pulumi schema from the Terraform provider:

```bash
make schema
```

This command:
- Builds the `pulumi-tfgen-duplocloud` code generator
- Runs the generator to create `provider/cmd/pulumi-resource-duplocloud/schema.json`
- Converts Terraform HCL examples to Pulumi examples
- Generates bridge metadata

**Expected output:**
- Schema file generated successfully
- Example conversion statistics (typically 80%+ success rate)
- Warnings about examples that couldn't be converted (this is normal)

### Step 4: Regenerate All Language SDKs

Generate SDKs for all supported languages:

```bash
make generate_sdks
```

This command generates:
- **Go SDK**: `sdk/go/duplocloud/`
- **Python SDK**: `sdk/python/pulumi_duplocloud/`
- **Node.js/TypeScript SDK**: `sdk/nodejs/`
- **.NET/C# SDK**: `sdk/dotnet/`

**What gets updated:**
- New resource types
- New data sources
- Updated resource properties
- Updated documentation
- Package metadata files

### Step 5: Build and Test (Optional but Recommended)

Build the provider and SDKs to ensure everything compiles:

```bash
# Build everything
make build

# Or build individual components
make provider          # Build the provider binary
make build_sdks       # Build all SDKs
```

Run tests to verify functionality:

```bash
# Run provider tests
make test_provider

# Run example tests (requires configured DuploCloud credentials)
make test
```

## Verification Steps

After upgrading, verify the changes:

### 1. Check Modified Files

```bash
git status
```

Expected changes:
- `provider/go.mod` - Updated version
- `provider/go.sum` - Updated checksums
- `provider/cmd/pulumi-resource-duplocloud/schema.json` - Updated schema
- `sdk/*/` - Updated SDK files for all languages

### 2. Review New Resources

Check what new resources or data sources were added:

```bash
# List new files in SDKs
git status --short | grep "^??" | grep -E "\.(go|py|ts|cs)$"
```

### 3. Check for Breaking Changes

Review the Terraform provider's changelog:
```bash
# Visit the upstream provider's releases
open https://github.com/duplocloud/terraform-provider-duplocloud/releases
```

### 4. Test Locally

If you have a test Pulumi project:

```bash
# Install the local provider
make install

# Test with your Pulumi project
cd /path/to/your/pulumi/project
pulumi up --preview
```

## Commit and Push Changes

Once verified, commit and push your changes:

```bash
# Add all changes
git add .

# Commit with descriptive message
git commit -m "Update to terraform-provider-duplocloud v0.11.31

- Updated provider dependency from v0.11.0 to v0.11.31
- Regenerated schema and all language SDKs
- Added new resources: [list key new resources]
- Added new data sources: [list key new data sources]
"

# Push to your branch
git push origin your-branch-name
```

## Troubleshooting

### Issue: `go mod tidy` fails with version not found

**Solution:** Verify the version exists in the upstream repository:
```bash
# Check available versions
open https://github.com/duplocloud/terraform-provider-duplocloud/tags
```

### Issue: Schema generation fails

**Solution:** 
1. Ensure the provider builds correctly:
   ```bash
   cd provider
   go build
   ```
2. Check for compilation errors in `provider/resources.go`
3. Verify the Terraform provider is compatible with the bridge version

### Issue: SDK generation produces errors

**Solution:**
1. Check that all required tools are installed:
   ```bash
   go version
   node --version
   python --version
   dotnet --version
   ```
2. Clean and regenerate:
   ```bash
   make clean
   make generate
   ```

### Issue: Examples fail to convert

**Solution:** This is normal. The Pulumi bridge cannot convert all Terraform HCL examples, especially those using:
- Complex Terraform functions
- Provider-specific features
- Cross-provider dependencies

These warnings can typically be ignored unless they affect critical resources.

## Example: Complete Manual Upgrade

Here's a complete example of upgrading from v0.11.0 to v0.11.31:

```bash
# 1. Update the version in go.mod
sed -i '' 's/terraform-provider-duplocloud v0.11.0/terraform-provider-duplocloud v0.11.31/' provider/go.mod

# 2. Update dependencies
cd provider && go mod tidy && cd ..

# 3. Regenerate schema
make schema

# 4. Regenerate SDKs
make generate_sdks

# 5. Verify changes
git status
git diff provider/go.mod

# 6. Build and test
make build
make test_provider

# 7. Commit changes
git add .
git commit -m "Update to terraform-provider-duplocloud v0.11.31"
git push
```

## Configuration Files

### `.upgrade-config.yml`

This file configures the automated upgrade workflow:

```yaml
upstream-provider-name: terraform-provider-duplocloud
upstream-provider-org: duplocloud
pulumi-infer-version: true
remove-plugins: true
```

- `upstream-provider-name`: Name of the Terraform provider
- `upstream-provider-org`: GitHub organization hosting the provider
- `pulumi-infer-version`: Automatically infer version from upstream
- `remove-plugins`: Clean up old plugins during upgrade

## Additional Resources

- [Pulumi Terraform Bridge Documentation](https://github.com/pulumi/pulumi-terraform-bridge)
- [DuploCloud Terraform Provider](https://github.com/duplocloud/terraform-provider-duplocloud)
- [Pulumi Provider Boilerplate](https://github.com/pulumi/pulumi-tf-provider-boilerplate)

## Recent Upgrade History

| Date | From Version | To Version | Notes |
|------|--------------|------------|-------|
| 2026-01-14 | v0.11.0 | v0.11.31 | Added Azure Cosmos DB, ECache Global Datastore, Tenant KMS resources |

---

**Note:** This document should be updated whenever the upgrade process changes or when significant upgrades are performed.
