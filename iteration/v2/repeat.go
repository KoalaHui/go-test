package iteration

func Repeat(str string, len int) (repeat string) {
	for i := 0; i < len; i++ {
		repeat += str
	}
	return

}
