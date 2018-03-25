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
	"github.com/gorilla/mux"
)

type ScrapeRequest struct {
	StartURL       string `json:"start_url"`
	MaxConcurrency int    `json:"max_concurrency"`
	CallbackURL    string `json:"callback_url"`
}

type DummyParser struct {
}

func (d DummyParser) ParsePage(*http.Response) crawlhub.ScrapeResult {
	return crawlhub.ScrapeResult{}
}

func StartScrapeEndpoint(w http.ResponseWriter, req *http.Request) {
	var scrapeJob ScrapeRequest
	json.NewDecoder(req.Body).Decode(&scrapeJob)
	parser := DummyParser{}
	go func(url string, parser DummyParser, concurrency int) {
		crawlhub.StandardCrawl(url, url, parser, concurrency)
	}(scrapeJob.StartURL, parser, scrapeJob.MaxConcurrency)
	fmt.Println(scrapeJob)
	json.NewEncoder(w).Encode(scrapeJob)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/scrape", StartScrapeEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

```
