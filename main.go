package main

import (
	"fmt"
	"log"
	pV "phoneValidation/api"
)

const apiKey = "4be4be9d04fa4775b7528e395033397f"

type phoneNumber string

func main() {
	pV.Start(scanNumber(), apiKey)
}

func scanNumber() string {
	var pNumber string

	_, err := fmt.Scan(&pNumber)
	if err != nil {
		log.Fatal(err)
	}

	return pNumber
}
