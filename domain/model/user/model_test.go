package user

import (
	"reflect"
	"testing"
)

func TestNewName(t *testing.T) {
	type args struct {
		first string
		last  string
	}
	tests := []struct {
		name    string
		args    args
		want    Name
		wantErr bool
	}{
		{
			name: "valid 1",
			args: args{
				first: "Taro",
				last:  "Yamada",
			},
			want: Name{
				first: "Taro",
				last:  "Yamada",
			},
			wantErr: false,
		},
		{
			name: "invalid 1",
			args: args{
				first: "Taro",
				last:  "",
			},
			wantErr: true,
		},
		{
			name: "invalid 2",
			args: args{
				first: "",
				last:  "Yamada",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewName(tt.args.first, tt.args.last)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_FullName(t *testing.T) {
	type fields struct {
		first string
		last  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "check mapping",
			fields: fields{
				first: "Taro",
				last:  "Yamada",
			},
			want: "Taro Yamada",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Name{
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := n.FullName(); got != tt.want {
				t.Errorf("Name.FullName() = %v, want %v", got, tt.want)
			}
		})
	}
}
