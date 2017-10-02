package addMap

import ("testing"
	"fmt"
	)

func TestAddkeyOnlyInMap(t *testing.T) {

var want bool 

if want = AddKeyonlyInMap("kannan"); want != true {
			t.Errorf("AddkeyOnlyInMap(%v) ",want)
}
if want = AddKeyonlyInMap("kannan"); want == true {
			t.Errorf("AddkeyOnlyInMap(%v) ",want)
}
if want = AddKeyonlyInMap("kannan1"); want != true {
			t.Errorf("AddkeyOnlyInMap(%v) ",want)
}
if want = AddKeyonlyInMap("kannan8"); want != true {
			t.Errorf("AddkeyOnlyInMap(%v) ",want)
}

if want = DeleteValueOnlyMap("kannan"); want != true {
			t.Errorf("AddkeyOnlyInMap(%v) ",want)
}
if want = DeleteValueOnlyMap("kannan"); want == true {
			t.Errorf("AddkeyOnlyInMap(%v) ",want)
}
fmt.Println("%T",data)

}
