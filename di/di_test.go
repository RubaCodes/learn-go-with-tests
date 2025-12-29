package di

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer,"Chris")

	got:= buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q",got,want)
	}
}

func Greet(buf io.Writer,name string){
	fmt.Fprintf(buf,"Hello, %s",name)
}