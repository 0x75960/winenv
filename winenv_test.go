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
