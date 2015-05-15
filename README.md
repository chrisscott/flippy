# Flippy
Flippy is a Slackbot that flips tables and text.

## Commands
Invoke flippy and have a table flip as a response:
`/flip`

Have flippy flip some text instead of a table:
`/flip dood`

# App Configuration
Flippy is meant to run on a go buildpack on something like Heroku or dokku. The only configuration necessary before deploying your app is to set the `FLIPP_SLACKTOKENS` environment variable to a 💩-separated lists of Slack tokens you want to answer to.

# Slack Configuration
TBD but the tl;dr is set up a bot and use the token in the App Configuration above and set the URL to where you have flippy running.
