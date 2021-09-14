package unity

import (
	"bufio"
	"regexp"

	"github.com/fatih/color"
)

const (
	IMPORT_ASSET_MATCHER = `Done\simporting\sasset:\s'(.*)'\s\(target\shash:\s'([a-zA-Z0-9]+)'\)\sin\s(.*)\sseconds`
)

type ImportAssetParser struct {
	color   *color.Color
	matcher *regexp.Regexp
}

func NewImportAssetParser() *ImportAssetParser {
	white := color.New(color.FgWhite)
	return &ImportAssetParser{
		color:   white,
		matcher: regexp.MustCompile(IMPORT_ASSET_MATCHER),
	}
}

func (parser *ImportAssetParser) Match(line string, reader *bufio.Reader) bool {

	if match := parser.matcher.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("Import asset: %s (%s) %s seconds\n", match[1], match[2], match[3])
		return true
	}

	return false
}
