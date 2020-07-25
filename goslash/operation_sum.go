package goslash

func Sum(i ...int) (int, error) {

	s := 0

	for _, val := range i {
		s += val
	}

	return s, nil
}
