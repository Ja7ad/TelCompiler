# TelCompiler
online code compiler telegram bot

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/Ja7ad/TelCompiler)

# Environment varaibles
- `ROVIDER` = sponsored user, channel or group show to end of bot message
- `PORT` = health and debug port
- `SENTRY` = sentry DSN address for tracking errors
- `TOKEN` = bot token

# How to run?
- build project with `go build .`
- set TOKEN, PORT environment variables or use flags
```shell
Usage: telcompiler [--token TOKEN] [--port PORT] [--provider PROVIDER] [--sentrydsn SENTRYDSN] [--paddr PADDR] [--puser PUSER] [--ppass PPASS]

Options:
  --token TOKEN, -t TOKEN
                         telegram bot token, you can get from BotFather [env: TOKEN]
  --port PORT, -p PORT   pprof port for check stack [default: 8080, env: PORT]
  --provider PROVIDER    set bot provider on bot answer [env: PROVIDER]
  --sentrydsn SENTRYDSN, -s SENTRYDSN
                         sentry dns address for tracking errors [env: SENTRY]
  --paddr PADDR          socks5 proxy address ip:port [env: PROXY_ADDRESS]
  --puser PUSER          socks5 proxy user [env: PROXY_USER]
  --ppass PPASS          socks5 proxy password [env: PROXY_PASS]
  --help, -h             display this help and exit

```
- run `./telcompiler`
