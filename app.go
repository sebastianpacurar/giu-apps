package main

import (
	"fmt"
	"github.com/AllenDang/giu"
	"imgui-based-app/components/giu-geography"
)

type MenuItems struct {
	isFirstTimeLoaded bool
}

var (
	MasterWidth     = 840
	MasterHeight    = 480
	FullHeight      float32
	MainMenuWidth   = float32(MasterWidth / 3)
	AppsWindowWidth = float32(MasterWidth) - MainMenuWidth
	MainApps        = map[string]bool{"Geography": false, "Quiz Game": false, "OS Info": false}
	QuizWindow      *giu.WindowWidget
	GeoWindow       *giu.WindowWidget
	titleFont         *giu.FontInfo
)

func loop() {

	/// This MUST BE RAN ONLY ONCE, at startup! so it can limit the number of requests
	///   will fix in the future, when a sqlite concept will be prototyped.
	if !giu_geography.CountryRef.IsUpdated {
		err := giu_geography.InitCountries()
		if err != nil {
			return
		}
	}

	size := giu.Context.GetPlatform().DisplaySize()
	FullHeight = size[1]

	if int(size[0]) >= 840 {
		MainMenuWidth = size[0] / 5
		AppsWindowWidth = size[0] - MainMenuWidth
	}

	/// The main window of the app
	// TODO: To be Updated, to prevent crazy overhead from so many conditional sentences
	giu.Window("Main Menu").
		///Size = LHN Menu like size and position
		Size(MainMenuWidth, FullHeight).
		Pos(0, 0).
		Flags(
			giu.WindowFlagsNoMove |
				giu.WindowFlagsNoResize |
				giu.WindowFlagsNoTitleBar,
		).
		Layout(
			giu.Label("test").Wrapped(true).Font(titleFont),
		)

	giu.Window("Apps Layout").
		Size(AppsWindowWidth, FullHeight).
		Pos(MainMenuWidth, 0).
		Flags(
			giu.WindowFlagsNoMove |
				giu.WindowFlagsNoResize |
				giu.WindowFlagsNoTitleBar,
		).
		Layout(
			giu.Label("test 2").Wrapped(true).Font(&giu.FontInfo{}),
		)

	/// These represent the separate App Windows
	for k, v := range MainApps {
		if v {
			switch window := k; window {
			case "Geography":
				GeoWindow = giu.Window(k).Size(size[0]/4, size[1]/4)
				GeoWindow.Flags(giu.WindowFlagsMenuBar).Layout(
					giu_geography.CountriesTable(),
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
	// Change the default font to sans and of 20 pixels height
	giu.SetDefaultFont("Sans.ttf", 20)

	// change titleFont to look larger
	titleFont = giu.AddFont("Sans.ttf", 28)

	win := giu.NewMasterWindow("Universal App", MasterWidth, MasterHeight, giu.MasterWindowFlagsMaximized)
	win.Run(loop)
}
