package main

import (
	"fmt"
	"github.com/AllenDang/giu"
	"image"
	"image/color"
)

type RoundFlatBtnWidget struct {
	id      string
	clicked func()
}

func RoundFlatBtn(id string, clicked func()) *RoundFlatBtnWidget {
	return &RoundFlatBtnWidget{
		id:      id,
		clicked: clicked,
	}
}

func (rfb *RoundFlatBtnWidget) Build() {
	w, h := giu.CalcTextSize(rfb.id)
	padding := float32(8.0)

	pos := giu.GetCursorPos()

	giu.InvisibleButton(rfb.id).
		Size(w+padding+50, h+padding+25).
		OnClick(rfb.clicked).
		Build()

	canvas := giu.GetCanvas()

	canvas.AddRect(
		pos.Add(image.Pt(0, 0)),
		pos.Add(image.Pt(75, 25)),
		color.RGBA{R: 200, G: 75, B: 75, A: 255},
		5,
		giu.DrawFlagsRoundCornersAll,
		2,
	)
	// canvas.AddRectFilled(pos.Add(image.Pt(220, 0)), pos.Add(image.Pt(320, 100)), color.RGBA{R: 200, G: 75, B: 75, A: 255}, 0, 0)

	canvas.AddText(image.Pt((75)/2, 25/2), color.RGBA{R: 255, G: 255, B: 255, A: 255}, rfb.id)

}

func onRoundFlatBtn() {
	fmt.Println("Round Flat Button")
}

func loop() {

	w := giu.SingleWindow()
	posX := giu.GetMousePos()
	w.Layout(
		giu.Label(fmt.Sprintf("Position: %v", posX)),
		giu.Row(RoundFlatBtn("Circle Button", onRoundFlatBtn)),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Custom Widget", 400, 300, giu.MasterWindowFlagsMaximized)
	wnd.Run(loop)
}
