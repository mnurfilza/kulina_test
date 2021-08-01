package kulina

func colorCount(color []int) map[int]int {

	tot := make(map[int]int)
	for _, v := range color {
		_, ok := tot[v]
		if ok {

			tot[v] += 1
		} else {
			tot[v] = 1

		}
	}

	return tot
}

func findPairColor(total int, color []int) int {
	colors := colorCount(color)
	sum := 0
	for _, v := range colors {
		modRes := v % 2
		if modRes != 0 {
			res := v - modRes
			if res != 0 {
				sum += res / 2
			}

		} else {
			sum += v / 2
		}
	}
	return 0
}
