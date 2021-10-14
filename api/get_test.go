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
			name: "err-1",
			args: args{"S59ABC"},
			wantCountryData: cty.Dat{
				CountryName:   "",
				PrimaryPrefix: "",
				AliasPrefix:   "",
				Continent:     "",
				CqZone:        "",
				ItuZone:       "",
				LatLon:        cty.LatLonDeg{},
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
