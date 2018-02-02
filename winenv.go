// Package winenv collect Windows environment informations at local
package winenv

import (
	"fmt"
	"strings"
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

// String for WindowsEnvInfo
func (e WindowsEnvInfo) String() (str string) {

	str = e.Product

	if e.ServicePack != "" {
		sp := strings.Replace(e.ServicePack, "Service Pack ", "SP", 1)
		str = fmt.Sprintf("%s %s", str, sp)
	}

	str = fmt.Sprintf("%s (%s) [ver. %s] [build %s]", str, e.Architecture, e.Version, e.Build)

	if e.Release != "" {
		str = fmt.Sprintf("%s [release %s]", str, e.Release)
	}

	return
}
