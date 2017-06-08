package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/0x75960/winenv"
)

func main() {

	env, err := winenv.NewWindowsEnvInfo()
	if err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(os.Stdout).Encode(&env)

}
