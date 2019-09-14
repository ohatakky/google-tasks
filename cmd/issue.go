package cmd

import (
	"log"
	"time"

	"github.com/ohatakky/tasks/pkg"
	"github.com/spf13/cobra"
	"google.golang.org/api/tasks/v1"
)

var (
	issueCmd = &cobra.Command{
		Use:   "issue",
		Short: "issue",
		Long:  `issue`,
		Run: func(cmd *cobra.Command, args []string) {
			err := issue(args)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	title    string
	afterDay int
)

func init() {
	issueCmd.Flags().StringVarP(&title, "title", "t", "", "")
	issueCmd.Flags().IntVarP(&afterDay, "afterDay", "d", 1, "")
	rootCmd.AddCommand(issueCmd)
}

func issue(args []string) error {
	mylistID := pkg.Config.ListID
	due := time.Now().AddDate(0, 0, afterDay).Format(time.RFC3339)
	task := &tasks.Task{
		Title: title,
		Due:   due,
	}
	_, err := pkg.Service.Tasks.Insert(mylistID, task).Do()
	if err != nil {
		return err
	}

	return nil
}
