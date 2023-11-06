package main

import (
	"flag"
	"log"

	"github.com/fazarmitrais/atm-simulation-stage-2/delivery/csv"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "Path to directory")
	flag.Parse()
	csv := csv.NewCSV(path)
	err := csv.Import()
	if err != nil {
		log.Fatal(err)
	}
}
