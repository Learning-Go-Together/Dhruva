package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed {
		return len(e.body) * 2
	}
	return len(e.body) * 5
}

func (e email) format() string {
	if e.isSubscribed {
		status := "Subscribed"
		return fmt.Sprintf("'%s' | %s", e.body, status)
	}

	status := "Not Subscribed"
	return fmt.Sprintf("'%s' | %s", e.body, status)
	// return fmt.Sprintf("Email Body | %s", e.body)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
