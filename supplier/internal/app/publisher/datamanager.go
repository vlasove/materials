package publisher

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/vlasove/materials/supplier/internal/app/models"
)

type OrderReader struct {
	file   *os.File
	orders []*models.Order
}

func NewOrderReader() *OrderReader {
	return &OrderReader{
		orders: []*models.Order{},
	}
}

func (o *OrderReader) Open(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	o.file = file
	return nil
}

func (o *OrderReader) Unmarshal() error {
	byteData, err := ioutil.ReadAll(o.file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, &o.orders)
	if err != nil {
		return err
	}
	return nil
}

func toByte(order *models.Order) ([]byte, error) {
	data, err := json.Marshal(order)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}
