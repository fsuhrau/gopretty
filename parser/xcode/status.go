package xcode

import (
	"bufio"
	"regexp"

	"github.com/fatih/color"
)

const (
	STATUS_MATCHER = `\*\*\s\w*\s(\w*)\s\*\*`
)

type StatusParser struct {
	greenColor  *color.Color
	redColor    *color.Color
	statusMatch *regexp.Regexp
}

func NewStatusParser() *StatusParser {
	green := color.New(color.FgGreen)
	green.Add(color.Bold)

	red := color.New(color.FgRed)
	red.Add(color.Bold)

	return &StatusParser{
		greenColor:  green,
		redColor:    red,
		statusMatch: regexp.MustCompile(STATUS_MATCHER),
	}
}

func (parser *StatusParser) Match(line string, reader *bufio.Reader) bool {
	if match := parser.statusMatch.FindStringSubmatch(line); len(match) > 0 {
		if match[1] == "SUCCEEDED" {
			parser.greenColor.Printf("%s\n", match[0])
		} else if match[1] == "FAILED" {
			parser.redColor.Printf("%s\n", match[0])
		}
		return true
	}

	return false
}
