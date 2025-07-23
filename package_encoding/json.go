package package_encoding

import (
	"encoding/json"
	"fmt"
)

type Orang struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

// DemoJSON mendemokan encoding dan decoding JSON
func DemoJSON() {
	obj := Orang{"Budi", 25}
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}
	fmt.Println("JSON Encoded:", string(jsonBytes))

	var obj2 Orang
	err = json.Unmarshal(jsonBytes, &obj2)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	fmt.Printf("JSON Decoded: %+v\n", obj2)
}
