package parser

import (
	"github.com.s51ds/ctydb/cty"
	"reflect"
	"testing"
)

func Test_parseCtyDatRecordErrorCases(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []cty.Dta
		wantErr        bool
	}{
		{
			name: "emptyRecord",
			args: args{
				ctyDatRecord: "",
			},
			wantCtyDatList: nil,
			wantErr:        true,
		},
		{
			name: "tooShorRecord",
			args: args{
				ctyDatRecord: "SloveniaCtyDat:                 15:  28:  EU: ",
			},
			wantCtyDatList: nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordSlovenia(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantDtaRecords []cty.Dta
		wantErr        bool
	}{
		{
			name: "SloveniaCtyDat",
			args: args{
				ctyDatRecord: SloveniaCtyDat,
			},
			wantDtaRecords: []cty.Dta{
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S5",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE15,
					ItuZone:       cty.ITUZONE28,
					LatLon: cty.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S51LGT/LH",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE15,
					ItuZone:       cty.ITUZONE28,
					LatLon: cty.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S52AL/YL",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE15,
					ItuZone:       cty.ITUZONE28,
					LatLon: cty.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S52L/LH",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE15,
					ItuZone:       cty.ITUZONE28,
					LatLon: cty.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S58U/LH",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE15,
					ItuZone:       cty.ITUZONE28,
					LatLon: cty.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S59HIJ/LH",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE15,
					ItuZone:       cty.ITUZONE28,
					LatLon: cty.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDtaRecords, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDtaRecords, tt.wantDtaRecords) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotDtaRecords, tt.wantDtaRecords)
			}
		})
	}
}

func Test_parseCtyDatRecordSweden(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []cty.Dta
		wantErr        bool
	}{
		{
			name: "SwedenCtyDat",
			args: args{
				ctyDatRecord: SwedenCtyDat,
			},
			wantCtyDatList: []cty.Dta{
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE14,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},

				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "7S",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE14,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "8S",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE14,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SL",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE14,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},

				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "8S8ODEN/MM",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE40,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SK7RN/LH",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE14,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM7AAL/S",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE14,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordAfricanItaly(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []cty.Dta
		wantErr        bool
	}{
		{
			name: "AfricanItalyCtyDat",
			args: args{
				ctyDatRecord: AfricanItalyCtyDat,
			},
			wantCtyDatList: []cty.Dta{
				{
					CountryName:   "African Italy",
					PrimaryPrefix: "IG9",
					AliasPrefix:   "IG9",
					Continent:     cty.AF,
					CqZone:        cty.CQZONE33,
					ItuZone:       cty.ITUZONE37,
					LatLon: cty.LatLonDeg{
						Lat: 35.67,
						Lon: 12.67,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "African Italy",
					PrimaryPrefix: "IG9",
					AliasPrefix:   "IH9",
					Continent:     cty.AF,
					CqZone:        cty.CQZONE33,
					ItuZone:       cty.ITUZONE37,
					LatLon: cty.LatLonDeg{
						Lat: 35.67,
						Lon: 12.67,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "African Italy",
					PrimaryPrefix: "IG9",
					AliasPrefix:   "IO9Y",
					Continent:     cty.AF,
					CqZone:        cty.CQZONE33,
					ItuZone:       cty.ITUZONE37,
					LatLon: cty.LatLonDeg{
						Lat: 35.67,
						Lon: 12.67,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "African Italy",
					PrimaryPrefix: "IG9",
					AliasPrefix:   "IY9A",
					Continent:     cty.AF,
					CqZone:        cty.CQZONE33,
					ItuZone:       cty.ITUZONE37,
					LatLon: cty.LatLonDeg{
						Lat: 35.67,
						Lon: 12.67,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordYemen(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []cty.Dta
		wantErr        bool
	}{
		{
			name: "YemenCtyDat",
			args: args{
				ctyDatRecord: YemenCtyDat,
			},
			wantCtyDatList: []cty.Dta{
				{
					CountryName:   "Yemen",
					PrimaryPrefix: "7O",
					AliasPrefix:   "7O",
					Continent:     cty.AS,
					CqZone:        cty.CQZONE21,
					ItuZone:       cty.ITUZONE39,
					LatLon: cty.LatLonDeg{
						Lat: 15.65,
						Lon: 48.12,
					},
					TimeOffset: "-3.0",
				},
				{
					CountryName:   "Yemen",
					PrimaryPrefix: "7O",
					AliasPrefix:   "7O2A",
					Continent:     cty.AS,
					CqZone:        cty.CQZONE37,
					ItuZone:       cty.ITUZONE48,
					LatLon: cty.LatLonDeg{
						Lat: 15.65,
						Lon: 48.12,
					},
					TimeOffset: "-3.0",
				},
				{
					CountryName:   "Yemen",
					PrimaryPrefix: "7O",
					AliasPrefix:   "7O6T",
					Continent:     cty.AS,
					CqZone:        cty.CQZONE37,
					ItuZone:       cty.ITUZONE48,
					LatLon: cty.LatLonDeg{
						Lat: 15.65,
						Lon: 48.12,
					},
					TimeOffset: "-3.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordPeter1Island(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []cty.Dta
		wantErr        bool
	}{
		{
			name: "Peter 1 Island",
			args: args{
				ctyDatRecord: Peter1IslandCtyDat,
			},
			wantCtyDatList: []cty.Dta{
				{
					CountryName:   "Peter 1 Island",
					PrimaryPrefix: "3Y/p",
					AliasPrefix:   "3Y/p",
					Continent:     cty.SA,
					CqZone:        cty.CQZONE12,
					ItuZone:       cty.ITUZONE72,
					LatLon: cty.LatLonDeg{
						Lat: -68.77,
						Lon: -90.58,
					},
					TimeOffset: "4.0",
				},
				{
					CountryName:   "Peter 1 Island",
					PrimaryPrefix: "3Y/p",
					AliasPrefix:   "3Y0X",
					Continent:     cty.SA,
					CqZone:        cty.CQZONE12,
					ItuZone:       cty.ITUZONE72,
					LatLon: cty.LatLonDeg{
						Lat: -68.77,
						Lon: -90.58,
					},
					TimeOffset: "4.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordBouvet(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []cty.Dta
		wantErr        bool
	}{
		{
			name: "BouvetCtyDat",
			args: args{
				ctyDatRecord: BouvetCtyDat,
			},
			wantCtyDatList: []cty.Dta{
				{
					CountryName:   "Bouvet",
					PrimaryPrefix: "3Y/b",
					AliasPrefix:   "3Y/b",
					Continent:     cty.AF,
					CqZone:        cty.CQZONE38,
					ItuZone:       cty.ITUZONE67,
					LatLon: cty.LatLonDeg{
						Lat: -54.42,
						Lon: 3.38,
					},
					TimeOffset: "-1.0",
				},

				{
					CountryName:   "Bouvet",
					PrimaryPrefix: "3Y/b",
					AliasPrefix:   "3Y/ZS6GCM",
					Continent:     cty.AF,
					CqZone:        cty.CQZONE38,
					ItuZone:       cty.ITUZONE67,
					LatLon: cty.LatLonDeg{
						Lat: -54.42,
						Lon: 3.38,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Bouvet",
					PrimaryPrefix: "3Y/b",
					AliasPrefix:   "3Y0C",
					Continent:     cty.AF,
					CqZone:        cty.CQZONE38,
					ItuZone:       cty.ITUZONE67,
					LatLon: cty.LatLonDeg{
						Lat: -54.42,
						Lon: 3.38,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Bouvet",
					PrimaryPrefix: "3Y/b",
					AliasPrefix:   "3Y0E",
					Continent:     cty.AF,
					CqZone:        cty.CQZONE38,
					ItuZone:       cty.ITUZONE67,
					LatLon: cty.LatLonDeg{
						Lat: -54.42,
						Lon: 3.38,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}
func Test_parseCtyDatRecordSloveniaCtyWtDat(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []cty.Dta
		wantErr        bool
	}{
		{
			name: "Slovenia Cty Wt Mod",
			args: args{
				ctyDatRecord: SloveniaWtModDat,
			},
			wantCtyDatList: []cty.Dta{
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S5",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE15,
					ItuZone:       cty.ITUZONE28,
					LatLon: cty.LatLonDeg{
						Lat: 46.00,
						Lon: 14.00,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordSwedenCtyWtDat(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []cty.Dta
		wantErr        bool
	}{
		{
			name: "Sweden Cty Wt Mod",
			args: args{
				ctyDatRecord: SwedenWtModDat,
			},
			wantCtyDatList: []cty.Dta{
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE14,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},

				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "7S",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE14,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM7",
					Continent:     cty.EU,
					CqZone:        cty.CQZONE14,
					ItuZone:       cty.ITUZONE18,
					LatLon: cty.LatLonDeg{
						Lat: 55.58,
						Lon: 13.10,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordUnitedStatesCtyWtDat(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []cty.Dta
		wantErr        bool
	}{
		{
			name: "USA Cty Wt Mod",
			args: args{
				ctyDatRecord: UnitedStatesWtModDat,
			},
			wantCtyDatList: []cty.Dta{
				{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "K",
					Continent:     cty.NA,
					CqZone:        cty.CQZONE5,
					ItuZone:       cty.ITUZONE8,
					LatLon: cty.LatLonDeg{
						Lat: 37.53,
						Lon: -91.67,
					},
					TimeOffset: "5.0",
				},

				{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "AA",
					Continent:     cty.NA,
					CqZone:        cty.CQZONE5,
					ItuZone:       cty.ITUZONE8,
					LatLon: cty.LatLonDeg{
						Lat: 37.53,
						Lon: -91.67,
					},
					TimeOffset: "5.0",
				},

				{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "K5ZD",
					Continent:     cty.NA,
					CqZone:        cty.CQZONE5,
					ItuZone:       cty.ITUZONE8,
					LatLon: cty.LatLonDeg{
						Lat: 42.27,
						Lon: -71.37,
					},
					TimeOffset: "5.0",
				},

				{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "AD1C",
					Continent:     cty.NA,
					CqZone:        cty.CQZONE4,
					ItuZone:       cty.ITUZONE7,
					LatLon: cty.LatLonDeg{
						Lat: 39.52,
						Lon: -105.20,
					},
					TimeOffset: "6.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_removeComments(t *testing.T) {
	type args struct {
		ctyDatRecords string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "comments-1",
			args: args{
				ctyDatRecords: TestInput1,
			},
			want: TestOutput1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeComments(tt.args.ctyDatRecords); got != tt.want {
				t.Errorf("removeComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCtyDatRecordsGo(t *testing.T) {
	type args struct {
		ctyDatRecords string
	}
	tests := []struct {
		name      string
		args      args
		wantMsize int
		wantErr   bool
	}{
		{
			name: "CtyDatRecords",
			args: args{
				ctyDatRecords: CtyDatRecords,
			},
			wantMsize: 21314,
			wantErr:   false,
		},
		{
			name: "CtyWtModDatRecords",
			args: args{
				ctyDatRecords: CtyWtModDatRecords,
			},
			wantMsize: 26991,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMsize, err := parseCtyDatRecordsForTest(tt.args.ctyDatRecords)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecordsForTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMsize != tt.wantMsize {
				t.Errorf("parseCtyDatRecordsForTest() = %v, want %v", gotMsize, tt.wantMsize)
			}
		})
	}
}
