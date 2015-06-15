# (╯°□°）╯
Flippy is a webhook responder for Slack that flips tables and text.

# Installation

## App Configuration
Flippy is meant to run on a go buildpack on something like Heroku or dokku on a public URL.

Two environment variables need to be set:

* `FLIPPY_SLACKTOKENS` - *required* - a :poop:-separated (`:poop:` not the emoji character since I have no idea how portable that would be or if it would even work) list of Slack tokens (set in the webhook below) you want to answer for.
* `FLIPPY_TRIGGERWORD` - optional - an optional trigger word for flippy to listen for. Defaults to `flip`.

## Slack Configuration

* Set up an Outgoing Webhook in Slack and set the following:
  * Channel: the channel to listen or on or all channels
  * Trigger Word(s): what you set `FLIPPY_TRIGGERWORD` in the above or `flip` if you didn't set it
  * URL(s): The site where you have flippy running followed by `/slack` (e.g `http://flippy.example.com/slack`)

# Commands
The examples below assume you are using the default trigger word `flip`. If you are using a custom trigger word use that instead.

## Invoke flippy and have a table flip as a response:

`flip`

`(╯°□°）╯︵ ┻━┻`

## Have flippy flip some text instead of a table:

`flip dood`

`(╯°□°）╯︵ poop`

# Troubleshooting

* Check all your settings
* Make sure you are typing the trigger word at the beginning of the line
* If you are using a custom trigger word be sure you are using that and not `flip`
* If you have restricted the webhook to a single channel make sure you are in that channel
* Flippy logs to `STDOUT`
