package cmd

import (
	"Gophercizes/task/students/animesh/task/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Update task status to done",

	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range args {
			id, err := strconv.Atoi(a)
			if err == nil {
				e := db.DoneTask(id)
				if e == nil {
					fmt.Printf("task with id %d done\n", id)
				} else {
					fmt.Printf("Error when doing task with id %d : %s\n", id, e)
				}
			} else {
				fmt.Println("Bad task id , Again a Gadha \n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
