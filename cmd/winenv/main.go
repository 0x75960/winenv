package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/0x75960/winenv"
)

//EnvInfo of Running OS
type EnvInfo struct {
	OSProductName  string `json:"os_product_name,omitempty"`
	OSVersion      string `json:"os_version,omitempty"`
	OSReleaseID    string `json:"os_release_id,omitempty"`
	OSBuild        string `json:"os_build,omitempty"`
	OSServicePack  string `json:"service_pack,omitempty"`
	OSArchitecture string `json:"os_architecture,omitempty"`
}

func main() {

	p, v, r, b, sp, err := winenv.GetWindowsVersion()
	if err != nil {
		log.Println(v, b, sp)
		log.Fatalln(err)
	}

	arch := winenv.GetWindowsArchitecture()

	env := EnvInfo{
		OSProductName:  p,
		OSVersion:      v,
		OSReleaseID:    r,
		OSBuild:        b,
		OSServicePack:  sp,
		OSArchitecture: arch,
	}

	json.NewEncoder(os.Stdout).Encode(&env)

}
