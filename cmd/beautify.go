package cmd

import (
	"bufio"
	"fmt"
	"github.com/fsuhrau/gopretty/parser"
	"github.com/fsuhrau/gopretty/parser/config"
	"github.com/fsuhrau/gopretty/parser/unity"
	"github.com/fsuhrau/gopretty/parser/xcode"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// beautifyCmd represents the beautify command
var beautifyCmd = &cobra.Command{
	Use:   "beautify",
	Short: "read from stdout and beautify it based on rules",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var matcher []parser.ParserInterface
		if len(config.Config.Matcher) == 0 {
			// use old style matcher
			matcher = make([]parser.ParserInterface, 19)
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
		} else {
			matcher = make([]parser.ParserInterface, len(config.Config.Matcher))
			for i, m := range config.Config.Matcher {
				matcher[i] = config.NewConfigurationParser(m)
			}
		}

		var reader *bufio.Reader

		if len(args) > 0 {
			file, err := os.Open(args[0])
			if err != nil {
				log.Fatalf("Failed to open file: %s", err)
			}
			defer file.Close()
			reader = bufio.NewReader(file)
		} else {
			reader = bufio.NewReader(os.Stdin)
		}

		errLogger := log.New(os.Stderr, "", 0)

		//logger := log.New(os.Stdout, "", 0)
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
			for {
				overflow := false
				for _, m := range matcher {
					if m.Match(text, reader, func(overflowLine string) {
						overflow = true
						text = overflowLine
					}) {
						matches = true
						break
					}
				}

				if !overflow {
					break
				}
			}

			if Verbose && !matches {
				fmt.Printf("NOMA: %s", text)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(beautifyCmd)
}
