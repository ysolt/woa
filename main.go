package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
)

type City struct {
	Name     string  `json:"name"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Distance int
}

type Cites struct {
	City []City
}

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

type ByDistance []City

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

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
		distance := int(distance(city[i].Lat, city[i].Lon, 47.497913, 19.040236, "K"))
		if distance < distanceFilter {
			city[i].Distance = distance
			citiesWithinDistance = append(citiesWithinDistance, city[i])
		}

		//fmt.Println("User Type: " + city[i].Name)
	}
	sort.Sort(ByDistance(citiesWithinDistance))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for i := 0; i < len(citiesWithinDistance); i++ {
		fmt.Fprintln(w, citiesWithinDistance[i].Name+" \t "+strconv.Itoa(citiesWithinDistance[i].Distance)+"\t")
	}
	//fmt.Println(city)
	w.Flush()

	defer jsonFile.Close()
}
