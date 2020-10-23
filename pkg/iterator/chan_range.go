package iterator

func NewChanIntRange(start, end int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := start + 1; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}
