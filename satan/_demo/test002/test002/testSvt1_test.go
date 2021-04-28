package test002

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

func TestStProto2(t *testing.T) {
	bf := protocol.NewStBuffer([]byte{})
	m := NewMapper()
	m.Gap = matrixInt(1000, 1000)
	for i:=0; i<1000; i++ {
		for j:=0; j<1000; j++ {
			m.Gap[i][j] = i*j
		}
	}
	m.HhTest["Gap"] = m.Gap
	if err := m.WriteDataBuf(bf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(bf.Bytes()))

	mm := NewMapper()
	if err := mm.ReadDataBuf(bf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mm.Gap[27][89])
}

func matrixInt(x, y int) [][]int {
	m1 := make([][]int, x)
	for i:=0; i<x; i++ {
		m2 := make([]int, y)
		m1[i] = m2
	}
	return m1
}