package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type request struct {
	apiKey  string
	pNumber string
}

type response struct {
	Phone    string      `json:"phone"`
	Valid    bool        `json:"valid"`
	Location string      `json:"location"`
	Type     string      `json:"type"`
	Carrier  string      `json:"carrier"`
	Country  countryData `json:"country"`
}

type countryData struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}

func Start(pNumber string, apiKey string) {
	var r request
	r.pNumber = pNumber
	r.apiKey = apiKey

	printData(getResponse(r.send()))
}

func (r request) send() *http.Response {
	url := "https://phonevalidation.abstractapi.com/v1/?api_key=" + r.apiKey + "&phone=" + r.pNumber

	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func printData(r response) {

	fmt.Println("")

	if !r.Valid {
		fmt.Println("Phone number is not valid!")
		return
	}

	fmt.Println("Result information:")
	fmt.Println("Phone number:", r.Phone)
	fmt.Println("Valid:", r.Valid)
	fmt.Println("Location:", r.Location)
	fmt.Println("Type:", r.Type)
	fmt.Println("Carrier:", r.Carrier)
	fmt.Println("")

	fmt.Println("Information about country:")
	fmt.Println("Code:", r.Country.Code)
	fmt.Println("Prefix:", r.Country.Prefix)
}

func getResponse(resp *http.Response) response {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r response

	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}

	return r
}
