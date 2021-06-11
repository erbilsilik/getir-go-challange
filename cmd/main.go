package main

import (
	"fmt"
	"github.com/erbilsilik/getir-go-challange/infrastructure/repository"
	"github.com/erbilsilik/getir-go-challange/pkg/mongodb"
	"github.com/erbilsilik/getir-go-challange/pkg/utilities"
	"github.com/erbilsilik/getir-go-challange/usecase/record"
	"log"
	"os"
)

func main()  {
	mongodb.New(
		os.Getenv("MONGODB_URI"),
		os.Getenv("MONGODB_DB"),
	)
	recordRepository := repository.NewRecordRepositoryMongoDB()
	recordService := record.NewService(recordRepository)

	layout := "2006-01-02"
	startDateParsed := utilities.ParseDate(layout, "2016-01-26")
	endDateParsed := utilities.ParseDate(layout, "2018-02-02")

	// TODO -> Use command line arguments
	query := record.CalculateRecordsTotalCountQuery{
		StartDate: startDateParsed,
		EndDate: endDateParsed,
		MinCount: 2700,
		MaxCount: 3000,
	}
	records, err := recordService.List(&query)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range records {
		fmt.Printf("%s %d %s \n", r.Key, r.TotalCount, r.CreatedAt)
	}
}
