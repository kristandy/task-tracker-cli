package main

import (
	"fmt"
	"os"
	"task-tracker-cli/task"
	"strconv"
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
			fmt.Println("Harus tambahkan deskripsi!")
			return
		}
		deksripsiTugas := os.Args[2]
		fmt.Println("deksripsiTugas", deksripsiTugas)
		task.AddTask(deksripsiTugas)
	case "list":
		task.ReadTasks()
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Tambahkan deskripsi baru dan id!")
			return
		}
		idString := os.Args[2]
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("ID harus berupa angka")
			fmt.Println("Input yang diberikan: ", idString)
		}
		deksripsiTugasBaru := os.Args[3]
		task.UpdateTask(id, deksripsiTugasBaru)
	case "mark-done", "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Tambahkan status dan id!")
			return
		}
		idString := os.Args[2]
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("ID harus berupa angka")
			fmt.Println("Input yang diberikan: ", idString)
		}
		
		statusTugasBaru := os.Args[1]
		if statusTugasBaru == "mark-done" {
			statusTugasBaru = "done"
		} else {
			statusTugasBaru = "in-progress"
		}
		task.UpdateStatus(id, statusTugasBaru)
	}

}