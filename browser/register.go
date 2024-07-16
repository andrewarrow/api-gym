package browser

import "github.com/andrewarrow/feedback/wasm"

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	if Global.Start == "welcome.html" {
		RegisterLoginEvents()
	} else if Global.Start == "gym.html" {
		RegisterGymEvents()
	} else if Global.Start == "endpoints.html" {
		RegisterEndpoints()
	}
}

func RegisterLoginEvents() {
}
