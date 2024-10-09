package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type testData struct {
	phones []int
	mu     sync.Mutex
}

func (td *testData) add() {
	td.mu.Lock() // Blocking
	td.phones = append(td.phones, td.randPhone())
	defer td.mu.Unlock() // Unblocking after adding a number
}

func generate(n int, td *testData) *testData {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			td.add()
		}()
	}
	wg.Wait() // Waiting for goroutines to finish
	return td
}

func main() {
	td := testData{}
	generate(100, &td)
	fmt.Println(len(td.phones))
}

func (td *testData) randPhone() int {
	return 89000000000 + rand.Intn(999999999-100000000) + 100000000
}
