package unity

import (
	"bufio"
	"github.com/fatih/color"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	GRADLE_FALURE_MATCHER          = `FAILURE:\s(.*)\n`
	GRADLE_WHAT_WENT_WRONG_MATCHER = `\*\sWhat\swent\swrong:\n`
	GRADLE_UNEXPECTED_MATCHER      = `.*unexpectedly exit.*\n`
)

type GradleParser struct {
	color           *color.Color
	failureMatch    *regexp.Regexp
	wwwMatch        *regexp.Regexp
	unexpectedMatch *regexp.Regexp
}

func NewGradleParser() *GradleParser {
	//white := color.New(color.FgWhite)
	red := color.New(color.FgRed)
	red.Add(color.Bold)
	return &GradleParser{
		color:           red,
		failureMatch:    regexp.MustCompile(GRADLE_FALURE_MATCHER),
		wwwMatch:        regexp.MustCompile(GRADLE_WHAT_WENT_WRONG_MATCHER),
		unexpectedMatch: regexp.MustCompile(GRADLE_UNEXPECTED_MATCHER),
	}
}

func (parser *GradleParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {

	if match := parser.failureMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("❌Gradle Falure: %s\n", match[1])
		return true
	}

	if match := parser.wwwMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s", match[0])
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					os.Exit(1)
				}
				return true
			}
			if text == "\n" {
				return true
			}
			parser.color.Println(strings.TrimRight(text, "\n"))
		}
		return true
	}

	if match := parser.unexpectedMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("❌%s", match[0])
		return true
	}

	return false
}
