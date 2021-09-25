package design

type TopBar struct {
	Geometry []float32
}

type SideBar struct {
	Geometry []float32
}

type BottomBar struct {
	Geometry []float32
}

var (
	TopBarS = &TopBar{
		Geometry: make([]float32, 4),
	}

	SideBarS = &SideBar{
		Geometry: make([]float32, 4),
	}

	BottomBarS = &BottomBar{
		Geometry: make([]float32, 4),
	}
)
