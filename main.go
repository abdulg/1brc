package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileName := flag.String("file", "measurements.txt", "a filename")
	flag.Parse()
	//fmt.Printf("processing %s\n", *fileName)
	result, err := process(*fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result.output()
}

func process(fileName string) (Stations, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	lineScanner := bufio.NewScanner(file)
	results := NewStations()
	for lineScanner.Scan() {
		line := lineScanner.Text()
		//fmt.Printf("processing: %s\n", line)
		tokens := strings.Split(line, ";")
		temperature, _ := strconv.ParseFloat(tokens[1], 64)
		if r, ok := results[tokens[0]]; ok {
			r.add(temperature)
			results[tokens[0]] = r
		} else {
			results[tokens[0]] = NewReading(temperature)
		}
	}

	return results, nil
}

type Stations map[string]*Reading

func NewStations() Stations {
	return make(map[string]*Reading)
}

func (s Stations) output() {
	locations := make([]string, 0, len(s))
	for location := range s {
		locations = append(locations, location)
	}
	sort.Strings(locations)

	fmt.Print("{")
	for index, location := range locations {
		if index > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s=%s", location, s[location].result())
	}
	fmt.Println("}")
}

type Reading struct {
	min   float64
	max   float64
	total float64
	count uint32
}

func NewReading(value float64) *Reading {
	return &Reading{
		min:   value,
		max:   value,
		total: value,
		count: 1,
	}
}
func (r *Reading) add(value float64) {
	r.min = math.Min(r.min, value)
	r.max = math.Max(r.max, value)
	r.total += value
	r.count += 1
}

func (r *Reading) result() string {
	return fmt.Sprintf("%.1f/%.1f/%.1f", r.min, r.mean(), r.max)
}

func (r *Reading) mean() float64 {
	return math.Round((r.total/(float64(r.count)))*10) / 10
}
