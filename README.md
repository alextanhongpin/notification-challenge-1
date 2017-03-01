# Notification team Coding Challenge
SEEK Asia notification team coding challenge.

You have to build a system that calls github API for latest updated repo and feed this information to an external Slack webhook. Slack incoming webhook API [documentation](https://api.slack.com/incoming-webhooks).

Github API [Documentation](https://developer.github.com/v3/search/)

Few caveats :
 - You can't feed the same repo that you sent to webhook previously. 
 - The process should run every 30 minutes automatically
 - You can customize the name and icon of your Incoming Webhook in the Integration Settings section. However, you can override the displayed name by sending "username": "new-bot-name" in your JSON payload. You can also override the bot icon either with "icon_url": "https://slack.com/img/icons/app-57.png" or "icon_emoji": ":ghost:".
