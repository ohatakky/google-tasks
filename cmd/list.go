package cmd

import (
	"fmt"
	"log"

	"github.com/ohatakky/tasks/pkg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `list`,
	Run: func(cmd *cobra.Command, args []string) {
		err := list()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() error {
	list, err := pkg.Service.Tasks.List(pkg.Config.ListID).Do()
	if err != nil {
		return err
	}

	for _, v := range list.Items {
		fmt.Printf("%s\n", v.Title)
	}

	return nil
}
