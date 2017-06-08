// Package winenv collect Windows environment informations at local
package winenv

import (
	"os"

	"golang.org/x/sys/windows/registry"
)

// WindowsEnvInfo to get
type WindowsEnvInfo struct {
	Product      string `json:"product"`
	Version      string `json:"version"`
	Release      string `json:"release"`
	Build        string `json:"build"`
	ServicePack  string `json:"service_pack"`
	Architecture string `json:"architecture"`
}

// NewWindowsEnvInfo from Environment variable and Registry
func NewWindowsEnvInfo() (envinfo WindowsEnvInfo, err error) {

	k, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		return WindowsEnvInfo{}, err
	}
	defer k.Close()

	// keys must be exists
	version, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		return WindowsEnvInfo{}, err
	}

	product, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return WindowsEnvInfo{}, err
	}

	// keys maybe not exisits
	release, _, _ := k.GetStringValue("ReleaseId")
	build, _, _ := k.GetStringValue("CurrentBuildNumber")
	servicepack, _, _ := k.GetStringValue("CSDVersion")

	// Get from environment variable
	arch := os.Getenv("Processor_Architecture")

	return WindowsEnvInfo{
		Version:      version,
		Product:      product,
		Release:      release,
		Build:        build,
		ServicePack:  servicepack,
		Architecture: arch,
	}, nil
}
