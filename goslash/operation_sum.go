package goslash

func Sum(i ...int) (int, error) {
	summary := 0

	for _, val := range i {
		summary += val
	}
	return summary, nil
}
