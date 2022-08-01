package bot

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"strings"
	"telcompiler/api/rextester"
	"telcompiler/global"
)

func InitBot() error {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:   os.Getenv("TOKEN"),
		Updates: 2000,
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
			if up.Message != nil {
				if up.Message.ReplyTo != nil {
					go processCompileCode(up.Message)
				}
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
	res, err := rextester.RequestCompileCode(languageCode, message.Text)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("error on request %v", err)
	}
	msg := normalizeReplayMessage(message, res)
	if _, err := global.Bot.Reply(message, msg, &telebot.SendOptions{ParseMode: telebot.ModeMarkdownV2}); err != nil {
		sentry.CaptureException(err)
		log.Printf("error on replay %v\n msg %v", err, msg)
	}
}

func normalizeReplayMessage(msg *telebot.Message, result *rextester.Result) string {
	return fmt.Sprintf(botResponseMessage(), result.Language, msg.Sender.Username, escapeSpecialChar(checkMessageSize(msg.Text)), escapeSpecialChar(resultCode(result)), result.Stats, os.Getenv("BOT_PROVIDER"))
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
