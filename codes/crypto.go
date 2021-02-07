package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

var secretKey = []byte("passphrasewhichneedstobe32bytes!")

func encrypt() {
	fmt.Println("Encryption Program v1")

	text := []byte("My Super Secret Message: Hello World!")

	//generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(secretKey)
	// if there are any errors, handle them

	if err != nil {
		log.Fatal(err)
	}
	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		log.Fatal(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())

	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	data := gcm.Seal(nonce, nonce, text, nil)

	err = ioutil.WriteFile("codes/files/encrypt.txt", data, 0777)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Message encrypted", data)


}

func decryption(){
	fmt.Println("Decryption Program v1")

	content, err := ioutil.ReadFile("codes/files/encrypt.txt")
	if err != nil {
		log.Fatal(err)
	}

	c, err := aes.NewCipher(secretKey)

	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatal(err)
	}

	nonceSize := gcm.NonceSize()
	if len(content) < nonceSize {
		log.Fatal(err)
	}

	nonce, context := content[:nonceSize], content[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, context, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message decrypted", string(plaintext))
}


func main() {
	//encrypt()
	decryption()
}
