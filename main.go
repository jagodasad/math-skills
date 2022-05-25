package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func readFile(s string) (val []int, err error) {
	a, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}

	eachline := strings.Split(string(a), "\n")

	val = make([]int, 0, len(eachline))

	for _, length := range eachline {

		if len(length) == 0 {
			continue
		}

		n, err := strconv.Atoi(length)
		if err != nil {
			return nil, err
		}
		val = append(val, n)
	}

	return val, nil
}

func main() {
	val, err := readFile("data.txt")
	if err != nil {
		panic(err)
	}
	

	sum := 0.0
	Average := 0.0
	Median := 0.0
	
	
	for i := 0; i < len(val); i++ {
		sum += float64(val[i])
		Average = sum / float64(len(val))

	}



	sort.Ints(val)
	M := len(val) / 2
	if len(val)%2 == 1 {
		Median = float64(val[M])
	} else {
		Median = (float64(val[M-1]) + float64(val[M])) / 2
	}


	sumSqErr := 0.0
	for _, number := range val {
		sumSqErr += math.Pow(float64(number)-float64(Average), 2)
	}
	variance := sumSqErr / float64(len(val))
	sd := math.Sqrt(variance)
	Vari := int(math.Round(variance))
	SD := int(math.Round(sd))
	aver := int(math.Round(Average))
	Med := int(math.Round(Median))

	


	fmt.Println("Average:", aver)
	fmt.Println("Median:", Med)
	fmt.Println("Variance:", Vari)
	fmt.Println("Standard Deviation:", SD)
	
}

