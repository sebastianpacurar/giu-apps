package main

import (
	"fmt"
	"github.com/AllenDang/giu"
	"image/color"
	"imgui-based-app/custom-widgets"
	"imgui-based-app/design"
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
		300,
		fullHeight - (design.TopBarS.Geometry[1] + design.BottomBarS.Geometry[1]),
		0,
		23,
	}

	if !design.SideMenuS.Toggled {
		design.SideMenuS.Geometry[0] = 0
	}

	design.AppLayoutS.Geometry = []float32{
		fullWidth - (design.SideMenuS.Geometry[0] + design.SideBarS.Geometry[0]),
		fullHeight - (design.TopBarS.Geometry[1] + design.BottomBarS.Geometry[1]),
		design.SideMenuS.Geometry[0],
		design.TopBarS.Geometry[1],
	}

	giu.Window("Apps Window").
		Size(design.AppLayoutS.Geometry[0], design.AppLayoutS.Geometry[1]).
		Pos(design.AppLayoutS.Geometry[2], design.AppLayoutS.Geometry[3]).
		Flags(defaultFlags | giu.WindowFlagsMenuBar).
		Layout(
			giu.RangeBuilder("App Screen", design.AppsI, func(i int, v interface{}) giu.Widget {
				currApp := &design.AppsS.AppsList[i]
				miniAppsTabs := make([]string, len(currApp.MiniApps))

				for i := range miniAppsTabs {
					miniAppsTabs[i] = currApp.MiniApps[i].Name
				}

				miniAppsI := make([]interface{}, len(currApp.MiniApps))
				for i := range miniAppsI {
					miniAppsI[i] = design.MiniAppI(currApp.MiniApps[i])
				}
				return giu.Condition(
					currApp.Active,
					giu.Layout{
						giu.MenuBar().
							Layout(
								giu.RangeBuilder("Sub Menu", miniAppsI, func(j int, v interface{}) giu.Widget {
									return giu.Condition(
										currApp.MiniApps[j].Active,
										giu.Layout{
											custom_widgets.MyTabItem(currApp.MiniApps[j].Name, currApp.MiniApps[j].Current, func() {
												fmt.Println(currApp.MiniApps[j].Name)
											}),
										},
										nil,
									)
								}),
							),
						giu.RangeBuilder("Content of Current Mini App", miniAppsI, func(j int, v interface{}) giu.Widget {
							return giu.Condition(
								currApp.MiniApps[j].Current,
								giu.Layout{
									giu.Child().Layout(
										giu.Label(currApp.MiniApps[j].Name),
									),
								}, nil,
							)
						}),
					}, nil,
				)
			}),
		)

	// Toggle Dashboard on start and when there are no apps selected
	giu.Window("Main").
		Size(design.AppLayoutS.Geometry[0], design.AppLayoutS.Geometry[1]).
		Pos(design.AppLayoutS.Geometry[2], design.TopBarS.Geometry[1]).
		Flags(defaultFlags).
		Layout()

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
	// "Main Menu" and "Apps AppLayout"
	if design.SideMenuS.Toggled {
		giu.Window("Main Menu").
			// Size = LHN Menu-like size and position
			Size(design.SideMenuS.Geometry[0], design.AppLayoutS.Geometry[1]).
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

						// APPS Menu
						giu.Table().
							Flags(giu.TableFlagsBorders).
							Columns(
								giu.TableColumn("Categories").Flags(giu.TableColumnFlagsWidthStretch),
								giu.TableColumn("Apps").Flags(giu.TableColumnFlagsWidthStretch),
							).Rows(
							buildAppsRows()...,
						),
					),
			)
	}
}

func buildAppsRows() []*giu.TableRowWidget {
	rows := make([]*giu.TableRowWidget, len(design.AppsS.AppsList))

	for i := range rows {
		miniAppsI := make([]interface{}, len(design.AppsS.AppsList[i].MiniApps))
		for j := range miniAppsI {
			miniAppsI[j] = design.MiniAppI(design.AppsS.AppsList[i].MiniApps[j])
		}
		rows[i] = giu.TableRow(
			giu.Label(design.AppsS.AppsList[i].Name),
			giu.RangeBuilder("Sub Menu", miniAppsI, func(j int, v interface{}) giu.Widget {
				return giu.Row(
					// checkbox which has green thick when checked
					giu.Condition(
						design.AppsS.AppsList[i].MiniApps[j].Active,
						giu.Layout{
							giu.Style().
								SetColor(giu.StyleColorText, color.RGBA{G: 255, A: 255}).
								To(
									giu.Selectable(design.AppsS.AppsList[i].MiniApps[j].Name).
										OnClick(func() {
											design.AppsS.AppsList[i].MiniApps[j].Active = !design.AppsS.AppsList[i].MiniApps[j].Active
										}).Selected(design.AppsS.AppsList[i].MiniApps[j].Active),
								),
						}, giu.Layout{
							giu.Selectable(design.AppsS.AppsList[i].MiniApps[j].Name).
								OnClick(func() {
									design.AppsS.AppsList[i].MiniApps[j].Active = !design.AppsS.AppsList[i].MiniApps[j].Active
								}).Selected(design.AppsS.AppsList[i].MiniApps[j].Active),
						},
					),
				)
			}),
		)
	}
	return rows
}

func main() {
	// Change the default font to sans and of 18 pixels height
	giu.SetDefaultFont("Sans.ttf", 18)

	titleFont = giu.AddFont("Sans.ttf", 28)
	smallFont = giu.AddFont("Sans.ttf", 15)

	win := giu.NewMasterWindow("Universal App", 960, 640, giu.MasterWindowFlagsMaximized)
	win.Run(loop)
}
