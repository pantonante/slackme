# Slackme BOT

I often need to receive notifications from a remote server whenever a long job/process ends.
That's why I build this embarrassingly simple bot, `slackme`!

## Prerequisite

- Golang

## Setup

You need to create a Slack Incoming Webhooks, follow the guide [here](https://api.slack.com/incoming-webhooks).

One you obtained your Incoming Webhook URL create the file `$HOME/.slackme.yaml` and put the following content
```
default:
  https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX
```
Replace the dummy URL with your, newle create Webook URL.

## Usage

Suppose you have to run the command `longprocess`, now run

```
longprocess; slackme message "Long Process finished. (Ret: $?)"
```

This command will send a message when the process finishes with its return code.