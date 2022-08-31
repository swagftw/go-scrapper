package crawler

import (
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Product struct {
	Title      string `json:"title"`
	Condition  string `json:"condition"`
	Price      string `json:"price"`
	ProductUrl string `json:"product_url"`
}

type Filter struct {
	Condition string
	MinPrice  float32
	MaxPrice  float32
}

var tmpDir = "./data"

func Crawl(url string, flr Filter) {
	// creat tmp dir
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		err = os.Mkdir(tmpDir, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}

	c := colly.NewCollector()

	// list of products after scraping
	products := make([]*Product, 0)

	// wait group to wait for all products to be scraped
	wg := new(sync.WaitGroup)

	c.OnHTML("li.sresult.lvresult.clearfix.li", func(e *colly.HTMLElement) {
		// if product class is matched
		wg.Add(1)

		// scrape product details and add to products list and create the file to store the product details
		go func(e *colly.HTMLElement, wg *sync.WaitGroup) {
			p := PopulateProducts(e, flr, wg)
			if p != nil {
				products = append(products, p)
			}

			wg.Done()
		}(e, wg)
	})

	log.Println("Scraping started")

	err := c.Visit(url)
	if err != nil {
		log.Fatalln(err)
	}

	wg.Wait()

	log.Println("Scraping completed")
}

// PopulateProducts scrapes the product details and creates the file in tmp directory
func PopulateProducts(e *colly.HTMLElement, flr Filter, wg *sync.WaitGroup) *Product {
	tmp := new(Product)

	title := e.ChildText("h3.lvtitle > a.vip")
	productUrl := e.ChildAttr("h3.lvtitle > a.vip", "href")
	condition := e.ChildText("div.lvsubtitle")
	price := strings.Split(e.ChildText("li.lvprice > span.bold"), "$")[1]

	if flr.Condition != "" && strings.EqualFold(condition, flr.Condition) {
		return nil
	}

	priceFloat, err := strconv.ParseFloat(price, 32)
	if err != nil {
		log.Fatalln(err)
	}

	if flr.MinPrice != 0 {
		if float32(priceFloat) < flr.MinPrice {
			return nil
		}
	}

	if flr.MaxPrice != 0 {
		if float32(priceFloat) > flr.MaxPrice {
			return nil
		}
	}

	tmp.Title = title
	tmp.Condition = condition
	tmp.Price = price
	tmp.ProductUrl = productUrl

	// create file for the product
	wg.Add(1)
	go func(p *Product) {
		createFile(*tmp)
		wg.Done()
	}(tmp)

	return tmp
}
