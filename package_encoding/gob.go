package package_encoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Data struct {
	X int
	Y string
}

// DemoGOB mendemokan encoding dan decoding GOB
func DemoGOB() {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	data := Data{42, "Go"}
	if err := enc.Encode(data); err != nil {
		fmt.Println("GOB Encode error:", err)
		return
	}
	fmt.Println("GOB Encoded Bytes:", buf.Bytes())

	dec := gob.NewDecoder(&buf)
	var data2 Data
	if err := dec.Decode(&data2); err != nil {
		fmt.Println("GOB Decode error:", err)
		return
	}
	fmt.Printf("GOB Decoded: %+v\n", data2)
}
