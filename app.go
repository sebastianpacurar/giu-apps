package main

import (
	"github.com/AllenDang/giu"
	"image/color"
)

type AppI interface{}
type MiniAppI interface{}

var (
	MasterWidth     = 840
	MasterHeight    = 480
	FullHeight      float32
	MainMenuWidth   = float32(MasterWidth / 3)
	AppsWindowWidth = float32(MasterWidth) - MainMenuWidth
	titleFont       *giu.FontInfo
)

// AppsS - The struct for the Menu
// AppsI - The AppsList[] as an Interface (to be used with RangeBuilder() as values param)
var (
	AppsI = make([]interface{}, len(AppsS.AppsList))
	AppsS = &Apps{
		AppsList: []App{
			{
				Name:   "Geography",
				Active: false,
				MiniApps: []MiniApp{
					{
						Name:   "Countries Table",
						Active: false,
					},
					{
						Name:   "Map",
						Active: false,
					},
				},
			},

			{
				Name:   "Dictionary",
				Active: false,
				MiniApps: []MiniApp{
					{
						Name:   "English Explicative Dictionary",
						Active: false,
					},
				},
			},
		},
	}
)

// Apps / TODO: Think of a way to work with structs rather than variables
type Apps struct {
	AppsList []App
}

type App struct {
	Name     string
	Active   bool
	MiniApps []MiniApp
}

type MiniApp struct {
	Name   string
	Active bool
}

func loop() {

	size := giu.Context.GetPlatform().DisplaySize()
	FullHeight = size[1]

	if int(size[0]) >= 840 {
		MainMenuWidth = size[0] / 5
		AppsWindowWidth = size[0] - MainMenuWidth
	}

	//
	for i := range AppsI {
		AppsI[i] = AppI(AppsS.AppsList[i])
	}
	// The app consists of 2 main windows:
	// "Main Menu" and "Apps Layout"
	giu.Window("Main Menu").
		// Size = LHN Menu like size and position
		Size(MainMenuWidth, FullHeight).
		Pos(0, 0).
		Flags(
			giu.WindowFlagsNoMove |
				giu.WindowFlagsNoResize |
				giu.WindowFlagsNoTitleBar,
		).
		Layout(
			giu.Child().
				Border(true).
				Layout(
					// This is the Title of the Main Menu. set Text Color to Cyan rgba(0, 255, 255, 255)
					// Also, use the titleFont of 28px sans
					giu.Style().
						SetColor(giu.StyleColorText, color.RGBA{G: 255, B: 255, A: 255}).
						To(
							giu.Label("Main Menu").Wrapped(true).Font(titleFont),
						),
					giu.Separator(),
					giu.TreeNode("Apps").
						Flags(giu.TreeNodeFlagsCollapsingHeader).
						Layout(
							// This is where the Main Menu items is generated
							giu.RangeBuilder("Menu", AppsI, func(i int, v interface{}) giu.Widget {
								currApp := &AppsS.AppsList[i]
								MiniAppsI := make([]interface{}, len(currApp.MiniApps))
								for i := range MiniAppsI {
									MiniAppsI[i] = MiniAppI(currApp.MiniApps[i])
								}
								return giu.TreeNode(currApp.Name).
									Layout(
										// This is where the Sub Menu for every Menu Item will be generated
										giu.RangeBuilder("Sub Menu", MiniAppsI, func(j int, v interface{}) giu.Widget {
											currMiniApp := &currApp.MiniApps[j]
											return giu.Row(
												giu.Checkbox("", &currMiniApp.Active).
													OnChange(func() {
														currMiniApp.Active = !currMiniApp.Active
													}),
												giu.Selectable(currMiniApp.Name).
													OnClick(func() {
														currMiniApp.Active = !currMiniApp.Active
													}).
													Selected(currMiniApp.Active),
											)
										}),
									)
							}),
						),
				),
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
}

func main() {
	// Change the default font to sans and of 20 pixels height
	giu.SetDefaultFont("Sans.ttf", 15)

	// change titleFont to look larger
	titleFont = giu.AddFont("Sans.ttf", 28)

	win := giu.NewMasterWindow("Universal App", MasterWidth, MasterHeight, giu.MasterWindowFlagsMaximized)
	win.Run(loop)
}
