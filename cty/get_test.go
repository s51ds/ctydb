package cty

import (
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
		wantCountryData Dat
		wantErr         bool
	}{
		{
			name: "err-1",
			args: args{"S59ABC"},
			wantCountryData: Dat{
				CountryName:   "",
				PrimaryPrefix: "",
				AliasPrefix:   "",
				Continent:     0,
				CqZone:        0,
				ItuZone:       0,
				LatLon:        LatLonDeg{},
				TimeOffset:    "",
			},
			wantErr: true,
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
