package addstring

import "fmt"

func concate(s1 string, s2 string) string {
	return s1 + " " + s2
}
func Print(s1 string, s2 string) {
	fmt.Println(concate(s1, s2))
}
