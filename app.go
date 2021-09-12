package main

import (
	"fmt"
	"github.com/AllenDang/giu"
	"imgui-based-app/components"
)

var (
	Apps = map[string]bool{"Geography": false, "Quiz Game": false}
)

func loop() {

	/// this MUST BE RAN ONLY ONCE, at startup! so it can limit the number of requests
	///   will fix in the future, when a sqlite concept will be prototyped.
	if !components.Countries.IsUpdated {
		err := components.InitCountries()
		if err != nil {
			return
		}
	}

	/// the main window of the app
	giu.SingleWindowWithMenuBar().Layout(
		giu.MenuBar().Layout(
			giu.Menu("Apps").Layout(
				giu.MenuItem("Geography").OnClick(func() {
					Apps["Geography"] = !Apps["Geography"]
				}),
				giu.MenuItem("Quiz Game").OnClick(func() {
					Apps["Quiz Game"] = !Apps["Quiz Game"]
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
		),
	)

	for k, v := range Apps {
		if v {
			switch window := k; window {
			case "Geography":
				geoWin := giu.Window(k)
				geoWin.IsOpen(&v).Flags(giu.WindowFlagsNone).Layout(
					components.CountriesTable(),
				)
				break
			case "Quiz Game":
				giu.Window(k).IsOpen(&v).Flags(giu.WindowFlagsNone).Layout(
					giu.Label(fmt.Sprintf("This is the %s Window", k)),
					giu.Button("toggle-window").OnClick(func() {
						Apps[k] = !v
					}),
				)
			}
		}
	}
}

func main() {
	win := giu.NewMasterWindow("Giu Apps", 1000, 800, 0)
	win.Run(loop)
}
