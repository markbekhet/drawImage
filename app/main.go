package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sync"
)

func help() {
	log.Fatalln(
		` To use this program you must enter a path to a file
			e.g: ./drawImage -path=file
		`)
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "path", "", "the path of the file that has the user's input")
	flag.Parse()
	if len(filePath) == 0 {
		help()
		os.Exit(-1)
	}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("there was an error reading the file %v", err)
		os.Exit(-1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error scanning the file %v", err)
	}

	long, short := compileRegexes()

	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
			os.Exit(-1)
		}
	}()

	parseChannel := make(chan []string)

	var parseWG sync.WaitGroup

	for scanner.Scan() {
		line := scanner.Text()
		parseWG.Add(1)
		go parseStatement(line, long, short, parseChannel, &parseWG)
	}
	objectChannel := make(chan DrawingObject)
	var constructWG sync.WaitGroup

	for data := range parseChannel {
		go constructDrawingObject(data, objectChannel, &constructWG)
	}
	parseWG.Wait()
	close(parseChannel)

}
