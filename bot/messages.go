package bot

import (
	"strings"
	"telcompiler/api/rextester"
)

const (
	MSG_START = "جهت دیدن راهنمای ربات کامپایلر دستور /help را بفرستید."
	MSG_HELP  = `
جهت استفاده از ربات شما می توانید با ارسال دستورات زیر کد مورد نظر خود را اجرا کرده و خروجی را بببینید:

<b>/go - اجرای کد زبان گولنگ</b>
<b>/py - اجرای کد زبان پایتون</b>
<b>/c - اجرای کد زبان سی</b>
<b>/cpp - اجرای کد زبان سی پلاس پلاس</b>
<b>/rust - اجرای کد زبان راست</b>
<b>/about - اطلاعات نویسنده ربات</b>
<b>/help - راهنمای ربات</b>
`
)

var escapeSpecialChars = strings.NewReplacer(
	"_", "\\_",
	"*", "\\*",
	"[", "\\[",
	"]", "\\]",
	"(", "\\(",
	")", "\\)",
	"~", "\\~",
	"`", "\\`",
	">", "\\>",
	"#", "\\#",
	"+", "\\+",
	"-", "\\-",
	"=", "\\=",
	"|", "\\|",
	"{", "\\{",
	"}", "\\}",
	".", "\\.",
	"!", "\\!",
)

func botResponseMessage() string {
	msg := "*زبان :* %s \n"
	msg += "*کاربر :* @%s \n"
	msg += "\n*کد ارسال شده :*\n"
	msg += "\n`%s`\n"
	msg += "*نتیجه :*\n"
	msg += "\n`%s`\n"
	msg += "*منابع مصرف شده :*\n"
	msg += "\n`%s`\n\n"
	msg += "%s"
	return msg
}

func checkMessageSize(code string) string {
	if len(code) > 500 {
		return "Character size is limited in Telegram"
	}
	return code
}

func escapeSpecialChar(msg string) string {
	return escapeSpecialChars.Replace(msg)
}

func resultCode(result *rextester.Result) string {
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
