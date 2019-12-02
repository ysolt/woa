package main

import (
	"fmt"
	"os"
)

func main() {

	distanceLimit := argParser()

	citiesWithinDistance, err := getDatabaseEntries(distanceLimit)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	displayCitiesWithinDistance(citiesWithinDistance)

}
