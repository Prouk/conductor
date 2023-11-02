package main

import (
	"github.com/Prouk/conductor/src"
	"log"
)

func main() {
	conductor := src.InitConductor()
	err := conductor.Config.ReadConf()
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	log.Printf("bot started on port: %s", conductor.Config.Conductor.Server.Port)
}
