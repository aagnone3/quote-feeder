package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type quote struct {
	Quote      string   `json:"quote"`
	Length     string   `json:"length"`
	Author     string   `json:"author"`
	Tags       []string `json:"tags"`
	Category   string   `json:"category"`
	Date       string   `json:"date"`
	Permalink  string   `json:"permalink"`
	Title      string   `json:"title"`
	Background string   `json:"background"`
	Id         string   `json:"id"`
}

type contents struct {
	Quotes    []quote `json:"quotes"`
	Copyright string  `json:"copyright"`
}

type success struct {
	Total int `json:"total"`
}

type queryResponse struct {
	Success  success  `json:"success"`
	Contents contents `json:"contents"`
}

func parseQuoteOfDay() error {
	var response queryResponse
	file, _ := ioutil.ReadFile("qod.json")
	json.Unmarshal([]byte(file), &response)
	quote := response.Contents.Quotes[0]
	fmt.Printf("%s\n- %s\n", quote.Quote, quote.Author)
	return nil
}

func queryQuoteOfDay() error {
	url := "http://quotes.rest/qod.json"
	getResponse, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer getResponse.Body.Close()
	contents, err := ioutil.ReadAll(getResponse.Body)
	if err != nil {
		panic(err)
	}

    var response queryResponse
    json.Unmarshal([]byte(contents), &response)
	quote := response.Contents.Quotes[0]
	fmt.Printf("%s\n- %s\n", quote.Quote, quote.Author)
	return nil
}
