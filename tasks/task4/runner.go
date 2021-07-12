// Package task4 Реализовать набор из N воркеров,
// которые читают из канала произвольные данные и выводят в stdout.
// Данные в канал пишутся из главного потока.
// Необходима возможность выбора кол-во воркеров при старте,
// а также способ завершения работы всех воркеров.
package task4

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Start with workerPoolSize param
func Start(workerPoolSize int) {
	consumer := Consumer{
		inputChan: make(chan int, workerPoolSize*10),
		jobsChan:  make(chan int, workerPoolSize),
	}

	//generator := Generator{callbackFunc: consumer.callbackFunc}
	generator := FiniteGenerator{consumer}
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	generator.start()
	//go generator.start(ctx)
	go consumer.startConsumer(ctx, cancelFunc)

	wg.Add(workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		go consumer.workerFunc(wg, i)
	}

	// chan for terminated signals
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)

	select {
	case <-termChan:
		// if terminated
		fmt.Println("=========Shutdown Signal=========")
		cancelFunc()
	case <-ctx.Done():
		// if normally exited
		fmt.Println("=========Normally exited==========")
	}
	// Wait until all workers gracefully interupted
	wg.Wait()

	fmt.Println("==============All workers done!========")
}
