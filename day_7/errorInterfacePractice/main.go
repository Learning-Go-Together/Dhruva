package errorinterfacepractice

import "fmt"

type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("An Error Occured in the Program %v", e.name)
}

func errorTest() error {
	return userError{name: "Muffin"}
}
