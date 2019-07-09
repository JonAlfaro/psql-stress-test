package cmd

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the psql stress test",
	Long: `run the psql stress test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		zzz := viper.Get("Config.Name")
		fmt.Println(zzz)
		spew.Dump(cfg)
		var clearCondition bool
		for !clearCondition {
			clearCondition = true
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
