package task4

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Consumer struct has two channels
// inputChan will filled by generator
// jobsChan should be filled by resending from inputChan and
// this chan will be listening by workers
type Consumer struct {
	inputChan chan int
	jobsChan  chan int
}

// callbackFunc ONLY FOR INFINITE GENERATOR
// func (c Consumer) callbackFunc(event int) {
// 	c.inputChan <- event
// }

// workerFunc for worker management
func (c *Consumer) workerFunc(wg *sync.WaitGroup, index int) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", index)

	for eventIndex := range c.jobsChan {

		fmt.Printf("Worker %d started job %d\n", index, eventIndex)
		// simulate work  taking between 1-3 seconds
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d finished processing job %d\n", index, eventIndex)
	}
	// if jobsChan closed - it's signal for GI of this worker
	fmt.Printf("Worker %d gracefully interrupted\n", index)
}

// startConsumer is a method for:
//
// * recieve data from inputChan and sending data to jobsChan
// in case of  inputChan was closed it will closed jobsChan too and call canecl() func
// this case is a realisation of normally program exited
//
// * listening ctx.Done() . It will be recieved if program terminated
func (c *Consumer) startConsumer(ctx context.Context, cancel context.CancelFunc) {
	for {
		select {
		case job, ok := <-c.inputChan:
			if ok {
				c.jobsChan <- job
			} else {
				fmt.Println("All data blocks recieved. Closing jobsChan")
				close(c.jobsChan)
				cancel()
				return
			}

		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.jobsChan)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}
