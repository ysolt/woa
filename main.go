package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	filename := "example.json"
	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened " + filename)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var city []City
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &city)

	var distanceFilter int
	if len(os.Args) > 1 {
		flag.Parse()
		s := flag.Arg(0)
		// string to int
		distanceFilter, err = strconv.Atoi(s)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
	} else {
		distanceFilter = 1000
		fmt.Println("Defaulting distance to 1000")
	}

	var citiesWithinDistance []City
	for i := 0; i < len(city); i++ {
		distance := int(calculateDistance(city[i].Lat, city[i].Lon, 47.497913, 19.040236, "K"))
		if distance < distanceFilter {
			city[i].Distance = distance
			citiesWithinDistance = append(citiesWithinDistance, city[i])
		}

		//fmt.Println("User Type: " + city[i].Name)
	}
	displayCitiesWithinDistance(citiesWithinDistance)

	defer jsonFile.Close()
}
