package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		"James",
		"Bond",
		[]string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
		//fmt.Errorf(err.Error())
		return
	}
	fmt.Println(err)
	fmt.Println(string(bs))
	fmt.Printf("type of : %T \n", bs)
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err == nil {
		//return []byte{}, fmt.Errorf("cannot unmarshaling %v %v ", a, err)
		//return []byte{}, fmt.Errorf("cannot unmarshaling %v ", a)
		return []byte{}, errors.New(fmt.Sprintf("there is an error.. %v", err))
	}
	return bs, nil
}
