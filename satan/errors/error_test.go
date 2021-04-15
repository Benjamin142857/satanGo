package errors

import (
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	fmt.Println(NewStError(1001))
	fmt.Println("-------")
	fmt.Println(NewStError(1002))
	fmt.Println("-------")
	fmt.Println(NewStError(1002, NewStError(1001), NewStError(1001)))
	append
}
