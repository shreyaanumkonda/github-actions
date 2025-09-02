# Hello World Workflow

## Overview
This is my first GitHub Actions workflow - a simple "Hello World" example that demonstrates the basic structure and functionality of GitHub Actions.

## File Location
`.github/workflows/hello.yml`

## What It Does
- Prints "Hello, world!" to the console
- Runs on every push to the main branch
- Demonstrates basic workflow syntax

## Workflow Structure

```yaml
name: Hello World
on: [push]
jobs:
  say-hello:
    runs-on: ubuntu-latest
    steps:
      - name: Print a hello message
        run: echo "Hello, world!"
```

## Key Learning Points

### 1. Basic Workflow Components
- **name**: Identifies the workflow
- **on**: Defines when the workflow triggers
- **jobs**: Contains the work to be performed
- **runs-on**: Specifies the runner environment
- **steps**: Individual tasks within a job

### 2. Trigger Configuration
- **on: [push]**: Workflow runs on every push to any branch
- Simple array syntax for basic triggers

### 3. Job Configuration
- **say-hello**: Job name (must be unique)
- **runs-on: ubuntu-latest**: Uses GitHub's Ubuntu runner
- **steps**: Array of individual actions

### 4. Step Execution
- **name**: Human-readable step description
- **run**: Shell command to execute

## How to Use

1. **Automatic Execution**: This workflow runs automatically on every push to the main branch
2. **View Results**: Go to Actions tab in GitHub to see execution logs
3. **Manual Testing**: Make a small change and push to trigger the workflow

## What I Learned

- Basic GitHub Actions YAML syntax
- Understanding workflow structure and components
- How triggers work in GitHub Actions
- The difference between jobs and steps
- How to use shell commands in workflows

## Next Steps

After mastering this basic workflow, I moved on to:
- Multiple job workflows
- Different trigger types (pull requests, manual dispatch)
- More complex step configurations
- Using pre-built actions from the marketplace

## Related Workflows

- [Good Morning Workflow](./goodmorning-workflow.md) - Advanced example with multiple jobs and triggers
