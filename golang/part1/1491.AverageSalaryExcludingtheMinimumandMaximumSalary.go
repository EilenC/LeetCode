package part1

func average(salary []int) float64 {
	nl := len(salary)
	max, min, total := salary[0], salary[0], 0
	for i := 0; i < nl; i++ {
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
		total += salary[i]
	}
	return float64(total-max-min) / float64(nl-2)
}
