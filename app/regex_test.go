package main

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const LONG_REGEX string = `([a-zA-Z]+)\(([a-zA-Z 0-9]+)\)communicates\(([a-zA-Z 0-9]+)\)to\(\[(.*)\]\)$`
const SHORT_REGEX string = `([a-zA-Z]+)\(([a-zA-Z 0-9]+)\)$`

func TestRegexWithNoOptinalGroups(t *testing.T) {
	Convey("Given a regex with multiple groups", t, func() {
		r, _ := regexp.Compile(SHORT_REGEX)
		Convey("When trying to match a phrase containing a match", func() {
			match := r.FindAllStringSubmatch("circle(item)", -1)
			Convey("Match should have all results", func() {
				// the first item is the whole string
				So(len(match[0]), ShouldEqual, 3)
				So(match[0][1], ShouldEqual, "circle")
				So(match[0][2], ShouldEqual, "item")
			})
		})
	})
}

func TestRegexRegexMatch(t *testing.T) {
	Convey("Given a long regex and a substring regex", t, func() {
		r2, _ := regexp.Compile(LONG_REGEX)
		r, _ := regexp.Compile(SHORT_REGEX)
		s1 := "circle(item)"
		s2 := "circle(item)communicates(http request)to([item1])"
		Convey(
			`A match for the long regex should be detected for the long string`, func() {
				match2 := r2.Match([]byte(s2))
				So(match2, ShouldBeTrue)
				/*
					the short regex doesn't detect a match
					because it has an end of patern character
				*/
				match1 := r.Match([]byte(s2))
				So(match1, ShouldBeFalse)

				groups2 := r2.FindAllStringSubmatch(s2, -1)
				So(len(groups2[0]), ShouldEqual, 5)

			})
		Convey(
			"A match for the short regex doesn't imply a match for the long regex", func() {
				match2 := r2.Match([]byte(s1))
				So(match2, ShouldBeFalse)

				match1 := r.Match([]byte(s1))
				So(match1, ShouldBeTrue)

				groups1 := r.FindAllStringSubmatch(s1, -1)
				So(len(groups1[0]), ShouldEqual, 3)
			})

	})
}
