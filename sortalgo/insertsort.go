package sortalgo

func doInsertSort(list []int) []int {
	var sorted []int
	for i := 0; i < len(list); i++ {
		sorted = insert(sorted, list[i])
	}
	return sorted
}

func insert(list []int, item int) []int {

	for k, v := range list {
		if item < v {
			return append(list[:k], append([]int{item}, list[k:]...)...)

		}
	}
	return append(list, item)
}
