package bot

import "telcompiler/global"

func Commands() {
	global.Bot.Handle("/start", start)
	global.Bot.Handle("/help", help)
	global.Bot.Handle("/go", golang)
	global.Bot.Handle("/py", python)
	global.Bot.Handle("/c", clang)
	global.Bot.Handle("/cpp", cppLang)
	global.Bot.Handle("/rust", rust)
}
