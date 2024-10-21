package main

import (
	"syscall/js"

	"arbfeedtowasm/jsutil"
)

func main() {
	js.Global().Set("CalculateTransactionsRootWithStartTx", js.FuncOf(jsutil.CalculateTransactionsRootWithStartTx))

	select {}
}
