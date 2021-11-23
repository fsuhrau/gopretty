package unity

import (
	"bufio"
	"regexp"

	"github.com/fatih/color"
)

const (
	PLATFROM_BUILD_MESSAGE = `={2,3}\s.*`
	PROGRESS_BAR           = `DisplayProgressbar:\s.*`
	SCRIPT_COMPILATION     = `\[ScriptCompilation\]\s.*`
	EXITING_BATCHMODE      = `Exiting\sbatchmode.*`
)

type BuildProcessParser struct {
	color                  *color.Color
	buildMatch             *regexp.Regexp
	progressBarMatch       *regexp.Regexp
	scriptCompilationMatch *regexp.Regexp
	exitingBatchmodeMatch  *regexp.Regexp
}

func NewBuildProcessParser() *BuildProcessParser {
	white := color.New(color.FgWhite)
	return &BuildProcessParser{
		color:                  white,
		buildMatch:             regexp.MustCompile(PLATFROM_BUILD_MESSAGE),
		progressBarMatch:       regexp.MustCompile(PROGRESS_BAR),
		scriptCompilationMatch: regexp.MustCompile(SCRIPT_COMPILATION),
		exitingBatchmodeMatch:  regexp.MustCompile(EXITING_BATCHMODE),
	}
}

func (parser *BuildProcessParser) Match(line string, reader *bufio.Reader) bool {

	if match := parser.buildMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("Build: %s \n", match[0])
		return true
	}

	if match := parser.progressBarMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("Build: %s \n", match[0])
		return true
	}

	if match := parser.scriptCompilationMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("Build: %s \n", match[0])
		return true
	}

	if match := parser.exitingBatchmodeMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("Build: %s \n", match[0])
		return true
	}

	return false
}
