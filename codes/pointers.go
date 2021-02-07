package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 3, 4, 19, 21, 0, 2}
	v1 := &arr // copy reference
	v2 := arr  // copy reference
	v3 := make([]int, len(arr))
	copy(v3, arr) // shadow copy, copy value
	fmt.Println(arr)
	fmt.Println(*v1)
	fmt.Println(v2)
	fmt.Println(v3)

	v2[0] = 10
	v3[0] = 100

	fmt.Printf("key: %pn value: %d \n", arr, arr)
	fmt.Printf("key: %pn value: %d \n", *v1, *v1)
	fmt.Printf("key: %pn value: %d \n", v2, v2)
	fmt.Printf("key: %pn value: %d \n", v3, v3)

}
