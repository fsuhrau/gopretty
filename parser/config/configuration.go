package config

import (
	"bufio"
	_const "github.com/fsuhrau/gopretty/const"
	"github.com/fsuhrau/gopretty/parser"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type ConfigurationParser struct {
	name             string
	exp              *regexp.Regexp
	color            *color.Color
	print            string
	multiline        bool
	multilineNumber  int
	multilineExp     *regexp.Regexp
	multilineEndline *string
	multilinePrint   string
}

func NewConfigurationParser(m matcher) parser.ParserInterface {
	c, ok := _const.Colors[m.Color]
	if !ok {
		c = _const.Colors["white"]
	}

	var multilineNumber int
	var multilineExp *regexp.Regexp
	var multilinePrint string
	var multilineEndline *string
	isMultiline := m.MultilineLines > 0 || len(m.MultilineRegex) > 0 || m.MultilineEndline != nil
	if isMultiline {
		multilineNumber = m.MultilineLines
		if m.MultilineRegex != "" {
			multilineExp, _ = regexp.Compile(m.MultilineRegex)
		}
		if m.MultilinePrint != "" {
			multilinePrint = m.MultilinePrint
		}
		if m.MultilineEndline != nil {
			multilineEndline = m.MultilineEndline
		}
	}

	return &ConfigurationParser{
		name:             m.Name,
		exp:              regexp.MustCompile(m.Regex),
		color:            c,
		print:            m.Print,
		multiline:        isMultiline,
		multilineNumber:  multilineNumber,
		multilineExp:     multilineExp,
		multilinePrint:   multilinePrint,
		multilineEndline: multilineEndline,
	}
}

func (parser *ConfigurationParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {

	if match := parser.exp.FindStringSubmatch(line); len(match) > 0 {
		if parser.print != "" {
			values := convertToExpressionsMap(parser.exp, match)
			_, _ = parser.color.Printf("%s\n", mapExpressionsToOutput(parser.print, values))
		} else {
			_, _ = parser.color.Printf("%s\n", strings.TrimRight(line, "\n"))
		}
		if parser.multiline {
			if parser.multilineNumber > 0 {
				for i := 0; i < parser.multilineNumber; i++ {
					text, err := reader.ReadString('\n')
					if err != nil {
						if err != io.EOF {
							os.Exit(1)
						}
						return true
					}

					_, _ = parser.color.Println(strings.TrimRight(text, "\n"))
				}
				return true
			} else if parser.multilineExp != nil {
				for {
					text, err := reader.ReadString('\n')
					if err != nil {
						if err != io.EOF {
							os.Exit(1)
						}
						return true
					}
					if subMatch := parser.multilineExp.FindStringSubmatch(text); len(subMatch) > 0 {
						if parser.multilinePrint != "" {
							subValues := convertToExpressionsMap(parser.multilineExp, subMatch)
							_, _ = parser.color.Printf("%s\n", mapExpressionsToOutput(parser.multilinePrint, subValues))
						} else {
							_, _ = parser.color.Printf("%s\n", strings.TrimRight(text, "\n"))
						}
						continue
					}
					overflowFunction(text)
					return true
				}
			} else if parser.multilineEndline != nil {
				for {
					text, err := reader.ReadString('\n')
					if err != nil {
						if err != io.EOF {
							os.Exit(1)
						}
						return true
					}
					text = strings.TrimRight(text, "\n")
					if text == *parser.multilineEndline {
						return true
					}
					_, _ = parser.color.Println(text)
				}
			}
			return true
		}
		return true
	}

	return false
}

func convertToExpressionsMap(exp *regexp.Regexp, match []string) map[string]string {
	paramsMap := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

func mapExpressionsToOutput(format string, values map[string]string) string {
	output := format
	for k, v := range values {
		output = strings.Replace(output, "{"+k+"}", v, -1)
	}
	return output
}
