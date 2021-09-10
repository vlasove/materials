package filemanager

import (
	"encoding/json"
	"io/ioutil"

	"github.com/vlasove/materials/test_src/generator/internal/app/models"
)

// FileManager ...
type FileManager struct {
	fileName string
}

// New ...
func New(fname string) *FileManager {
	return &FileManager{fname}
}

// Dump ...
func (fm *FileManager) Dump(data []*models.Order) error {
	m := map[string][]*models.Order{"orders": data}
	byteDate, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fm.fileName, byteDate, 0644)
	if err != nil {
		return err
	}
	return nil
}
