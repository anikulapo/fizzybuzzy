// fizzybuzzy project main.go
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	MIN := flag.Int("min", 1, "min value of print range")
	MAX := flag.Int("max", 100, "max value of print range")
	STEP := flag.Int("step", 1, "step value for defined range")
	FILE := flag.String("file", "", ".csv with int,string pairs. ex 3,fizz")
	flag.Parse()

	if *MAX < *MIN {
		fmt.Printf("max value %v cannot be less than min value %v\n", *MAX, *MIN)
		return
	}

	fizzMap := readValues(*FILE)

	for i := *MIN; i <= *MAX; i += *STEP {
		var r string
		for k, v := range fizzMap {
			if i%k == 0 {
				r += v
			}
		}

		if len(r) > 0 {
			println(r)
		} else {
			println(i)
		}
	}
}

//takes a file name of a .csv, returns a map of contents
func readValues(f string) map[int]string {

	m := make(map[int]string)

	file, _ := os.Open(f)
	defer file.Close()

	reader := csv.NewReader(file)

	for {

		line, err := reader.Read()
		if err != nil {
			break
		}

		//ParseInt always comes back with int64
		key, _ := strconv.ParseInt(line[0], 10, 0)

		m[int(key)] = line[1]
	}

	// classic values. we should default these
	// if we can't populate our map
	if len(m) <= 0 {

		m[3] = "Fizz"
		m[5] = "Buzz"

	}

	return m
}
