package options

import "github.com/alexflint/go-arg"

type Options struct {
	Token        string `arg:"env:TOKEN, -t" help:"telegram bot token, you can get from BotFather"`
	Port         string `arg:"env:PORT, -p" default:"8080" help:"pprof port for check stack"`
	Provider     string `arg:"env:PROVIDER, --provider" help:"set bot provider on bot answer"`
	SentryDSN    string `arg:"env:SENTRY, -s" help:"sentry dns address for tracking errors"`
	ProxyAddress string `arg:"env:PROXY_ADDRESS, --paddr" help:"socks5 proxy address ip:port"`
	ProxyUser    string `arg:"env:PROXY_USER, --puser" help:"socks5 proxy user"`
	ProxyPass    string `arg:"env:PROXY_PASS, --ppass" help:"socks5 proxy password"`
}

func InitOptions() *Options {
	opt := &Options{}
	arg.MustParse(opt)
	return opt
}
