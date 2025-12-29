package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T){
	sum:= Add(2,2)
	expected := 4

	if sum !=expected {
		t.Errorf("expected '%d' but got '%d'", expected,sum)
	}
}
// Add takes to integers and returns the sum of them
func Add(x,y int)int{
	return x+y
}

//TESTABLE EXAMPLES
//While the example will always be compiled,
//adding this comment means the example will also be executed. 
func ExampleAdd(){
	sum:= Add(1,5)
	fmt.Println(sum)
	//Output: 6
}