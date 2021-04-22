package cmd

import (
	"Gophercizes/task/students/animesh/task/db"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the task description, effort (-e), urgency (-u), criticality (-c)",

	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range args {
			id, err := strconv.Atoi(a)
			if err == nil {
				t, err := db.FindTask(id)
				if err != nil {
					fmt.Println(err)
					return
				}
				if effort >= 0 {
					e := db.UpdateEffort(id, effort)
					ack(id, e)
				}
				if criticality >= 0 {
					e := db.UpdateCriticality(id, criticality)
					ack(id, e)
				}
				if urgency >= 0 {
					e := db.UpdateUrgency(id, urgency)
					ack(id, e)
				}
				if effort >= 0 || urgency >= 0 || criticality >= 0 {
					return
				}
				
				var input string
				fmt.Println("Enter new description for :", t.Desc)
				fmt.Scanf("%q", &input)
				if strings.Trim(input, " ") == "" {
					return
				}
				e := db.UpdateTask(id, input)
				ack(id, e)
			} else {
				fmt.Println("Bad task id , HuH! man you are a down and out Gadha \n", id)
			}
		}
	},
}

func ack(id int, e error) {
	if e == nil {
		fmt.Printf("task with id %d updated\n", id)
	} else {
		fmt.Printf("Error when updating task with id %d : %s\n", id, e)
	}
}

var effort, criticality, urgency int

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&effort, "effort", "e", -1, "The task effort (0 = easy < 1 = complex < 2 = complicate < 3 = complex and complicate)")
	updateCmd.Flags().IntVarP(&criticality, "criticality", "c", -1, "The task criticality (Criticality 0 = wished < 1 = wanted < 2 = needed)")
	updateCmd.Flags().IntVarP(&urgency, "urgency", "u", -1, "The task urgency (0 = intime < 1 = waited < 2 = urgent)")
}
