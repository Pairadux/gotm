package storage

// IMPORTS {{{
import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Pairadux/gotm/internal/models"
	"github.com/spf13/viper"
) // }}}

func PrintJson(allTasks *models.TaskList) {
	data, err := json.MarshalIndent(allTasks, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func ToJson(allTasks *models.TaskList) []byte {
	data, err := json.Marshal(allTasks)
	if err != nil {
		panic(err)
	}
	return data
}

func SaveTasksToFile(filename string, allTasks *models.TaskList) error {
	return os.WriteFile(filename, ToJson(allTasks), 0o644)
}

func LoadTasks() (*models.TaskList, error) {
	filename := viper.GetString("json_path")

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	t := models.TaskList{}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
