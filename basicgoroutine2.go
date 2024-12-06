package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= n; i++ {
		fmt.Println(i)
		amt := time.Duration(rand.Intn(500))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	var n int
	var wg sync.WaitGroup

	wg.Add(1)
	fmt.Println("Enter an integer :  ")
	fmt.Scanln(&n)
	go f(n, &wg)

	wg.Wait()
}
