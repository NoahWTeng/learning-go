package main

import (
	"fmt"
	"sync"
	"time"
)

// Go routine When a new Goroutine is started, the goroutine call returns
// immediately. Unlike functions, the control does not wait for the Goroutine to
// finish executing. The control returns immediately to the next line of code
// after the Goroutine call and any return values from the Goroutine are ignored.
// The main Goroutine should be running for any other Goroutines to run. If the
// main Goroutine terminates then the program will be terminated and no other
// Goroutine will run.

func hello() {
	fmt.Println("Hello World!")
}

// Multi
func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func multiRoutines() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

// Chan

/*
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}

	b := make(chan int)
	data := <- b // read from channel a
	b <- data // write to channel a
*/
func helloChan(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func withChannel() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go helloChan(done)
	<-done
	fmt.Println("Main received data")
}

// Example channel 01
type Channels struct {
	num     int
	channel chan int
}

func (s Channels) calcSqr() {
	sum := 0
	for s.num != 0 {
		digit := s.num % 10
		sum += digit * digit
		s.num /= 10
	}
	s.channel <- sum
}

func (s Channels) calcCub() {
	sum := 0
	for s.num != 0 {
		digit := s.num % 10
		sum += digit * digit * digit
		s.num /= 10
	}
	s.channel <- sum
}

func oneExaChannel() {

	sq := Channels{589, make(chan int)}
	cub := Channels{453, make(chan int)}

	go sq.calcSqr()
	go cub.calcCub()
	a, b := <-sq.channel, <-cub.channel
	fmt.Println("Channel sq, cub: ", a, b)
}

// Example channel 02

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func twoExaChannel() {
	ch := make(chan int)

	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

}

// WaitGroup

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func waitGroupF() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

// Select

func processing(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func selectFun() {
	ch := make(chan string)
	go processing(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}

// Mutex

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func lostNum() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

var y = 0

func increment1(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}

func solvingLostNum() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment1(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of y", y)
}

var c = 0

func increment2(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	c = c + 1
	<-ch
	wg.Done()
}

func solvingWithChannel() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment2(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of c", c)

}

func mutexFun() {

	//lostNum()

	//solvingLostNum()

	solvingWithChannel()
}

func main() {

	//go hello()
	//time.Sleep(1 * time.Second)
	//fmt.Println("Finally main")
	//multiRoutines()

	//withChannel()
	//oneExaChannel()
	//twoExaChannel()

	//waitGroupF()

	//selectFun()

	mutexFun()

}
