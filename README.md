# GitHub Actions Workflows

This repository contains various GitHub Actions workflows for automation and testing.

## Workflows

### 1. Comment Bot (`comment-bot.yml`)

A smart comment bot that responds to issue comments and updates its reply when you edit your comment.

**Features:**
- Responds to new comments on issues
- Updates its previous reply when you edit your comment (no duplicate comments)
- Ignores its own comments to prevent infinite loops
- Uses a unique identifier to track and update the same comment

**How it works:**
1. When you post a comment → Bot creates a reply
2. When you edit your comment → Bot updates its previous reply
3. The bot uses `identifier: chatops-bot-reply` to find and update the same comment

**Trigger:** `issue_comment` events (created, edited)

### 2. Issue Comment Test (`issue-comment.yml`)

A simple workflow that prints issue comment details to the console.

**Features:**
- Logs issue number and comment body
- Useful for testing issue comment events

**Trigger:** `issue_comment` events (created)

### 3. Good Morning Workflow (`goodmorning-workflow.md`)

Documentation for a good morning automation workflow.

### 4. Hello Workflow (`hello-workflow.md`)

Documentation for a hello world workflow.

## Setup

1. Clone this repository
2. The workflows are automatically available in the `.github/workflows/` directory
3. They will trigger based on their respective event conditions

## Testing

To test the Comment Bot:
1. Create an issue in this repository
2. Post a comment on the issue
3. Edit your comment to see the bot update its reply

## Permissions

The Comment Bot workflow requires the following permissions:
- `issues: write` - To create and update comments
- `pull-requests: write` - To handle PR comments

## Contributing

Feel free to submit issues and enhancement requests!