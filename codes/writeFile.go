package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

func createFile()  {
	f, err := os.Create("codes/files/write.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("Hello world!")

	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	defer func() {
		if err = f.Close()
			err != nil {
			log.Fatal(err)
		}
	}()
}

func createBytesFile()  {
	f, err := os.Create("codes/files/write.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}

	n2, err := f.Write(d2)

	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createMultipleLine()  {
	f, err := os.Create("codes/files/write.txt")
	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

	for _, v := range d {
		_, err = fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

func appendToLine()  {
	f, err := os.OpenFile("codes/files/write.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}

// concurrently

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}


func consume(data chan int, done chan bool) {
	f, err := os.Create("codes/files/write.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			_ = f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func concurrentlyWrite()  {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}

func main()  {

	//createFile()
	//createBytesFile()
	//createMultipleLine()
	//appendToLine()

	//concurrentlyWrite()
}

