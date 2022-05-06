package application

import "testing"

func Test_birthdayInts(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   int
		wantErr bool
	}{
		{
			name: "valid 1",
			args: args{
				s: "1990/12/31",
			},
			want:    1990,
			want1:   12,
			want2:   31,
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
			got, got1, got2, err := birthdayInts(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("birthdayInts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("birthdayInts() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("birthdayInts() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("birthdayInts() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
