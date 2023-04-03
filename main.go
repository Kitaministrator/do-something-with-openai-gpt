package main

import (
	"log"
)

func main() {
	// Load config from file or env
	err := loadConfig()
	if err != nil {
		panic(err)
	}
	// printout all key-value pairs in ClientConfig
	log.Println(ClientCfg.BaseURL)
}
