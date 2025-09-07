package task

import (
	"encoding/json"
	"fmt"
	"os"
	"errors"
)

type Task struct {
	id          int `json:"id"`
	description string `json:"description"`
	status      string `json:"status"`
	createdAt   string `json:"createdAt"`
	updatedAt   string `json:"updatedAt"`
}
func readTaskFromJSON() {
	var tasks Task
	byteValue, err := os.ReadFile("task.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File json tidak ditemukan, buat file baru")
			dataJson := []byte("{}")
			writeJson := os.WriteFile("task.json", "dataJson", 0644)
			if writeJson != nil {
				fmt.Println("Gagal saat membuat file json", writeJson)
			}
			fmt.Println("File json telah berhasil dibuat")
		} else {
			fmt.Println("Error saat baca file", err)
			return
		}
	} else {
		err = json.Unmarshal(byteValue, &tasks)
		if err != nil {
			fmt.Println("Gagal decode file JSON", err)
			return
		}
	}
	
	fmt.Println(tasks.id)
	fmt.Println(tasks.description)
	fmt.Println(tasks.status)
	fmt.Println(tasks.createdAt)
	fmt.Println(tasks.updatedAt)

}
