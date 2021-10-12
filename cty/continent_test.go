package cty

import "testing"

func TestContinent(t *testing.T) {
	type args struct {
		continentAbbreviation string
	}
	tests := []struct {
		name    string
		args    args
		want    ContinentEnum
		wantErr bool
	}{
		{
			name: "NA",
			args: args{
				continentAbbreviation: "NA",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "AN",
			args: args{
				continentAbbreviation: "AN",
			},
			want:    7,
			wantErr: false,
		},
		{
			name: "err",
			args: args{
				continentAbbreviation: "EUR",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Continent(tt.args.continentAbbreviation)
			if (err != nil) != tt.wantErr {
				t.Errorf("Continent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Continent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
