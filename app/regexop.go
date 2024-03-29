package main

import (
	"fmt"
	"regexp"
	"sync"
)

const LONG_REGEX string = `([a-zA-Z]+)\(([a-zA-Z 0-9]+)\)communicates\(([a-zA-Z 0-9]+)\)to\(\[(.*)\]\)$`
const SHORT_REGEX string = `([a-zA-Z]+)\(([a-zA-Z 0-9]+)\)$`

func compileRegexes() (*regexp.Regexp, *regexp.Regexp) {
	longRegex, _ := regexp.Compile(LONG_REGEX)
	shortRegex, _ := regexp.Compile(SHORT_REGEX)
	return longRegex, shortRegex
}

func parseStatement(line string, long, short *regexp.Regexp, channel chan []string, wg *sync.WaitGroup) {
	// try to parse using the short regex
	// if it fails try the long regex
	// If both regexes fail we have a problem

	match := short.FindAllStringSubmatch(line, -1)
	if match == nil {
		match = long.FindAllStringSubmatch(line, -1)
		if match == nil {
			panic(fmt.Sprintf("The format of the line %v is not known", line))
		}
	}
	channel <- match[0][1:]
	wg.Done()

}
