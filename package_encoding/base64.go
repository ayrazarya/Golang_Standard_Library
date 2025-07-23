package package_encoding

import (
	"encoding/base64"
	"fmt"
)

// DemoBase64 mendemokan encoding dan decoding base64
func DemoBase64() {
	data := "Belajar Go!"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Base64 Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}
	fmt.Println("Base64 Decoded:", string(decoded))
}
