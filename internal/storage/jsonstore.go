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

func PrintJson(workspaces map[string]*models.Workspace) {
	data, err := json.MarshalIndent(workspaces, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func ToJson(workspaces map[string]*models.Workspace) []byte {
	data, err := json.Marshal(workspaces)
	if err != nil {
		panic(err)
	}
	return data
}

func SaveTasksToFile(filename string, workspaces map[string]*models.Workspace) error {
	return os.WriteFile(filename, ToJson(workspaces), 0o644)
}

func LoadWorkspaces() (map[string]*models.Workspace, error) {
	filename := viper.GetString("json_path")

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return make(map[string]*models.Workspace), nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var workspaces map[string]*models.Workspace
	if err := json.Unmarshal(data, &workspaces); err != nil {
		return nil, err
	}

	return workspaces, nil
}
