package main

import (
	"fmt"
	"os"
	"task-tracker-cli/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Perintah tidak ditemukan")
		return
	}
	perintah := os.Args[1]

	switch perintah {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Description must be added!")
			return
		}
		deksripsiTugas := os.Args[2]
		task.AddTask(deksripsiTugas)
	case "list":
		task.ReadTaskFromJSON()
	}

}