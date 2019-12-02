package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func argParser() (float64, float64, int) {
	var parsedArgs [3]float64
	var err error

	if len(os.Args) > 2 {
		flag.Parse()

		for i := 0; i < 3; i++ {
			s := flag.Arg(i)
			parsedArgs[i], err = strconv.ParseFloat(s, 64)
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
		}

	} else {
		fmt.Println("Not enough parameters!\n" +
			"Usage: woa <latitude> <longitude> <distance_in_km>\n" +
			"E.g ./woa 47.497913 19.040236 1000")
		os.Exit(1)
	}
	return parsedArgs[0], parsedArgs[1], int(parsedArgs[2])
}
