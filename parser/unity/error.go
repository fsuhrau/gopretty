package unity

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

const (
	ERROR_MATCHER = `(.*)(\(\d+,\d+\)):\serror\s([A-Z0-9]+):\s(.*)`
	ABORT         = `Aborting\sbatchmode\sdue\sto\sfailure:`
	EXCEPTION     = `(.*)Exception:\s(.*)`
)

type ErrorParser struct {
	color          *color.Color
	matcher        *regexp.Regexp
	abortMatch     *regexp.Regexp
	exceptionMatch *regexp.Regexp
}

func NewErrorParser() *ErrorParser {
	red := color.New(color.FgRed)
	red.Add(color.Bold)

	return &ErrorParser{
		color:          red,
		matcher:        regexp.MustCompile(ERROR_MATCHER),
		abortMatch:     regexp.MustCompile(ABORT),
		exceptionMatch: regexp.MustCompile(EXCEPTION),
	}
}

func (parser *ErrorParser) Match(line string, reader *bufio.Reader) bool {

	if match := parser.matcher.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}

	if match := parser.exceptionMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s%s \n", emoji.Sprint(":x: "), match[0])
		for i := 0; i < 1; i++ {
			text, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					os.Exit(1)
				}
				return true
			}
			parser.color.Println(strings.TrimRight(text, "\n"))
		}
		return true
	}

	if match := parser.abortMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("Build Abort: %s \n", match[0])
		for i := 0; i < 1; i++ {
			text, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					os.Exit(1)
				}
				return true
			}
			parser.color.Println(strings.TrimRight(text, "\n"))
		}

		return true
	}

	return false
}
