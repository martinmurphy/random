package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var filename string
	var numToChoose int
	var onlyChoices bool

	processCommandLineArgs(&filename, &numToChoose, &onlyChoices)

	buf := readFile(filename)

	items := strings.Split(strings.TrimSpace(string(buf)), "\n")

	numItems := len(items)

	if !onlyChoices {
		fmt.Printf("Choosing %d items from %d in file: %s\n", numToChoose, numItems, filename)
	}

	winners := choose(numToChoose, numItems)

	for i, winner := range winners {
		if onlyChoices {
			fmt.Printf("%s\n", items[winner])
		} else {
			fmt.Printf("place: %d, '%s' (%d)\n", i+1, items[winner], winner+1) // most people like seeing 1 based choices
		}
	}
}

func processCommandLineArgs(filename *string, numToChoose *int, onlyChoices *bool) {
	flag.StringVar(filename, "file", "items.txt", "Name of file containing choices")
	flag.IntVar(numToChoose, "c", 3, "Number to Choose")
	flag.BoolVar(onlyChoices, "onlyChoices", false, "only return the choices")
	flag.Parse()
}

func readFile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}
