package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/icholy/holidayapi"
)

var (
	key     string
	timeout time.Duration
	params  holidayapi.Params
)

func init() {
	flag.StringVar(&key, "key", "", "api key")
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "request timeout")
	flag.StringVar(&params.Country, "country", "CA", "comma separated countries")
	flag.IntVar(&params.Year, "year", time.Now().Year(), "year")
	flag.IntVar(&params.Month, "month", 0, "1 or 2 digit month [1-12]")
	flag.IntVar(&params.Day, "day", 0, "1 or 2 digit day, [1-31]")
	flag.BoolVar(&params.Previous, "previous", false, "return previous holidays based on the date")
	flag.BoolVar(&params.Upcoming, "upcoming", false, "return upcoming holidays based on the date")
	flag.BoolVar(&params.Public, "public", false, "return only public holidays")
	flag.StringVar(&params.Language, "language", "", "ISO 639-1 format")
	flag.Parse()
}

func main() {
	client := holidayapi.New(key)
	client.HTTPClient = &http.Client{
		Timeout: timeout,
	}
	holidays, err := client.Holidays(params)
	if err != nil {
		log.Fatal(err)
	}
	for _, h := range holidays {
		fmt.Println(h)
	}
}
