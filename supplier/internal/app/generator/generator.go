package generator

import (
	"github.com/vlasove/materials/supplier/internal/app/filemanager"
	"github.com/vlasove/materials/supplier/internal/app/models"
)

// GenerateRandomOrders ...
func GenerateRandomOrders(n int) []*models.Order {
	os := make([]*models.Order, n)
	for i := 0; i < n; i++ {
		os[i] = GenerateOrder()
	}
	return os
}

// GenerateOrder ...
func GenerateOrder() *models.Order {
	o := &models.Order{}
	o.FillRandomOrder()
	return o
}

func Start(amount int, filePath string) error {
	os := GenerateRandomOrders(amount)
	fm := filemanager.New(filePath)
	if err := fm.Dump(os); err != nil {
		return err
	}
	return nil
}
