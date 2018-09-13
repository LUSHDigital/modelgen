package main

import (
	"log"

	"github.com/LUSHDigital/modelgen/cmd"
)

var (
	version string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	cmd.Execute()
}
