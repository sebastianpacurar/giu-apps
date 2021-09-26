package design

type Layout struct {
	Geometry                                  []float32 // [width, height, posX, posY]
	TypesIndex, WindowsIndex, DirectionsIndex int32     // combo box
	ComboTypesOptions                         []string  // combo box
	ComboWindowsOptions                       []string  // combo box
	ComboDirectionOptions                     []string  // combo box
	CurrWindowsNo                             int       // current active window(s) number
	CurrType                                  string    // current layout type (window/splitter)
	CurrDirection                             string    // current orientation (vertical/horizontal/grid)
	CurrCombination                           []string  // current combination (type/windowsNo/orientation)
	PrevCombination                           []string  // previous combination (types/windowsNo/orientation)
	RunningWindows                            []*Window // current batch of active windows
	IsDashboardView                           bool      // in case there are no active windows
	IsButtonTriggered                         bool      // toggle if button gets clicked
	IsButtonDisabled                          bool      // disable button if currCombo != prevComo
}

type Window struct {
	Title      string
	Geometry   []float32
	LayoutSlot int
}

var (
	LayoutS = &Layout{
		Geometry: make([]float32, 4),
		//ComboTypesOptions:     []string{"Window", "Splitter"},
		ComboTypesOptions:   []string{"Window"},
		ComboWindowsOptions: []string{"1", "2"},
		//ComboDirectionOptions: []string{"Vertical", "Horizontal", "Grid"},
		ComboDirectionOptions: []string{"Vertical"},
		CurrType:              "Window",
		CurrDirection:         "Vertical",
		CurrWindowsNo:         0,
		PrevCombination:       make([]string, 3),
		CurrCombination:       []string{"Window", "1", "Vertical"},
		IsDashboardView:       true,
		IsButtonTriggered:     false,
		IsButtonDisabled:      false,

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
)
