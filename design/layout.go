package design

type AppLayout struct {
	Geometry                      []float32 // [width, height, posX, posY]
	WindowsIndex, DirectionsIndex int32     // combo box
	ComboWinLayoutsOptions        []string  // combo box
	ComboDirectionOptions         []string  // combo box
	CurrWindowsNo                 int       // current active window(s) number
	CurrDirection                 string    // current orientation (vertical/horizontal/grid)
	CurrCombination               []string  // current combination (type/windowsNo/orientation)
	PrevCombination               []string  // previous combination (types/windowsNo/orientation)
	ActiveWindows                 []*Window // current batch of active windows
	IsDashboardView               bool      // in case there are no active windows
	IsButtonTriggered             bool      // toggle if button gets clicked
	IsButtonDisabled              bool      // disable button if currCombo != prevComo
}

type Window struct {
	Title       string
	WindowIndex int
	layout      WindowLayout
}

type WindowLayout struct {
	activeApps *Apps
	Screens    int
}

var (
	AppLayoutS = &AppLayout{
		Geometry:               make([]float32, 4),
		ComboWinLayoutsOptions: []string{"1", "2"},
		ComboDirectionOptions:  []string{"Vertical"},
		CurrDirection:          "Vertical",
		CurrWindowsNo:          1,
		CurrCombination:        []string{"1", "Vertical"},
		IsDashboardView:        true,
		IsButtonTriggered:      false,

		// TODO: fixing in progress
		ActiveWindows: []*Window{
			{
				// first element is the initial setup
				Title:       "Dashboard",
				WindowIndex: 0,
				layout: WindowLayout{
					Screens:    1,
					activeApps: &Apps{},
				},
			},
		},
	}
)
