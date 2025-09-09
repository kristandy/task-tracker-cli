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
	var newId int
	task := []Task{}
	status := "todo"
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

	if len(task) == 0 {
		newId = 1
	} else {
		lastElement := task[len(task)-1]
		newId = lastElement.Id + 1
	}

	newTask := &Task{
		Id: newId,
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
	fmt.Printf("Task berhasil ditambahkan (ID: %d)", newId)
}

func ReadTasks() {
	var task []Task
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

	for _, values := range task {
		fmt.Println(values.Id)
		fmt.Println(values.Description)
		fmt.Println(values.Status)
		fmt.Println(values.CreatedAt)
		fmt.Println(values.UpdatedAt)

	}
}

func ReadTaskByStatus(status string) {
	var task []Task
	var listTask []Task
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
	
	for _, values := range task {
		if values.Status == status {
			listTaskBaru := append(listTask, values)
			for _, values := range listTaskBaru {
				fmt.Println(values.Id)
				fmt.Println(values.Description)
				fmt.Println(values.Status)
				fmt.Println(values.CreatedAt)
				fmt.Println(values.UpdatedAt)
			}
		} 
	}
}

func UpdateStatus (id int, newStatus string) {
	task := []Task{}
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
	targetIndex := -1
	for index, value := range task {
		if value.Id == id {
			targetIndex = index
		}
	}

	if targetIndex != -1 {
		task[targetIndex].Status = newStatus
		task[targetIndex].UpdatedAt = waktuWib
	}

	taskBytes, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Gagal encode", err)
	}

	err = os.WriteFile("task.json", taskBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Status task berhasil diupdate (ID: %d)", id)
}

func UpdateTask(id int, desc string) {
	task := []Task{}
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
	targetIndex := -1
	for index, value := range task {
		fmt.Println("task skrg", value)
		if value.Id == id {
			targetIndex = index
			fmt.Println("deskripsi baruy", desc)
			fmt.Println("yang dah direplace", value.Description)
		}
	}

	if targetIndex != -1 {
		task[targetIndex].Description = desc
		task[targetIndex].UpdatedAt = waktuWib
	}

	taskBytes, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Gagal encode", err)
	}

	err = os.WriteFile("task.json", taskBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Task berhasil diupdate (ID: %d)", id)
}