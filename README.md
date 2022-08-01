# TelCompiler
online code compiler telegram bot

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/Ja7ad/TelCompiler)

# Environment varaibles
- `BOT_PROVIDER` = sponsored user, channel or group show to end of bot message
- `PORT` = health and debug port
- `SENTRY_DSN` = sentry DSN address for tracking errors
- `TOKEN` = bot token

# How to run?
- build project with `go build .`
- set TOKEN, PORT environment variables
- run `./telcompiler`
