package test001

import (
	"fmt"
	"satanGo/satan/protocol"
	"testing"
)

func TestStProto(t *testing.T) {
	bf := protocol.NewStBuffer([]byte{})
	p1 := NewPerson()
	p2 := NewPerson()

	p1.Name = "Benjamin"
	p1.Hobby = []string{"a", "bb", "ccc"}
	p2.Name = "Stella"
	p2.Hobby = []string{"aaa", "b", "新陈代谢"}
	p1.Family["wife"] = p2
	if err := p1.WriteDataBuf(bf); err!=nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bf.Bytes())

	pp1 := NewPerson()
	if err := pp1.ReadDataBuf(bf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pp1.Name)
	fmt.Println(pp1.Hobby)
	t1 := *pp1.Family["wife"]
	t1.Hobby[0] = "Fuck"
	fmt.Println(pp1.Family["wife"])
}
