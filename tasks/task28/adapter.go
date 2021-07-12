// Package task28 Реализовать паттерн адаптер на любом примере.
package task28

import (
	"fmt"
)

// client ...
type client struct{}

// SendMessageToService ...
func (c *client) SendMessageToService(srv service) {
	fmt.Println("Client send message to service:", srv.GetServiceName())
	srv.SendMessage()
}

// service ...
type service interface {
	SendMessage()
	GetServiceName() string
}

// TelegramService ...
type TelegramService struct {
	name string
}

// SendPopMessage ...
func (tlg *TelegramService) SendPopMessage() {
	fmt.Println("Telegram send message")
}

// TelegramAdapter ...
type TelegramAdapter struct {
	telegramService *TelegramService
}

// SendMessage ...
func (tlgAdd *TelegramAdapter) SendMessage() {
	fmt.Println("Connect tlg adapter")
	tlgAdd.telegramService.SendPopMessage()
}

// GetServiceName ...
func (tlgAdd *TelegramAdapter) GetServiceName() string {
	return tlgAdd.telegramService.name
}

// Start ...
func Start() {
	client := &client{}
	telegramServ := &TelegramService{"telegram 2.0.x for ultra lollers"}
	telegramAdapter := &TelegramAdapter{telegramServ}

	client.SendMessageToService(telegramAdapter)
}
