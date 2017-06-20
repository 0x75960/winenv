package winenv

import "testing"

var Win7 = WindowsEnvInfo{
	Version:      "6.1",
	Build:        "7601",
	Product:      "Windows 7 Home Premium",
	ServicePack:  "Service Pack 2",
	Architecture: "x86",
}

var Win10 = WindowsEnvInfo{
	Version:      "6.3",
	Build:        "15063",
	Release:      "1703",
	Product:      "Windows 10 Pro",
	Architecture: "AMD64",
}

func TestWindowsEnvInfo_Equals(t *testing.T) {
	type args struct {
		other WindowsEnvInfo
	}
	tests := []struct {
		name string
		self WindowsEnvInfo
		args args
		want bool
	}{
		{
			name: "case1 Win7 with Win7",
			self: Win7,
			args: args{
				other: Win7,
			},
			want: true,
		},
		{
			name: "case 2 Win7 with Win10",
			self: Win7,
			args: args{
				other: Win10,
			},
			want: false,
		},
		{
			name: "case 3 Win10 with Win7",
			self: Win10,
			args: args{
				other: Win7,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Equals(tt.args.other); got != tt.want {
				t.Errorf("WindowsEnvInfo.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindowsEnvInfo_Newer(t *testing.T) {
	type args struct {
		other WindowsEnvInfo
	}
	tests := []struct {
		name string
		self WindowsEnvInfo
		args args
		want bool
	}{
		{
			name: "case1 Win7 with Win7",
			self: Win7,
			args: args{
				other: Win7,
			},
			want: false,
		},
		{
			name: "case 2 Win7 with Win10",
			self: Win7,
			args: args{
				other: Win10,
			},
			want: false,
		},
		{
			name: "case 3 Win10 with Win7",
			self: Win10,
			args: args{
				other: Win7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Newer(tt.args.other); got != tt.want {
				t.Errorf("WindowsEnvInfo.Newer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindowsEnvInfo_Older(t *testing.T) {
	type args struct {
		other WindowsEnvInfo
	}
	tests := []struct {
		name string
		self WindowsEnvInfo
		args args
		want bool
	}{
		{
			name: "case1 Win7 with Win7",
			self: Win7,
			args: args{
				other: Win7,
			},
			want: false,
		},
		{
			name: "case 2 Win7 with Win10",
			self: Win7,
			args: args{
				other: Win10,
			},
			want: true,
		},
		{
			name: "case 3 Win10 with Win7",
			self: Win10,
			args: args{
				other: Win7,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Older(tt.args.other); got != tt.want {
				t.Errorf("WindowsEnvInfo.Older() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindowsEnvInfo_NewerOrEqual(t *testing.T) {
	type args struct {
		other WindowsEnvInfo
	}
	tests := []struct {
		name string
		self WindowsEnvInfo
		args args
		want bool
	}{
		{
			name: "case1 Win7 with Win7",
			self: Win7,
			args: args{
				other: Win7,
			},
			want: true,
		},
		{
			name: "case 2 Win7 with Win10",
			self: Win7,
			args: args{
				other: Win10,
			},
			want: false,
		},
		{
			name: "case 3 Win10 with Win7",
			self: Win10,
			args: args{
				other: Win7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.NewerOrEqual(tt.args.other); got != tt.want {
				t.Errorf("WindowsEnvInfo.NewerOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindowsEnvInfo_OlderOrEqual(t *testing.T) {
	type args struct {
		other WindowsEnvInfo
	}
	tests := []struct {
		name string
		self WindowsEnvInfo
		args args
		want bool
	}{
		{
			name: "case1 Win7 with Win7",
			self: Win7,
			args: args{
				other: Win7,
			},
			want: true,
		},
		{
			name: "case 2 Win7 with Win10",
			self: Win7,
			args: args{
				other: Win10,
			},
			want: true,
		},
		{
			name: "case 3 Win10 with Win7",
			self: Win10,
			args: args{
				other: Win7,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.OlderOrEqual(tt.args.other); got != tt.want {
				t.Errorf("WindowsEnvInfo.OlderOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
