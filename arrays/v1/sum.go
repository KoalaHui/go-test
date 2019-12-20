package arrys

func Sum(arr [5]int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}
