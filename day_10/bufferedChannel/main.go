package bufferedchannel

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	go func() {
		for _, e := range emails {
			ch <- e
		}
		// close(ch)
	}()
	return ch
}
