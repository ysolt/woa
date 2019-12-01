package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func argParser() int {
	var distanceLimit int
	var err error

	if len(os.Args) > 1 {
		flag.Parse()
		s := flag.Arg(0)
		// string to int
		distanceLimit, err = strconv.Atoi(s)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
	} else {
		distanceLimit = 1000
		fmt.Println("Defaulting distance limiter to to " + strconv.Itoa(distanceLimit))
	}
	return distanceLimit
}
