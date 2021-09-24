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
	fullWidth, fullHeight float32
	titleFont, smallFont  *giu.FontInfo
	defaultFlags          = giu.WindowFlagsNoMove | giu.WindowFlagsNoResize | giu.WindowFlagsNoTitleBar
)

// appsS - The struct of the Menu
// appsI - The appsList[] as an Interface (to be used with RangeBuilder() as values param)
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
		geometry:              make([]float32, 4),
		comboTypesOptions:     []string{"Window", "Splitter"},
		comboWindowsOptions:   []string{"1", "2"},
		comboDirectionOptions: []string{"Vertical", "Horizontal", "Grid"},
		currType:              "Window",
		currDirection:         "Vertical",
		currWindowsNo:         0,
		prevCombination:       []string{"Window", "1", "Vertical"},
		currCombination:       []string{"Window", "1", "Vertical"},
		isDashboardView:       true,

		// TODO: currently on hold
		//runningWindows: []*Window{
		//	{
		//		// first element is the initial setup
		//		title:      "Dashboard",
		//		geometry:   make([]float32, 4),
		//		layoutSlot: 1,
		//	},
		//},
	}

	sideMenuS = &SideMenu{
		geometry: make([]float32, 4),
		toggled:  true,
	}

	// TODO: Leave menuBarS hardcoded for now
	menuBarS = &MenuBar{
		geometry: []float32{
			fullWidth,
			23,
			0,
			0,
		},
	}
)

type SideMenu struct {
	geometry []float32
	toggled  bool
}

type MenuBar struct {
	geometry []float32
}

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

type Window struct {
	title      string
	geometry   []float32
	layoutSlot int
}

type Layout struct {
	geometry                                  []float32
	typesIndex, windowsIndex, directionsIndex int32
	comboTypesOptions                         []string
	comboWindowsOptions                       []string
	comboDirectionOptions                     []string
	currWindowsNo                             int
	currType                                  string
	currDirection                             string
	currCombination                           []string
	prevCombination                           []string
	runningWindows                            []*Window
	isDashboardView                           bool
}

// conditionedArrowBtn - is used to swap directions of the arrow after each click
func conditionedArrowBtn() giu.Widget {
	var arrowBtn *giu.ArrowButtonWidget
	if sideMenuS.toggled {
		arrowBtn = giu.ArrowButton("close menu", giu.DirectionLeft).OnClick(func() {
			sideMenuS.toggled = false
		})
	} else {
		arrowBtn = giu.ArrowButton("open menu", giu.DirectionRight).OnClick(func() {
			sideMenuS.toggled = true
		})
	}
	return arrowBtn
}

func loop() {
	size := giu.Context.GetPlatform().DisplaySize()
	fullWidth = size[0]
	fullHeight = size[1]
	sideMenuS.geometry[0] = fullWidth / 4
	sideMenuS.geometry[1] = fullHeight

	// For sizes bigger than 990px use responsive width
	// If the Main Menu is closed, then stretch Apps Window to full width
	if int(fullWidth) <= 990 {
		sideMenuS.geometry[0] = 250
	}

	if !sideMenuS.toggled {
		sideMenuS.geometry[0] = 0
	}

	layoutS.geometry = []float32{
		fullWidth - sideMenuS.geometry[0],
		fullHeight - menuBarS.geometry[1],
		sideMenuS.geometry[0],
		menuBarS.geometry[1],
	}

	// Create a list of interfaces converted from struct
	for i := range appsI {
		appsI[i] = AppI(appsS.appsList[i])
	}

	// TODO: Currently on hold
	// Toggle Dashboard on start and when there are no apps selected
	//if layoutS.isDashboardView {
	//	giu.Window("Dashboard").
	//		Size(layoutS.geometry[0], layoutS.geometry[1]).
	//		Pos(layoutS.geometry[2], menuBarS.geometry[1]).
	//		Flags(defaultFlags).
	//		Layout(
	//			giu.Label("Dashboard"),
	//		)
	//} else {
	//	for i := range layoutS.runningWindows {
	//		currWin := layoutS.runningWindows[i]
	//		giu.Window(currWin.title).
	//			Size(currWin.geometry[0], currWin.geometry[1]).
	//			Pos(currWin.geometry[2], currWin.geometry[3]).
	//			Flags(defaultFlags).
	//			Layout(
	//				giu.Label(currWin.title),
	//			)
	//	}
	//}

	giu.Window("Dashboard").
		Size(layoutS.geometry[0], layoutS.geometry[1]).
		Pos(layoutS.geometry[2], menuBarS.geometry[1]).
		Flags(defaultFlags).
		Layout(
			giu.Label("Dashboard"),
		)

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
	if sideMenuS.toggled {
		giu.Window("Main Menu").
			// Size = LHN Menu-like size and position
			Size(sideMenuS.geometry[0], layoutS.geometry[1]).
			Pos(0, menuBarS.geometry[1]).
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
														giu.Combo("", layoutS.comboTypesOptions[layoutS.typesIndex], layoutS.comboTypesOptions, &layoutS.typesIndex).
															Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
															Size((sideMenuS.geometry[0]/3)-18).
															OnChange(func() {
																layoutS.currType = layoutS.comboTypesOptions[layoutS.typesIndex]
															}),

														giu.Combo("", layoutS.comboWindowsOptions[layoutS.windowsIndex], layoutS.comboWindowsOptions, &layoutS.windowsIndex).
															Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
															Size((sideMenuS.geometry[0]/3)-18).
															OnChange(func() {
																layoutS.currWindowsNo = int(layoutS.windowsIndex) + 1
															}),

														giu.Combo("", layoutS.comboDirectionOptions[layoutS.directionsIndex], layoutS.comboDirectionOptions, &layoutS.directionsIndex).
															Size((sideMenuS.geometry[0]/3)-18).
															Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
															OnChange(func() {
																layoutS.currDirection = layoutS.comboDirectionOptions[layoutS.directionsIndex]
															}),
													),
												),
										),
									// The Button below triggers buildAppsLayout function,
									// And will appear as Disabled if the combination maps are the same
									giu.Button("waiting fix").
										Size(giu.Auto, 25).
										// TODO: currently on hold
										//OnClick(buildAppsLayout).
										Disabled(true),
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
}

func isBuildLayoutBtnDisabled() bool {
	res := true
	for i := range layoutS.currCombination {
		if layoutS.currCombination[i] == layoutS.prevCombination[i] {
			res = false
			break
		}
	}
	return res
}

// TODO: currently on hold
func buildAppsLayout() {
	if layoutS.currCombination != nil {
		for i := 0; i < 3; i++ {
			layoutS.prevCombination[i] = layoutS.currCombination[i]
		}
	}

	layoutS.currCombination = []string{
		layoutS.currType,
		strconv.Itoa(layoutS.currWindowsNo),
		layoutS.currDirection,
	}

	if layoutS.currWindowsNo > 0 {
		layoutS.isDashboardView = false
	}

	layoutS.runningWindows = make([]*Window, layoutS.currWindowsNo)
	for i := range layoutS.runningWindows {
		layoutS.runningWindows[i] = &Window{}
	}

	switch layoutType := layoutS.currType; layoutType {
	case "Window":
		switch count := layoutS.currWindowsNo; count {
		case 1:
			for i := 0; i < count; i++ {
				layoutS.runningWindows[i].geometry = []float32{
					layoutS.geometry[0],
					layoutS.geometry[1],
					layoutS.geometry[2],
					layoutS.geometry[3],
				}
				layoutS.runningWindows[i].layoutSlot = 1
				layoutS.runningWindows[i].title = "Dashboard"
			}

		case 2:
			for i := range layoutS.runningWindows {
				layoutS.runningWindows[i] = &Window{}
			}
			switch layoutDirection := layoutS.currDirection; layoutDirection {

			case "Vertical":
				for i := range layoutS.runningWindows {
					if i%2 == 0 {
						layoutS.runningWindows[i].title = fmt.Sprintf("Window %d", i+1)
						layoutS.runningWindows[i].geometry = []float32{
							layoutS.geometry[0] / 2,
							layoutS.geometry[1],
							sideMenuS.geometry[0],
							menuBarS.geometry[1],
						}
						layoutS.runningWindows[i].layoutSlot = i + 1
					} else {
						layoutS.runningWindows[i].title = fmt.Sprintf("Window %d", i+1)
						layoutS.runningWindows[i].geometry = []float32{
							layoutS.runningWindows[0].geometry[0],
							layoutS.runningWindows[0].geometry[1],
							layoutS.geometry[0] - layoutS.runningWindows[0].geometry[0],
							menuBarS.geometry[1],
						}
						layoutS.runningWindows[i].layoutSlot = i + 1
					}
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
