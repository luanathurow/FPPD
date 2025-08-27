package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func pessoa(id int, banheiro chan int) {
	fmt.Printf("Pessoa %d quer usar banheiro \n", id)
	//banheiro <- id // entra no banheiro
	select {
	case banheiro <- id:
		fmt.Printf("Pessoa %d esta usando o banheiro\n", id)
	case <-time.After(2 * time.Second):
		fmt.Printf("Pessoa %d desistiu de usar o banheiro", id)
		return
	}
	time.Sleep(5 * time.Second)
	banheiro <- id // saiu do banheiro
	fmt.Printf("Pessoa %d saiu do banheiro\n", id)
}

func main() {
	banheiro := make(chan int, 3)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go pessoa(i, banheiro)
	}
	wg.Wait()
}
