package cmd

import (
	"Gophercizes/task/students/animesh/task/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the TODO list",

	Run: func(cmd *cobra.Command, args []string) {
		s := strings.Join(args, " ")
		id, err := db.AddTask(s)
		if err != nil {
			fmt.Printf("Unable to add task %s : %v\n", s, err)
		}
		fmt.Println("Added", s, "with id", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
