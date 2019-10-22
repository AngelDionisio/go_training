package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

func asyncOne() {
	time.Sleep(time.Second * 2)
	fmt.Println("I am from asyncOne")
	wg.Done()
}

func asyncTwo() {
	fmt.Println("I am from asyncTwo")
	wg.Done()
}

func asyncThree() {
	time.Sleep(time.Second * 3)
	fmt.Println("I am from asyncThree")
	wg.Done()
}

func main() {
	defer fmt.Println("all functions finished executing")
	wg.Add(3)
	
	fmt.Println("num CPUs:", runtime.NumCPU())
	fmt.Println("START num Goroutines:", runtime.NumGoroutine())

	go asyncOne()
	go asyncTwo()
	go asyncThree()
	
	fmt.Println("MID num Goroutines:", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("END num Goroutines:", runtime.NumGoroutine())

}

