package main

import (
	"github.com/AllenDang/giu"
	"image/color"
)

type AppI interface{}
type MiniAppI interface{}

var (
	FullHeight      float32
	FullWidth       float32
	MainMenuWidth   float32
	AppsWindowPosX  int
	AppsWindowWidth float32
	isMenuToggled   = true
	MenuBarHeight   = float32(23)
	titleFont       *giu.FontInfo
)

// AppsS - The struct of the Menu
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
						Name:   "English",
						Active: false,
					},
				},
			},
			{
				Name:   "Text Handler",
				Active: false,
				MiniApps: []MiniApp{
					{
						Name:   "Bash Console",
						Active: false,
					},
					{
						Name:   "JSON Formatter",
						Active: false,
					},
					{
						Name:   "Text Editor",
						Active: false,
					},
				},
			},
		},
	}
)

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

// ConditionedArrowBtn - is used to swap directions of the arrow after each click
func ConditionedArrowBtn() giu.Widget {
	var arrowBtn *giu.ArrowButtonWidget
	if isMenuToggled {
		arrowBtn = giu.ArrowButton("close menu", giu.DirectionLeft).OnClick(func() {
			isMenuToggled = false
		})
	} else {
		arrowBtn = giu.ArrowButton("open menu", giu.DirectionRight).OnClick(func() {
			isMenuToggled = true
		})
	}
	return arrowBtn
}

func loop() {

	size := giu.Context.GetPlatform().DisplaySize()
	FullWidth = size[0]
	FullHeight = size[1]

	// For sizes bigger than 990px use responsive width
	// If the Main Menu is closed, then stretch Apps Window to full width
	if int(size[0]) >= 990 {
		MainMenuWidth = size[0] / 4
		AppsWindowPosX = int(MainMenuWidth)
	} else {
		MainMenuWidth = 250
		AppsWindowPosX = 250
	}
	if !isMenuToggled {
		AppsWindowWidth = FullWidth
		AppsWindowPosX = 0
		MainMenuWidth = 0
	}
	AppsWindowWidth = FullWidth - MainMenuWidth

	// Create a list of interfaces converted from struct
	for i := range AppsI {
		AppsI[i] = AppI(AppsS.AppsList[i])
	}

	giu.Window("Menu Bar").
		Pos(0, 0).
		Flags(
			giu.WindowFlagsNoMove |
				giu.WindowFlagsNoResize |
				giu.WindowFlagsNoTitleBar,
		).Layout(
		giu.MainMenuBar().Layout(
			ConditionedArrowBtn(),
		),
	)

	// The app consists of 2 main windows:
	// "Main Menu" and "Apps Layout"
	if isMenuToggled {
		giu.Window("Main Menu").
			// Size = LHN Menu-like size and position
			Size(MainMenuWidth, FullHeight-MenuBarHeight).
			Pos(0, MenuBarHeight).
			Flags(
				giu.WindowFlagsNoMove |
					giu.WindowFlagsNoResize |
					giu.WindowFlagsNoTitleBar,
			).
			Layout(
				giu.Child().
					Border(true).
					Layout(
						// This is the Title of the20 Main Menu. set Text Color to Cyan rgba(0, 255, 255, 255)
						// Also, use the titleFont of 28px sans
						giu.Row(
							giu.Style().
								SetColor(giu.StyleColorText, color.RGBA{G: 255, B: 255, A: 255}).
								To(
									giu.Label("Main Menu").Wrapped(true).Font(titleFont),
								),
						),
						giu.Separator(),
						// LAYOUT Menu
						giu.TreeNode("Layout").
							Flags(giu.TreeNodeFlagsCollapsingHeader).
							Layout(giu.Label("test")),

						// APPS Menu
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
										Flags(giu.TreeNodeFlagsSpanFullWidth).
										Layout(
											// This is where the Sub Menu for every Menu Item will be generated
											giu.RangeBuilder("Sub Menu", MiniAppsI, func(j int, v interface{}) giu.Widget {
												currMiniApp := &currApp.MiniApps[j]
												return giu.Row(
													giu.Checkbox("", &currMiniApp.Active),
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
	}
	giu.Window("Apps").
		Size(AppsWindowWidth, FullHeight-MenuBarHeight).
		Pos(float32(AppsWindowPosX), MenuBarHeight).
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
	// Change the default font to sans and of 18 pixels height
	giu.SetDefaultFont("Sans.ttf", 18)

	// Change titleFont to look larger
	titleFont = giu.AddFont("Sans.ttf", 28)

	win := giu.NewMasterWindow("Universal App", 960, 640, giu.MasterWindowFlagsMaximized)
	win.Run(loop)
}
