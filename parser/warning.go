package parser

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
	// warnings
	COMPILE_WARNING_MATCHER         = `(\/.+\/(.*):.*:.*):\swarning:\s(.*)`
	LD_WARNING_MATCHER              = `(ld: )warning: (.*)`
	GENERIC_WARNING_MATCHER         = `warning:\s(.*)`
	WILL_NOT_BE_CODE_SIGNED_MATCHER = `(.* will not be code signed because .*)`
	XCODE_WARNING_MATCHER           = `---.*WARNING:\s(.*)`
)

type WarningParser struct {
	yellowColor            *color.Color
	compileWarningMatcher  *regexp.Regexp
	linkerWarningMatcher   *regexp.Regexp
	genericWarningMatcher  *regexp.Regexp
	codeSignWarningMatcher *regexp.Regexp
	xcodeWarningMatch      *regexp.Regexp
}

func NewWarningParser() *WarningParser {
	yellow := color.New(color.FgYellow)
	yellow.Add(color.Bold)

	return &WarningParser{
		yellowColor:            yellow,
		compileWarningMatcher:  regexp.MustCompile(COMPILE_WARNING_MATCHER),
		linkerWarningMatcher:   regexp.MustCompile(LD_WARNING_MATCHER),
		genericWarningMatcher:  regexp.MustCompile(GENERIC_WARNING_MATCHER),
		codeSignWarningMatcher: regexp.MustCompile(WILL_NOT_BE_CODE_SIGNED_MATCHER),
		xcodeWarningMatch:      regexp.MustCompile(XCODE_WARNING_MATCHER),
	}
}

func (parser *WarningParser) Match(line string, reader *bufio.Reader) bool {

	if match := parser.compileWarningMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.yellowColor.Printf("%s%s\n", emoji.Sprint(":warning: "), match[0])
		// read 2 additonal lines
		for i := 0; i < 2; i++ {
			text, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					os.Exit(1)
				}
				return true
			}
			parser.yellowColor.Println(strings.TrimRight(text, "\n"))
		}
		return true
	}

	if match := parser.linkerWarningMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.yellowColor.Println(match[0])
		return true
	}

	if match := parser.genericWarningMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.yellowColor.Println(match[0])
		return true
	}

	if match := parser.codeSignWarningMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.yellowColor.Println(match[0])
		return true
	}

	if match := parser.xcodeWarningMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.yellowColor.Println(match[0])
		return true
	}

	return false
}
