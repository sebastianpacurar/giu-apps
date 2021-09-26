package main

import (
	"fmt"
	"github.com/AllenDang/giu"
	"image/color"
	"imgui-based-app/design"
	"strconv"
)

// Data related to the App Layout handling
var (
	fullWidth, fullHeight float32
	titleFont, smallFont  *giu.FontInfo
	defaultFlags          = giu.WindowFlagsNoMove | giu.WindowFlagsNoResize | giu.WindowFlagsNoTitleBar
)

// conditionedArrowBtn - is used to swap directions of the arrow after each click
func conditionedArrowBtn() giu.Widget {
	var arrowBtn *giu.ArrowButtonWidget
	if design.SideMenuS.Toggled {
		arrowBtn = giu.ArrowButton("close menu", giu.DirectionLeft).OnClick(func() {
			design.SideMenuS.Toggled = false
		})
	} else {
		arrowBtn = giu.ArrowButton("open menu", giu.DirectionRight).OnClick(func() {
			design.SideMenuS.Toggled = true
		})
	}
	return arrowBtn
}

func loop() {
	size := giu.Context.GetPlatform().DisplaySize()
	fullWidth = size[0]
	fullHeight = size[1]

	// Geometry = [width, height, positionX, positionY]
	design.TopBarS.Geometry = []float32{
		fullWidth,
		23,
		0,
		0,
	}
	design.BottomBarS.Geometry = []float32{
		fullWidth,
		40,
		0,
		fullHeight - 40,
	}
	design.SideBarS.Geometry = []float32{
		40,
		fullHeight - (design.TopBarS.Geometry[1] + design.BottomBarS.Geometry[1]),
		fullWidth - 40,
		23,
	}
	design.SideMenuS.Geometry = []float32{
		fullWidth / 4,
		fullHeight - (design.TopBarS.Geometry[1] + design.BottomBarS.Geometry[1]),
		0,
		23,
	}

	// For sizes bigger than 990px use responsive width
	// If the Main Menu is closed, then stretch Apps Window to full width
	if int(fullWidth) <= 990 {
		design.SideMenuS.Geometry[0] = 250
	} else if int(fullWidth) >= 1200 {
		design.SideMenuS.Geometry[0] = 300
	}

	if !design.SideMenuS.Toggled {
		design.SideMenuS.Geometry[0] = 0
	}

	//TODO: very broken and easy to fix!
	for i := range design.LayoutS.CurrCombination {
		if design.LayoutS.CurrCombination[i] != design.LayoutS.PrevCombination[i] {
			design.LayoutS.IsButtonDisabled = false
		}
	}

	design.LayoutS.Geometry = []float32{
		fullWidth - (design.SideMenuS.Geometry[0] + design.SideBarS.Geometry[0]),
		fullHeight - (design.TopBarS.Geometry[1] + design.BottomBarS.Geometry[1]),
		design.SideMenuS.Geometry[0],
		design.TopBarS.Geometry[1],
	}

	if design.LayoutS.IsButtonTriggered {
		buildAppsLayout()
	}

	// Toggle Dashboard on start and when there are no apps selected
	if design.LayoutS.IsDashboardView {
		giu.Window("Dashboard").
			Size(design.LayoutS.Geometry[0], design.LayoutS.Geometry[1]).
			Pos(design.LayoutS.Geometry[2], design.TopBarS.Geometry[1]).
			Flags(defaultFlags).
			Layout(
				giu.Label("Dashboard"),
			)
	} else {
		for i := range design.LayoutS.RunningWindows {
			currWin := design.LayoutS.RunningWindows[i]
			giu.Window(currWin.Title).
				//TODO this should be the actual size!!
				Size(currWin.Geometry[0], currWin.Geometry[1]).
				Pos(currWin.Geometry[2], currWin.Geometry[3]).
				Flags(defaultFlags).
				Layout(
					giu.Label(currWin.Title),
				)
		}
	}

	giu.Window("Bottom Bar").
		Size(design.BottomBarS.Geometry[0], design.BottomBarS.Geometry[1]).
		Pos(design.BottomBarS.Geometry[2], design.BottomBarS.Geometry[3]).
		Flags(defaultFlags).
		Layout(
			giu.Row(
				giu.Button("test").Size(100, 20),
				giu.Button("test").Size(100, 20),
				giu.Dummy(design.BottomBarS.Geometry[0]-270, 0),
				giu.ImageWithFile("icons/home_white_icon_48dp.png").
					Size(24, 24),
			),
		)

	giu.Window("Side Bar").
		Size(design.SideBarS.Geometry[0], design.SideBarS.Geometry[1]).
		Pos(design.SideBarS.Geometry[2], design.SideBarS.Geometry[3]).
		Flags(defaultFlags).
		Layout(
			giu.Column(
				giu.ImageWithFile("icons/home_white_icon_48dp.png").
					Size(24, 24),
			),
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

	// Create a list of interfaces converted from struct
	for i := range design.AppsI {
		design.AppsI[i] = design.AppI(design.AppsS.AppsList[i])
	}

	// The app consists of 2 main windows:
	// "Main Menu" and "Apps Layout"
	if design.SideMenuS.Toggled {
		giu.Window("Main Menu").
			// Size = LHN Menu-like size and position
			Size(design.SideMenuS.Geometry[0], design.LayoutS.Geometry[1]).
			Pos(0, design.TopBarS.Geometry[1]).
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
														giu.Combo("", design.LayoutS.ComboTypesOptions[design.LayoutS.TypesIndex], design.LayoutS.ComboTypesOptions, &design.LayoutS.TypesIndex).
															Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
															Size((design.SideMenuS.Geometry[0]/3)-18).
															OnChange(func() {
																design.LayoutS.CurrType = design.LayoutS.ComboTypesOptions[design.LayoutS.TypesIndex]
															}),

														giu.Combo("", design.LayoutS.ComboWindowsOptions[design.LayoutS.WindowsIndex], design.LayoutS.ComboWindowsOptions, &design.LayoutS.WindowsIndex).
															Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
															Size((design.SideMenuS.Geometry[0]/3)-18).
															OnChange(func() {
																design.LayoutS.CurrWindowsNo = int(design.LayoutS.WindowsIndex) + 1
															}),

														giu.Combo("", design.LayoutS.ComboDirectionOptions[design.LayoutS.DirectionsIndex], design.LayoutS.ComboDirectionOptions, &design.LayoutS.DirectionsIndex).
															Size((design.SideMenuS.Geometry[0]/3)-18).
															Flags(giu.ComboFlagHeightSmall|giu.ComboFlagNoArrowButton).
															OnChange(func() {
																design.LayoutS.CurrDirection = design.LayoutS.ComboDirectionOptions[design.LayoutS.DirectionsIndex]
															}),
													),
												),
										),
									// The Button below triggers buildAppsLayout function,
									// And will appear as Disabled if the combination maps are the same
									giu.Button("Build Layout").
										Size(giu.Auto, 25).
										// TODO: currently on hold
										OnClick(func() {
											design.LayoutS.IsButtonTriggered = !design.LayoutS.IsButtonTriggered
											design.LayoutS.IsButtonDisabled = true
										}).
										Disabled(design.LayoutS.IsButtonDisabled),
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
								giu.RangeBuilder("Menu", design.AppsI, func(i int, v interface{}) giu.Widget {
									currApp := &design.AppsS.AppsList[i]
									miniAppsI := make([]interface{}, len(currApp.MiniApps))
									for i := range miniAppsI {
										miniAppsI[i] = design.MiniAppI(currApp.MiniApps[i])
									}
									return giu.TreeNode(currApp.Name).
										Flags(giu.TreeNodeFlagsSpanFullWidth).
										Layout(
											// This is where the Sub Menu for every Menu Item will be generated
											giu.RangeBuilder("Sub Menu", miniAppsI, func(j int, v interface{}) giu.Widget {
												currMiniApp := &currApp.MiniApps[j]
												return giu.Row(
													// checkbox which has green thick when checked
													giu.Style().
														SetColor(giu.StyleColorCheckMark, color.RGBA{G: 255, A: 255}).
														To(
															giu.Checkbox("", &currMiniApp.Active),
														),
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
}

//TODO: very broken and easy to fix!
func buildAppsLayout() {
	if design.LayoutS.CurrCombination != nil {
		for i := 0; i < 3; i++ {
			design.LayoutS.PrevCombination[i] = design.LayoutS.CurrCombination[i]
		}
	}

	design.LayoutS.CurrCombination = []string{
		design.LayoutS.CurrType,
		strconv.Itoa(design.LayoutS.CurrWindowsNo),
		design.LayoutS.CurrDirection,
	}

	if design.LayoutS.CurrWindowsNo > 0 {
		design.LayoutS.IsDashboardView = false
	}

	design.LayoutS.RunningWindows = make([]*design.Window, design.LayoutS.CurrWindowsNo)
	for i := range design.LayoutS.RunningWindows {
		design.LayoutS.RunningWindows[i] = &design.Window{}
	}

	switch layoutType := design.LayoutS.CurrType; layoutType {
	case "Window":
		switch count := design.LayoutS.CurrWindowsNo; count {
		case 1:
			for i := 0; i < count; i++ {
				design.LayoutS.RunningWindows[i].Geometry = []float32{
					design.LayoutS.Geometry[0],
					design.LayoutS.Geometry[1],
					design.LayoutS.Geometry[2],
					design.LayoutS.Geometry[3],
				}
				design.LayoutS.RunningWindows[i].LayoutSlot = 1
				design.LayoutS.RunningWindows[i].Title = "Dashboard"
			}

		case 2:
			for i := range design.LayoutS.RunningWindows {
				design.LayoutS.RunningWindows[i] = &design.Window{}
			}
			switch layoutDirection := design.LayoutS.CurrDirection; layoutDirection {

			case "Vertical":
				for i := range design.LayoutS.RunningWindows {
					if i%2 == 0 {
						design.LayoutS.RunningWindows[i].Title = fmt.Sprintf("Window %d", i+1)
						design.LayoutS.RunningWindows[i].Geometry = []float32{
							(fullWidth - (design.SideMenuS.Geometry[0] + design.SideBarS.Geometry[0])) / 2,
							design.LayoutS.Geometry[1],
							design.SideMenuS.Geometry[0],
							design.TopBarS.Geometry[1],
						}
						design.LayoutS.RunningWindows[i].LayoutSlot = i + 1
					} else {
						design.LayoutS.RunningWindows[i].Title = fmt.Sprintf("Window %d", i+1)
						design.LayoutS.RunningWindows[i].Geometry = []float32{
							design.LayoutS.RunningWindows[0].Geometry[0],
							design.LayoutS.RunningWindows[0].Geometry[1],
							design.LayoutS.Geometry[0] - design.LayoutS.RunningWindows[0].Geometry[0],
							design.TopBarS.Geometry[1],
						}
						design.LayoutS.RunningWindows[i].LayoutSlot = i + 1
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
