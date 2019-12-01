package main

import (
	"fmt"
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

func calculateDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
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

func displayCitiesWithinDistance(citiesWithinDistance []City) {
	sort.Sort(ByDistance(citiesWithinDistance))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for i := 0; i < len(citiesWithinDistance); i++ {
		fmt.Fprintln(w, citiesWithinDistance[i].Name+" \t "+strconv.Itoa(citiesWithinDistance[i].Distance)+"\t")
	}
	w.Flush()
}

type ByDistance []City

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }
