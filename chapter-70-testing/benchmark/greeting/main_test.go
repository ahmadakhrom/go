package greeting

import (
	"fmt"
	"testing"
)

//just call "go test" to run it
func TestBirthday(t *testing.T) {

	s := Birthday("John William")
	if s != "happy birthday John William" {
		t.Errorf("got %s, unexpected happy birthday John William", s)
	}
}

//for showing into godoc via godoc -http=:8080
func ExampleBirthday() {

	fmt.Println(Birthday("John William"))
}

//just call "go test -bench <func name>" to run benchmark all it
func BenchmarkBirthday(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Birthday("John William")
	}
}
