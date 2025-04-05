package xcode

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

const (
	// errors
	COMPILE_ERROR_MATCHER                    = `(\/.+\/(.*):.*:.*):\s(?:fatal\s)?error:\s(.*)`
	CLANG_ERROR_MATCHER                      = `(clang: error:.*)`
	CHECK_DEPENDENCIES_ERRORS_MATCHER        = `(Code\s?Sign error:.*|Code signing is required for product type .* in SDK .*|No profile matching .* found:.*|Provisioning profile .* doesn't .*|Swift is unavailable on .*|.?Use Legacy Swift Language Version.*)`
	PROVISIONING_PROFILE_REQUIRED_MATCHER    = `(.*requires a provisioning profile.*)`
	NO_CERTIFICATE_MATCHER                   = `(No certificate matching.*)`
	FATAL_ERROR_MATCHER                      = `(fatal error:.*)`
	FILE_MISSING_ERROR_MATCHER               = `<unknown>:0:\s(error:\s.*)\s'(\/.+\/.*\..*)'`
	LINKER_DUPLICATE_SYMBOLS_MATCHER         = `(duplicate symbol .*):`
	LINKER_UNDEFINED_SYMBOL_LOCATION_MATCHER = `(.* in .*\.o)`
	LINKER_UNDEFINED_SYMBOLS_MATCHER         = `(Undefined symbols for architecture .*):`
	PODS_ERROR_MATCHER                       = `(error:\s.*)`
	SYMBOL_REFERENCED_FROM_MATCHER           = `/\s+"(.*)", referenced from:`
	MODULE_INCLUDES_ERROR_MATCHER            = `\<module-includes\>:.*?:.*?:\s(?:fatal\s)?(error:\s.*)`
)

type ErrorParser struct {
	redColor                              *color.Color
	compileErrorMatcher                   *regexp.Regexp
	clangErrorMatcher                     *regexp.Regexp
	dependenciesErrorMatcher              *regexp.Regexp
	provisioningProfileMatcher            *regexp.Regexp
	noCertificateMatcher                  *regexp.Regexp
	fatalErrorMatcher                     *regexp.Regexp
	fileMissingMatcher                    *regexp.Regexp
	linkerDuplicateSymbolsMatcher         *regexp.Regexp
	linkerUndefinedSymbolsLocationMatcher *regexp.Regexp
	linkerUndefinedSymbolsMatcher         *regexp.Regexp
	podsErrorMatcher                      *regexp.Regexp
	symboleReferencedMatcher              *regexp.Regexp
	modulesIncludesErrorMatcher           *regexp.Regexp
}

func NewErrorParser() *ErrorParser {
	red := color.New(color.FgRed)
	red.Add(color.Bold)

	return &ErrorParser{
		redColor:                              red,
		compileErrorMatcher:                   regexp.MustCompile(COMPILE_ERROR_MATCHER),
		clangErrorMatcher:                     regexp.MustCompile(CLANG_ERROR_MATCHER),
		dependenciesErrorMatcher:              regexp.MustCompile(CHECK_DEPENDENCIES_ERRORS_MATCHER),
		provisioningProfileMatcher:            regexp.MustCompile(PROVISIONING_PROFILE_REQUIRED_MATCHER),
		noCertificateMatcher:                  regexp.MustCompile(NO_CERTIFICATE_MATCHER),
		fatalErrorMatcher:                     regexp.MustCompile(FATAL_ERROR_MATCHER),
		fileMissingMatcher:                    regexp.MustCompile(FILE_MISSING_ERROR_MATCHER),
		linkerDuplicateSymbolsMatcher:         regexp.MustCompile(LINKER_DUPLICATE_SYMBOLS_MATCHER),
		linkerUndefinedSymbolsLocationMatcher: regexp.MustCompile(LINKER_UNDEFINED_SYMBOL_LOCATION_MATCHER),
		linkerUndefinedSymbolsMatcher:         regexp.MustCompile(LINKER_UNDEFINED_SYMBOLS_MATCHER),
		podsErrorMatcher:                      regexp.MustCompile(PODS_ERROR_MATCHER),
		symboleReferencedMatcher:              regexp.MustCompile(SYMBOL_REFERENCED_FROM_MATCHER),
		modulesIncludesErrorMatcher:           regexp.MustCompile(MODULE_INCLUDES_ERROR_MATCHER),
	}
}

func (parser *ErrorParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {

	if match := parser.compileErrorMatcher.FindStringSubmatch(line); len(match) > 0 {

		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])

		// read 2 additonal lines
		for i := 0; i < 2; i++ {
			text, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					os.Exit(1)
				}
				return true
			}
			parser.redColor.Println(strings.TrimRight(text, "\n"))
		}
		return true
	}

	if match := parser.clangErrorMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.dependenciesErrorMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.provisioningProfileMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.noCertificateMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.fatalErrorMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.fileMissingMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.linkerDuplicateSymbolsMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.linkerUndefinedSymbolsLocationMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.linkerUndefinedSymbolsMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.podsErrorMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.symboleReferencedMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}
	if match := parser.modulesIncludesErrorMatcher.FindStringSubmatch(line); len(match) > 0 {
		parser.redColor.Printf("%s%s\n", emoji.Sprint(":x: "), match[0])
		return true
	}

	return false
}
