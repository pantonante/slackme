# Slackme BOT

I often need to receive notifications from a remote server whenever a long job/process ends.
That's why I build this embarrassingly simple bot, `slackme`!

## Prerequisite

- Golang

### Install

Just run `go get github.com/pantonante/slackme`.

## Setup

You need to create a Slack Incoming Webhooks, follow the guide [here](https://api.slack.com/incoming-webhooks).

One you obtained your Incoming Webhook URL create the file `$HOME/.slackme.yaml` and put the following content
```
default:
  https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX
direct:
  https://hooks.slack.com/services/T00000000/B00000000/YYYYYYYYYYYYYYYYYYYYYYYY
```
Replace the dummy URL with your, newly createt Webook URL.

## Usage

Suppose you have to run the command `longprocess`, now run
```
longprocess; slackme message "Long Process finished. (Ret: $?)"
```
This command will send a message to the `default` webhook when the process finishes with its return code.

You can also specify a channel name, e.g.,
```
longprocess; slackme message -c direct "Long Process finished. (Ret: $?)" 
```
