package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id          int `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func AddTask(desc string) {
	task := []Task{}
	status := "in-progress"
	zonaWaktu, _ := time.LoadLocation("Asia/Jakarta")
	waktuWib := time.Now().In(zonaWaktu)
	byteValue, err := os.ReadFile("task.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File json tidak ditemukan, buat file baru")
			dataJson := []byte{}
			writeJson := os.WriteFile("task.json", dataJson, 0644)
			if writeJson != nil {
				fmt.Println("Gagal saat membuat file json", writeJson)
			}
			fmt.Println("File json telah berhasil dibuat")
		} else {
			fmt.Println("Error saat baca file", err)
			return
		}
	} else {
		err = json.Unmarshal(byteValue, &task)
		if err != nil {
			fmt.Println("Gagal decode file JSON", err)
			return
		}
	}

	newTask := &Task{
		Id: 1,
		Description: desc,
		Status: status,
		CreatedAt: waktuWib,
		UpdatedAt: waktuWib,
	}

	task = append(task, *newTask)

	taskBytes, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Gagal encode", err)
	}

	err = os.WriteFile("task.json", taskBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func readTaskFromJSON() {
	var task Task
	byteValue, err := os.ReadFile("task.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File json tidak ditemukan, buat file baru")
			dataJson := []byte("{}")
			writeJson := os.WriteFile("task.json", dataJson, 0644)
			if writeJson != nil {
				fmt.Println("Gagal saat membuat file json", writeJson)
			}
			fmt.Println("File json telah berhasil dibuat")
		} else {
			fmt.Println("Error saat baca file", err)
			return
		}
	} else {
		err = json.Unmarshal(byteValue, &task)
		if err != nil {
			fmt.Println("Gagal decode file JSON", err)
			return
		}
	}
	
	fmt.Println(task.Id)
	fmt.Println(task.Description)
	fmt.Println(task.Status)
	fmt.Println(task.CreatedAt)
	fmt.Println(task.UpdatedAt)
}

func writeTaskToJSON() {}