package main

import (
	"log"
	"os"

	"github.com/nissy/trfile"
)

// Golang time format
var fileNameFormat = "logs/2006/01/02/halo.log" //=> logs/2016/09/13/halo.log

func main() {
	w, err := trfile.NewFormat(fileNameFormat)

	if err != nil {
		os.Exit(1)
	}

	log.SetOutput(w)
	log.Println("Halo World!")
}
