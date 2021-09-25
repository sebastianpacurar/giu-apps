package design

type AppI interface{}
type MiniAppI interface{}

type Apps struct {
	AppsList []App
}

type App struct {
	Name     string
	Active   bool
	MiniApps []MiniApp
}

type MiniApp struct {
	Name   string
	Active bool
}

// appsS - The struct of the Menu
// appsI - The AppsList[] as an Interface (to be used with RangeBuilder() as values param)
var (
	AppsI = make([]interface{}, len(AppsS.AppsList))
	AppsS = &Apps{
		AppsList: []App{
			{
				Name:   "Geography",
				Active: false,
				MiniApps: []MiniApp{
					{
						Name:   "All Countries",
						Active: false,
					},
					{
						Name:   "Map",
						Active: false,
					},
				},
			},

			{
				Name:   "Dictionary",
				Active: false,
				MiniApps: []MiniApp{
					{
						Name:   "English",
						Active: false,
					},
				},
			},

			{
				Name:   "Maths",
				Active: false,
				MiniApps: []MiniApp{
					{
						Name:   "Calculator",
						Active: false,
					},
					{
						Name:   "Geometry",
						Active: false,
					},
					{
						Name:   "Trigonometry",
						Active: false,
					},
				},
			},

			{
				Name:   "Text Handler",
				Active: false,
				MiniApps: []MiniApp{
					{
						Name:   "Bash Console",
						Active: false,
					},
					{
						Name:   "JSON Formatter",
						Active: false,
					},
					{
						Name:   "Text Editor",
						Active: false,
					},
				},
			},
		},
	}
)
