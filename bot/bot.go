package bot

import (
	"fmt"
	"github.com/Ja7ad/telebot"
	"github.com/getsentry/sentry-go"
	"github.com/ja7ad/telcompiler/api/rextester"
	"github.com/ja7ad/telcompiler/global"
	"log"
	"strings"
)

func InitBot(token string, proxyAddr, proxyUser, proxyPass string) error {
	settings := telebot.Settings{
		Token:   token,
		Updates: 2000,
	}
	if len(proxyAddr) != 0 {
		settings.Proxy = &telebot.Proxy{
			Address:  proxyAddr,
			UserName: proxyUser,
			Password: proxyPass,
		}
	}
	bot, err := telebot.NewBot(settings)
	if err != nil {
		return err
	}
	global.Bot = bot
	return nil
}

func ProcessUpdate(provider string) {
	for update := range global.Bot.Updates {
		global.Bot.ProcessUpdate(update)
		if update.Message != nil {
			if update.Message.ReplyTo != nil {
				go func(msg *telebot.Message, provider string) {
					processCompileCode(msg, provider)
				}(update.Message, provider)
			}
		}
	}
}

func processCompileCode(message *telebot.Message, provider string) {
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
	msg := normalizeReplayMessage(message, res, provider)
	if _, err := global.Bot.Reply(message, msg, &telebot.SendOptions{ParseMode: telebot.ModeMarkdownV2}); err != nil {
		sentry.CaptureException(err)
		log.Printf("error on replay %v\n msg %v", err, msg)
	}
}

func normalizeReplayMessage(msg *telebot.Message, result *rextester.Result, provider string) string {
	return fmt.Sprintf(botResponseMessage(), result.Language, msg.Sender.Username, escapeSpecialChar(checkMessageSize(msg.Text)), escapeSpecialChar(resultCode(result)), result.Stats, provider)
}

func getLanguageCode(msg string) int {
	if len(msg) > 100 {
		return 0
	}
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
