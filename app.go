package main

import (
	"github.com/AllenDang/giu"
	"image/color"
)

type AppI interface{}
type MiniAppI interface{}

var (
	fullWidth, fullHeight          float32
	sideMenuWidth, appsWindowWidth float32
	appsWindowPosX                 int
	isSideMenuOpen                 = true
	menuBarHeight                  = float32(23)
	titleFont                      *giu.FontInfo
	smallFont                      *giu.FontInfo
	largeFont                      *giu.FontInfo
)

// appS - The struct of the Menu
// appsI - The appsList[] as an Interface (to be used with RangeBuilder() as values param)
var (
	appsI = make([]interface{}, len(appS.appsList))
	appS  = &Apps{
		appsList: []App{
			{
				name:   "Geography",
				active: false,
				miniApps: []MiniApp{
					{
						name:   "All Countries",
						active: false,
					},
					{
						name:   "Map",
						active: false,
					},
				},
			},

			{
				name:   "Dictionary",
				active: false,
				miniApps: []MiniApp{
					{
						name:   "English",
						active: false,
					},
				},
			},

			{
				name:   "Maths",
				active: false,
				miniApps: []MiniApp{
					{
						name:   "Calculator",
						active: false,
					},
					{
						name:   "Geometry",
						active: false,
					},
					{
						name:   "Trigonometry",
						active: false,
					},
				},
			},

			{
				name:   "Text Handler",
				active: false,
				miniApps: []MiniApp{
					{
						name:   "Bash Console",
						active: false,
					},
					{
						name:   "JSON Formatter",
						active: false,
					},
					{
						name:   "Text Editor",
						active: false,
					},
				},
			},
		},
	}
	layoutS = &Layout{
		types:      []string{"Window", "Splitter"},
		windows:    []string{"1", "2", "3", "4"},
		directions: []string{"Wrap", "Vertical", "Horizontal"},
	}
)

type Apps struct {
	appsList []App
}

type App struct {
	name     string
	active   bool
	miniApps []MiniApp
}

type MiniApp struct {
	name   string
	active bool
}

type Layout struct {
	typesIndex, windowsIndex, directionsIndex int32
	types, windows, directions                []string
	currWindowsNo                             int
	currType                                  string
	currDirection                             string
	activeWindows                             []*giu.WindowWidget
}

// conditionedArrowBtn - is used to swap directions of the arrow after each click
func conditionedArrowBtn() giu.Widget {
	var arrowBtn *giu.ArrowButtonWidget
	if isSideMenuOpen {
		arrowBtn = giu.ArrowButton("close menu", giu.DirectionLeft).OnClick(func() {
			isSideMenuOpen = false
		})
	} else {
		arrowBtn = giu.ArrowButton("open menu", giu.DirectionRight).OnClick(func() {
			isSideMenuOpen = true
		})
	}
	return arrowBtn
}

func loop() {

	size := giu.Context.GetPlatform().DisplaySize()
	fullWidth = size[0]
	fullHeight = size[1]

	// For sizes bigger than 990px use responsive width
	// If the Main Menu is closed, then stretch Apps Window to full width
	if int(size[0]) >= 990 {
		sideMenuWidth = size[0] / 4
		appsWindowPosX = int(sideMenuWidth)
	} else {
		sideMenuWidth = 250
		appsWindowPosX = 250
	}
	if !isSideMenuOpen {
		appsWindowWidth = fullWidth
		appsWindowPosX = 0
		sideMenuWidth = 0
	}
	appsWindowWidth = fullWidth - sideMenuWidth

	// Create a list of interfaces converted from struct
	for i := range appsI {
		appsI[i] = AppI(appS.appsList[i])
	}

	giu.Window("Menu Bar").
		Pos(0, 0).
		Flags(
			giu.WindowFlagsNoMove |
				giu.WindowFlagsNoResize |
				giu.WindowFlagsNoTitleBar,
		).
		Layout(
			giu.MainMenuBar().Layout(
				// Either left or right as direction
				conditionedArrowBtn(),
			),
		)

	// The app consists of 2 main windows:
	// "Main Menu" and "Apps Layout"
	if isSideMenuOpen {
		giu.Window("Main Menu").
			// Size = LHN Menu-like size and position
			Size(sideMenuWidth, fullHeight-menuBarHeight).
			Pos(0, menuBarHeight).
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
							Flags(giu.TreeNodeFlagsCollapsingHeader|giu.TreeNodeFlagsDefaultOpen).
							Layout(
								giu.Style().
									SetFont(smallFont).
									To(
										giu.Table().
											Size(sideMenuWidth-30, 45).
											Flags(
												giu.TableFlagsScrollX|
													giu.TableFlagsBorders,
											).
											Columns(
												giu.TableColumn("Type").Flags(giu.TableColumnFlagsWidthStretch),
												giu.TableColumn("Windows").Flags(giu.TableColumnFlagsWidthStretch),
												giu.TableColumn("Direction").Flags(giu.TableColumnFlagsWidthStretch),
											).
											Rows(
												giu.TableRow(
													giu.Combo("", layoutS.types[layoutS.typesIndex], layoutS.types, &layoutS.typesIndex).
														Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
														Size((sideMenuWidth/3)-18).
														OnChange(func() {
															layoutS.currType = string(layoutS.windows[layoutS.typesIndex][0])
														}),

													giu.Combo("", layoutS.windows[layoutS.windowsIndex], layoutS.windows, &layoutS.windowsIndex).
														Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
														Size((sideMenuWidth/3)-18).
														OnChange(func() {
															layoutS.currWindowsNo = int(layoutS.windows[layoutS.windowsIndex][0])
														}),

													giu.Combo("", layoutS.directions[layoutS.directionsIndex], layoutS.directions, &layoutS.directionsIndex).
														Size((sideMenuWidth/3)-18).
														Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
														OnChange(func() {
															layoutS.currDirection = string(layoutS.windows[layoutS.directionsIndex][0])
														}),
												),
											),
									),
							),

						giu.Separator(),

						// APPS Menu
						giu.TreeNode("Apps").
							Flags(giu.TreeNodeFlagsCollapsingHeader).
							Layout(
								// This is where the Main Menu items is generated
								giu.RangeBuilder("Menu", appsI, func(i int, v interface{}) giu.Widget {
									currApp := &appS.appsList[i]
									miniAppsI := make([]interface{}, len(currApp.miniApps))
									for i := range miniAppsI {
										miniAppsI[i] = MiniAppI(currApp.miniApps[i])
									}
									return giu.TreeNode(currApp.name).
										Flags(giu.TreeNodeFlagsSpanFullWidth).
										Layout(
											// This is where the Sub Menu for every Menu Item will be generated
											giu.RangeBuilder("Sub Menu", miniAppsI, func(j int, v interface{}) giu.Widget {
												currMiniApp := &currApp.miniApps[j]
												return giu.Row(
													// checkbox which has green thick when checked
													giu.Style().
														SetColor(giu.StyleColorCheckMark, color.RGBA{G: 255, A: 255}).
														To(
															giu.Checkbox("", &currMiniApp.active),
														),
													giu.Selectable(currMiniApp.name).
														OnClick(func() {
															currMiniApp.active = !currMiniApp.active
														}).
														Selected(currMiniApp.active),
												)
											}),
										)
								}),
							),
					),
			)
	}

	giu.Window("Apps").
		Size(appsWindowWidth, fullHeight-menuBarHeight).
		Pos(float32(appsWindowPosX), menuBarHeight).
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

	titleFont = giu.AddFont("Sans.ttf", 28)
	smallFont = giu.AddFont("Sans.ttf", 15)
	largeFont = giu.AddFont("Sans.ttf", 22)

	win := giu.NewMasterWindow("Universal App", 960, 640, giu.MasterWindowFlagsMaximized)
	win.Run(loop)
}
