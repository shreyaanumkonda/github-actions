# Go API CI/CD Workflow

## Overview

The Go API CI/CD Workflow demonstrates how to build, test, and deploy Go applications using GitHub Actions. This workflow showcases community actions for Go development and automated testing.

## File Location

- **Workflow File**: `.github/workflows/community-actions.yml`
- **Documentation**: `workflows/go-api-workflow.md`

## Features

- **Automated Go Setup**: Uses community actions to install Go
- **Code Building**: Compiles Go application automatically
- **Testing**: Runs Go tests in CI environment
- **Multiple Triggers**: Push, PR, and manual dispatch support

## How It Works

### Trigger Events
```yaml
on:
  push:
    branches: [main]
  pull_request:
```

### Key Components

1. **Community Actions Used**:
   - `actions/checkout@v4` - Downloads repository code
   - `actions/setup-go@v4` - Installs Go runtime

2. **Build Process**:
   ```yaml
   - name: Build the code 
     run: go build ./...
   - name: Test the code
     run: go test ./...
   ```

## What We Learned About Community Actions

### Community Actions vs Official Actions

**Official Actions** (by GitHub):
- `actions/checkout` - Built by GitHub
- `actions/setup-go` - Built by GitHub
- Always reliable and well-maintained

**Community Actions** (by developers):
- `peter-evans/create-or-update-comment` - Built by Peter Evans
- Third-party solutions for specific needs
- Can be more specialized but need evaluation

### How to Choose Community Actions

1. **Check Stars & Usage**: Popular actions are more reliable
2. **Read Documentation**: Good docs = good maintenance
3. **Check Last Updated**: Recent updates = active development
4. **Review Source Code**: Understand what it does
5. **Test Thoroughly**: Always test in non-production first

### Best Practices

- **Pin Versions**: Use specific versions (`@v4`) not `@latest`
- **Check Permissions**: Review what permissions the action needs
- **Monitor Updates**: Keep actions updated for security
- **Have Fallbacks**: Plan for when actions break

## Workflow Behavior

### When Code is Pushed
- Workflow triggers automatically
- Installs Go 1.21
- Builds the application
- Runs tests

### When PR is Created
- Same process as push
- Validates code before merge

## Technical Details

- **Go Version**: 1.21 (matches go.mod)
- **Build Command**: `go build ./...`
- **Test Command**: `go test ./...`
- **Runner**: Ubuntu Latest

## The Go Application

The workflow builds a simple HTTP server (`main.go`) that:
- Listens on port 8080
- Responds to GET requests with "Hello, World!"
- Rejects non-GET requests with "Method not allowed"

## Use Cases

- **CI/CD Pipeline**: Automated testing and building
- **Go Development**: Standard workflow for Go projects
- **Learning**: Understanding community action integration
- **Quality Assurance**: Ensures code builds and tests pass

## Testing

To test the Go API CI/CD Workflow:

1. **Push changes** to main branch
2. **Create a PR** from feature branch
3. **Check Actions tab** to see workflow running
4. **Verify build success** and test results

This workflow demonstrates the power of community actions in creating robust CI/CD pipelines for Go applications.
