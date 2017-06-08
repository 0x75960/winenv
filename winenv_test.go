package winenv

import (
	"testing"
)

func TestNewWindowsEnvInfo(t *testing.T) {

	_, err := NewWindowsEnvInfo()
	if err != nil {
		t.Fatalf("error happend caused by %s (this test always fail if this environment is not windows)\n", err)
	}

}

func TestWindowsEnvInfo_String(t *testing.T) {

	tests := []struct {
		name    string
		e       WindowsEnvInfo
		wantStr string
	}{
		{
			name: "case1: win10",
			e: WindowsEnvInfo{
				Version:      "6.3",
				Build:        "15063",
				Release:      "1703",
				Product:      "Windows 10 Pro",
				Architecture: "AMD64",
			},
			wantStr: "Windows 10 Pro (AMD64) [ver. 6.3] [build 15063] [release 1703]",
		},
		{
			name: "case2: win7",
			e: WindowsEnvInfo{
				Version:      "6.1",
				Build:        "7601",
				Product:      "Windows 7 Home Premium",
				ServicePack:  "Service Pack 2",
				Architecture: "x86",
			},
			wantStr: "Windows 7 Home Premium SP2 (x86) [ver. 6.1] [build 7601]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStr := tt.e.String(); gotStr != tt.wantStr {
				t.Errorf("WindowsEnvInfo.String() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}
