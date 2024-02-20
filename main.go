package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	APIKEY = "gfnYJdIavCcbaPYCANqwYTXbEpiFf8L4"
)

// The simple spell correction
type Correction struct {
	Text          string   `json:"Text"`
	BestCandidate string   `json:"best_candidate"`
	Candidates    []string `json:"candidates"`
}

// The Api Response with all corrections
type FixingResponse struct {
	OriginalText string       `json:"original_text"`
	Corrections  []Correction `json:"corrections"`
}

// Get the inputs form command params
func (fr *FixingResponse) GetInputs() *FixingResponse{
	if len(os.Args) < 2 {
		log.Fatal("Provide a text!")
	}

	paramsString := strings.Join(os.Args[1:], " ")
	fr.OriginalText = paramsString
	return fr
}


// Sending Request with text
func (fr *FixingResponse)SendRequest() *FixingResponse {
	client := &http.Client{}
	urlText := url.QueryEscape(fr.OriginalText)
	url := fmt.Sprintf("https://api.apilayer.com/spell/spellchecker?q=%s", urlText)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", APIKEY)

	if err != nil {
		log.Fatal("Error")
	}
	res, err := client.Do(req)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}

	if err != nil {
		log.Fatal("Error")
	}
	CorrResp := new(FixingResponse)
	if err := json.NewDecoder(res.Body).Decode(CorrResp); err != nil {
		log.Fatal("Error")
	}

	return CorrResp
}

// Fix the Errors by replacing
func (fr *FixingResponse) Fix() {
	original := fr.OriginalText

	for _, corr := range fr.Corrections {
		original = strings.Replace(original, corr.Text, corr.BestCandidate, -1)
	}

	fmt.Println(original)
}

func main() {
	new(FixingResponse).
	GetInputs().
	SendRequest().
	Fix()
}
