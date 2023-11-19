package cmd

import (
	"fmt"
	"os"

	"github.com/chethan0707/gowc/internals/helpers"
	"github.com/chethan0707/gowc/internals/interfaces"
	"github.com/spf13/cobra"
)

var (
	lineFlag bool
	wordFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "gowc",
	Short: "wc is a word, line, and character count tool",
	Long:  `wc is a word, line, and character count tool that reads from the standard input or from a file and outputs the number of lines, words, and characters`,

	Run: func(_ *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No File name passed")
			os.Exit(1)
		}

		var output string
		var result interfaces.Result = helpers.ReadTextFile(args[0])

		if lineFlag {
			output += fmt.Sprintf("%8d", result.Lines)
		}

		// if wordFlag {
		// 	output += helpers.ReadTextFile(args[0])
		// }

		output += fmt.Sprintf("  %s", args[0])
		fmt.Println(output)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&lineFlag, "lines", "l", false, "Count number of lines")
	rootCmd.Flags().BoolVarP(&wordFlag, "words", "w", false, "Count number of words")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
