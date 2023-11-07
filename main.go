package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fazarmitrais/atm-simulation-stage-2/delivery/rest"
	accountCsv "github.com/fazarmitrais/atm-simulation-stage-2/domain/account/repository/csv"
	accountMapdatastore "github.com/fazarmitrais/atm-simulation-stage-2/domain/account/repository/mapDatastore"
	"github.com/fazarmitrais/atm-simulation-stage-2/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	envInit()
	var path string
	c := context.Background()
	flag.StringVar(&path, "path", "", "Path to directory")
	flag.Parse()
	acctMap := accountMapdatastore.NewMapDatastore()
	acctCsv := accountCsv.NewCSV(path)
	svc := service.NewService(acctMap, acctCsv)
	if path != "" {
		err := svc.Import(c)
		if err != nil {
			log.Fatalf("Error importing : %v", err)
		}
		fmt.Println("Successfully imported data from csv: ", path)
	}
	re := rest.New(svc)
	m := mux.NewRouter()
	re.Register(m)
	fmt.Println("App is running on port 8080")
	http.ListenAndServe(":8080", m)
}

func envInit() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}
