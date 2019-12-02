package main

import (
	"fmt"
	"os"
)

func main() {

	lat, lon, distanceLimit := argParser()

	citiesWithinDistance, err := getDatabaseEntriesFor(lat, lon, distanceLimit)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	displayCitiesWithinDistance(citiesWithinDistance)

}
