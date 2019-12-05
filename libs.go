package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
)

type QueryResult struct {
	Bookmark  string `json:"bookmark"`
	totalRows int    `json:"total_rows"`
	Rows      []Doc  `json:"rows"`
}

type Doc struct {
	Id    string    `json:"id"`
	Order []float64 `json:"order"`
	City  City      `json:"fields"`
}

type City struct {
	Name     string  `json:"name"`
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	Distance int
}

func calculateDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
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
	dist = dist * 60 * 1.1515 * 1.609344

	return dist
}

func calculateQueryFilter(dist int, lat1 float64, lon1 float64) (float64, float64, float64, float64) {
	const PI float64 = 3.141592653589793

	radlat1 := PI * lat1 / 180

	angularDistance := float64(dist) / 60 / 1.1515 / 1.609344 * PI / 180

	lat2North := math.Asin(math.Sin(radlat1)*math.Cos(angularDistance)+math.Cos(radlat1)*math.Sin(angularDistance)*math.Cos(0)) * 180 / PI
	lat2South := math.Asin(math.Sin(radlat1)*math.Cos(angularDistance)+math.Cos(radlat1)*math.Sin(angularDistance)*math.Cos(PI)) * 180 / PI

	tmp := math.Asin(math.Sin(radlat1)*math.Cos(angularDistance) + math.Cos(radlat1)*math.Sin(angularDistance)*math.Cos(PI/2))
	lon2East := lon1 + math.Atan2(math.Sin(PI/2)*math.Sin(angularDistance)*math.Cos(radlat1), math.Cos(angularDistance)-math.Sin(radlat1)*math.Sin(tmp))*180/PI

	tmp2 := math.Asin(math.Sin(radlat1)*math.Cos(angularDistance) + math.Cos(radlat1)*math.Sin(angularDistance)*math.Cos(-PI/2))
	lon2West := lon1 + math.Atan2(math.Sin(-PI/2)*math.Sin(angularDistance)*math.Cos(radlat1), math.Cos(angularDistance)-math.Sin(radlat1)*math.Sin(tmp2))*180/PI

	return lat2North, lat2South, lon2East, lon2West
}

func displayCitiesWithinDistance(citiesWithinDistance []City) {
	sort.Sort(ByDistance(citiesWithinDistance))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "City name \t Distance \t Latitude \t Longitude ")
	fmt.Fprintln(w, "-------------\t ---------\t -------------\t ---------\t")

	for i := 0; i < len(citiesWithinDistance); i++ {
		fmt.Fprintln(w, citiesWithinDistance[i].Name+" \t "+strconv.Itoa(citiesWithinDistance[i].Distance)+"\t", citiesWithinDistance[i].Lat, "\t", citiesWithinDistance[i].Lon)
	}
	w.Flush()
}

type ByDistance []City

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }
