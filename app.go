package main

import (
	"fmt"
	"github.com/AllenDang/giu"
	"imgui-based-app/components/geography"
)

type MenuItems struct {
	isFirstTimeLoaded bool
}

// MainApps /TODO: Remember to switch to dynamic allocation instead of hardcoded logic
// MenuItemsRef - To be able to verify the boolean easier during Widget rendering iterations
var (
	MainApps     = map[string]bool{"Geography": false, "Quiz Game": false, "OS Info": false}
	MenuItemsRef = &MenuItems{}
	QuizWindow   *giu.WindowWidget
	GeoWindow    *giu.WindowWidget
)

func loop() {

	/// This MUST BE RAN ONLY ONCE, at startup! so it can limit the number of requests
	///   will fix in the future, when a sqlite concept will be prototyped.
	if !geography.CountryRef.IsUpdated {
		err := geography.InitCountries()
		if err != nil {
			return
		}
	}

	/// The main window of the app
	// TODO: To be Updated, to prevent crazy overhead from so many conditional sentences
	giu.SingleWindowWithMenuBar().Layout(
		giu.MenuBar().Layout(
			giu.Menu("Apps").Layout(
				giu.RangeBuilder("main-apps", []interface{}{
					"Geography", "Quiz Game", "OS Info",
				}, func(i int, v interface{}) giu.Widget {
					currentApp := v.(string)

					// if false, then set it to true, until the application closes entirely.
					if MenuItemsRef.isFirstTimeLoaded {
						return giu.MenuItem(currentApp).OnClick(func() {
							if MainApps[currentApp] {
								switch window := currentApp; window {
								case "Geography":
									GeoWindow.BringToFront()
									break
								case "Quiz Game":
									QuizWindow.BringToFront()
									break
								default:
									break
								}
							} else {
								MainApps[currentApp] = !MainApps[currentApp]
							}
						}).Selected(MainApps[currentApp])
					} else {
						MenuItemsRef.isFirstTimeLoaded = true
						return giu.MenuItem(currentApp).OnClick(func() {
							MainApps[currentApp] = false
						}).Selected(MainApps[currentApp])
					}
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

	/// These represent the separate App Windows
	for k, v := range MainApps {
		if v {
			switch window := k; window {
			case "Geography":
				GeoWindow = giu.Window(k)
				GeoWindow.IsOpen(&v).Flags(giu.WindowFlagsAlwaysUseWindowPadding).Layout(
					geography.CountriesTable(),
				)
				break
			case "Quiz Game":
				QuizWindow = giu.Window(k)
				QuizWindow.IsOpen(&v).Flags(giu.WindowFlagsAlwaysUseWindowPadding).Layout(
					giu.Label(fmt.Sprintf("This is the %s Window", k)),
					giu.Button("toggle-window").OnClick(func() {
						MainApps[k] = !v
					}),
				)
			}
		}
	}
}

func main() {
	win := giu.NewMasterWindow("Giu Apps", 840, 480, giu.MasterWindowFlagsMaximized)
	win.Run(loop)
}
