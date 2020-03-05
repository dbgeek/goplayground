package sortalgo

func doBubbleSort(v []int) []int {

	for idx := 0; idx < len(v); idx++ {
		//swap := false
		for i := 0; i < len(v)-1-idx; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				//swap = true
			}
		}
		/*if !swap {
			break
		}*/
	}
	return v
}
