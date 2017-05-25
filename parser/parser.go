package parser

import "bufio"

type ParserInterface interface {
	// Match the frist line and read more required data when needed
	Match(line string, reader *bufio.Reader) bool
}
