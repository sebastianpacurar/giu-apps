package design

type SideMenu struct {
	Geometry []float32
	Toggled  bool
}

var (
	SideMenuS = &SideMenu{
		Geometry: make([]float32, 4),
		Toggled:  true,
	}
)
