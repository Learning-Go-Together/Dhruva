package maps

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(phoneNumbers) != len(names) {
		return nil, errors.New("invalid sizes")
	}
	mydict := make(map[string]user)
	for i := 0; i < len(names); i++ {
		mydict[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
	}

	return mydict, nil
}

type user struct {
	name        string
	phoneNumber int
}
