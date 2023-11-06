package main

import (
	"flag"
	"fmt"
	"log"

	accountCsv "github.com/fazarmitrais/atm-simulation-stage-2/domain/account/repository/csv"
	accountMapdatastore "github.com/fazarmitrais/atm-simulation-stage-2/domain/account/repository/mapDatastore"
	"github.com/fazarmitrais/atm-simulation-stage-2/service"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "Path to directory")
	flag.Parse()
	acctMap := accountMapdatastore.NewMapDatastore()
	acctCsv := accountCsv.NewCSV(path)
	svc := service.NewService(acctMap, acctCsv)
	err := svc.Import()
	if err != nil {
		log.Fatalf("Error importing : %v", err)
	}
	res, _ := svc.GetAll()
	for _, r := range res {
		fmt.Printf("%s %s %0.f %s \n", r.Name, r.PIN, r.Balance, r.AccountNumber)
	}
}
