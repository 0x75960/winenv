// Package winenv collect Windows environment informations at local
package winenv

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/windows/registry"
)

// GetWindowsVersion from registry
func GetWindowsVersion() (product, version, release, build, servicepack string, err error) {

	k, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`,
		registry.QUERY_VALUE,
	)

	if err != nil {
		return product, version, release, build, servicepack, err
	}
	defer k.Close()

	succeeded := false

	// もしかしたら将来的になくなるかもしれないキーなので、基本的にはエラーにしない
	product, _, err = k.GetStringValue("ProductName")
	if err == nil {
		succeeded = true
	} else {
		log.Println("failed to get Product", err)
	}

	version, _, err = k.GetStringValue("CurrentVersion")
	if err == nil {
		succeeded = true
	} else {
		log.Println("failed to get CurrentVersion", err)
	}

	release, _, err = k.GetStringValue("ReleaseId")
	if err == nil {
		succeeded = true
	} else {
		log.Println("failed to get ReleaseId", err)
	}

	build, _, err = k.GetStringValue("CurrentBuildNumber")
	if err == nil {
		succeeded = true
	} else {
		log.Println("failed to get CurrentBuildNumer", err)
	}

	servicepack, _, err = k.GetStringValue("CSDVersion")
	if err == nil {
		succeeded = true
	} else {
		log.Println("failed to get CSDVersion", err)
	}

	if !succeeded {
		// すべて取得失敗した場合のみ、エラーとする
		return "", "", "", "", "", fmt.Errorf("failed to get environment informations")
	}

	return product, version, release, build, servicepack, nil
}

// GetWindowsArchitecture from environment variable
func GetWindowsArchitecture() (s string) {
	arch := os.Getenv("Processor_Architecture")
	switch arch {
	case "x86":
		return "x86"
	case "AMD64":
		return "x64"
	default:
		return arch
	}
}
