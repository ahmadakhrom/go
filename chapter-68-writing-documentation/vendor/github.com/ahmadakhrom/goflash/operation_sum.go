package goflash

//sum adds an unlimited number of valuesof type int
func Sum(i ...int) (int, error) {

	s := 0

	for _, val := range i {
		s += val
	}

	return s, nil
}
