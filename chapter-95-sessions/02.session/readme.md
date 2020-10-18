## why in 02.session/main.go at line 49 these code works fine

```

package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	//m := map[string]int{}

    // while map filled a value(s), this map return 2 params (value(s) & bool)
	m["mcleod"] = 45  

	toddAge, ok := m["mcleod"]
	if ok == true{
		fmt.Println(toddAge, ok)
	}
	return
}

//it will print
//45 true

```