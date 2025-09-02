# Issue Comment Workflow

## Overview
This workflow demonstrates how to trigger GitHub Actions based on issue comments. It's a great example of event-driven automation that responds to user interactions with your repository.

## File Location
`.github/workflows/issue-comment.yml`

## What It Does
- Triggers when someone comments on an issue
- Prints the issue number and comment body to the console
- Demonstrates how to access GitHub event data in workflows

## Workflow Structure

```yaml
name: Test Issue Comment
on:
    issue_comment:
        types: [created]
jobs:
    Issue-comment:
        runs-on: ubuntu-latest
        steps:
            - name: Print comment
              run: echo "New comment on issue ${{ github.event.issue.number }} - ${{ github.event.comment.body }}"
```

## Key Learning Points

### 1. Issue Comment Events
- **issue_comment**: Triggers when comments are added to issues
- **types: [created]**: Only triggers on new comments (not edits or deletions)
- Valid types: `created`, `edited`, `deleted`

### 2. GitHub Event Context
- **github.event.issue.number**: Access the issue number
- **github.event.comment.body**: Access the comment content
- **github.event.issue.title**: Access the issue title
- **github.event.comment.user.login**: Access the commenter's username

### 3. Event-Driven Automation
- Workflows can respond to user interactions
- Useful for automated responses, notifications, or processing
- Enables interactive workflows based on repository activity

## How to Use

### Testing the Workflow
1. **Create an issue** in your repository
2. **Add a comment** to the issue
3. **Check the Actions tab** to see the workflow execution
4. **View the logs** to see the printed comment details

### Example Output
```
New comment on issue 1 - This is a test comment
```

## What I Learned

- How to configure issue comment triggers
- Understanding GitHub event context variables
- Event-driven workflow automation
- Accessing issue and comment data in workflows
- YAML syntax considerations for echo commands

## Advanced Use Cases

### 1. Automated Responses
```yaml
- name: Respond to comment
  uses: actions/github-script@v6
  with:
    script: |
      github.rest.issues.createComment({
        issue_number: context.issue.number,
        owner: context.repo.owner,
        repo: context.repo.repo,
        body: 'Thanks for your comment!'
      })
```

### 2. Comment Validation
- Check if comments contain specific keywords
- Validate comment format or content
- Trigger different actions based on comment content

### 3. Notification Systems
- Send notifications when specific users comment
- Alert maintainers about important discussions
- Integrate with external services (Slack, Discord, etc.)

## GitHub Event Context Variables

### Issue Data
- `github.event.issue.number` - Issue number
- `github.event.issue.title` - Issue title
- `github.event.issue.body` - Issue description
- `github.event.issue.state` - Issue state (open/closed)

### Comment Data
- `github.event.comment.body` - Comment content
- `github.event.comment.user.login` - Commenter username
- `github.event.comment.created_at` - Comment timestamp
- `github.event.comment.html_url` - Comment URL

### Repository Data
- `github.event.repository.name` - Repository name
- `github.event.repository.owner.login` - Repository owner
- `github.event.repository.full_name` - Full repository name

## Troubleshooting

### Common Issues
- **YAML syntax errors**: Be careful with colons in echo strings
- **Event context**: Make sure you're using the correct event variables
- **Permissions**: Some actions require specific permissions

### Best Practices
- Use descriptive step names
- Test with simple comments first
- Be mindful of rate limits for API calls
- Consider security implications of automated responses

## Next Steps

This workflow demonstrates basic issue comment handling. Future learning areas include:
- Pull request comment workflows
- Automated issue labeling
- Comment-based command processing
- Integration with external services
- Advanced GitHub API interactions

## Related Workflows

- [Hello World Workflow](./hello-workflow.md) - Basic workflow introduction
- [Good Morning Workflow](./goodmorning-workflow.md) - Multiple jobs and triggers
