package iteration

func Repeat(x string, amount int) (repeatedString string) {
	for i := 0; i < amount; i++ {
		repeatedString += x
	}

	return
}
