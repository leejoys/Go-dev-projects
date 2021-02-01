package sort

func bubbleSort(ar []int) {
	arLen := len(ar)
	for i := 0; true; i++ {
		sorted := true
		for i := 1; i < arLen; i++ {
			if ar[i-1] > ar[i] {
				ar[i-1], ar[i] = ar[i], ar[i-1]
				sorted = false
			}
		}
		if sorted {
			return
		}
		arLen--
	}
}
