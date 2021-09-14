package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fsuhrau/gopretty/parser"
	"github.com/fsuhrau/gopretty/parser/unity"
	"github.com/fsuhrau/gopretty/parser/xcode"
)

func main() {
	debug := false
	if len(os.Args) > 1 {
		debug = os.Args[1] == "debug"
	}
	matcher := make([]parser.ParserInterface, 19)
	matcher[0] = xcode.NewInfoParser()
	matcher[1] = xcode.NewSigningParser()
	matcher[2] = xcode.NewCompileParser()
	matcher[3] = xcode.NewStatusParser()
	matcher[4] = xcode.NewErrorParser()
	matcher[5] = xcode.NewWarningParser()
	matcher[6] = xcode.NewLinkerParser()
	matcher[7] = xcode.NewPackagingParser()
	matcher[8] = xcode.NewTestParser()
	matcher[9] = xcode.NewCopyParser()
	matcher[10] = xcode.NewPhasesParser()
	matcher[11] = unity.NewErrorParser()
	matcher[12] = unity.NewWarningParser()
	matcher[13] = unity.NewRemoteAssetCacheParser()
	matcher[14] = unity.NewScriptCompiledParser()
	matcher[15] = unity.NewFinishedILPostProcessor()
	matcher[16] = unity.NewImportAssetParser()
	matcher[17] = unity.NewBuildProcessParser()
	matcher[18] = unity.NewDotEnvParser()

	errLogger := log.New(os.Stderr, "", 0)
	//logger := log.New(os.Stdout, "", 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				errLogger.Println(err.Error())
				os.Exit(1)
			}
			return
		}

		matches := false
		for _, m := range matcher {
			if m.Match(text, reader) {
				matches = true
				break
			}
		}
		if debug && !matches {
			fmt.Printf("NOMA: %s", text)
		}
	}
}
