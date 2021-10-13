package cty

import "fmt"

// Dat represent country data as define in https://www.country-files.com/cty-dat-format/
type Dat struct {
	CountryName   string //Country Name
	PrimaryPrefix string
	AliasPrefix   string        //Primary or Alias DXCC Prefix without optional * indicator
	Continent     ContinentEnum //2-letter Continent abbreviation
	CqZone        CqzoneEnum    //CQ Zone
	ItuZone       ItuzoneEnum   //ITU Zone
	LatLon        LatLonDeg     //Latitude in degrees, + for North; Longitude in degrees, + for West
	TimeOffset    string        //Local time offset from GMT
}

type LatLonDeg struct {
	Lat, Lon float64
}

func (a *LatLonDeg) String() string {
	//	return fmt.Sprintf("Lat=%.4f, Lon=%.4f", a.Lat, a.Lon)
	return fmt.Sprintf("{%.6f %.6f}", a.Lat, a.Lon)
}
