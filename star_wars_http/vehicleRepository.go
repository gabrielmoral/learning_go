package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getVehicles(url string) vehiclesResult {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	vehiclesRes := vehiclesResult{}
	jsonErr := json.Unmarshal(body, &vehiclesRes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return vehiclesRes
}

func getPeople(url string) peopleResult {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	peopleRes := peopleResult{}
	jsonErr := json.Unmarshal(body, &peopleRes)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return peopleRes
}
