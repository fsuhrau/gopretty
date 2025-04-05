package xcode

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

const (
	CHECK_DEPENDENCIES_MATCHER = `Check dependencies`
	CLEAN_REMOVE_MATCHER       = `Clean.Remove`
	// compile proicess
	ANALYZE_MATCHER = `Analyze(?:Shallow)?\s(.*\/(.*\.(?:m|mm|cc|cpp|c|cxx)))\s*`
	// 1 = target
	// 2 = project
	// 3 = configuration
	BUILD_TARGET_MATCHER = `=== BUILD TARGET\s(.*)\sOF PROJECT\s(.*)\sWITH.*CONFIGURATION\s(.*)\s===`
	// 1 = target
	// 2 = project
	// 3 = configuration
	AGGREGATE_TARGET_MATCHER = `=== BUILD AGGREGATE TARGET\s(.*)\sOF PROJECT\s(.*)\sWITH.*CONFIGURATION\s(.*)\s===`
	// 1 = target
	// 2 = project
	// 3 = configuration
	ANALYZE_TARGET_MATCHER = `=== ANALYZE TARGET\s(.*)\sOF PROJECT\s(.*)\sWITH.*CONFIGURATION\s(.*)\s===`
	// 1 = target
	// 2 = project
	// 3 = configuration
	CLEAN_TARGET_MATCHER = `=== CLEAN TARGET\s(.*)\sOF PROJECT\s(.*)\sWITH.*CONFIGURATION\s(.*)\s===`
)

type InfoParser struct {
	whiteColor       *color.Color
	dependencyMatch  *regexp.Regexp
	cleanRemoteMatch *regexp.Regexp
	analyseMatch     *regexp.Regexp
	aggregateMatch   *regexp.Regexp
	targetMatch      *regexp.Regexp
	cleanMatch       *regexp.Regexp
}

func NewInfoParser() *InfoParser {
	// white foreground
	whiteColor := color.New(color.FgWhite)
	whiteColor.Add(color.Bold)

	return &InfoParser{
		whiteColor:       whiteColor,
		dependencyMatch:  regexp.MustCompile(CHECK_DEPENDENCIES_MATCHER),
		cleanRemoteMatch: regexp.MustCompile(CLEAN_REMOVE_MATCHER),
		analyseMatch:     regexp.MustCompile(ANALYZE_MATCHER),
		aggregateMatch:   regexp.MustCompile(AGGREGATE_TARGET_MATCHER),
		targetMatch:      regexp.MustCompile(BUILD_TARGET_MATCHER),
		cleanMatch:       regexp.MustCompile(CLEAN_TARGET_MATCHER),
	}
}

func (parser *InfoParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {

	if match := parser.dependencyMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.whiteColor.Println(match[0])
		return true
	}

	if match := parser.cleanRemoteMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.whiteColor.Println(match[0])
		return true
	}

	if match := parser.analyseMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("Analyse %s with Configuration %s\n", parser.whiteColor.Sprintf(match[2]), parser.whiteColor.Sprintf(match[3]))
		return true
	}

	if match := parser.aggregateMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("Build Aggregate %s with Configuration %s\n", parser.whiteColor.Sprintf(match[2]), parser.whiteColor.Sprintf(match[3]))
		return true
	}

	if match := parser.targetMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("Build Project %s with Configuration %s\n", parser.whiteColor.Sprintf(match[2]), parser.whiteColor.Sprintf(match[3]))
		return true
	}

	if match := parser.cleanMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("Clean Project %s for Configuration %s\n", parser.whiteColor.Sprintf(match[2]), parser.whiteColor.Sprintf(match[3]))
		return true
	}

	return false
}
