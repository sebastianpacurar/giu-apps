package custom_widgets

import (
	"fmt"
	"github.com/AllenDang/giu"
	"image"
	"image/color"
)

type TabItemWidget struct {
	Name     string
	IsActive bool
	Clicked  func()
}

// MyTabItem - Takes the Name of the tab, a bool if the tab is active or not, and a function
func MyTabItem(Name string, IsActive bool, Clicked func()) *TabItemWidget {
	return &TabItemWidget{
		Name:     Name,
		IsActive: IsActive,
		Clicked:  Clicked,
	}
}

// Build - based on the width and height of the text. p = padding
// pos = the place where the item's position should start
// Construct 2 grey lines on the left and right side of the text
// Construct one bottom green (if current tab) or blue (if not current tab)
func (ti *TabItemWidget) Build() {

	w, h := giu.CalcTextSize(ti.Name)
	p := float32(4.0)
	pos := giu.GetCursorPos()

	canvas := giu.GetCanvas()
	canvas.AddLine(
		pos.Add(image.Pt(0, 0)),
		pos.Add(image.Pt(0, int(h+p))),
		color.RGBA{R: 46, G: 49, B: 49, A: 255},
		1,
	)

	canvas.AddLine(
		pos.Add(image.Pt(int(w+p), 0)),
		pos.Add(image.Pt(int(w+p), int(h+p))),
		color.RGBA{R: 46, G: 49, B: 49, A: 255},
		1,
	)

	// Add the text in the middle of the tab
	canvas.AddText(pos.Add(image.Pt(int(w/12+p), int(h)/10)), color.RGBA{R: 255, G: 255, B: 255, A: 255}, ti.Name)

	giu.InvisibleButton(ti.Name).Size(w+p, h+p).OnClick(ti.Clicked).Build()
	if ti.IsActive {
		canvas.AddLine(
			pos.Add(image.Pt(0, int(h+p))),
			pos.Add(image.Pt(int(w+p), int(h+p))),
			color.RGBA{G: 255, A: 255},
			1,
		)
	} else {
		canvas.AddLine(
			pos.Add(image.Pt(0, int(h+p))),
			pos.Add(image.Pt(int(w+p), int(h+p))),
			color.RGBA{B: 255, A: 255},
			1,
		)
	}
}

func loop() {

	w := giu.SingleWindow()
	posX := giu.GetMousePos()
	w.Layout(
		giu.Row(giu.Label(fmt.Sprintf("Position: %v", posX))),
		giu.Row(giu.Label(fmt.Sprintf("Position: %v", posX))),
		giu.Row(giu.Label(fmt.Sprintf("Position: %v", posX))),
		giu.Row(
			MyTabItem("test", false, func() {
				fmt.Println("1 Clicked")
			}),
			MyTabItem("test", false, func() {
				fmt.Println("2 Clicked")
			}),
			MyTabItem("test", true, func() {
				fmt.Println("3 Clicked")
			}),
			MyTabItem("test", false, func() {
				fmt.Println("4 Clicked")
			}),
		),
	)
}
