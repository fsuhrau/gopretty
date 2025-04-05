package xcode

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

const (
	// 1 =
	CODESIGN_MATCHER = `CodeSign\s((?:\\ |[^ ])*)+`
	// 1 = framework name
	CODESIGN_FRAMEWORK_MATCHER = `CodeSign\s((?:\\ |[^ ])*.framework)\/Versions/`

	// 1 = signing identity
	CODESIGN_IDENTITY_MATCHER = `\s*CODE_SIGN_IDENTITY\s=\s(.*)`
	// 1 = provisioning profile
	PROVISIONING_PROFILE_MATCHER = `\s*PROVISIONING_PROFILE\s=\s(.*)`

	// 1 = signing identity
	SIGNING_IDENTITY_MATCHER = `Signing Identity:\s*"(.*)"`
	// 1 = provisioning profile
	PROV_PROFILE_MATCHER = `Provisioning Profile:\s*"(.*)"`
)

type SigningParser struct {
	whiteColor               *color.Color
	codeSignMatch            *regexp.Regexp
	codeSignFrameworkMatch   *regexp.Regexp
	signingIdentityMatch     *regexp.Regexp
	provisioningProfileMatch *regexp.Regexp
	signIdentMatch           *regexp.Regexp
	provProfileMatch         *regexp.Regexp
}

func NewSigningParser() *SigningParser {
	// white foreground
	whiteColor := color.New(color.FgWhite)
	whiteColor.Add(color.Bold)

	return &SigningParser{
		whiteColor:               whiteColor,
		codeSignMatch:            regexp.MustCompile(CODESIGN_MATCHER),
		codeSignFrameworkMatch:   regexp.MustCompile(CODESIGN_FRAMEWORK_MATCHER),
		signingIdentityMatch:     regexp.MustCompile(CODESIGN_IDENTITY_MATCHER),
		provisioningProfileMatch: regexp.MustCompile(PROVISIONING_PROFILE_MATCHER),
		signIdentMatch:           regexp.MustCompile(SIGNING_IDENTITY_MATCHER),
		provProfileMatch:         regexp.MustCompile(PROV_PROFILE_MATCHER),
	}
}

func (parser *SigningParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {
	if match := parser.codeSignMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s '%s'\n", parser.whiteColor.Sprint("CodeSign:"), match[1])
		return true
	}

	if match := parser.codeSignFrameworkMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("%s '%s'\n", parser.whiteColor.Sprint("CodeSign Framework:"), match[1])
		return true
	}

	if match := parser.signingIdentityMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("Using Signing Identity '%s'\n", parser.whiteColor.Sprintf(match[1]))
		return true
	}

	if match := parser.provisioningProfileMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("Using Provisioning Profile '%s'\n", parser.whiteColor.Sprintf(match[1]))
		return true
	}

	if match := parser.signIdentMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("Signing Identity '%s'\n", parser.whiteColor.Sprintf(match[1]))
		return true
	}

	if match := parser.provProfileMatch.FindStringSubmatch(line); len(match) > 0 {
		fmt.Printf("Provisioning Profile '%s'\n", parser.whiteColor.Sprintf(match[1]))
		// read 1 additonal line
		for i := 0; i < 1; i++ {
			text, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					os.Exit(1)
				}
				return true
			}
			parser.whiteColor.Println(strings.TrimRight(text, "\n"))
		}
		return true
	}
	return false
}
