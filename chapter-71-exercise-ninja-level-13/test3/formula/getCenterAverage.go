package formula

import (
	"sort"
)

func GetCenterAverage(numbers []int) float64 {
	sort.Ints(numbers)
	as := numbers[1:]

	n := 0
	for _, v := range as {
		n += v
	}

	res := float64(n) / float64(len(as))
	return res
}

// func main() {
// 	d := []int{5, 10}
// 	s := GetAverage(d)
// 	fmt.Println(s)
// }

//[]int{1,3,4,2,343,45,4,3,4,5,4,34,} <== how to get average from this case
