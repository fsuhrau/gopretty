package unity

import (
	"bufio"
	"regexp"

	"github.com/fatih/color"
)

const (
	REMOTE_ASSET_CACHE_MATCHER = `RemoteAssetCache\s-\s.*\s-\ssuccess:(true|false),\s.*key:([a-zA-Z0-9]+)\s.*extension:'(.*)',\stime\selapsed:\s(.*)\sseconds`
)

type RemoteAssetCacheParser struct {
	color   *color.Color
	matcher *regexp.Regexp
}

func NewRemoteAssetCacheParser() *RemoteAssetCacheParser {
	white := color.New(color.FgWhite)
	return &RemoteAssetCacheParser{
		color:   white,
		matcher: regexp.MustCompile(REMOTE_ASSET_CACHE_MATCHER),
	}
}

func (parser *RemoteAssetCacheParser) Match(line string, reader *bufio.Reader, overflowFunction func(overflowLine string)) bool {

	if match := parser.matcher.FindStringSubmatch(line); len(match) > 0 {
		parser.color.Printf("%s %s %s success: %s\n", match[2], match[3], match[4], match[1])
		return true
	}

	return false
}
