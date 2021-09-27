package design

type AppLayout struct {
	Geometry []float32 // [width, height, posX, posY]
	//WindowsIndex, DirectionsIndex int32     // combo box
	//ComboWinLayoutsOptions        []string  // combo box
	//ComboDirectionOptions         []string  // combo box
	CurrWindowsNo     int      // current active window(s) number
	CurrDirection     string   // current orientation (vertical/horizontal/grid)
	CurrCombination   []string // current combination (type/windowsNo/orientation)
	IsDashboardView   bool     // in case there are no active windows
	IsButtonTriggered bool     // toggle if button gets clicked
	Apps
}

var (
	AppLayoutS = &AppLayout{
		Geometry: make([]float32, 4),
		//ComboWinLayoutsOptions: []string{"1", "2"},
		//ComboDirectionOptions:  []string{"Vertical"},
		//CurrDirection:          "Vertical",
		//CurrCombination:        []string{"1", "Vertical"},
		IsDashboardView:   true,
		IsButtonTriggered: false,
	}
)
