/*package main

import (
	"fmt"
	"time"
)

func ImprimeCrescente(ch chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	ch <- true
	fmt.Println("imprimeCrescente retornou")
}

func imprimeDecrescente(ch chan bool) {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	ch <- true
	fmt.Println("imprimeDecrescente retornou")
}

func main() {
	ch := make(chan bool, 2)
	go ImprimeCrescente(ch)
	go imprimeDecrescente(ch)
	time.Sleep(15 * time.Second)
	<-ch
	<-ch
}
