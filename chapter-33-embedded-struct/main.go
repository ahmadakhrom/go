package main

import "fmt"

type student struct {
	id    int
	name  string
	class int
}

type school struct {
	student
	status bool
}

func main() {

	student1 := school{
		student: student{ // <<== line applied embedded struct
			id:    1001,
			name:  "andy nugraha",
			class: 12,
		},
		status: true,
	}

	student2 := student{
		id:    1002,
		name:  "sarah amelia",
		class: 11,
	}

	fmt.Println(student1)
	fmt.Println(student2, "\n")

	fmt.Println(student1.id, ",", student1.name, ",", student1.class, ", ", student1.status)
	fmt.Println(student2.id, ",", student2.name, ",", student2.class)

}
