package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
    layoutISO = "2006-01-02"
    layoutUS  = "January 2, 2006"
	annualizedMonthStart = "01"
	annualizedMonthEnd = "12"
	annualizedDayStart = "01"
	annualizedDayEnd = "31"
)

type Portfolio struct {
	Stocks []Stock
	Profit float64
}

type Stock struct {
	Date 	string	`json:"date"`
	Price 	float64	`json:"price"`
}

func (p *Portfolio) calculateProfit(startDate time.Time, endDate time.Time) float64 {
	fmt.Println("Serching between: ", startDate.Format(layoutUS), 
				" and ", endDate.Format(layoutUS))

	// iterate through every stock within the pfs array and
    // calculate the profit between a period
	for _, stock := range p.Stocks {
		// calculate the days between the start and end date
		// to iterate in that range
		days := endDate.Sub(startDate).Hours() / 24

		// assign the start date to a variable to be able
		// to iterate over the entire collection
		startDateTemp := startDate
		for i := 0; i < int(days); i++ {
			rangeValid := startDateTemp.Before(endDate)

			if rangeValid {
				p.Profit += stock.calculatePrice(startDateTemp)
			}
			startDateTemp = startDateTemp.AddDate(0, 0, 1)
		}
	}
	return p.Profit
}

func (s *Stock) calculatePrice(date time.Time) float64 {
	t, _ := time.Parse(layoutISO, s.Date)
	if t.Format(layoutUS) == date.Format(layoutUS) {
		return s.Price
	}
	return 0.0
}

func getDate(splitDate []string) time.Time {
	year, _ := strconv.Atoi(splitDate[0])
	month, _ := strconv.Atoi(splitDate[1])
	day, _ := strconv.Atoi(splitDate[2])
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func main() {
	// Open jsonFile with data dummy
    jsonFile, err := os.Open("stocks.json")
    // if os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened stocks.json")
    // defer the closing of jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read open jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // initialize Portfolio class with an stocks arrays
    var pfs Portfolio

    // we unmarshal the byteArray which contains
    // jsonFile content into 'pfs' which we defined above
    json.Unmarshal(byteValue, &pfs)

    // assign pfs to temporal portfolio for calculate profits 
	p := pfs

	// receive startDate and endDate from the execution arguments
	// and generate split with year, month and day
	fmt.Println("Here")
	splitStartDate := strings.Split(os.Args[1], "-")
    splitEndDate := strings.Split(os.Args[2], "-")

	// call the function Profit between given dates
	fmt.Println("The profit is: ", p.calculateProfit(getDate(splitStartDate), getDate(splitEndDate)))
	
	// initialize dates for annualized profit
	splitStartDate = []string{splitStartDate[0], annualizedMonthStart, annualizedDayStart}
	splitEndDate = []string{splitStartDate[0], annualizedMonthEnd, annualizedDayEnd}

	// The current profit is reset to zero in order 
	// to obtain the annualized profit
	p.Profit = 0.0

	// call the function Profit for annualized period
	annualized := strconv.FormatFloat(p.calculateProfit(getDate(splitStartDate), getDate(splitEndDate)), 'f', -1, 64)
	fmt.Println("and Profit annualized ", annualized)
}