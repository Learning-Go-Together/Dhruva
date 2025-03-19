package rangechannel

func concurrentFib(n int) []int {
	ch := make(chan int, n)
	series := make([]int, 0)
	go fibonacci(n, ch)
	for num := range ch {
		series = append(series, num)
	}
	return series
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
