package say_test

import (
	"fmt"
	"github.com/cheesiewalters/say"
	"testing"
)

func ExampleHello() {
	greeting := say.Hello("Conor")
	fmt.Println(greeting)
}

func TestHello(t *testing.T) {
	n := "Conor"
	expected := "Hello Conor."
	actual := say.Hello(n)

	if expected != actual {
		t.Logf("Hello: expected [%s] go [%s]", expected, actual)
		t.Fail()
	}
}
