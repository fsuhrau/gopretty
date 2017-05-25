package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/fsuhrau/gopretty/parser"
)

func main() {
	matcher := make([]parser.ParserInterface, 11)
	matcher[0] = parser.NewInfoParser()
	matcher[1] = parser.NewSigningParser()
	matcher[2] = parser.NewCompileParser()
	matcher[3] = parser.NewStatusParser()
	matcher[4] = parser.NewErrorParser()
	matcher[5] = parser.NewWarningParser()
	matcher[6] = parser.NewLinkerParser()
	matcher[7] = parser.NewPackagingParser()
	matcher[8] = parser.NewTestParser()
	matcher[9] = parser.NewCopyParser()
	matcher[10] = parser.NewPhasesParser()

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

		for _, m := range matcher {
			if m.Match(text, reader) {
				break
			}
		}
	}
}
