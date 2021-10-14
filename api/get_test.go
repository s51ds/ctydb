package api

import (
	"github.com.s51ds/ctydb/cty"
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
	}
	for _, tt := range tests {
		Load()
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
