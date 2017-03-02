# Notification team Coding Challenge

You are the new member of SEEKAsia technical team, and your first task is to build a system that work as the following:
Every 30 minutes it runs and calls github API to get the latest 10 updated repos and show this feed in Slack channel via Slack incoming webhook. 

However, there are few things to consider:
 - Dublications are not allowed, you can't feed the same repo that you sent to webhook previously. 
 - You need to customize the displayed name and the icon of the webhook to be showing your full name and any picture you choose.
 
 
You can [invite](https://notification-team-challenge.herokuapp.com/) yourself to SEEKAsia slack channel that we have setup for testing. 

Channel specific webhook url : https://hooks.slack.com/services/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

### Useful Links:
- Slack incoming webhook API [documentation](https://api.slack.com/incoming-webhooks).
- Github API [documentation](https://developer.github.com/v3/search/)
