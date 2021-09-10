package generator

import "github.com/vlasove/materials/test_src/generator/internal/app/models"

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
