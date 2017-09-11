package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

const googleURL = "https://www.google.com/finance/historical?q=GOOG&startdate=2017-7-01&enddate=2017-7-31&output=csv"

type GoogleDataStruct struct {
	Date   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int64
}

func GoogleData() {
	// data := getData()
	out, err := http.Get(googleURL)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Body.Close()

	w := csv.NewReader(out.Body)
	records, err := w.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)
	// record, _ := w.Read()
	// googVal := GoogleDataStruct{record[0], atoi(record[1]), record[2], record[3], record[4], record[5]}
}
