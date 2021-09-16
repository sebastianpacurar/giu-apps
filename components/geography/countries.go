package geography

import (
	"encoding/json"
	"fmt"
	"github.com/AllenDang/giu"
	"image/color"
	"imgui-based-app/APIs"
	"log"
)

var (
	CountryRef = &Country{}
	Details    = CountryRef.Details
)

type Country struct {
	Details   []CountryInfo
	IsUpdated bool
}

// CountryInfo - All details about a country, grabbed through InitCountries() function
type CountryInfo struct {
	Name           string            `json:"name"`
	TopLevelDomain []string          `json:"topLevelDomain"`
	Alpha2Code     string            `json:"alpha2code"`
	Alpha3Code     string            `json:"alpha3code"`
	CallingCodes   []string          `json:"callingCodes"`
	Capital        string            `json:"capital"`
	AltSpellings   []string          `json:"altSpellings"`
	Region         string            `json:"region"`
	Subregion      string            `json:"subregion"`
	Population     int32             `json:"population"`
	Flag           string            `json:"flag"`
	LatLng         []float64         `json:"latlng"`
	Demonym        string            `json:"demonym"`
	Area           float64           `json:"area"`
	Gini           float64           `json:"gini"`
	Timezones      []string          `json:"timezones"`
	Borders        []string          `json:"borders"`
	NativeName     string            `json:"nativeName"`
	NumericCode    string            `json:"numericCode"`
	Currencies     []Currency        `json:"currencies"`
	Languages      []Language        `json:"languages"`
	Translations   map[string]string `json:"translations"`
}

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Language struct {
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

// buildCountryRow - represents one country as a row in the table, the first row being the table head
///Returns all the countries as the rows of the table,
func buildCountryRow() []*giu.TableRowWidget {
	entries := make([]*giu.TableRowWidget, len(Details))
	for i := range entries {
		if i == 0 {
			entries[i] = giu.TableRow(
				giu.Label("Count"),
				giu.Label("Name"),
				giu.Label("Capital"),
				giu.Label("Alpha 2 Code"),
				giu.Label("Subregion"),
				giu.Label("Numeric Code"),
			).Flags(giu.TableRowFlagsHeaders).MinHeight(80).
				BgColor(&(color.RGBA{R: 200, G: 100, B: 100, A: 255}))
		} else {
			entries[i] = giu.TableRow(
				giu.Label(fmt.Sprintf("%d", i)),
				giu.Label(Details[i].Name),
				giu.Label(Details[i].Capital),
				giu.Label(Details[i].Alpha2Code),
				giu.Label(Details[i].Subregion),
				giu.Label(Details[i].NumericCode),
			).MinHeight(80)
		}
	}
	return entries
}

// CountriesTable - builds the table in one shot using the spread operator,
///and the TableRowWidgets obtained from buildCountryRow
func CountriesTable() *giu.TableWidget {
	return giu.Table().
		Flags(giu.TableFlagsSortable).Freeze(0, 1).FastMode(true).Rows(buildCountryRow()...)
}

// InitCountries - runs only once at startup, to populate the Details list of structs.
///after it gets triggered, at the end of the function, Country.IsUpdated
func InitCountries() error {
	data, err := APIs.FetchCountries("all")
	if err != nil {
		log.Fatalln("Eroare la fetch data spre RETST EU", err.Error())
		return err
	}

	if err := json.Unmarshal(data, &Details); err != nil {
		log.Fatalln("Eroare la json Unmarshal pe initCountriesPage()", err.Error())
		return err
	}

	CountryRef.IsUpdated = true
	return nil
}
