package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	ID   int
	Name string
	Age  int
}

func main() {

	resp := `[
   {
     "ID": 1,
     "Name": "Andy Nugraha",
     "Age": 27
   },
   {
     "ID": 2,
     "Name": "Budi Budiman",
     "Age": 26
   }
 ]`

	//var respStudent []student
	respStudent :=  []student{}
	err := json.Unmarshal([]byte(resp), &respStudent)

	if err != nil {
		var err error
		fmt.Println(err)
	}

	fmt.Println("printed of json respon raw :", respStudent)

	var ID, Age int
	var Name string

	for _, val := range respStudent {
		ID = val.ID
		Name = val.Name
		Age = val.Age
		 fmt.Println("\n------ID","Age","Name")
		 fmt.Println(val.ID, val.Age, val.Name)
		 fmt.print

	}
	fmt.Println("\nID\tName\t\tAge\n", ID, "\t", Name, "\t", Age) //---ID      Name            Age
	//---------------------------------------------------------------2       Budi Budiman    26   //selalu print value ID terakhir, why???

}
