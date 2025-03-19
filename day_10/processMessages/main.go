package processmessages

import (
	"time"
)

func processMessages(messages []string) []string {
	ch := make(chan string, len(messages))
	// var wg sync.WaitGroup
	done := make(chan struct{})

	// var result []string
	// wg.Add(1)
	for _, msg := range messages {
		go func(m string) {
			ch <- msg
		}(msg)
	}
	go func() {
		for i := 0; i < len(messages); i++ {
			<-ch
		}
		close(ch)
		done <- struct{}{}
		// close(done)

	}()
	var result []string
	for processmsg := range ch {
		result = append(result, processmsg)
	}

	<-done

	return result
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
