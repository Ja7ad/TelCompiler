package global

import (
	"github.com/valyala/fasthttp"
	"gopkg.in/telebot.v3"
)

var (
	Bot    *telebot.Bot
	Client *fasthttp.Client
)
