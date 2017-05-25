package parser

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

const (
	LIBTOOL_MATCHER = `Libtool.*\/(.*\.a)`
	LINKING_MATCHER = `Ld \/?.*\/(.*?) (.*) (.*)`
)

type LinkerParser struct {
	whiteColor     *color.Color
	libToolMatcher *regexp.Regexp
	linkingMatcher *regexp.Regexp
}

func NewLinkerParser() *LinkerParser {
	// white foreground
	whiteColor := color.New(color.FgWhite)
	whiteColor.Add(color.Bold)

	return &LinkerParser{
		whiteColor:     whiteColor,
		libToolMatcher: regexp.MustCompile(LIBTOOL_MATCHER),
		linkingMatcher: regexp.MustCompile(LINKING_MATCHER),
	}
}

func (parser *LinkerParser) Match(line string, reader *bufio.Reader) bool {
	if match := parser.libToolMatcher.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}

	if match := parser.linkingMatcher.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s %s %s %s\n", parser.whiteColor.Sprint("Ld:"), match[1], match[2], match[3])
		return true
	}

	return false
}
