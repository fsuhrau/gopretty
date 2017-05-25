package parser

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

const (
	GENERATE_DSYM_MATCHER           = `GenerateDSYMFile \/.*\/(.*\.dSYM)`
	CREATE_UNIVERSAL_BINARY_MATCHER = `CreateUniversalBinary.*`
)

type PackagingParser struct {
	whiteColor     *color.Color
	dsymMatch      *regexp.Regexp
	universalMatch *regexp.Regexp
}

func NewPackagingParser() *PackagingParser {
	// white foreground
	whiteColor := color.New(color.FgWhite)
	whiteColor.Add(color.Bold)

	return &PackagingParser{
		whiteColor:     whiteColor,
		dsymMatch:      regexp.MustCompile(GENERATE_DSYM_MATCHER),
		universalMatch: regexp.MustCompile(CREATE_UNIVERSAL_BINARY_MATCHER),
	}
}

func (parser *PackagingParser) Match(line string, reader *bufio.Reader) bool {

	if match := parser.dsymMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s %s\n", parser.whiteColor.Sprint("GenerateDSYM File:"), match[1])
		return true
	}

	if match := parser.universalMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", parser.whiteColor.Sprint("CreateUniversalBinary"))
		return true
	}

	return false
}
