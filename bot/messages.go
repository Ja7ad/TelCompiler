package bot

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
	MSG_CODE = `
<b>زبان :</b> %s
<b>کاربر :</b> <i>@%s</i>

<b>کد ارسال شده :</b> 

<code>%s</code>

<b>نتیجه :</b> 

<code>%s</code>

<b>منابع مصرف شده : </b> 

<code>%s</code>

%s

`
)
