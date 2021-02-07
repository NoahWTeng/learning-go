package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	text := "My little secret"

	tEnc := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println("Encode 64", tEnc)

	dEnc, _ := base64.StdEncoding.DecodeString(tEnc)
	fmt.Println("Decode 64", string(dEnc))

}
