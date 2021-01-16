package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	tplFile string
	valFile string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&tplFile, "template-file", "t", "", "template file")
	rootCmd.MarkFlagRequired("template-file")
	rootCmd.PersistentFlags().StringVarP(&valFile, "value-file", "v", "", "value file")
	rootCmd.MarkFlagRequired("value-file")
}

var rootCmd = &cobra.Command{
	Use:   "spade",
	Short: "Spade implements data-driven templates for generating textual output.",
	Run: func(cmd *cobra.Command, args []string) {
		val, err := parseVal()
		if err != nil {
			panic(err)
		}

		renderTpl(val)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}