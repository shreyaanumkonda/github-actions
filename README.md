# GitHub Actions Learning Repository

Welcome to my GitHub Actions learning journey! This repository contains practical examples and workflows I've created while learning GitHub Actions from scratch.

## Learning Objectives

- Understanding GitHub Actions workflow syntax and best practices
- Learning different trigger types and event handling
- Exploring job configurations, steps, and action integrations
- Building practical automation workflows for various use cases
- Solving real-world challenges through iterative development

## Current Focus Areas

- **Workflow Automation**: Scheduled and event-driven workflows
- **Comment Management**: Advanced bot interactions and comment handling
- **CI/CD Patterns**: Continuous integration and deployment workflows
- **Custom Actions**: Building reusable workflow components

## Workflow Examples

This repository contains various workflow examples demonstrating different GitHub Actions patterns and use cases. Each workflow includes detailed documentation explaining its purpose, implementation, and learning objectives.

### Current Workflows

| Workflow | Type | Purpose | Documentation |
|----------|------|---------|---------------|
| Hello World | Basic | Introduction to GitHub Actions | [hello-workflow.md](./workflows/hello-workflow.md) |
| Good Morning | Basic | Scheduled automation workflow | [goodmorning-workflow.md](./workflows/goodmorning-workflow.md) |
| Issue Comment | Basic | Event-driven comment response | [issue-comment-workflow.md](./workflows/issue-comment-workflow.md) |
| Comment Bot | Intermediate | Comment management with updates | [comment-bot-workflow.md](./workflows/comment-bot-workflow.md) |
| REST API | Intermediate | Direct GitHub API integration | [rest-api-workflow.md](./workflows/rest-api-workflow.md) |
| Go API CI/CD | Intermediate | Go application with automated testing | [go-api-workflow.md](./workflows/go-api-workflow.md) |

### Workflow Categories

- **Basic Workflows**: Simple examples for learning fundamentals
- **Intermediate Workflows**: Event handling and conditional logic
- **Future Workflows**: Complex automation and custom actions (coming soon)

*This list grows as new workflows are added to the repository.*

## Setup

1. Clone this repository
2. The workflows are automatically available in the `.github/workflows/` directory
3. They will trigger based on their respective event conditions

## Testing

Each workflow can be tested based on its trigger conditions:

- **Hello World**: Push changes to the main branch
- **Good Morning**: Wait for the scheduled time (9 AM UTC daily)
- **Issue Comment**: Create an issue and post a comment
- **Comment Bot**: Create an issue, post a comment, then edit it
- **REST API**: Create a PR to trigger API comment creation
- **Go API CI/CD**: Push changes or create PR to trigger build/test

See individual workflow documentation for detailed testing instructions.

## Permissions

Workflows that interact with GitHub resources require specific permissions:

- **Comment-based workflows**: `issues: write`, `pull-requests: write`
- **Repository workflows**: `contents: read` (default)
- **Custom permissions**: Defined in each workflow file

## Contributing

Feel free to submit issues and enhancement requests!