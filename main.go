package main

import (
	_ "embed"
	"gostick/internal/runtime"
	"os"

	"github.com/getlantern/systray"
)

//go:embed assets/gostick.ico
var icon []byte

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {

	// Start GoStick runtime
	go runtime.Start()

	// Set tray icon
	systray.SetIcon(icon)

	// Tray metadata
	systray.SetTitle("GoStick")
	systray.SetTooltip("GoStick Running")

	// Menu items
	mQuit := systray.AddMenuItem(
		"Quit",
		"Quit GoStick",
	)

	// Quit handler
	go func() {

		<-mQuit.ClickedCh

		systray.Quit()
		os.Exit(0)
	}()
}

func onExit() {

}
