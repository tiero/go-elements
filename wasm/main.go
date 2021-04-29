package main

import (
	"encoding/hex"
	"syscall/js"

	"github.com/vulpemventures/go-elements/pegin"
)

// Main function: it sets up our Wasm application
func main() {
	// Define the function in the JavaScript scope
	js.Global().Set("getPeginAddress", GetPeginAddress())
	// Prevent the function from returning, which is required in a wasm module
	select {}
}

func GetPeginAddress() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		pubKey := h2b(args[0].String())
		fedPegInfo := args[1].String()
		net := args[2].Int()
		isDynaFed := args[3].Bool()
		contract := h2b(args[4].String())
		pegin.GetAddressInfo(
			pubKey,
			fedPegInfo,
			net,
			isDynaFed,
			contract,
		)
		return map[string]interface{}{
			"hello":  "world",
			"answer": 42,
		}
	})
}

func b2h(buf []byte) string {
	return hex.EncodeToString(buf)
}

func h2b(str string) []byte {
	buf, _ := hex.DecodeString(str)
	return buf
}
