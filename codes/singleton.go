// Singleton creational design pattern restricts the instantiation of a type to a single object.
package main

import (
	"fmt"
	"sync"
)


type singleton map[string]string

var (
	once sync.Once

	instance singleton
)


func getSingleton() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
func main()  {

	s := getSingleton()
	fmt.Println(s)
	s["hello"] = "world!"
	fmt.Println(s)
	s2 := getSingleton()
	fmt.Println("This is", s2["hello"])

	s3 := getSingleton()

	s3["hello"] = "bye!"

	fmt.Println("This is", s3["hello"])

}


// Singleton pattern represents a global state and most of the time reduces testability.