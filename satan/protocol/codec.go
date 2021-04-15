package protocol

import (
	"bytes"
	"encoding/binary"
	"math"
	"satanGo/satan/errors"
)

type StProtocolType byte

const (
	Unknown StProtocolType = iota
	Byte
	Bool
	Int
	Long
	Float
	Double
	String
	List
	Map
	Struct
)

var StTypeMap = map[byte]StProtocolType{
	1:  Byte,
	2:  Bool,
	3:  Int,
	4:  Long,
	5:  Float,
	6:  Double,
	7:  String,
	8:  List,
	9:  Map,
	10: Struct,
}

type StStruct interface {
	WriterDataBuf(bf *StBuffer) error
	ReadDataBuf(bf *StBuffer) error
}

type StBuffer struct {
	buf *bytes.Buffer
}

func (bf *StBuffer) WriteTag(tag byte) error {
	return bf.writeByte(tag)
}

func (bf *StBuffer) WriteDataType(tp StProtocolType) error {
	return bf.writeByte(byte(tp))
}

func (bf *StBuffer) WriteLength(l int) error {
	var bs []byte
	if l < 0 {
		return errors.NewStError(1003)
	} else if l == 0 {
		bs = make([]byte, 1)
		bs[0] = 0
	} else if l <= (1<<8 - 1) {
		bs = make([]byte, 2)
		bs[0] = 1
		bs[1] = byte(l)
	} else if l <= (1<<16 - 1) {
		bs = make([]byte, 3)
		bs[0] = 2
		bs[1] = byte(l >> 8)
		bs[2] = byte(l)
	} else if l <= (1<<24 - 1) {
		bs = make([]byte, 4)
		bs[0] = 3
		bs[1] = byte(l >> 16)
		bs[2] = byte(l >> 8)
		bs[3] = byte(l)
	} else if l <= (1<<32 - 1) {
		bs = make([]byte, 5)
		bs[0] = 4
		bs[1] = byte(l >> 24)
		bs[2] = byte(l >> 16)
		bs[3] = byte(l >> 8)
		bs[4] = byte(l)
	} else if l <= (1<<40 - 1) {
		bs = make([]byte, 6)
		bs[0] = 5
		bs[1] = byte(l >> 32)
		bs[2] = byte(l >> 24)
		bs[3] = byte(l >> 16)
		bs[4] = byte(l >> 8)
		bs[5] = byte(l)
	} else if l <= (1<<48 - 1) {
		bs = make([]byte, 7)
		bs[0] = 6
		bs[1] = byte(l >> 40)
		bs[2] = byte(l >> 32)
		bs[3] = byte(l >> 24)
		bs[4] = byte(l >> 16)
		bs[5] = byte(l >> 8)
		bs[6] = byte(l)
	} else {
		return errors.NewStError(1005)
	}

	if _, err := bf.buf.Write(bs); err != nil {
		return errors.NewStError(1003, err)
	}
	return nil
}

func (bf *StBuffer) WriteDataBuf(tp StProtocolType, d interface{}) error {
	switch tp {
	case Byte:
		_d, ok := d.(byte)
		if !ok {
			return errors.NewStError(1002)
		}
		return bf.writeByte(_d)
	case Bool:
		_d, ok := d.(bool)
		if !ok {
			return errors.NewStError(1002)
		}
		return bf.writeBool(_d)
	case Int:
		_d, ok := d.(int)
		if !ok {
			return errors.NewStError(1002)
		}
		return bf.writeInt(int32(_d))
	case Long:
		_d, ok := d.(int64)
		if !ok {
			return errors.NewStError(1002)
		}
		return bf.writeLong(_d)
	case Float:
		_d, ok := d.(float32)
		if !ok {
			return errors.NewStError(1002)
		}
		return bf.writeFloat(_d)
	case Double:
		_d, ok := d.(float64)
		if !ok {
			return errors.NewStError(1002)
		}
		return bf.writeDouble(_d)
	case String:
		_d, ok := d.(string)
		if !ok {
			return errors.NewStError(1002)
		}
		// length
		if err := bf.WriteLength(len(_d)); err != nil {
			return err
		}
		// data
		return bf.writeString(_d)
	case Struct:
		_d, ok := d.(StStruct)
		if !ok {
			return errors.NewStError(1002)
		}
		return _d.WriterDataBuf(bf)
	default:
		return errors.NewStError(1001)
	}
}

func (bf *StBuffer) ReadTag() (tg byte, err error) {
	return bf.readByte()
}

func (bf *StBuffer) ReadDataType() (tp StProtocolType, err error) {
	_d, err := bf.readByte()
	if StTypeMap[_d] == Unknown {
		return Unknown, errors.NewStError(1002)
	}
	return StTypeMap[_d], err
}

func (bf *StBuffer) ReadLength() (l int, err error) {
	bl, err := bf.readByte()
	if err != nil {
		return 0, err
	}
	var bs []byte
	switch bl {
	case 0:
		l=0
	case 1:
		bs = make([]byte, 1)
		if _, err := bf.buf.Read(bs); err != nil {
			return 0, errors.NewStError(1004, err)
		}
		l = int(bs[0])
	case 2:
		bs = make([]byte, 2)
		if _, err := bf.buf.Read(bs); err != nil {
			return 0, errors.NewStError(1004, err)
		}
		l = int(bs[1]) | int(bs[0])<<8
	case 3:
		bs = make([]byte, 3)
		if _, err := bf.buf.Read(bs); err != nil {
			return 0, errors.NewStError(1004, err)
		}
		l = int(bs[2]) | int(bs[1])<<8 | int(bs[0])<<16
	case 4:
		bs = make([]byte, 4)
		if _, err := bf.buf.Read(bs); err != nil {
			return 0, errors.NewStError(1004, err)
		}
		l = int(bs[3]) | int(bs[2])<<8 | int(bs[1])<<16 | int(bs[0])<<24
	case 5:
		bs = make([]byte, 5)
		if _, err := bf.buf.Read(bs); err != nil {
			return 0, errors.NewStError(1004, err)
		}
		l = int(bs[4]) | int(bs[3])<<8 | int(bs[2])<<16 | int(bs[1])<<24 | int(bs[0])<<32
	case 6:
		bs = make([]byte, 6)
		if _, err := bf.buf.Read(bs); err != nil {
			return 0, errors.NewStError(1004, err)
		}
		l = int(bs[5]) | int(bs[4])<<8 | int(bs[3])<<16 | int(bs[2])<<24 | int(bs[1])<<32 | int(bs[0])<<40
	default:
		return 0, errors.NewStError(1005)
	}

	return l, err
}

func (bf *StBuffer) ReadDataBuf(tp StProtocolType) (d interface{}, err error) {
	switch tp {
	case Byte:
		return bf.readByte()
	case Bool:
		return bf.readBool()
	case Int:
		_d, err := bf.readInt()
		return int(_d), err
	case Long:
		return bf.readLong()
	case Float:
		return bf.readFloat()
	case Double:
		return bf.readDouble()
	case String:
		// length
		l, err := bf.ReadLength()
		if err != nil {
			return "", err
		}
		// data
		return bf.readString(l)
	case Struct:
		_d, ok := d.(StStruct)
		if !ok {
			return nil, errors.NewStError(1002)
		}
		err = _d.ReadDataBuf(bf)
		d = _d
	default:
		return nil, errors.NewStError(1001)
	}
	return d, nil
}

func (bf *StBuffer) writeByte(d byte) error {
	if err := bf.buf.WriteByte(d); err != nil {
		return errors.NewStError(1003, err)
	}
	return nil
}

func (bf *StBuffer) writeBool(d bool) error {
	if d {
		if err := bf.buf.WriteByte(1); err != nil {
			return errors.NewStError(1003, err)
		}
	} else {
		if err := bf.buf.WriteByte(0); err != nil {
			return errors.NewStError(1003, err)
		}
	}
	return nil
}

func (bf *StBuffer) writeInt(d int32) error {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(d))
	if _, err := bf.buf.Write(bs); err != nil {
		return errors.NewStError(1003, err)
	}
	return nil
}

func (bf *StBuffer) writeLong(d int64) error {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(d))
	if _, err := bf.buf.Write(bs); err != nil {
		return errors.NewStError(1003, err)
	}
	return nil
}

func (bf *StBuffer) writeFloat(d float32) error {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, math.Float32bits(d))
	if _, err := bf.buf.Write(bs); err != nil {
		return errors.NewStError(1003, err)
	}
	return nil
}

func (bf *StBuffer) writeDouble(d float64) error {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, math.Float64bits(d))
	if _, err := bf.buf.Write(bs); err != nil {
		return errors.NewStError(1003, err)
	}
	return nil
}

func (bf *StBuffer) writeString(d string) error {
	if _, err := bf.buf.WriteString(d); err != nil {
		return errors.NewStError(1003, err)
	}
	return nil
}

func (bf *StBuffer) readByte() (d byte, err error) {
	d, err = bf.buf.ReadByte()
	if err != nil {
		return 0, errors.NewStError(1004, err)
	}
	return d, nil
}

func (bf *StBuffer) readBool() (d bool, err error) {
	_d, err := bf.buf.ReadByte()
	if err != nil {
		return false, errors.NewStError(1004, err)
	}
	if _d == 1 {
		d = true
	} else if _d == 0 {
		d = false
	} else {
		return false, errors.NewStError(1004, err)
	}
	return d, nil
}

func (bf *StBuffer) readInt() (d int32, err error) {
	bs := make([]byte, 4)
	if _, err = bf.buf.Read(bs); err != nil {
		return d, errors.NewStError(1004, err)
	}
	d = int32(binary.BigEndian.Uint32(bs))
	return d, nil
}

func (bf *StBuffer) readLong() (d int64, err error) {
	bs := make([]byte, 8)
	if _, err = bf.buf.Read(bs); err != nil {
		return d, errors.NewStError(1004, err)
	}
	d = int64(binary.BigEndian.Uint64(bs))
	return d, nil
}

func (bf *StBuffer) readFloat() (d float32, err error) {
	bs := make([]byte, 4)
	if _, err = bf.buf.Read(bs); err != nil {
		return d, errors.NewStError(1004, err)
	}
	d = math.Float32frombits(binary.BigEndian.Uint32(bs))
	return d, nil
}

func (bf *StBuffer) readDouble() (d float64, err error) {
	bs := make([]byte, 8)
	if _, err = bf.buf.Read(bs); err != nil {
		return d, errors.NewStError(1004, err)
	}
	d = math.Float64frombits(binary.BigEndian.Uint64(bs))
	return d, nil
}

func (bf *StBuffer) readString(l int) (d string, err error) {
	bs := make([]byte, l)
	if _, err = bf.buf.Read(bs); err != nil {
		return "", errors.NewStError(1004, err)
	}
	return string(bs), nil
}

func (bf *StBuffer) Bytes() []byte {
	return bf.buf.Bytes()
}

func (bf *StBuffer) ReSet(bs []byte) {
	bf.buf.Reset()
	bf.buf = bytes.NewBuffer(bs)
}

func NewStBuffer(bs []byte) *StBuffer {
	return &StBuffer{
		buf: bytes.NewBuffer(bs),
	}
}
