package main

import (
	"expvar"
	"fmt"
	"github.com/getsentry/sentry-go"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"telcompiler/api"
	"telcompiler/bot"
	"telcompiler/global"
	"time"
)

func main() {
	if len(os.Getenv("SENTRY_DSN")) != 0 {
		if err := initSentry(); err != nil {
			log.Fatalf("error on initilize sentry %v", err)
		}
	}
	if err := bot.InitBot(); err != nil {
		sentry.CaptureException(err)
		log.Fatalf("error on initilize bot %v", err)
	}
	api.InitAPIClient()
	global.Bot.Use(middleware.Logger())
	bot.Commands()
	go bot.ProcessUpdate()
	log.Println("bot started")
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), pprofService()); err != nil {
			log.Fatal(err)
		}
	}()
	global.Bot.Start()

}

func initSentry() error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	}); err != nil {
		return err
	}
	defer sentry.Flush(2 * time.Second)
	return nil
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
