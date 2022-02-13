package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Pita",
	RunE: func(cmd *cobra.Command, args []string) error {
		v := "0.1.0"

		f, err := os.ReadFile("VERSION")
		if err != nil {
			return err
		}
		v = strings.TrimSuffix(string(f), "\n")
		fmt.Println(v)
		return nil
	},
}
