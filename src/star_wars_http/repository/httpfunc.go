package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func get(url string, v interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &v)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}
