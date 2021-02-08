package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func hashed()  {
	s := "Hello World!"

	hmd5 := md5.Sum([]byte(s))
	hsha1 := sha1.Sum([]byte(s))
	hsha2 := sha256.Sum256([]byte(s))

	fmt.Printf("MD5: %x\n", hmd5)
	fmt.Printf("SHA1: %x\n", hsha1)
	fmt.Printf("SHA256: %x\n", hsha2)
}

func fileHash(){
	f, err := os.Open("codes/files/hashed.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan(){
		t := s.Text()
		hsha2 := sha256.Sum256([]byte(t))
		fmt.Printf("SHA256: %x\n", hsha2)
	}

}


func main()  {
	//hashed()
	fileHash()
}
