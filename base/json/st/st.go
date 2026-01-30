package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type NotifyMsg struct {
	Event_id          string  `json:"event_id"`
	event_timestamp   int64   `json:"event_timestamp"`
	event_type        string  `json:"event_type"`
	transaction_id    string  `json:"transaction_id"`
	order_id          string  `json:"order_id"`
	request_id        string  `json:"request_id"`
	vendor            string  `json:"vendor"`
	status            string  `json:"status"`
	status_message    string  `json:"status_message"`
	status_updated_at string  `json:"status_updated_at"`
	amount            float64 `json:"amount"`
	currency          string  `json:"currency"`
	created_at        int64   `json:"created_at"`
	payment_method    string  `json:"payment_method"`
}

func main() {
	nm := &NotifyMsg{
		Event_id:        "1",
		event_timestamp: time.Now().Unix(),
		event_type:      "2",
		transaction_id:  "3",
		order_id:        "4",
		request_id:      "5",
	}
	ret, _ := json.Marshal(nm)
	fmt.Println(string(ret))
}
