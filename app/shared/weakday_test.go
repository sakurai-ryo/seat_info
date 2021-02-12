package shared

import (
	"testing"
	"time"
)

// TODO: 異常系とかのテストケースの追加

// 厳密なユニットテストではないが、ひとまず
func TestIsOpen(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "return true",
			args: args{
				t: time.Date(2021, 2, 14, 12, 5, 6, 0, time.Local),
			},
			want: true,
		},
		{
			name: "return false due to third sunday",
			args: args{
				t: time.Date(2021, 2, 21, 12, 5, 6, 0, time.Local),
			},
			want: false,
		},
		{
			name: "return false due to regular day before open",
			args: args{
				t: time.Date(2021, 2, 22, 6, 59, 6, 0, time.Local),
			},
			want: false,
		},
		{
			name: "return false due to regular day after open",
			args: args{
				t: time.Date(2021, 2, 22, 19, 0, 6, 0, time.Local),
			},
			want: false,
		},
		{
			name: "return false due to holiday before open",
			args: args{
				t: time.Date(2021, 2, 20, 10, 59, 6, 0, time.Local),
			},
			want: false,
		},
		{
			name: "return false due to holiday before open",
			args: args{
				t: time.Date(2021, 2, 22, 19, 0, 6, 0, time.Local),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOpen(tt.args.t); got != tt.want {
				t.Errorf("IsOpen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSaturday(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "return true if saturday",
			args: args{
				t: time.Date(2021, 2, 6, 4, 5, 6, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSaturday(tt.args.t); got != tt.want {
				t.Errorf("isSaturday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSunday(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "return true if sunday",
			args: args{
				t: time.Date(2021, 2, 7, 4, 5, 6, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSunday(tt.args.t); got != tt.want {
				t.Errorf("isSunday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsThirdSunday(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "return true if third saturday",
			args: args{
				t: time.Date(2021, 2, 21, 4, 5, 6, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isThirdSunday(tt.args.t); got != tt.want {
				t.Errorf("isThirdSunday() = %v, want %v", got, tt.want)
			}
		})
	}
}
