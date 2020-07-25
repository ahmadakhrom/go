package goslash

func Sum(i ...int) (int, error) {
	summary := 0

	for val, _ := range i {
		summary += val
	}
	return summary, nil
}
