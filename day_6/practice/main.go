package practice

func Practice() {
	i := any(true)
	retrievedInt, ok := i.(int)
	if ok {
		println("value is Interger", retrievedInt)
	} else {
		println("value type is not integer ", retrievedInt)
	}
}
