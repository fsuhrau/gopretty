package parser

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

const (
	COPY_HEADER_MATCHER  = `CpHeader\s(.*\.h)\s(.*\.h)`
	COPY_PLIST_MATCHER   = `CopyPlistFile\s(.*\.plist)\s(.*\.plist)`
	COPY_STRINGS_MATCHER = `CopyStringsFile.*\/(.*.strings)`
	CPRESOURCE_MATCHER   = `CpResource\s(.*)\s\/`
)

type CopyParser struct {
	whiteColor           *color.Color
	copyHeaderMatcher    *regexp.Regexp
	copyPlistMatcher     *regexp.Regexp
	copyStringsMatcher   *regexp.Regexp
	copyResourcesMatcher *regexp.Regexp
}

func NewCopyParser() *CopyParser {
	// white foreground
	whiteColor := color.New(color.FgWhite)
	whiteColor.Add(color.Bold)

	return &CopyParser{
		whiteColor:           whiteColor,
		copyHeaderMatcher:    regexp.MustCompile(COPY_HEADER_MATCHER),
		copyPlistMatcher:     regexp.MustCompile(COPY_PLIST_MATCHER),
		copyStringsMatcher:   regexp.MustCompile(COPY_STRINGS_MATCHER),
		copyResourcesMatcher: regexp.MustCompile(CPRESOURCE_MATCHER),
	}
}

func (parser *CopyParser) Match(line string, reader *bufio.Reader) bool {

	if match := parser.copyHeaderMatcher.FindStringSubmatch(line); len(match) > 0 {
		fmt.Println(match[0])
		return true
	}

	if match := parser.copyPlistMatcher.FindStringSubmatch(line); len(match) > 0 {
		fmt.Println(match[0])
		return true
	}

	if match := parser.copyStringsMatcher.FindStringSubmatch(line); len(match) > 0 {
		fmt.Println(match[0])
		return true
	}

	if match := parser.copyResourcesMatcher.FindStringSubmatch(line); len(match) > 0 {
		fmt.Println(match[0])
		return true
	}

	return false
}
