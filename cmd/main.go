package main

import (
	"flag"
	"github.com/swagftw/crawler/pkg/crawler"
	"log"
	"strconv"
	"strings"
)

func main() {
	// url is optional argument with default value
	url := flag.String("url", "https://www.ebay.com/sch/garlandcomputer/m.html", "url to scrape")

	// all the filters are optional arguments with default values
	condition := flag.String("condition", "", "Condition of the product")
	minPrice := flag.String("min_price", "0", "Minimum price of the product")
	maxPrice := flag.String("max_price", "0", "Maximum price of the product")

	flag.Parse()

	min, err := strconv.ParseFloat(*minPrice, 32)
	if err != nil {
		log.Fatalln(err)
	}

	max, err := strconv.ParseFloat(*maxPrice, 32)
	if err != nil {
		log.Fatalln(err)
	}

	flr := crawler.Filter{
		Condition: strings.TrimSpace(*condition),
		MinPrice:  float32(min),
		MaxPrice:  float32(max),
	}

	crawler.Crawl(*url, flr)
}
