package main

import (
	"fmt"
	"os"
)

func main() {
	var queryresult QueryResult
	err := getDatabaseEntries(&queryresult)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

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

}
