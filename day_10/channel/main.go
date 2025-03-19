package channel

func MakeChannel(num int) chan int {
	ch := make(chan int, 1)
	go func() {
		ch <- num
		ch <- num + 1
	}()
	return ch
}
