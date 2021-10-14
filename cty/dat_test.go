package cty

import "testing"

func TestDat_String(t *testing.T) {
	type fields struct {
		CountryName   string
		PrimaryPrefix string
		AliasPrefix   string
		Continent     string
		CqZone        string
		ItuZone       string
		LatLon        LatLonDeg
		TimeOffset    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty",
			fields: fields{
				CountryName:   "",
				PrimaryPrefix: "",
				AliasPrefix:   "",
				Continent:     "",
				CqZone:        "",
				ItuZone:       "",
				LatLon: LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				TimeOffset: "",
			},
			want: `{"CountryName":"","PrimaryPrefix":"","AliasPrefix":"","Continent":"","CqZone":"","ItuZone":"","LatLon":{"Lat":0,"Lon":0},"TimeOffset":""}`,
		},
		{
			name: "S5",
			fields: fields{
				CountryName:   "Slovenia",
				PrimaryPrefix: "S5",
				AliasPrefix:   "",
				Continent:     "EU",
				CqZone:        "15",
				ItuZone:       "28",
				LatLon: LatLonDeg{
					Lat: 46.00,
					Lon: 14.00,
				},
				TimeOffset: "-1.0",
			},
			want: `{"CountryName":"Slovenia","PrimaryPrefix":"S5","AliasPrefix":"","Continent":"EU","CqZone":"15","ItuZone":"28","LatLon":{"Lat":46,"Lon":14},"TimeOffset":"-1.0"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Dat{
				CountryName:   tt.fields.CountryName,
				PrimaryPrefix: tt.fields.PrimaryPrefix,
				AliasPrefix:   tt.fields.AliasPrefix,
				Continent:     tt.fields.Continent,
				CqZone:        tt.fields.CqZone,
				ItuZone:       tt.fields.ItuZone,
				LatLon:        tt.fields.LatLon,
				TimeOffset:    tt.fields.TimeOffset,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
