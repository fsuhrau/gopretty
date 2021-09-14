package unity

import (
	"bufio"
	"regexp"

	"github.com/fatih/color"
)

const (
	FINISHED_ILPOSTPROCESSOR_MATCHER = `-\sFinished\sILPostProcessor\s'(.*)'\son\s(.*)\sin\s(.*)\sseconds`
)

type FinishedILPostProcessor struct {
	color   *color.Color
	matcher *regexp.Regexp
}

func NewFinishedILPostProcessor() *FinishedILPostProcessor {
	white := color.New(color.FgWhite)
	return &FinishedILPostProcessor{
		color:   white,
		matcher: regexp.MustCompile(FINISHED_ILPOSTPROCESSOR_MATCHER),
	}
}

func (parser *FinishedILPostProcessor) Match(line string, reader *bufio.Reader) bool {

	if match := parser.matcher.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("ILPostProcessor: %s on %s %s seconds\n", match[1], match[2], match[3])
		return true
	}

	return false
}
