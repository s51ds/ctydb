package parser

import (
	"github.com.s51ds/ctydb/ad1c"
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
		wantCtyDatList []cty.Dat
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
				ctyDatRecord: "SloveniaCtyDat4test:                 15:  28:  EU: ",
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
		wantDtaRecords []cty.Dat
		wantErr        bool
	}{
		{
			name: "SloveniaCtyDat4test",
			args: args{
				ctyDatRecord: ad1c.SloveniaCtyDat4test,
			},
			wantDtaRecords: []cty.Dat{
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S5",
					Continent:     "EU",
					CqZone:        "15",
					ItuZone:       "28",
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
					Continent:     "EU",
					CqZone:        "15",
					ItuZone:       "28",
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
					Continent:     "EU",
					CqZone:        "15",
					ItuZone:       "28",
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
					Continent:     "EU",
					CqZone:        "15",
					ItuZone:       "28",
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
					Continent:     "EU",
					CqZone:        "15",
					ItuZone:       "28",
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
					Continent:     "EU",
					CqZone:        "15",
					ItuZone:       "28",
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
		wantCtyDatList []cty.Dat
		wantErr        bool
	}{
		{
			name: "SwedenCtyDat4test",
			args: args{
				ctyDatRecord: ad1c.SwedenCtyDat4test,
			},
			wantCtyDatList: []cty.Dat{
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM",
					Continent:     "EU",
					CqZone:        "14",
					ItuZone:       "18",
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
					Continent:     "EU",
					CqZone:        "14",
					ItuZone:       "18",
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
					Continent:     "EU",
					CqZone:        "14",
					ItuZone:       "18",
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
					Continent:     "EU",
					CqZone:        "14",
					ItuZone:       "18",
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
					Continent:     "EU",
					CqZone:        "40",
					ItuZone:       "18",
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
					Continent:     "EU",
					CqZone:        "14",
					ItuZone:       "18",
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
					Continent:     "EU",
					CqZone:        "14",
					ItuZone:       "18",
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
		wantCtyDatList []cty.Dat
		wantErr        bool
	}{
		{
			name: "AfricanItalyCtyDat4test",
			args: args{
				ctyDatRecord: ad1c.AfricanItalyCtyDat4test,
			},
			wantCtyDatList: []cty.Dat{
				{
					CountryName:   "African Italy",
					PrimaryPrefix: "IG9",
					AliasPrefix:   "IG9",
					Continent:     "AF",
					CqZone:        "33",
					ItuZone:       "37",
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
					Continent:     "AF",
					CqZone:        "33",
					ItuZone:       "37",
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
					Continent:     "AF",
					CqZone:        "33",
					ItuZone:       "37",
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
					Continent:     "AF",
					CqZone:        "33",
					ItuZone:       "37",
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
		wantCtyDatList []cty.Dat
		wantErr        bool
	}{
		{
			name: "YemenCtyDat4test",
			args: args{
				ctyDatRecord: ad1c.YemenCtyDat4test,
			},
			wantCtyDatList: []cty.Dat{
				{
					CountryName:   "Yemen",
					PrimaryPrefix: "7O",
					AliasPrefix:   "7O",
					Continent:     "AS",
					CqZone:        "21",
					ItuZone:       "39",
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
					Continent:     "AS",
					CqZone:        "37",
					ItuZone:       "48",
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
					Continent:     "AS",
					CqZone:        "37",
					ItuZone:       "48",
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
		wantCtyDatList []cty.Dat
		wantErr        bool
	}{
		{
			name: "Peter 1 Island",
			args: args{
				ctyDatRecord: ad1c.Peter1IslandCtyDat4test,
			},
			wantCtyDatList: []cty.Dat{
				{
					CountryName:   "Peter 1 Island",
					PrimaryPrefix: "3Y/p",
					AliasPrefix:   "3Y/p",
					Continent:     "SA",
					CqZone:        "12",
					ItuZone:       "72",
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
					Continent:     "SA",
					CqZone:        "12",
					ItuZone:       "72",
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
		wantCtyDatList []cty.Dat
		wantErr        bool
	}{
		{
			name: "BouvetCtyDat4test",
			args: args{
				ctyDatRecord: ad1c.BouvetCtyDat4test,
			},
			wantCtyDatList: []cty.Dat{
				{
					CountryName:   "Bouvet",
					PrimaryPrefix: "3Y/b",
					AliasPrefix:   "3Y/b",
					Continent:     "AF",
					CqZone:        "38",
					ItuZone:       "67",
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
					Continent:     "AF",
					CqZone:        "38",
					ItuZone:       "67",
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
					Continent:     "AF",
					CqZone:        "38",
					ItuZone:       "67",
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
					Continent:     "AF",
					CqZone:        "38",
					ItuZone:       "67",
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
		wantCtyDatList []cty.Dat
		wantErr        bool
	}{
		{
			name: "Slovenia Cty Wt Mod",
			args: args{
				ctyDatRecord: ad1c.SloveniaWtModDat4test,
			},
			wantCtyDatList: []cty.Dat{
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S5",
					Continent:     "EU",
					CqZone:        "15",
					ItuZone:       "28",
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
		wantCtyDatList []cty.Dat
		wantErr        bool
	}{
		{
			name: "Sweden Cty Wt Mod",
			args: args{
				ctyDatRecord: ad1c.SwedenWtModDat4test,
			},
			wantCtyDatList: []cty.Dat{
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM",
					Continent:     "EU",
					CqZone:        "14",
					ItuZone:       "18",
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
					Continent:     "EU",
					CqZone:        "14",
					ItuZone:       "18",
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
					Continent:     "EU",
					CqZone:        "14",
					ItuZone:       "18",
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
		wantCtyDatList []cty.Dat
		wantErr        bool
	}{
		{
			name: "USA Cty Wt Mod",
			args: args{
				ctyDatRecord: ad1c.UnitedStatesWtModDat4test,
			},
			wantCtyDatList: []cty.Dat{
				{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "K",
					Continent:     "NA",
					CqZone:        "05",
					ItuZone:       "08",
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
					Continent:     "NA",
					CqZone:        "05",
					ItuZone:       "08",
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
					Continent:     "NA",
					CqZone:        "05",
					ItuZone:       "08",
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
					Continent:     "NA",
					CqZone:        "04",
					ItuZone:       "07",
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
				ctyDatRecords: ad1c.TestInput1,
			},
			want: ad1c.TestOutput1,
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

func Test_parseCtyDatRecords(t *testing.T) {
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
			name: "CtyDatFile4test",
			args: args{
				ctyDatRecords: ad1c.CtyDatFile4test,
			},
			wantMsize: 21314,
			wantErr:   false,
		},
		{
			name: "CtyWtModDatFile4test",
			args: args{
				ctyDatRecords: ad1c.CtyWtModDatFile4test,
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
