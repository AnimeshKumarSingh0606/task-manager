package cmd

import (
	"Gophercizes/task/students/animesh/task/db"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a task from DB",

	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range args {
			id, err := strconv.Atoi(a)
			if err == nil {
				t, err := db.FindTask(id)
				if err != nil {
					fmt.Println(err)
					return
				}
				var input string
				fmt.Printf("Really delete task [%s]? [y/n]", t.Desc)
				fmt.Scanln(&input)
				if strings.Trim(input, " ") != "y" {
					return
				}
				err = db.DeleteTask(id)
				if err == nil {
					fmt.Printf("task with id %d deleted\n", id)
				} else {
					fmt.Printf("Error when removing task with id %d : %s\n", id, err)
				}
			} else {
				fmt.Println("Bad task id Gadha kahinka \n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
