package unity

import (
	"bufio"
	"regexp"

	"github.com/fatih/color"
)

const (
	LICENSE_CLIENT = `[LicensingClient|Licensing::Module|]\s.*`
	LICENSE_SYSTEM = `LICENSE SYSTEM.*`
)

type EditorParser struct {
	color              *color.Color
	licenseClientMatch *regexp.Regexp
	licenseSystemMatch *regexp.Regexp
}

func NewEditorParser() *EditorParser {
	white := color.New(color.FgWhite)
	return &EditorParser{
		color:              white,
		licenseClientMatch: regexp.MustCompile(LICENSE_CLIENT),
		licenseSystemMatch: regexp.MustCompile(LICENSE_SYSTEM),
	}
}

func (parser *EditorParser) Match(line string, reader *bufio.Reader) bool {

	if match := parser.licenseClientMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s\n", match[0])
		return true
	}

	if match := parser.licenseSystemMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s \n", match[0])
		return true
	}

	return false
}
