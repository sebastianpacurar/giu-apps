package main

import (
	"fmt"
	"github.com/AllenDang/giu"
	"image/color"
	"strconv"
)

type AppI interface{}
type MiniAppI interface{}

// Data related to the App Layout handling
var (
	fullWidth, fullHeight                            float32
	sideMenuWidth, appsWindowWidth, appsWindowHeight float32
	menuBarHeight                                    = float32(23)
	appsWindowPosX                                   int
	isSideMenuOpen                                   = true
	titleFont, smallFont                             *giu.FontInfo
	defaultFlags                                     = giu.WindowFlagsNoMove | giu.WindowFlagsNoResize | giu.WindowFlagsNoTitleBar

	// windowsGeom - contains size and position of all running apps windows
	// currWinGeom - this is the window under iteration, which will eventually
	windowsGeom = map[string]map[string][]float32{}
	currWinGeom = map[string][]float32{}
)

// AppS - The struct of the Menu
// AppsI - The appsList[] as an Interface (to be used with RangeBuilder() as values param)
var (
	appsI = make([]interface{}, len(appsS.appsList))
	appsS = &Apps{
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
		types:         []string{"Window", "Splitter"},
		windows:       []string{"1", "2"},
		directions:    []string{"Vertical", "Horizontal", "Grid"},
		currType:      "Window",
		currDirection: "Grid",
		currWindowsNo: 0,
		prevCombination: map[string]string{
			"type":      "Window",
			"count":     "1",
			"direction": "Vertical",
		},
		currCombination: map[string]string{
			"type":      "Window",
			"count":     "1",
			"direction": "Vertical",
		},
		isDashboardView: true,
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
	currCombination                           map[string]string
	prevCombination                           map[string]string
	runningWindows                            []*giu.WindowWidget
	isDashboardView                           bool
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
	appsWindowHeight = fullHeight - menuBarHeight
	// Create a list of interfaces converted from struct
	for i := range appsI {
		appsI[i] = AppI(appsS.appsList[i])
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
			Size(sideMenuWidth, appsWindowHeight).
			Pos(0, menuBarHeight).
			Flags(defaultFlags).
			Layout(
				giu.Child().
					Border(true).
					Layout(
						// This is the Title of the20 Main Menu. set Text Color to Cyan rgba(0, 255, 255, 255)
						// Also, use the titleFont of 28px sans
						giu.Row(
							giu.Style().
								SetColor(giu.StyleColorText, color.RGBA{G: 255, B: 255, A: 255}).
								SetFont(titleFont).
								To(
									giu.Label("Main Menu").Wrapped(true),
								),
						),

						giu.Style().
							SetColor(giu.StyleColorSeparator, color.RGBA{G: 255, B: 255, A: 255}).
							To(
								giu.Separator(),
							),

						// LAYOUT Menu
						giu.TreeNode("Layout").
							Flags(giu.TreeNodeFlagsCollapsingHeader|giu.TreeNodeFlagsDefaultOpen).
							Layout(
								giu.Column(
									giu.Style().
										SetFont(smallFont).
										To(
											giu.Table().
												Size(giu.Auto, 45).
												Flags(
													giu.TableFlagsScrollX|
														giu.TableFlagsBorders,
												).
												Columns(
													giu.TableColumn("Type").Flags(giu.TableColumnFlagsWidthStretch),
													giu.TableColumn("Windows").Flags(giu.TableColumnFlagsWidthStretch),
													giu.TableColumn("Orientation").Flags(giu.TableColumnFlagsWidthStretch),
												).
												Rows(

													// TODO: Implement Iterative way to avoid redundancy
													giu.TableRow(
														giu.Combo("", layoutS.types[layoutS.typesIndex], layoutS.types, &layoutS.typesIndex).
															Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
															Size((sideMenuWidth/3)-18).
															OnChange(func() {
																layoutS.currType = layoutS.types[layoutS.typesIndex]
															}),

														giu.Combo("", layoutS.windows[layoutS.windowsIndex], layoutS.windows, &layoutS.windowsIndex).
															Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
															Size((sideMenuWidth/3)-18).
															OnChange(func() {
																layoutS.currWindowsNo = int(layoutS.windowsIndex)
															}),

														giu.Combo("", layoutS.directions[layoutS.directionsIndex], layoutS.directions, &layoutS.directionsIndex).
															Size((sideMenuWidth/3)-18).
															Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
															OnChange(func() {
																layoutS.currDirection = layoutS.directions[layoutS.directionsIndex]
															}),
													),
												),
										),
									// The Button below triggers buildLayout function,
									// And will appear as Disabled if the combination maps are the same
									giu.Button("Build Layout").
										Size(giu.Auto, 25).
										OnClick(buildLayout).
										Disabled(isBuildLayoutBtnDisabled()),
								),
							),

						giu.Style().
							SetColor(giu.StyleColorSeparator, color.RGBA{G: 255, B: 255, A: 255}).
							To(
								giu.Separator(),
							),

						// APPS Menu
						giu.TreeNode("Apps").
							Flags(giu.TreeNodeFlagsCollapsingHeader).
							Layout(
								// This is where the Main Menu items is generated
								giu.RangeBuilder("Menu", appsI, func(i int, v interface{}) giu.Widget {
									currApp := &appsS.appsList[i]
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

	dashboard := giu.Window("Dashboard").
		Size(appsWindowWidth, appsWindowHeight).
		Pos(float32(appsWindowPosX), menuBarHeight).
		Flags(defaultFlags).IsOpen(&layoutS.isDashboardView)

	// Toggle Dashboard only when there i
	if layoutS.currWindowsNo == 0 {
		layoutS.isDashboardView = true
	} else {
		layoutS.isDashboardView = false
	}
	dashboard.IsOpen(&layoutS.isDashboardView).BringToFront()
}

func isBuildLayoutBtnDisabled() bool {
	res := true
	for k, _ := range layoutS.currCombination {
		if layoutS.currCombination[k] == layoutS.prevCombination[k] {
			res = false
			break
		}
	}
	return res
}

func buildLayout() {
	currWinGeom = map[string][]float32{
		"size": {fullWidth - sideMenuWidth, fullHeight - menuBarHeight},
		"pos":  {sideMenuWidth, menuBarHeight},
	}

	windowsGeom = map[string]map[string][]float32{
		"w1": make(map[string][]float32, 2),
		"w2": make(map[string][]float32, 2),
		"w3": make(map[string][]float32, 2),
		"w4": make(map[string][]float32, 2),
	}

	for k, _ := range windowsGeom {
		windowsGeom[k] = map[string][]float32{
			"size": {fullWidth - sideMenuWidth, fullHeight - menuBarHeight},
			"pos":  {sideMenuWidth, menuBarHeight},
		}
	}

	if layoutS.currCombination != nil {
		layoutS.prevCombination = layoutS.currCombination
	}

	// TODO: This needs to be placed before all the Layout Combo Boxes
	layoutS.currCombination = map[string]string{
		"type":      layoutS.currType,
		"count":     strconv.Itoa(layoutS.currWindowsNo),
		"direction": layoutS.currDirection,
	}

	switch layoutType := layoutS.currType; layoutType {
	case "Window":
		switch layoutDirection := layoutS.currDirection; layoutDirection {
		case "Vertical":
			if !isSideMenuOpen {
				currWinGeom["size"][0] = fullWidth
				currWinGeom["pos"][0] = 0
			}
			layoutS.runningWindows = make([]*giu.WindowWidget, layoutS.currWindowsNo)

			switch count := layoutS.currWindowsNo; count {
			case 1:
				giu.Window("w1").
					Size(currWinGeom["size"][0], currWinGeom["size"][1]).
					Pos(currWinGeom["pos"][0], currWinGeom["pos"][1])
			case 2:
				windowsGeom["w1"]["size"] = []float32{currWinGeom["size"][0] / 2, currWinGeom["size"][1]}
				windowsGeom["w1"]["pos"] = []float32{currWinGeom["pos"][0], currWinGeom["pos"][1]}

				windowsGeom["w2"]["size"] = []float32{currWinGeom["size"][0] / 2, currWinGeom["size"][1]}
				windowsGeom["w2"]["pos"] = []float32{sideMenuWidth, currWinGeom["pos"][1]}

				for i := 1; i <= count; i++ {
					windowID := fmt.Sprintf("w%d", i)
					giu.Window(windowID).
						Size(windowsGeom[windowID]["size"][0], windowsGeom[windowID]["size"][1]).
						Pos(windowsGeom[windowID]["pos"][0], windowsGeom[windowID]["pos"][1]).
						Flags(defaultFlags)
				}
			}
		}
	}
}

func main() {
	// Change the default font to sans and of 18 pixels height
	giu.SetDefaultFont("Sans.ttf", 18)

	titleFont = giu.AddFont("Sans.ttf", 28)
	smallFont = giu.AddFont("Sans.ttf", 15)

	win := giu.NewMasterWindow("Universal App", 960, 640, giu.MasterWindowFlagsMaximized)
	win.Run(loop)
}
