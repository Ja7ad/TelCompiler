package global

import (
	"github.com/Ja7ad/telebot"
	"github.com/valyala/fasthttp"
)

var (
	Bot    *telebot.Bot
	Client *fasthttp.Client
)
