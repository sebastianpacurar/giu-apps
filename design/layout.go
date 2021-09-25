package design

type Layout struct {
	Geometry                                  []float32
	TypesIndex, WindowsIndex, DirectionsIndex int32
	ComboTypesOptions                         []string
	ComboWindowsOptions                       []string
	ComboDirectionOptions                     []string
	CurrWindowsNo                             int
	CurrType                                  string
	CurrDirection                             string
	CurrCombination                           []string
	PrevCombination                           []string
	RunningWindows                            []*Window
	IsDashboardView                           bool
}

type Window struct {
	Title      string
	Geometry   []float32
	LayoutSlot int
}

var (
	LayoutS = &Layout{
		Geometry:              make([]float32, 4),
		ComboTypesOptions:     []string{"Window", "Splitter"},
		ComboWindowsOptions:   []string{"1", "2"},
		ComboDirectionOptions: []string{"Vertical", "Horizontal", "Grid"},
		CurrType:              "Window",
		CurrDirection:         "Vertical",
		CurrWindowsNo:         0,
		PrevCombination:       []string{"Window", "1", "Vertical"},
		CurrCombination:       []string{"Window", "1", "Vertical"},
		IsDashboardView:       true,

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
