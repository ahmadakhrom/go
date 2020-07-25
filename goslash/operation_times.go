package goslash

func Times(i ...int) (int, error) {
	s := 1
	for _, val := range i {
		s *= val
	}
	return s, nil
}