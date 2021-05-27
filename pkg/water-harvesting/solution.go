package water_harvesting

func getBucketWater(start int, in []int) (int, int) {
	if len(in) - start < 3 {
		return 0, len(in)
	}

	startHeight := 0
	totalWater := make([]int, len(in))
	lastHeight := 0
	lastHeightPos := 0

	for i := start; i < len(in); i++ {
		if in[i] > lastHeight && startHeight > 0 {
			lastHeight = in[i]
			lastHeightPos = i
		}
		if in[i] < startHeight {
			totalWater[i] = (startHeight - in[i])
		} else {
			startHeight = in[i]
		}
		if lastHeight > startHeight {
			break
		}
	}
	sum := 0
	diff := startHeight - lastHeight
	for i := start + 1; i < lastHeightPos; i++ {
		sum += totalWater[i] - diff
	}
	return sum, lastHeightPos
}

func howMuchWater(in []int) int {
	i := 0
	var res int
	sum := 0
	for i < len(in) {
		res, i = getBucketWater(i, in)
		sum += res
	}
	return sum
}