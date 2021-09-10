package models

// Order ...
type Order struct {
	OrderUID          string  `json:"order_uid"`
	Entry             string  `json:"entry"`
	InternalSignature string  `json:"internal_signature"`
	Payment           Payment `json:"payment"`
	Items             []Items `json:"items"`
	Locale            string  `json:"locale"`
	CustomerID        string  `json:"customer_id"`
	TrackNumber       string  `json:"track_number"`
	DeliveryService   string  `json:"delivery_service"`
	Shardkey          string  `json:"shardkey"`
	SmID              int     `json:"sm_id"`
}

// Orders ...
type Orders struct {
	Orders []Order `json:"orders"`
}
