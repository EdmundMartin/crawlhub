# crawlhub
Toy
```golang
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EdmundMartin/crawlhub"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

type ScrapeRequest struct {
	StartURL       string `json:"start_url"`
	MaxConcurrency int    `json:"max_concurrency"`
	CallbackURL    string `json:"callback_url"`
}

type DummyParser struct {
}

func (d DummyParser) ParsePage(doc *goquery.Document) crawlhub.ScrapeResult {
	scrape := crawlhub.ScrapeResult{}
	scrape.PageTitle = doc.Find("title").First().Text()
	scrape.PrimaryH1 = doc.Find("h1").First().Text()
	return scrape
}

func StartScrapeEndpoint(w http.ResponseWriter, req *http.Request) {
	var scrapeJob ScrapeRequest
	json.NewDecoder(req.Body).Decode(&scrapeJob)
	parser := DummyParser{}
	go func(url, callback string, parser DummyParser, concurrency int) {
		baseDomain, _ := crawlhub.ParseBaseURL(url)
		crawlhub.StandardCrawl(baseDomain, url, callback, parser, concurrency)
	}(scrapeJob.StartURL, scrapeJob.CallbackURL, parser, scrapeJob.MaxConcurrency)
	fmt.Println(scrapeJob)
	json.NewEncoder(w).Encode(scrapeJob)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/scrape", StartScrapeEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

```
# Example Post
```json
{"start_url": "https://www.vox.com/", "max_concurrency": 2, "callback_url": "http://127.0.0.1:5000/api/"}
```
