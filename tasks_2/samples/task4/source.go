// Что выведет программа? Объяснить вывод программы.
package task4

import "fmt"

func Start() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// после выполнения цикла канал останется открытым
		//close(ch) - ДОБАВИТЬ, ЧТОБЫ ПОЛЕЧИТЬ deadlock
	}()
	// range ch будет пытаться читать из канала до тех пор, пока его
	// не закроют. При закрытии канала ему значение поменяется на nil
	// и range заканчивает работу
	for n := range ch { // в исходной постановке - будет deadlock
		fmt.Println(n)
	}
}
