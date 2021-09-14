// - Finished compile Library/ScriptAssemblies/UnityEngine.TestRunner.dll in 4.451196 seconds

package unity

import (
	"bufio"
	"regexp"

	"github.com/fatih/color"
)

const (
	FINISHED_SCRIPT_COMPILE_MATCHER = `-\sFinished\scompile\s(.*)\sin\s(.*)\sseconds`
)

type ScriptCompiledParser struct {
	color   *color.Color
	matcher *regexp.Regexp
}

func NewScriptCompiledParser() *ScriptCompiledParser {
	white := color.New(color.FgWhite)
	return &ScriptCompiledParser{
		color:   white,
		matcher: regexp.MustCompile(FINISHED_SCRIPT_COMPILE_MATCHER),
	}
}

func (parser *ScriptCompiledParser) Match(line string, reader *bufio.Reader) bool {

	if match := parser.matcher.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("Script Compiled: %s %s seconds\n", match[1], match[2])
		return true
	}

	return false
}
