package main

import "fmt"

func main() {
	//rodar duas funcoes em paralelo

	// Inicia duas goroutines em paralelo
	//go tarefa("Tarefa 1", 500*time.Millisecond)
	//go tarefa("Tarefa 2", 700*time.Millisecond)

	// Sem isso, o programa terminaria antes de rodar as goroutines.
	// Aqui apenas esperamos um tempo suficiente para elas terminarem.
	//time.Sleep(5 * time.Second)

	//fmt.Println("Fim do programa")

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan []int)

	// roda em paralelo
	go contaPares(numeros, ch)
	go contaImpares(numeros, ch)

	// recebe do canal (bloqueia até ter resultado)
	pares := <-ch
	impares := <-ch
	fmt.Println("Pares:", pares)
	fmt.Println("Pares:", impares)
}

/*func tarefa(nome string, tempo time.Duration) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s - execução %d\n", nome, i)
		time.Sleep(tempo)
	}
}*/

// o programa só finalize quando as duas goroutines terminarem (sync.WaitGroup)
func contaPares(nums []int, ch chan []int) {
	var pares []int
	for _, n := range nums {
		if n%2 == 0 {
			pares = append(pares, n)
		}
	}
	ch <- pares // envia o resultado pelo canal
}

func contaImpares(nums []int, ch chan []int) {
	var impares []int
	for _, n := range nums {
		if n%2 != 0 {
			impares = append(impares, n)
		}
	}
	ch <- impares // envia o resultado pelo canal
}
