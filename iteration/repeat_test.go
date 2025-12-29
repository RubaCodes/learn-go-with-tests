package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T){
	repeated := Repeat("a",5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q",expected,repeated)
	}
}

func ExampleRepeat(){
	res := Repeat("a",5)
	fmt.Println(res)
	//Output: aaaaa
}
func Repeat(char string,times int) string {
	var repeated strings.Builder
	for i:=0;i< times;i++{
		repeated.WriteString(char)
	}
	return  repeated.String()
}

//BENCHMARKING
// go test -bench=. -benchmem
func BenchmarkRepeat(b* testing.B){
	for b.Loop(){
		Repeat("A",5)
	}
}