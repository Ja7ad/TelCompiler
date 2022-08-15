package main

import (
	"crypto/tls"
	"expvar"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/ja7ad/telcompiler/bot"
	"github.com/ja7ad/telcompiler/global"
	"github.com/ja7ad/telcompiler/options"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	opt := options.InitOptions()
	if len(opt.SentryDSN) != 0 {
		if err := initSentry(opt.SentryDSN); err != nil {
			log.Fatalf("error on initilize sentry %v", err)
		}
	}
	if err := bot.InitBot(opt.Token, opt.ProxyAddress, opt.ProxyUser, opt.ProxyPass); err != nil {
		sentry.CaptureException(err)
		log.Fatalf("error on initilize bot %v", err)
	}
	initClient()
	bot.Commands()
	go func(provider string) {
		bot.ProcessUpdate(provider)
	}(opt.Provider)
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%s", opt.Port), pprofService()); err != nil {
			log.Fatal(err)
		}
	}()
	shutdownBot()
	log.Println("bot started")
	global.Bot.Start()

}

func initSentry(dsn string) error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
	}); err != nil {
		return err
	}
	defer sentry.Flush(2 * time.Second)
	return nil
}

func initClient() {
	global.Client = &fasthttp.Client{
		Name:                     "telCompiler",
		NoDefaultUserAgentHeader: true,
		TLSConfig:                &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost:          5000,
		MaxIdleConnDuration:      5 * time.Second,
	}
}

func pprofService() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", healthCheck)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())
	return mux
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("telcompiler is okey!!"))
}

func shutdownBot() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		log.Println("shutting down bot")
		global.Bot.Stop()
	}()
}
