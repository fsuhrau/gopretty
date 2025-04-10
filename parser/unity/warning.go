package unity

import (
	"bufio"
	"regexp"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

const (
	WARNING_MATCHER              = `(.*)(\(\d+,\d+\)):\swarning\s([A-Z0-9]+):\s(.*)`
	FIND_PLAYER_ASSEMBLY_WARNING = `Unable\sto\sfind\splayer\sassembly:.*`
	UNITY_WARNING                = `Warning!\s.*`
)

type WarningParser struct {
	color                  *color.Color
	warningMatch           *regexp.Regexp
	findPlayerWarningMatch *regexp.Regexp
	unityWarningMatch      *regexp.Regexp
}

func NewWarningParser() *WarningParser {
	yellow := color.New(color.FgYellow)
	yellow.Add(color.Bold)

	return &WarningParser{
		color:                  yellow,
		warningMatch:           regexp.MustCompile(WARNING_MATCHER),
		findPlayerWarningMatch: regexp.MustCompile(FIND_PLAYER_ASSEMBLY_WARNING),
		unityWarningMatch:      regexp.MustCompile(UNITY_WARNING),
	}
}

func (parser *WarningParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {

	if match := parser.warningMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s%s\n", emoji.Sprint(":warning: "), match[0])
		return true
	}

	if match := parser.findPlayerWarningMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s%s \n", emoji.Sprint(":warning: "), match[0])
		return true
	}

	if match := parser.unityWarningMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s%s \n", emoji.Sprint(":warning: "), match[0])
		return true
	}

	return false
}
