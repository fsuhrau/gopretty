package xcode

import (
	"bufio"
	"fmt"
	"regexp"
)

const (
	SHELL_COMMAND_MATCHER = `\s{4}(cd|setenv|(?:[\w\/:\\\s\-.]+?\/)?[\w\-]+)\s(.*)`
)

type ShellParser struct {
	shellMatch *regexp.Regexp
}

func NewShellParser() *ShellParser {
	return &ShellParser{
		shellMatch: regexp.MustCompile(SHELL_COMMAND_MATCHER),
	}
}

func (parser *ShellParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {
	if match := parser.shellMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Println(match[0])
		return true
	}

	return false
}
