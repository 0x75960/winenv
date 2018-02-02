package lwinenv

import (
	"os"

	"github.com/0x75960/winenv"
	"golang.org/x/sys/windows/registry"
)

// NewWindowsEnvInfo from Environment variable and Registry
func NewWindowsEnvInfo() (envinfo winenv.WindowsEnvInfo, err error) {

	k, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		return envinfo, err
	}
	defer k.Close()

	product, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return envinfo, err
	}

	// keys must be exists
	version, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		return envinfo, err
	}

	build, _, err := k.GetStringValue("CurrentBuildNumber")
	if err != nil {
		return envinfo, err
	}

	// keys maybe not exisits
	release, _, _ := k.GetStringValue("ReleaseId")
	servicepack, _, _ := k.GetStringValue("CSDVersion")

	// Get from environment variable
	arch := os.Getenv("Processor_Architecture")

	return winenv.WindowsEnvInfo{
		Version:      version,
		Product:      product,
		Release:      release,
		Build:        build,
		ServicePack:  servicepack,
		Architecture: arch,
	}, nil
}
