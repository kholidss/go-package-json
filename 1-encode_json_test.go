package gopackagejson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// encode JSON (dari objek ke bentuk JSON)
func logJson(data interface{}) {
	byte, error := json.Marshal(data)
	if error != nil {
		panic(error)
	}

	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	logJson("Joko")
	logJson(1)
	logJson(true)
	logJson([]string{"Joko", "Abdul", "Manap"})

}
