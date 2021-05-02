package test002

import (
	"satanGo/satan/errors"
	"satanGo/satan/protocol"
)

type Person struct {
	Name   string             `json:"name"`
	Hobby  []string           `json:"hobby"`
	Family map[string]*Person `json:"family"`
	Lover  *Person            `json:"lover"`
}

func (st *Person) WriteDataBuf(bf *protocol.StBuffer) error {
	if err := bf.WriteStructLength(4); err != nil {
		return err
	}

	if err := bf.WriteTag(0); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return err
	}
	if err := bf.WriteDataBuf(protocol.String, st.Name); err != nil {
		return err
	}

	if err := bf.WriteTag(1); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.List); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return err
	}
	if err := bf.WriteLength(len(st.Hobby)); err != nil {
		return err
	}
	for _, e1 := range st.Hobby {
		if err := bf.WriteDataBuf(protocol.String, e1); err != nil {
			return err
		}
	}

	if err := bf.WriteTag(2); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.Map); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.Struct); err != nil {
		return err
	}
	if err := bf.WriteLength(len(st.Family)); err != nil {
		return err
	}
	for k1, v1 := range st.Family {
		if err := bf.WriteDataBuf(protocol.String, k1); err != nil {
			return err
		}
		if err := v1.WriteDataBuf(bf); err != nil {
			return err
		}
	}

	if err := bf.WriteTag(3); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.Struct); err != nil {
		return err
	}
	if err := st.Lover.WriteDataBuf(bf); err != nil {
		return err
	}

	return nil
}
func (st *Person) ReadDataBuf(bf *protocol.StBuffer) error {
	l, err := bf.ReadStructLength()
	if err != nil {
		return err
	}

	for i := byte(0); i < l; i++ {
		tg, err := bf.ReadTag()
		if err != nil {
			return err
		}
		if _, err := bf.ReadDataType(); err != nil {
			return err
		}

		switch tg {
		case byte(0):
			_d1, err := bf.ReadDataBuf(protocol.String)
			if err != nil {
				return err
			}
			d1, ok := _d1.(string)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.Name = d1
		case byte(1):
			if _, err := bf.ReadDataType(); err != nil {
				return err
			}
			l1, err := bf.ReadLength()
			if err != nil {
				return err
			}
			d1 := make([]string, l1)
			for i1 := 0; i1 < l1; i1++ {
				_e2, err := bf.ReadDataBuf(protocol.String)
				if err != nil {
					return err
				}
				e2, ok := _e2.(string)
				if !ok {
					return errors.ErrDecodeBuf
				}
				d1[i1] = e2
			}
			st.Hobby = d1
		case byte(2):
			d1 := make(map[string]*Person)
			if _, err := bf.ReadDataType(); err != nil {
				return err
			}
			if _, err := bf.ReadDataType(); err != nil {
				return err
			}
			l1, err := bf.ReadLength()
			if err != nil {
				return err
			}
			for i1 := 0; i1 < l1; i1++ {
				_k2, err := bf.ReadDataBuf(protocol.String)
				if err != nil {
					return err
				}
				k2, ok := _k2.(string)
				if !ok {
					return errors.ErrDecodeBuf
				}
				v2 := NewPerson()
				if err := v2.ReadDataBuf(bf); err != nil {
					return err
				}
				d1[k2] = v2
			}
			st.Family = d1
		case byte(3):
			d1 := NewPerson()
			if err := d1.ReadDataBuf(bf); err != nil {
				return err
			}
			st.Lover = d1
		}

	}
	return nil
}
func NewPerson() *Person {
	return &Person{
		Name:   "",
		Hobby:  make([]string, 0),
		Family: make(map[string]*Person),
		Lover:  nil,
	}
}

type Mapper struct {
	Gap    [][]int            `json:"gap"`
	HhTest map[string][][]int `json:"hhTest"`
}

func (st *Mapper) WriteDataBuf(bf *protocol.StBuffer) error {
	if err := bf.WriteStructLength(2); err != nil {
		return err
	}

	if err := bf.WriteTag(0); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.List); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.List); err != nil {
		return err
	}
	if err := bf.WriteLength(len(st.Gap)); err != nil {
		return err
	}
	for _, e1 := range st.Gap {
		if err := bf.WriteDataType(protocol.Int); err != nil {
			return err
		}
		if err := bf.WriteLength(len(e1)); err != nil {
			return err
		}
		for _, e2 := range e1 {
			if err := bf.WriteDataBuf(protocol.Int, e2); err != nil {
				return err
			}
		}
	}

	if err := bf.WriteTag(1); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.Map); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return err
	}
	if err := bf.WriteDataType(protocol.List); err != nil {
		return err
	}
	if err := bf.WriteLength(len(st.HhTest)); err != nil {
		return err
	}
	for k1, v1 := range st.HhTest {
		if err := bf.WriteDataBuf(protocol.String, k1); err != nil {
			return err
		}
		if err := bf.WriteDataType(protocol.List); err != nil {
			return err
		}
		if err := bf.WriteLength(len(v1)); err != nil {
			return err
		}
		for _, e2 := range v1 {
			if err := bf.WriteDataType(protocol.Int); err != nil {
				return err
			}
			if err := bf.WriteLength(len(e2)); err != nil {
				return err
			}
			for _, e3 := range e2 {
				if err := bf.WriteDataBuf(protocol.Int, e3); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
func (st *Mapper) ReadDataBuf(bf *protocol.StBuffer) error {
	l, err := bf.ReadStructLength()
	if err != nil {
		return err
	}

	for i := byte(0); i < l; i++ {
		tg, err := bf.ReadTag()
		if err != nil {
			return err
		}
		if _, err := bf.ReadDataType(); err != nil {
			return err
		}

		switch tg {
		case byte(0):
			if _, err := bf.ReadDataType(); err != nil {
				return err
			}
			l1, err := bf.ReadLength()
			if err != nil {
				return err
			}
			d1 := make([][]int, l1)
			for i1 := 0; i1 < l1; i1++ {
				if _, err := bf.ReadDataType(); err != nil {
					return err
				}
				l2, err := bf.ReadLength()
				if err != nil {
					return err
				}
				e2 := make([]int, l2)
				for i2 := 0; i2 < l2; i2++ {
					_e3, err := bf.ReadDataBuf(protocol.Int)
					if err != nil {
						return err
					}
					e3, ok := _e3.(int)
					if !ok {
						return errors.ErrDecodeBuf
					}
					e2[i2] = e3
				}
				d1[i1] = e2
			}
			st.Gap = d1
		case byte(1):
			d1 := make(map[string][][]int)
			if _, err := bf.ReadDataType(); err != nil {
				return err
			}
			if _, err := bf.ReadDataType(); err != nil {
				return err
			}
			l1, err := bf.ReadLength()
			if err != nil {
				return err
			}
			for i1 := 0; i1 < l1; i1++ {
				_k2, err := bf.ReadDataBuf(protocol.String)
				if err != nil {
					return err
				}
				k2, ok := _k2.(string)
				if !ok {
					return errors.ErrDecodeBuf
				}
				if _, err := bf.ReadDataType(); err != nil {
					return err
				}
				l2, err := bf.ReadLength()
				if err != nil {
					return err
				}
				v2 := make([][]int, l2)
				for i2 := 0; i2 < l2; i2++ {
					if _, err := bf.ReadDataType(); err != nil {
						return err
					}
					l3, err := bf.ReadLength()
					if err != nil {
						return err
					}
					e3 := make([]int, l3)
					for i3 := 0; i3 < l3; i3++ {
						_e4, err := bf.ReadDataBuf(protocol.Int)
						if err != nil {
							return err
						}
						e4, ok := _e4.(int)
						if !ok {
							return errors.ErrDecodeBuf
						}
						e3[i3] = e4
					}
					v2[i2] = e3
				}
				d1[k2] = v2
			}
			st.HhTest = d1
		}

	}
	return nil
}
func NewMapper() *Mapper {
	return &Mapper{
		Gap:    make([][]int, 0),
		HhTest: make(map[string][][]int),
	}
}
