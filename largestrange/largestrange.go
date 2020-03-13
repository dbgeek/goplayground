package largestrange

func largestRange(array []int) []int {
	m := make(map[int]bool)
	var longest []int
	var longestRange int

	for _, v := range array {
		m[v] = true
	}

	for _, i := range array {
		if !m[i] {
			continue
		}
		left := i - 1
		right := i + 1
		currentRange := 1

		for m[left] {
			currentRange++
			left--
		}

		for m[right] {
			currentRange++
			right++
		}

		if currentRange > longestRange {
			longestRange = currentRange
			longest = []int{left + 1, right - 1}
		}
	}
	return longest
}
