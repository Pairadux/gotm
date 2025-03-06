/* LICENSE {{{
Copyright © 2025 Austin Gause <a.gause@outlook.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/ // }}}

package storage

// IMPORTS {{{
import (
	"encoding/json"
	"io"
	"os"

	"github.com/Pairadux/gotm/internal/models"
) // }}}

func ToJson(taskState *models.TaskState) []byte {
	data, err := json.Marshal(taskState)
	if err != nil {
		panic(err)
	}
	return data
}

func SaveTasksToFile(filename string, taskState *models.TaskState) error {
	return os.WriteFile(filename, ToJson(taskState), 0o644)
}

func Load(filename string) (*models.TaskState, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return &models.TaskState{
			Workspaces: make(map[string]*models.Workspace),
		}, nil
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

	var taskState models.TaskState
	if err := json.Unmarshal(data, &taskState); err != nil {
		return nil, err
	}

	return &taskState, nil
}
