# REST API Workflow

## Overview

The REST API Workflow demonstrates how to use GitHub's REST API directly in workflows instead of relying on community actions. This workflow creates comments on pull requests using direct API calls with `curl`.

## File Location

- **Workflow File**: `.github/workflows/commen-RESTapi.yml`
- **Documentation**: `workflows/rest-api-workflow.md`

## Features

- **Direct API Integration**: Uses GitHub REST API endpoints directly
- **Custom Comment Logic**: Full control over comment creation without third-party actions
- **Permission Management**: Demonstrates proper GitHub token permissions
- **Multiple Triggers**: Responds to both PR opening and new commits

## How It Works

### Trigger Events
```yaml
on:
  pull_request:
    types: [opened, synchronize]
```

### Key Components

1. **Permission Management**
   ```yaml
   permissions:
     pull-requests: write
   ```

2. **REST API Call**
   ```bash
   curl -X POST \
     -H "Authorization: Bearer ${{ secrets.GITHUB_TOKEN }}" \
     -H "Content-Type: application/json" \
     -d '{"body": "Hello, this is a comment via REST API"}' \
     "https://api.github.com/repos/${{ github.repository }}/issues/${{ github.event.pull_request.number }}/comments"
   ```

## The Learning Challenge

This workflow teaches how to:
- Use GitHub's REST API directly in workflows
- Handle authentication with `GITHUB_TOKEN`
- Set proper permissions for API access
- Make HTTP requests with `curl` in workflow steps

## Workflow Behavior

### When PR is Opened
- Workflow triggers on `opened` event
- Creates a comment via REST API

### When New Commits are Pushed
- Workflow triggers on `synchronize` event
- Creates another comment (can be modified to update existing)

## Technical Details

- **API Endpoint**: `POST /repos/{owner}/{repo}/issues/{issue_number}/comments`
- **Authentication**: `GITHUB_TOKEN` with `pull-requests: write` permission
- **Response Format**: JSON with comment body
- **HTTP Method**: POST with JSON payload

## Key Learning Points

1. **Permissions are Critical**: `GITHUB_TOKEN` needs explicit permissions
2. **API Endpoints**: Understanding GitHub's REST API structure
3. **Authentication**: Using Bearer tokens in API calls
4. **Event Context**: Accessing PR number via `${{ github.event.pull_request.number }}`

## Testing

To test the REST API Workflow:

1. Create a new branch
2. Make a change and push to the branch
3. Create a Pull Request
4. Observe the workflow create a comment via REST API
5. Push additional commits to see `synchronize` trigger

## Use Cases

- **Custom Comment Logic**: When community actions don't meet your needs
- **API Integration Learning**: Understanding GitHub's API directly
- **Advanced Workflows**: Building complex automation with full control
- **Educational Purposes**: Learning REST API concepts in GitHub Actions

This workflow demonstrates the power of using GitHub's REST API directly for maximum flexibility and control over workflow behavior.
