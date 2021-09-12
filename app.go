package main

import (
	"fmt"
	"github.com/AllenDang/giu"
)

var (
	apps = map[string]bool{"Geography": false, "Quiz Game": false}
)

func loop() {
	giu.SingleWindowWithMenuBar().Layout()
	giu.MainMenuBar().Layout(
		giu.Menu("Apps").Layout(
			giu.MenuItem("Geography").OnClick(func() {
				apps["Geography"] = !apps["Geography"]
			}),
			giu.MenuItem("Quiz Game").OnClick(func() {
				apps["Quiz Game"] = !apps["Quiz Game"]
			}),
			giu.Menu("Text Editor").Layout(
				giu.Menu("New").Layout(
					giu.MenuItem("Text Document"),
					giu.MenuItem("Excel"),
				),
				giu.MenuItem("Open"),
			),
			giu.Menu("Help").Layout(
				giu.MenuItem("About"),
			),
		),
	).Build()

	for k, v := range apps {
		if !v {
			giu.Window(k).IsOpen(&v).Flags(giu.WindowFlagsNone).Layout(
				giu.Label(fmt.Sprintf("This is the %s Window", k)),
				giu.Button("toggle-window").OnClick(func() {
					apps[k] = !v
				}),
			)
		}
	}
}

func main() {
	win := giu.NewMasterWindow("Giu Apps", 1000, 800, 0)
	win.Run(loop)
}
