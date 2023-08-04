package main

import (
	"fmt"
	"time"
)

// Função para processar uma tarefa simulada
func processTask(id int) {
	// Simulando um tempo de processamento
	time.Sleep(time.Second)
	fmt.Printf("Tarefa %d processada\n", id)
}

func main() {
	// Número total de tarefas a serem processadas
	totalTasks := 10

	// Tamanho do bugger do canal (limite de tarefas simultâneas)
	bufferSize := 3

	// Criação do canal com buffer
	taskQueue := make(chan int, bufferSize)

	go func() {
		for i := 1; i <= totalTasks; i++ {
			taskQueue <- i // Adiciona a tarefa ao canal
		}
		close(taskQueue) // Fecha o canal após todas as tarefas serem adicionadas
	}()

	for i := 1; i < bufferSize; i++ {
		go func(workerID int) {
			for taskID := range taskQueue {
				fmt.Printf("Trabalhador %d iniciou tarefa %d\n", workerID, taskID)
				processTask(taskID)
				fmt.Printf("Trabalhador %d concluiu tarefa %d\n", workerID, taskID)
			}
		}(i)
	}

	// Aguarda todas as goroutines de processamento de tarefas terminarem
	time.Sleep(time.Second * 5)

}
