package xcode

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

const (
	// 1 = file
	// 2 = achitecture
	// 3 = language
	COMPILE_MATCHER            = `Compile\w*\s[/:\w:\-\.]*\s([/:\w:\-\.]*\.(?:m|mm|c|cc|cpp|cxx|swift))\s\w*\s(\w*)\s([\w:+:-]*)\s.*`
	COMPILE_COMMAND_MATCHER    = `\s*(.*\/bin\/clang\s.*\s\-c\s(.*\.(?:m|mm|c|cc|cpp|cxx))\s.*\.o)$/`
	COMPILE_XIB_MATCHER        = `CompileXIB\s(.*\/(.*\.xib))/`
	COMPILE_STORYBOARD_MATCHER = `CompileStoryboard\s(.*\/([^\/].*\.storyboard))/`
)

type CompileParser struct {
	whiteColor             *color.Color
	compileMatch           *regexp.Regexp
	compileCommandMatch    *regexp.Regexp
	compileXibMatch        *regexp.Regexp
	compileStoryboardMatch *regexp.Regexp
}

func NewCompileParser() *CompileParser {
	// white foreground
	whiteColor := color.New(color.FgWhite)
	whiteColor.Add(color.Bold)

	return &CompileParser{
		whiteColor:             whiteColor,
		compileMatch:           regexp.MustCompile(COMPILE_MATCHER),
		compileCommandMatch:    regexp.MustCompile(COMPILE_COMMAND_MATCHER),
		compileXibMatch:        regexp.MustCompile(COMPILE_XIB_MATCHER),
		compileStoryboardMatch: regexp.MustCompile(COMPILE_STORYBOARD_MATCHER),
	}
}

func (parser *CompileParser) Match(line string, reader *bufio.Reader) bool {

	if match := parser.compileMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Println(color.WhiteString(fmt.Sprintf("%s %s %s", parser.whiteColor.Sprint("Compiling:"), match[2], match[1])))
		return true
	}

	if match := parser.compileCommandMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Println(color.WhiteString(fmt.Sprintf("%s %s %s", parser.whiteColor.Sprint("Compiling:"), match[2], match[1])))
		return true
	}

	if match := parser.compileXibMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Println(color.WhiteString(fmt.Sprintf("%s %s %s", parser.whiteColor.Sprint("Compiling:"), match[2], match[1])))
		return true
	}

	if match := parser.compileStoryboardMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Println(color.WhiteString(fmt.Sprintf("%s %s %s", parser.whiteColor.Sprint("Compiling:"), match[2], match[1])))
		return true
	}

	return false
}
