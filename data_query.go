package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchData() ([]byte, error) {
	url := "https://mikerhodes.cloudant.com/airportdb/_design/view1/_search/geo?limit=200&q=lon:[-90%20TO%2090]%20AND%20lat:[-90%20TO%2090]"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return nil, err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}
	//fmt.Println(string(body))
	return body, err
}

func getDatabaseEntries(queryresult *QueryResult) error {
	byteValue, err := fetchData()
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &queryresult)
	return err
}
