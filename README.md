# Notification Team Coding Challenge

As a member of notification engineering team serving notifications to millions of users, you have to build processes that calls various APIs and handles processing of the data to prepare it for consumption by other services.
 
The code challenge below is designed to test your basic understandings of some of these problems you may face in notification  team. 

### Problem:

You are the new member of SEEKAsia technical team, and your first task is to build a system that work as the following:
Every 30 minutes it runs and calls github API to get the latest 5 updated repos and show this feed in Slack channel via Slack incoming webhook. 

However, there are few things to consider:
 - Duplications are not allowed, you can't feed the same repo that you sent to webhook previously. 
 - You need to customize the displayed name and the icon of the webhook to be showing your full name and any picture you choose.
 
 
You can [invite](https://notification-team-challenge.herokuapp.com/) yourself to SEEKAsia slack channel that we have setup for testing. 

Channel specific webhook url : https://hooks.slack.com/services/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

### Useful Links:
- Slack incoming webhook API [documentation](https://api.slack.com/incoming-webhooks).
- Github API [documentation](https://developer.github.com/v3/search/)


### Judging criteria:

- **Code Quality**: how well-written and readable the code produced is, and how well-tested the critical paths in the code are.
- **Technology Choices**: how programming languages, libraries and tools were chosen and how tradeoffs inherent in those choices were handled.
- **Correctness & Efficiency**: An evaluation of whether the program generates correct output, and how fast and scalable the code is.
