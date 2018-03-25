package crawlhub

import (
	"net/http"
)

type ScrapeResult struct {
	PageTitle     string   `json:"page_title"`
	PrimaryH1     string   `json:"primary_h1"`
	ExtractedInfo []string `json:"extracted_info"`
}

type Parser interface {
	ParsePage(*http.Response) ScrapeResult
}
