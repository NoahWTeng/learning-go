package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFile() {

	content, err := ioutil.ReadFile("codes/files/write.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", content)

}

func readFileOS() {
	f, err := os.Open("codes/files/write.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close()
			err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)

	for s.Scan() {
		fmt.Println(s.Text())
	}

	err = s.Err()
	if err != nil{
		log.Fatal(err)
	}
}

func main() {

	//readFile()
	//readFileOS()

}
