package task4

// FiniteGenerator struct of generator for main producing
type FiniteGenerator struct {
	Consumer
}

// start will send finite amount of data
// to consumer input channel
func (fg *FiniteGenerator) start() {
	for i := 0; i < cap(fg.Consumer.inputChan); i++ {
		fg.Consumer.inputChan <- i
	}
	close(fg.Consumer.inputChan)
}
