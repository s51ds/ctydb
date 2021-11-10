package api

import (
	"github.com/s51ds/ctydb/cty"
	"os"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		callSign string
	}
	tests := []struct {
		name            string
		args            args
		wantCountryData cty.Dat
		wantErr         bool
	}{
		{
			name: "s59abc",
			args: args{"s59abc"},
			wantCountryData: cty.Dat{
				CountryName:   "Slovenia",
				PrimaryPrefix: "S5",
				AliasPrefix:   "S5",
				Continent:     "EU",
				CqZone:        "15",
				ItuZone:       "28",
				LatLon: cty.LatLonDeg{
					Lat: 46,
					Lon: 14,
				},
				TimeOffset: "-1.0",
			},
			wantErr: false,
		},
		{
			name: "S59ABC",
			args: args{"S59ABC"},
			wantCountryData: cty.Dat{
				CountryName:   "Slovenia",
				PrimaryPrefix: "S5",
				AliasPrefix:   "S5",
				Continent:     "EU",
				CqZone:        "15",
				ItuZone:       "28",
				LatLon: cty.LatLonDeg{
					Lat: 46,
					Lon: 14,
				},
				TimeOffset: "-1.0",
			},
			wantErr: false,
		},

		{
			name: "S51DS/P",
			args: args{"S51DS/P"},
			wantCountryData: cty.Dat{
				CountryName:   "Slovenia",
				PrimaryPrefix: "S5",
				AliasPrefix:   "S5",
				Continent:     "EU",
				CqZone:        "15",
				ItuZone:       "28",
				LatLon: cty.LatLonDeg{
					Lat: 46,
					Lon: 14,
				},
				TimeOffset: "-1.0",
			},
			wantErr: false,
		},
		{
			name: "9A/S51DS/P",
			args: args{"9A/S51DS/P"},
			wantCountryData: cty.Dat{
				CountryName:   "Croatia",
				PrimaryPrefix: "9A",
				AliasPrefix:   "9A",
				Continent:     "EU",
				CqZone:        "15",
				ItuZone:       "28",
				LatLon: cty.LatLonDeg{
					Lat: 45.18,
					Lon: 15.3,
				},
				TimeOffset: "-1.0",
			},
			wantErr: false,
		},
		{
			name: "9A/S51DS",
			args: args{"9A/S51DS"},
			wantCountryData: cty.Dat{
				CountryName:   "Croatia",
				PrimaryPrefix: "9A",
				AliasPrefix:   "9A",
				Continent:     "EU",
				CqZone:        "15",
				ItuZone:       "28",
				LatLon: cty.LatLonDeg{
					Lat: 45.18,
					Lon: 15.3,
				},
				TimeOffset: "-1.0",
			},
			wantErr: false,
		},

		{
			name: "S51DS/MM",
			args: args{"S51DS/MM"},
			wantCountryData: cty.Dat{
				CountryName:   "Slovenia",
				PrimaryPrefix: "S5",
				AliasPrefix:   "S5",
				Continent:     "EU",
				CqZone:        "15",
				ItuZone:       "28",
				LatLon: cty.LatLonDeg{
					Lat: 46,
					Lon: 14,
				},
				TimeOffset: "-1.0",
			},
			wantErr: false,
		},
		{
			name: "S51DS/M",
			args: args{"S51DS/M"},
			wantCountryData: cty.Dat{
				CountryName:   "Slovenia",
				PrimaryPrefix: "S5",
				AliasPrefix:   "S5",
				Continent:     "EU",
				CqZone:        "15",
				ItuZone:       "28",
				LatLon: cty.LatLonDeg{
					Lat: 46,
					Lon: 14,
				},
				TimeOffset: "-1.0",
			},
			wantErr: false,
		},

		{
			name: "OE1ILW/3",
			args: args{"OE1ILW/3"},
			wantCountryData: cty.Dat{
				CountryName:   "Austria",
				PrimaryPrefix: "OE",
				AliasPrefix:   "OE",
				Continent:     "EU",
				CqZone:        "15",
				ItuZone:       "28",
				LatLon: cty.LatLonDeg{
					Lat: 47.33,
					Lon: 13.33,
				},
				TimeOffset: "-1.0",
			},
			wantErr: false,
		},

		{
			name: "AC2AI/KH2",
			args: args{"AC2AI/KH2"},
			wantCountryData: cty.Dat{
				CountryName:   "Guam",
				PrimaryPrefix: "KH2",
				AliasPrefix:   "KH2",
				Continent:     "OC",
				CqZone:        "27",
				ItuZone:       "64",
				LatLon: cty.LatLonDeg{
					Lat: 13.37,
					Lon: 144.7,
				},
				TimeOffset: "-10.0",
			},
			wantErr: false,
		},
		{
			name: "KH2/AC2AI",
			args: args{"KH2/AC2AI"},
			wantCountryData: cty.Dat{
				CountryName:   "Guam",
				PrimaryPrefix: "KH2",
				AliasPrefix:   "KH2",
				Continent:     "OC",
				CqZone:        "27",
				ItuZone:       "64",
				LatLon: cty.LatLonDeg{
					Lat: 13.37,
					Lon: 144.7,
				},
				TimeOffset: "-10.0",
			},
			wantErr: false,
		},
		{
			name: "9M6/OH2YY",
			args: args{"9M6/OH2YY"},
			wantCountryData: cty.Dat{
				CountryName:   "Spratly Islands",
				PrimaryPrefix: "1S",
				AliasPrefix:   "9M6/OH2YY",
				Continent:     "AS",
				CqZone:        "26",
				ItuZone:       "50",
				LatLon: cty.LatLonDeg{
					Lat: 9.88,
					Lon: 114.23,
				},
				TimeOffset: "-8.0",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			gotCountryData, err := Get(tt.args.callSign)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCountryData, tt.wantCountryData) {
				t.Errorf("Get() gotCountryData = %v, want %v", gotCountryData, tt.wantCountryData)
			}
		})
	}
}

func TestMain(m *testing.M) {
	Load()
	os.Exit(m.Run())
}
