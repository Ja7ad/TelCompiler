package bot

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"strings"
	"telcompiler/api"
	"telcompiler/global"
	"time"
)

func InitBot() error {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:   os.Getenv("TOKEN"),
		Poller:  &telebot.LongPoller{Timeout: 5 * time.Second},
		Updates: 1000,
	})
	if err != nil {
		return err
	}
	global.Bot = bot
	return nil
}

func ProcessUpdate() {
	for {
		select {
		case up := <-global.Bot.Updates:
			if up.Message.ReplyTo != nil {
				go processCompileCode(up.Message)
			}
		}
	}
}

func processCompileCode(message *telebot.Message) {
	languageCode := getLanguageCode(message.ReplyTo.Text)
	log.Printf("new compile code language %v by user %v requested", languageCode, message.Sender.Username)
	if languageCode == 0 {
		return
	}
	res, err := api.RequestCompileCode(languageCode, message.Text)
	if err != nil {
		sentry.CaptureException(err)
	}
	if _, err := global.Bot.Reply(message, normalizeReplayMessage(message, res), &telebot.SendOptions{ParseMode: telebot.ModeHTML}); err != nil {
		sentry.CaptureException(err)
	}
}

func normalizeReplayMessage(msg *telebot.Message, result *api.Result) string {
	return fmt.Sprintf(MSG_CODE, result.Language, msg.Sender.Username, checkCodeSize(msg.Text), resultCode(result), result.Stats, os.Getenv("BOT_PROVIDER"))
}

func getLanguageCode(msg string) int {
	if strings.Contains(msg, "Go") {
		return 20
	} else if strings.Contains(msg, "Cpp") {
		return 7
	} else if strings.Contains(msg, "C") {
		return 6
	} else if strings.Contains(msg, "Python") {
		return 24
	} else if strings.Contains(msg, "Rust") {
		return 46
	} else {
		return 0
	}
}

func checkCodeSize(code string) string {
	if len(code) > 500 {
		return "Character size is limited in Telegram"
	}
	return code
}

func resultCode(result *api.Result) string {
	res := ""
	if len(result.Result) != 0 {
		res = result.Result
	}
	if len(result.Errors) != 0 {
		res = result.Errors
	}
	if len(result.Warnings) != 0 {
		res = result.Warnings
	}
	return res
}
