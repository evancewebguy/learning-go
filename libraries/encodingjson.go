package main

import (
	"encoding/json"
	"time"
)

/**
	NOTE: rest apis have json as the standard way to communicate between services and Go standard library.
          includes support for converting Go data types to and from JSON.
	marshaling - means converting from a Go data type to an encoding.
	unmarshalling - means converting tio a Go data type.
*/

//{
//	"id":"12345",
//	"date_ordered":"2020-05-01T13:01:02Z",
//	"customer_id":"3",
//	"items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}]
//}

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (o Order) orderHandler() {
	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		return
	}
}

func main() {

}
