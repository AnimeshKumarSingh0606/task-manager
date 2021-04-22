package main

import (
	"Gophercizes/task/students/animesh/task/cmd"
	"Gophercizes/task/students/animesh/task/db"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gophercizes/go-homedir"
)

func main() {
	h, err := homedir.Dir()
	must(err)
	home := filepath.Join(h, "tasks.db")
	must(db.InitDb(home))
	defer db.CloseDb()

	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
