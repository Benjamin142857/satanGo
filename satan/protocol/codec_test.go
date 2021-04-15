package protocol

import (
	"fmt"
	"testing"
)

var bf = NewStBuffer([]byte{})

func TestStBuffer_WriteDataBuf(t *testing.T) {
	var bs []byte
	var d interface{}

	// test write byte
	d = byte('a')
	if err := bf.WriteDataBuf(Byte, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode byte(%v): %v\n", d, bs)
	bf.ReSet([]byte{})

	// test write int
	d = 256*256*3 + 256*16 + 37
	if err := bf.WriteDataBuf(Int, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode int(%v): %v\n", d, bs)
	bf.ReSet([]byte{})

	// test write long
	d = int64(256*256*3 + 256*16 + 37)
	if err := bf.WriteDataBuf(Long, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode long(%v): %v\n", d, bs)
	bf.ReSet([]byte{})

	// test write string
	d = "Satan - Benjamin142857, 嘿嘿嘿"
	if err := bf.WriteDataBuf(String, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode string(%v): %v\n", d, bs)
	bf.ReSet([]byte{})

	// test write float
	d = float32(142.857)
	if err := bf.WriteDataBuf(Float, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode float(%v): %v\n", d, bs)
	bf.ReSet([]byte{})

	// test write all
	d = byte('a')
	if err := bf.WriteDataBuf(Byte, d); err != nil {
		t.Error(err)
	}
	d = 256*256*3 + 256*16 + 37
	if err := bf.WriteDataBuf(Int, d); err != nil {
		t.Error(err)
	}
	d = int64(256*256*3 + 256*16 + 37)
	if err := bf.WriteDataBuf(Long, d); err != nil {
		t.Error(err)
	}
	d = "Satan - Benjamin142857, 嘿嘿嘿"
	if err := bf.WriteDataBuf(String, d); err != nil {
		t.Error(err)
	}
	d = float32(142.857)
	if err := bf.WriteDataBuf(Float, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode all: %v\n", bs)
}

func TestStBuffer_ReadDataBuf(t *testing.T) {
	var bs []byte
	var err error
	var d interface{}
	var _d interface{}

	// test write && read byte
	d = byte('a')
	if err := bf.WriteDataBuf(Byte, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode byte(%v): %v\n", d, bs)
	_d, err = bf.ReadDataBuf(Byte)
	_ = _d.(byte)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode byte(%v): %v\n", d, _d)
	bf.ReSet([]byte{})

	// test write && read int
	d = 256*256*3 + 256*16 + 37
	if err := bf.WriteDataBuf(Int, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode int(%v): %v\n", d, bs)
	_d, err = bf.ReadDataBuf(Int)
	_ = _d.(int)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode int(%v): %v\n", d, _d)
	bf.ReSet([]byte{})

	// test write && read long
	d = int64(256*256*3 + 256*16 + 37)
	if err := bf.WriteDataBuf(Long, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode long(%v): %v\n", d, bs)
	_d, err = bf.ReadDataBuf(Long)
	_ = _d.(int64)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode long(%v): %v\n", d, _d)
	bf.ReSet([]byte{})

	// test write && read string
	d = "Satan - Benjamin142857, 嘿嘿嘿"
	if err := bf.WriteDataBuf(String, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode string(%v): %v\n", d, bs)
	_d, err = bf.ReadDataBuf(String)
	_ = _d.(string)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode string(%v): %v\n", d, _d)
	bf.ReSet([]byte{})

	// test write && read float
	d = float32(142.857)
	if err := bf.WriteDataBuf(Float, d); err != nil {
		t.Error(err)
	}
	bs = bf.Bytes()
	fmt.Printf("encode float(%v): %v\n", d, bs)
	_d, err = bf.ReadDataBuf(Float)
	_ = _d.(float32)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode float(%v): %v\n", d, _d)
	bf.ReSet([]byte{})

	// test write && read all
	fmt.Printf("decode all----------------------------\n")
	d = byte('a')
	if err := bf.WriteDataBuf(Byte, d); err != nil {
		t.Error(err)
	}
	d = 256*256*3 + 256*16 + 37
	if err := bf.WriteDataBuf(Int, d); err != nil {
		t.Error(err)
	}
	d = int64(256*256*3 + 256*16 + 37)
	if err := bf.WriteDataBuf(Long, d); err != nil {
		t.Error(err)
	}
	d = "Satan - Benjamin142857, 嘿嘿嘿"
	if err := bf.WriteDataBuf(String, d); err != nil {
		t.Error(err)
	}
	d = float32(142.857)
	if err := bf.WriteDataBuf(Float, d); err != nil {
		t.Error(err)
	}
	_d, err = bf.ReadDataBuf(Byte)
	_ = _d.(byte)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode byte: %v\n", _d)
	_d, err = bf.ReadDataBuf(Int)
	_ = _d.(int)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode int: %v\n", _d)
	_d, err = bf.ReadDataBuf(Long)
	_ = _d.(int64)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode long: %v\n", _d)
	_d, err = bf.ReadDataBuf(String)
	_ = _d.(string)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode string: %v\n", _d)
	_d, err = bf.ReadDataBuf(Float)
	_ = _d.(float32)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("decode float: %v\n", _d)
}