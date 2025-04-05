package xcode

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

const (
	PHASE_SCRIPT_EXECUTION_MATCHER = `PhaseScriptExecution\s((\\\ |\S)*)\s`
	PROCESS_PCH_MATCHER            = `ProcessPCH.*\s.*\s(.*.pch)`
	// PROCESS_PCH_COMMAND_MATCHER    = `\s*.*\/usr\/bin\/clang\s.*\s\-c\s(.*)\s\-o\s.*`
	PREPROCESS_MATCHER         = `Preprocess\s(?:(?:\\ |[^ ])*)\s((?:\\ |[^ ])*)`
	PBXCP_MATCHER              = `PBXCp\s((?:\\ |[^ ])*)`
	PROCESS_INFO_PLIST_MATCHER = `ProcessInfoPlistFile\s.*\.plist\s(.*\/+(.*\.plist))`

	TIFFUTIL_MATCHER      = `TiffUtil\s(.*)`
	TOUCH_MATCHER         = `Touch\s(.*\/(.+))`
	WRITE_FILE_MATCHER    = `write-file\s(.*)`
	WRITE_AUXILIARY_FILES = `Write auxiliary files`
)

type PhasesParser struct {
	whiteColor       *color.Color
	phaseScriptMatch *regexp.Regexp
	processPchMatch  *regexp.Regexp
	// processPchCommandMatch *regexp.Regexp
	preprocessMatch       *regexp.Regexp
	pbxcpMatch            *regexp.Regexp
	processInfoPListMatch *regexp.Regexp
	tiffUtilMatch         *regexp.Regexp
	touchMatch            *regexp.Regexp
	writeFileMatch        *regexp.Regexp
	writeAuxilaryMatch    *regexp.Regexp
}

func NewPhasesParser() *PhasesParser {
	// white foreground
	whiteColor := color.New(color.FgWhite)
	whiteColor.Add(color.Bold)

	return &PhasesParser{
		whiteColor:       whiteColor,
		phaseScriptMatch: regexp.MustCompile(PHASE_SCRIPT_EXECUTION_MATCHER),
		processPchMatch:  regexp.MustCompile(PROCESS_PCH_MATCHER),
		// processPchCommandMatch: regexp.MustCompile(PROCESS_PCH_COMMAND_MATCHER),
		preprocessMatch:       regexp.MustCompile(PREPROCESS_MATCHER),
		pbxcpMatch:            regexp.MustCompile(PBXCP_MATCHER),
		processInfoPListMatch: regexp.MustCompile(PROCESS_INFO_PLIST_MATCHER),
		tiffUtilMatch:         regexp.MustCompile(TIFFUTIL_MATCHER),
		touchMatch:            regexp.MustCompile(TOUCH_MATCHER),
		writeFileMatch:        regexp.MustCompile(WRITE_FILE_MATCHER),
		writeAuxilaryMatch:    regexp.MustCompile(WRITE_AUXILIARY_FILES),
	}
}

func (parser *PhasesParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {

	if match := parser.phaseScriptMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s %s\n", parser.whiteColor.Sprint("PhaseScriptExecution:"), match[1])
		return true
	}
	if match := parser.processPchMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s %s\n", parser.whiteColor.Sprint("ProcessPCH:"), match[1])
		return true
	}
	// if match := parser.processPchCommandMatch.FindStringSubmatch(line); len(match) > 0 {
	// 	fmt.Printf("3 %s\n", match[0])
	// 	return true
	// }
	if match := parser.preprocessMatch.FindStringSubmatch(line); len(match) > 0 {
		// fmt.Printf("%s %s\n", parser.whiteColor.Sprint("ProcessPCH:"), match[1])
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.pbxcpMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", match[0])
		return true
	}
	if match := parser.processInfoPListMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s %s\n", parser.whiteColor.Sprint("ProcessInfoPlistFile:"), match[1])
		return true
	}

	if match := parser.tiffUtilMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s %s\n", parser.whiteColor.Sprint("TiffUtil:"), match[1])
		return true
	}

	if match := parser.touchMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s %s\n", parser.whiteColor.Sprint("Touch:"), match[1])
		return true
	}

	if match := parser.writeFileMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s %s\n", parser.whiteColor.Sprint("WriteFile:"), match[1])
		return true
	}

	if match := parser.writeAuxilaryMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s\n", parser.whiteColor.Sprint("WriteAuxilary:"))
		return true
	}

	return false
}
