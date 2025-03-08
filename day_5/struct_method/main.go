package structmethod

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	// auth := "Authorization: Basic " + a.username + ":" + a.password
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}
