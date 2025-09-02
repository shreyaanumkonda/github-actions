# Good Morning Workflow

## Overview
This is an advanced GitHub Actions workflow that demonstrates multiple jobs, different trigger types, and more complex workflow configurations. It builds upon the basic concepts learned in the Hello World workflow.

## File Location
`.github/workflows/goodmorning.yml`

## What It Does
- Runs two parallel jobs: `say-goodnight` and `say-goodmorning`
- Each job prints different messages and performs different actions
- Demonstrates multiple trigger types and job dependencies

## Workflow Structure

```yaml
name: Good Morning
on:
    push:
        branches:
            - main
    pull_request:
        branches:
            - main
    workflow_dispatch:
jobs:
    say-goodnight:
        runs-on: ubuntu-latest
        steps:
            - name: Print a good morning message
              run: echo "Good morning!"
            - uses: actions/checkout@v2
            - run: echo "Good night! This is the end of the Job-1"
    say-goodmorning:
        runs-on: ubuntu-latest
        steps:
            - name: Print a good morning message
              run: echo "Good morning!"
            - uses: actions/checkout@v2
            - run: echo "Good night! This is the end of the Job-2"
              shell: bash
```

## Key Learning Points

### 1. Multiple Trigger Types
- **push**: Triggers on pushes to main branch
- **pull_request**: Triggers on pull requests to main branch
- **workflow_dispatch**: Allows manual execution from GitHub UI

### 2. Branch-Specific Triggers
```yaml
push:
    branches:
        - main
```
- Workflow only runs on pushes to the main branch
- Can specify multiple branches or use patterns

### 3. Multiple Jobs
- **say-goodnight**: First job with its own steps
- **say-goodmorning**: Second job running in parallel
- Jobs run independently and can have different configurations

### 4. Using Pre-built Actions
- **actions/checkout@v2**: Checks out the repository code
- Demonstrates how to use actions from the GitHub marketplace

### 5. Shell Configuration
- **shell: bash**: Explicitly specifies the shell for a step
- Useful for ensuring consistent behavior across different runners

## How to Use

### Automatic Triggers
1. **Push to main**: Workflow runs automatically
2. **Create pull request**: Workflow runs on PR creation/updates

### Manual Execution
1. Go to Actions tab in GitHub
2. Click on "Good Morning" workflow
3. Click "Run workflow" button
4. Select branch and click "Run workflow"

## What I Learned

- How to configure multiple trigger types
- Understanding job parallelism in GitHub Actions
- Using pre-built actions from the marketplace
- Branch-specific workflow triggers
- Manual workflow dispatch functionality
- Shell configuration for steps
- The difference between jobs and steps in complex workflows

## Advanced Concepts Demonstrated

### 1. Workflow Dispatch
- Manual trigger capability
- Useful for testing and on-demand execution
- Can be triggered from GitHub UI

### 2. Job Parallelism
- Multiple jobs run simultaneously
- Each job has its own runner environment
- Jobs can have different configurations

### 3. Action Usage
- Using `actions/checkout@v2` to access repository code
- Understanding action versioning (@v2)
- Integrating marketplace actions into workflows

## Troubleshooting

### Common Issues
- **Job name conflicts**: Each job must have a unique name
- **Action versioning**: Always specify action versions (@v2, @v3, etc.)
- **Shell compatibility**: Different runners may have different default shells

### Best Practices
- Use descriptive job and step names
- Specify action versions for reproducibility
- Test workflows with manual dispatch before relying on automatic triggers

## Next Steps

This workflow demonstrates intermediate GitHub Actions concepts. Future learning areas include:
- Environment variables and secrets
- Job dependencies and conditional execution
- Matrix builds and parallel testing
- Deployment workflows
- Custom actions and composite actions

## Related Workflows

- [Hello World Workflow](./hello-workflow.md) - Basic introduction to GitHub Actions
