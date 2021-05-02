package client

import (
	"fmt"
	"testing"
)

func TestStClient_regEndPoint(t *testing.T) {
	s := "172.21.23.106:100"
	fmt.Println(regEndPoint.MatchString(s))
}
