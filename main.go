package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := "resources/cloudant_response_example.json"
	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened " + filename)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var queryresult QueryResult
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &queryresult)

	distanceLimit := argParser()

	var citiesWithinDistance []City
	for i := 0; i < len(queryresult.Rows); i++ {
		distance := int(calculateDistance(queryresult.Rows[i].City.Lat, queryresult.Rows[i].City.Lon, 47.497913, 19.040236, "K"))
		if distance < distanceLimit {
			queryresult.Rows[i].City.Distance = distance
			citiesWithinDistance = append(citiesWithinDistance, queryresult.Rows[i].City)
		}

	}
	displayCitiesWithinDistance(citiesWithinDistance)

	defer jsonFile.Close()
}
