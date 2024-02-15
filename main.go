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

// The simple spell correction
type Correction struct {
	Text          string   `json:"Text"`
	BestCandidate string   `json:"best_candidate"`
	Candidates    []string `json:"candidates"`
}

// The Api Response with all corrections
type CorrectionsResponse struct {
	OriginalText string       `json:"original_text"`
	Corrections  []Correction `json:"corrections"`
}

// Sending Request with text
func SendRequest(text string) (*CorrectionsResponse, error) {
	client := &http.Client{}
	urlText := url.QueryEscape(text)
	url := fmt.Sprintf("https://api.apilayer.com/spell/spellchecker?q=%s", urlText)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", "gfnYJdIavCcbaPYCANqwYTXbEpiFf8L4")

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}
	CorrResp := new(CorrectionsResponse)
	if err := json.NewDecoder(res.Body).Decode(CorrResp); err != nil {
		return nil, err
	}

	return CorrResp, nil
}

// Fix the Errors by replacing
func (cr *CorrectionsResponse) Fix() (original string) {
	original = cr.OriginalText

	for _, corr := range cr.Corrections {
		original = strings.Replace(original, corr.Text, corr.BestCandidate, -1)
	}

	return
}

// Get the inputs form command params
func getInputs() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("no text provided")
	}

	paramsString := strings.Join(os.Args[1:], " ")
	return paramsString, nil
}

func main() {
	//GetInputs
	text, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}

	//SendRequest
	corrResp, err := SendRequest(text)
	if err != nil {
		log.Fatal(err)
	}

	//Fix!
	fmt.Println(corrResp.Fix())
}
