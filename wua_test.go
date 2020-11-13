package wua

import "testing"

func TestRun(t *testing.T) {
	Run(Settings{Title: "Hello",
		Left:      200,
		Top:       50,
		Width:     400,
		Height:    510,
		Resizable: true,
		Debug:     true})
}
