package main

import (
	"github.com/getsentry/sentry-go"
	"log"
	"net/http"
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
	bot.Commands()
	go bot.ProcessUpdate()
	go global.Bot.Start()
	log.Println("bot started")
	_ = http.ListenAndServe(os.Getenv("PORT"), http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("ok"))
		},
	))
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
