# (╯°□°）╯
Flippy is a Slackbot that flips tables and text.

## Commands
Invoke flippy and have a table flip as a response:

`/flip`

`(╯°□°）╯︵ ┻━┻`

Have flippy flip some text instead of a table:

`/flip dood`

`(╯°□°）╯︵ poop`

# App Configuration
Flippy is meant to run on a go buildpack on something like Heroku or dokku. The only configuration necessary before deploying your app is to set the `FLIPPY_SLACKTOKENS` environment variable to a :poop:-separated (`:poop:` not the emoji character since I have no idea how portable that would be or if it would even work) list of Slack tokens you want to answer to.

# Slack Configuration
TBD but the tl;dr is:

* Set up a Slack bot and set the URL to where you have flippy running
* Use the bot's token to set the `FLIPPY_SLACKTOKENS` environment variable in the above section.
* Run the bot
* In a channel the bot is listening on type `/flip` and you should get a `(╯°□°）╯︵ ┻━┻` response.

# Troubleshooting
Flippy logs to `STDOUT`.
