package bot

import "github.com/Ja7ad/telebot"

func start(ctx telebot.Context) error {
	return ctx.Send(MSG_START)
}

func help(ctx telebot.Context) error {
	return ctx.Send(MSG_HELP, &telebot.SendOptions{ParseMode: telebot.ModeHTML})
}

func golang(ctx telebot.Context) error {
	return ctx.Send("اوکی, لطفا کدی که به زبان Go هست را برایم بفرستید")
}

func python(ctx telebot.Context) error {
	return ctx.Send("اوکی, لطفا کدی که به زبان Python 3 هست را برایم بفرستید")
}

func clang(ctx telebot.Context) error {
	return ctx.Send("اوکی, لطفا کدی که به زبان C هست را برایم بفرستید")
}

func cppLang(ctx telebot.Context) error {
	return ctx.Send("اوکی, لطفا کدی که به زبان Cpp هست را برایم بفرستید")
}

func rust(ctx telebot.Context) error {
	return ctx.Send("اوکی, لطفا کدی که به زبان Rust هست را برایم بفرستید")
}
