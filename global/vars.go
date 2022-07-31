package global

import (
	"gopkg.in/resty.v1"
	"gopkg.in/telebot.v3"
)

var (
	Bot    *telebot.Bot
	Client *resty.Client
)
