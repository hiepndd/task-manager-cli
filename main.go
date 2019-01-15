package main

import (
	"fmt"
	"os"
	"path/filepath"

	"git.dwarvesf.com/task-manager-cli/cmd"
	"git.dwarvesf.com/task-manager-cli/db"
	homedir "github.com/mitchellh/go-homedir"
)

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}
