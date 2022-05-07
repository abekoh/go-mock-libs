package application

import (
	"reflect"
	"testing"

	"github.com/abekoh/go-mock-libs/domain/types"
)

func Test_parseDate(t *testing.T) {
	validDate, _ := types.NewDate(1990, 12, 31)

	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    types.Date
		wantErr bool
	}{
		{
			name: "valid 1",
			args: args{
				s: "1990/12/31",
			},
			want:    validDate,
			wantErr: false,
		},
		{
			name: "invalid 1",
			args: args{
				s: "12/31",
			},
			wantErr: true,
		},
		{
			name: "invalid 2",
			args: args{
				s: "hoge/12/31",
			},
			wantErr: true,
		},
		{
			name: "invalid 3",
			args: args{
				s: "1990/hoge/31",
			},
			wantErr: true,
		},
		{
			name: "invalid 4",
			args: args{
				s: "1990/12/hoge",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDate(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
