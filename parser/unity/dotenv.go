package unity

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

const (
	DOTENVFILE      = `DotEnvFile\s.*`
	DOTENV_OVERRIDE = `DotEnv\soverride.*`
	DOTENV_CONTAINS = `DotEnv\sfile\scontains:`
)

type DotEnvParser struct {
	color               *color.Color
	dotEnvMatch         *regexp.Regexp
	dotEnvOverrideMatch *regexp.Regexp
	dotEnvContainsMatch *regexp.Regexp
}

func NewDotEnvParser() *DotEnvParser {
	white := color.New(color.FgWhite)
	return &DotEnvParser{
		color:               white,
		dotEnvMatch:         regexp.MustCompile(DOTENVFILE),
		dotEnvOverrideMatch: regexp.MustCompile(DOTENV_OVERRIDE),
		dotEnvContainsMatch: regexp.MustCompile(DOTENV_CONTAINS),
	}
}

func (parser *DotEnvParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {

	if match := parser.dotEnvMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s\n", match[0])
		return true
	}

	if match := parser.dotEnvOverrideMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s \n", match[0])
		return true
	}

	if match := parser.dotEnvContainsMatch.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s \n", match[0])
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					os.Exit(1)
				}
				return true
			}
			if text == "\n" {
				return true
			}
			parser.color.Println(strings.TrimRight(text, "\n"))
		}
	}

	return false
}
