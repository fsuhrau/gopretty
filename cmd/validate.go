package cmd

import (
	"fmt"
	_const "github.com/fsuhrau/gopretty/const"
	"github.com/fsuhrau/gopretty/parser/config"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate a configuration file",
	Long:  `validate a configuration file and make sure the provided regex paser work`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, matcher := range config.Config.Matcher {
			_, err := regexp.Compile(matcher.Regex)
			if err != nil {
				fmt.Printf("[ERROR] \"%s\" regex: %v\n", matcher.Name, err)
				continue
			}

			if len(matcher.Color) > 0 {
				if _, ok := _const.Colors[matcher.Color]; !ok {
					fmt.Printf("[ERROR] \"%s\" color: %v does not exist\n", matcher.Name, matcher.Color)
					continue
				}
			}
			if len(matcher.MultilineRegex) > 0 {
				_, err := regexp.Compile(matcher.MultilineRegex)
				if err != nil {
					fmt.Printf("[ERROR] \"%s\" multiline_regex: %v\n", matcher.Name, err)
					continue
				}
			}
			if matcher.MultilineLines < 0 {
				fmt.Printf("[ERROR] \"%s\" multiline_lines: Lines Must be positive\n", matcher.Name)
				continue
			}

			if matcher.MultilineEndline != nil {
				if strings.Contains(*matcher.MultilineEndline, "\\n") {
					fmt.Printf("[ERROR] \"%s\" multiline_endline: Newlines are not supported\n", matcher.Name)
					continue
				}
			}

			fmt.Printf("[OK] \"%s\"\n", matcher.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
