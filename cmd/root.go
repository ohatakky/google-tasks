package cmd

import (
	"fmt"
	"os"

	"github.com/ohatakky/tasks/pkg"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "taks",
	Long:  `taks`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	pkg.InitConfig()
	pkg.InitTasks()
}
