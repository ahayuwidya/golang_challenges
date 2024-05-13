package main

import (
	"fmt"
	"sync"
)

func bisabisa(wg *sync.WaitGroup) {
	defer wg.Done()
	var arraybisa = [3]string{"bisa1", "bisa2", "bisa3"}
	for i := 1; i <= 4; i++ {
		fmt.Println(arraybisa, i)
	}
}

func cobacoba(wg *sync.WaitGroup) {
	defer wg.Done()
	var arraycoba = [3]string{"coba1", "coba2", "coba3"}
	for i := 1; i <= 4; i++ {
		fmt.Println(arraycoba, i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go bisabisa(&wg)
	go cobacoba(&wg)

	wg.Wait()
}
