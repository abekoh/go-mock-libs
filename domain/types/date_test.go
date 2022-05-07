package types

import (
	"reflect"
	"testing"
	"time"
)

func TestNewBirthday(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		want    Birthday
		wantErr bool
	}{
		{
			name: "valid 1",
			args: args{
				year:  1990,
				month: 12,
				day:   31,
			},
			want: Birthday{
				birthdayTime: time.Date(1990, 12, 31, 0, 0, 0, 0, time.Local),
			},
			wantErr: false,
		},
		{
			name: "invalid 1",
			args: args{
				year:  -1,
				month: 12,
				day:   31,
			},
			wantErr: true,
		},
		{
			name: "invalid 2",
			args: args{
				year:  1990,
				month: 13,
				day:   31,
			},
			wantErr: true,
		},
		{
			name: "invalid 3",
			args: args{
				year:  1990,
				month: 12,
				day:   32,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBirthday(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBirthday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBirthday() = %v, want %v", got, tt.want)
			}
		})
	}
}
