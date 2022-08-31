# E-Bay Scrapper

## Description

- This is a scrapper written in Golang.
- It scrapes(by default) data from [this](https://www.ebay.com/sch/garlandcomputer/m.html?rt=nc&_dmd=1) URL's product listings on Ebay.
- This particular version uses [Go Colly](https://github.com/gocolly/colly) for scrapping.
- End goal is to find the products listed on the link, take out the product details and save each product in a separate json file.

## Run the Scrapper
- Clone the repository 
```bash
git clone https://github.com/swagftw/go-scrapper.git
```

- Go to the root of the project and run the following command:

```bash
go run cmd/main.go
```

- This scrapper supports filters. Use can use them as below.

```bash
go run cmd/main.go --url https://www.ebay.com/sch/garlandcomputer/m.html --condition pre-owned --min_price 100 --max_price 200
```

- check the scrapped products in the `data` folder.
