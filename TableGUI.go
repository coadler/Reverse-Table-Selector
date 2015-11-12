package main

import (
	"github.com/unixpickle/gogui"
	"math"
	"os"
)

func main() {
	gogui.RunOnMain(openWindow)
	gogui.Main(&gogui.AppInfo{Name: "Random Table Selector"})
}

func openWindow() {
	// Create the window.
	w, _ := gogui.NewWindow(gogui.Rect{0, 0, 400, 400})
	w.SetTitle("Demo")
	w.Center()
	w.Show()
	w.SetCloseHandler(func() {
		os.Exit(0)
	})
}
