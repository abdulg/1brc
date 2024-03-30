package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("file", "measurements.txt", "a filename")
	flag.Parse()
	fmt.Printf("processing %s\n", *fileName)
	err := process(*fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func process(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
