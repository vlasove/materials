package helpers

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/vlasove/materials/test_src/publisher/internal/app/models"
)

// Load ...
func Load(filePath string) (*models.Orders, error) {
	ords := new(models.Orders)
	file, err := os.Open(filePath)
	if err != nil {
		return ords, err
	}
	defer file.Close()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		return ords, err
	}

	err = json.Unmarshal(byteData, ords)
	if err != nil {
		return ords, err
	}
	return ords, nil

}

// UnmarhsallAndSleep ...
func UnmarhsallAndSleep(d models.Order) ([]byte, error) {
	v, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
	return v, err
}
