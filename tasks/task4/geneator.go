package task4

// Generator ...
// type Generator struct {
// 	callbackFunc func(event int)
// }

// func (g Generator) start(ctx context.Context) {
// 	eventId := 0
// 	for {
// 		g.callbackFunc(eventId)
// 		eventId++
// 		time.Sleep(time.Millisecond * 500)
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Generator stopped after 1 sec ....")
// 			time.Sleep(1 * time.Second)
// 			fmt.Println("Generator successfully stopped!")
// 			return
// 		default:
// 			continue

// 		}
// 	}
// }
