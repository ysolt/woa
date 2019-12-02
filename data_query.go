package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func fetchData(bookmarkKey string) ([]byte, error) {

	url := "https://mikerhodes.cloudant.com/airportdb/_design/view1/_search/geo?limit=200&q=lon:[-90%20TO%2090]%20AND%20lat:[-90%20TO%2090]"
	if bookmarkKey != "" {
		url += "&bookmark=" + bookmarkKey
	}

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
		return nil, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	return body, readErr
}

func getDatabaseEntriesFor(lat float64, lon float64, distanceLimit int) ([]City, error) {
	key := ""
	var err error
	var byteValue []byte
	var tmpResults QueryResult
	var citiesWithinDistance []City

	for {
		byteValue, err = fetchData(key)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(byteValue, &tmpResults)

		for i := 0; i < len(tmpResults.Rows); i++ {
			distance := int(calculateDistance(tmpResults.Rows[i].City.Lat, tmpResults.Rows[i].City.Lon, lat, lon, "K"))
			if distance < distanceLimit {
				tmpResults.Rows[i].City.Distance = distance
				citiesWithinDistance = append(citiesWithinDistance, tmpResults.Rows[i].City)
			}

		}

		//fmt.Println(key, len(tmpResults.Rows))
		if key == tmpResults.Bookmark {
			break
		}
		key = tmpResults.Bookmark
	}
	return citiesWithinDistance, err
}
