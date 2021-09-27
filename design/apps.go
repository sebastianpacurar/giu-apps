package design

import "github.com/AllenDang/giu"

type AppI interface{}
type MiniAppI interface{}

type Apps struct {
	AppsList []App
}

type App struct {
	Name        string
	Active      bool
	WindowIndex int
	Screens     int
	Current     bool
	MiniApps    []MiniApp
	Layout      giu.Layout
}

type MiniApp struct {
	Name   string
	Active bool
	Screen int
}

// appsS - The struct of the Menu
// appsI - The AppsList[] as an Interface (to be used with RangeBuilder() as values param)
var (
	AppsI = make([]interface{}, len(AppsS.AppsList))
	AppsS = &Apps{
		AppsList: []App{
			{
				Name:    "Text Handler",
				Active:  true,
				Screens: 1,
				MiniApps: []MiniApp{
					{
						Name:   "Text Editor",
						Active: true,
						Screen: 1,
					},
					{
						Name:   "Bash Console",
						Active: false,
						Screen: 1,
					},
					{
						Name:   "JSON Formatter",
						Active: false,
						Screen: 1,
					},
				},
			},

			{
				Name:    "Geography",
				Active:  false,
				Screens: 1,
				MiniApps: []MiniApp{
					{
						Name:   "All Countries",
						Active: false,
						Screen: 1,
					},
					{
						Name:   "Map",
						Active: false,
						Screen: 1,
					},
				},
			},

			{
				Name:    "Dictionary",
				Active:  false,
				Screens: 1,
				MiniApps: []MiniApp{
					{
						Name:   "English",
						Active: false,
						Screen: 1,
					},
				},
			},

			{
				Name:    "Maths",
				Active:  false,
				Screens: 1,
				MiniApps: []MiniApp{
					{
						Name:   "Calculator",
						Active: false,
						Screen: 1,
					},
					{
						Name:   "Geometry",
						Active: false,
						Screen: 1,
					},
					{
						Name:   "Trigonometry",
						Active: false,
						Screen: 1,
					},
				},
			},
		},
	}
)
