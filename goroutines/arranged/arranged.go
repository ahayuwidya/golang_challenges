package main

import (
	"fmt"
	"time"
)

func cobacoba(ch chan<- int) {
	var arraycoba = [3]string{"coba1", "coba2", "coba3"}
	for i := 1; i < 5; i++ {
		fmt.Println(arraycoba, i)
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func bisabisa(ch <-chan int) {
	var arraybisa = [3]string{"bisa1", "bisa2", "bisa3"}
	for value := range ch {
		fmt.Println(arraybisa, value)
	}
}

func main() {
	ch := make(chan int)
	go cobacoba(ch)
	go bisabisa(ch)
	time.Sleep(5 * time.Second)
}
