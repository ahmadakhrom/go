package goflash

//sum time an unlimited number of valuesof type int
func Times(i ...int) (int, error) {
	s := 1
	for _, val := range i {
		s *= val
	}
	return s, nil
}
