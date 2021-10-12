package gopackagejson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Kholid",
		LastName:   "Saifulloh",
		Age:        23,
		Married:    false,
	}

	byte, _ := json.Marshal(customer)
	fmt.Println(string(byte))
}
