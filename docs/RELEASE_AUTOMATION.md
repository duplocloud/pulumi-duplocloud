# Pulumi Provider Release Automation

This document describes the complete automation flow for releasing the Pulumi DuploCloud provider when a new upstream Terraform provider is released.

## Overview

The release process is fully automated and consists of 4 main workflows that work together:

1. **upgrade-provider.yml** - Detects new upstream provider and creates PR
2. **auto-release.yml** - Creates git tags and version
3. **release.yml** - Builds and publishes packages to all registries
4. **notify-release.yml** - Sends Slack notification

```
┌─────────────────────────────────────────────────────────────┐
│  Upstream Terraform Provider Release                        │
│  (github.com/duplocloud/terraform-provider-duplocloud)      │
└────────────────────┬────────────────────────────────────────┘
                     │
                     ↓
        ┌────────────────────────┐
        │  upgrade-provider.yml  │  ⏰ Daily at 3 AM UTC OR Manual dispatch
        │  (Workflow Dispatch    │
        │   or Scheduled)        │
        └────────────┬───────────┘
                     │
         1. Check for new upstream version
         2. Create PR with changes
         3. Add release label (patch/minor/major)
         4. Merge PR
                     │
                     ↓
        ┌────────────────────────┐
        │  auto-release.yml      │  ⏰ Triggered by PR merge
        │  (on PR closed)        │
        └────────────┬───────────┘
                     │
         1. Check for needs-release/* label
         2. Determine version from provider/go.mod
         3. Create git tag v*.*.*
         4. Push tag to GitHub
                     │
                     ↓
        ┌────────────────────────┐
        │  release.yml           │  ⏰ Triggered by git tag push
        │  (on tag v*.*.*)       │
        └────────────┬───────────┘
                     │
         1. Build provider binary
         2. Build SDKs (Go, Python, Node.js, .NET)
         3. Run tests and linting
         4. Publish to registries:
            - PyPI (Python)
            - npm (Node.js)
            - NuGet (.NET)
            - GitHub Releases (Go)
                     │
                     ↓
        ┌────────────────────────┐
        │  notify-release.yml    │  ⏰ Triggered by release published
        │  (on release published)│
        └────────────────────────┘
                     │
         Send Slack notification with links
         to all packages and release page
```

---

## Detailed Workflow Descriptions

### 1. upgrade-provider.yml

**File:** `.github/workflows/upgrade-provider.yml`

**Trigger:**
- **Daily Schedule:** 3 AM UTC (8 PM PDT / 7 PM PST)
- **Manual Dispatch:** Anytime via GitHub Actions UI

**Process:**

1. **Check Upstream Version**
   - Uses `upgrade-provider` tool to check if new version exists in terraform-provider-duplocloud
   - If no new version found, workflow exits gracefully

2. **Create Pull Request**
   - Uses `pulumi/pulumi-upgrade-provider-action` to:
     - Update `provider/go.mod` with new terraform provider version
     - Regenerate Pulumi schema from Terraform provider
     - Regenerate all SDK files (Go, Python, Node.js, .NET)
   - Creates PR with title: `Sync Pulumi provider with terraform-provider-duplocloud v*.*.*`
   - PR includes detailed changelog and links to upstream release
   - **Important:** Set `automerge: false` so PR is not merged immediately

3. **Add Release Label**
   - Determines release type by comparing versions:
     - **Major:** Version major component changed (e.g., 0.x.x → 1.0.0)
     - **Minor:** Version minor component changed (e.g., 0.11.x → 0.12.0)
     - **Patch:** Version patch component changed (e.g., 0.11.30 → 0.11.31)
   - Adds appropriate label: `needs-release/major`, `needs-release/minor`, or `needs-release/patch`

4. **Merge PR**
   - Finds the open PR with release label
   - Automatically merges PR with squash strategy
   - Ensures label is present when merged (required for next workflow)

**Key Configuration:**
```yaml
automerge: false  # Don't auto-merge immediately
allow-missing-docs: true  # Allow release even if docs incomplete
```

**Notifications:** None (moved to notify-release.yml)

---

### 2. auto-release.yml

**File:** `.github/workflows/auto-release.yml`

**Trigger:** PR merged with label `needs-release/patch`, `needs-release/minor`, or `needs-release/major`

**Process:**

1. **Verify Merge Condition**
   - Checks: `github.event.pull_request.merged == true`
   - Checks: Has one of the release labels
   - Can also be manually triggered with `workflow_dispatch` input

2. **Extract Version**
   - If manual input provided: Use input version
   - Otherwise: Extract version from `provider/go.mod`
   - Version format: `v*.*.*` (e.g., v0.12.1)

3. **Create and Push Git Tag**
   - Configures git identity as `github-actions[bot]`
   - Creates annotated tag: `git tag -a "v{VERSION}" -m "{MESSAGE}"`
   - Tag message includes:
     - Release note: `Release v*.*.*`
     - Reference to PR that triggered it
     - Link to PR

4. **Verify Tag Push**
   - Confirms tag exists on remote: `git ls-remote --tags origin`
   - Exits with error if tag push failed
   - Logs success message indicating release.yml will trigger

**Key Configuration:**
```yaml
permissions:
  contents: write  # Needed to push tags
```

**Why this matters:** The git tag push automatically triggers `release.yml` via the repository's webhook configuration.

---

### 3. release.yml

**File:** `.github/workflows/release.yml`

**Status:** ⚠️ **Autogenerated by Pulumi ci-mgmt** - Do not modify directly

**Trigger:** Push of tags matching pattern `v*.*.*` (excludes pre-releases with `-` suffix)

**Process:**

1. **Prerequisites**
   - Extracts version from tag
   - Sets up build environment
   - Runs security/license checks

2. **Build Provider**
   - Compiles Pulumi provider binary
   - Validates provider schema
   - Runs acceptance tests

3. **Build SDKs**
   - **Go:** Creates SDK and publishes to GitHub Releases
   - **Python:** Generates PyPI-compatible package
   - **Node.js:** Generates npm-compatible package
   - **.NET:** Generates NuGet-compatible package

4. **Testing & Validation**
   - Runs unit tests across all SDKs
   - Runs linting checks
   - Validates license compliance

5. **Publish Packages**
   - **PyPI:** `pip install pulumi-duplocloud`
   - **npm:** `npm install @duplocloud/pulumi`
   - **NuGet:** `dotnet add package DuploCloud.Pulumi`
   - **GitHub:** Releases page with Go SDK and assets

**Important Notes:**
- This workflow is regenerated automatically by Pulumi ci-mgmt
- Modifications to this file will be overwritten
- To customize: Configure ci-mgmt upstream or use separate workflows

---

### 4. notify-release.yml

**File:** `.github/workflows/notify-release.yml`

**Trigger:** GitHub release published (created by release.yml)

**Process:**

1. **Extract Version**
   - Gets version from release tag (e.g., `v0.12.1` → `0.12.1`)

2. **Send Slack Notification**
   - Sends formatted message to Slack webhook
   - Includes:
     - Release version and announcement header
     - Direct links to:
       - Pulumi provider GitHub release
       - Terraform provider upstream release
       - npm package (@duplocloud/pulumi)
       - PyPI package (pulumi-duplocloud)
       - NuGet package (DuploCloud.Pulumi)
       - GitHub release page

**Slack Message Example:**
```
✅ Pulumi DuploCloud Provider Released! 🚀
Release Version: v0.12.1

Links:
- Pulumi Provider: https://github.com/duplocloud/pulumi-duplocloud/releases/tag/v0.12.1
- npm: https://www.npmjs.com/package/@duplocloud/pulumi/v/0.12.1
- PyPI: https://pypi.org/project/pulumi-duplocloud/0.12.1
- NuGet: https://www.nuget.org/packages/DuploCloud.Pulumi/0.12.1
```

---

## Critical Dependency: Release Labels

The entire workflow depends on the `needs-release/*` label being present when the PR is merged.

### Label Types:

| Label | When to Use | Example |
|-------|-----------|---------|
| `needs-release/patch` | Bug fixes, security patches | 0.11.30 → 0.11.31 |
| `needs-release/minor` | New features, minor changes | 0.11.x → 0.12.0 |
| `needs-release/major` | Breaking changes | 0.x.x → 1.0.0 |

### How Labels Are Applied:

1. **upgrade-provider.yml** determines the release type by comparing versions
2. Adds the appropriate label to the PR via `gh pr edit`
3. **auto-release.yml** checks for these labels before creating a tag
4. If label is missing, auto-release.yml will not trigger

---

## Manual Triggering

### Trigger Upgrade Check Manually

```bash
gh workflow run upgrade-provider.yml \
  -f version=0.12.3  # Optional: specify version, otherwise auto-detect
```

### Trigger Release Manually

```bash
gh workflow run auto-release.yml \
  -f version=0.12.3  # Specify version to release
```

---

## Troubleshooting

### Issue: auto-release.yml not triggered after PR merge

**Checklist:**
- [ ] Is the PR merged (not just closed)?
- [ ] Does the PR have one of the release labels: `needs-release/patch`, `needs-release/minor`, `needs-release/major`?
- [ ] Check PR details: `gh pr view {PR_NUMBER}`

**Debug:**
```bash
# Check if auto-release workflow ran
gh run list --workflow=auto-release.yml

# Check recent workflow runs
gh run list --limit=10

# View specific run details
gh run view {RUN_ID} --log
```

### Issue: release.yml not triggered after tag push

**Checklist:**
- [ ] Is the tag in the correct format: `v*.*.*`?
- [ ] Does the tag NOT contain `-` (excludes pre-releases)?
- [ ] Was tag pushed to origin?

**Debug:**
```bash
# List all tags
git tag -l

# Check if tag exists on remote
git ls-remote --tags origin | grep {TAG_NAME}

# Manually verify tag format
git describe --tags
```

### Issue: Slack notification not sent

**Checklist:**
- [ ] Is the release published (not just created)?
- [ ] Is `SLACK_WEBHOOK_URL` secret configured?
- [ ] Check notify-release.yml run logs

---

## Environment Variables & Secrets

### Required Secrets:

| Secret | Used In | Purpose |
|--------|---------|---------|
| `PULUMI_GPAT` | auto-release.yml | GitHub PAT for tag operations |
| `PULUMI_PROVIDER_AUTOMATION_TOKEN` | upgrade-provider.yml | Bot token for PR operations |
| `GITHUB_TOKEN` | All workflows | GitHub API access |
| `SLACK_WEBHOOK_URL` | notify-release.yml | Slack webhook for notifications |
| `NPM_TOKEN` | release.yml | npm publish permission |
| `PYPI_API_TOKEN` | release.yml | PyPI publish permission |
| `NUGET_PUBLISH_KEY` | release.yml | NuGet publish permission |

### Permissions Required:

```yaml
contents: write        # Push tags, create releases
pull-requests: write   # Merge PRs, add labels
issues: write          # Add/remove labels
```

---

## Related Documentation

- [Pulumi ci-mgmt](https://github.com/pulumi-labs/ci-mgmt) - Auto-generates release.yml (moved from pulumi/ci-mgmt)
- [pulumi-upgrade-provider-action](https://github.com/pulumi/pulumi-upgrade-provider-action) - Creates upgrade PRs
- [Terraform DuploCloud Provider](https://github.com/duplocloud/terraform-provider-duplocloud) - Upstream provider
- [CI_MGMT_GUIDE.md](./CI_MGMT_GUIDE.md) - How to customize workflows using ci-mgmt

---

## Quick Reference

### Files Modified

- `.github/workflows/upgrade-provider.yml` - Detects and creates upgrade PR
- `.github/workflows/auto-release.yml` - Creates git tags
- `.github/workflows/release.yml` - Builds and publishes (autogenerated)
- `.github/workflows/notify-release.yml` - Sends Slack notifications

### Version Files

- `provider/go.mod` - Terraform provider version reference
- `Pulumi.yaml` - Pulumi provider version
- SDK versions auto-generated from these

### Key Branches

- `main` - Production releases
- `upgrade/terraform-provider-*` - Temporary upgrade branches (auto-deleted)

---

## Monitoring & Alerts

### Check Workflow Status

```bash
# Get last 10 workflow runs
gh run list --limit=10

# Watch upgrade-provider runs
gh run list --workflow=upgrade-provider.yml --limit=5

# Get run details
gh run view {RUN_ID} --log
```

### Subscribe to Notifications

- **GitHub:** Watch repository → Custom → Releases
- **Slack:** Subscribe to webhook in your Slack workspace
- **Email:** GitHub notifications for workflow failures
