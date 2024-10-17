package jsutil

import (
	"syscall/js"
)

func logToConsole(message string) {
	js.Global().Get("console").Call("log", message)
}
