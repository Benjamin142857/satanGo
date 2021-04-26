package test001

import (
	"satanGo/satan/errors"
	"satanGo/satan/protocol"
)

type Person struct {
	Name string `json:"name"`
	Hobby []string `json:"hobby"`
	Family map[string]*Person `json:"family"`
}

func (st *Person) WriteDataBuf(bf *protocol.StBuffer) error {
	if err := bf.WriteStructLength(3); err != nil {
		return err
	}

	// tag 0
	if err := bf.WriteTag(0); err != nil  {
		return err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return err
	}
	if err := bf.WriteDataBuf(protocol.String, st.Name); err != nil {
		return err
	}

	// tag 1
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

	// tag 2
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
			d, err := bf.ReadDataBuf(protocol.String)
			if err != nil {
				return err
			}
			_d, ok := d.(string)
			if !ok {
				return errors.NewStError(1004)
			}
			st.Name = _d
		case byte(1):
			sl1 := make([]string, 0)
			if _, err := bf.ReadDataType(); err != nil {
				return err
			}
			l1, err := bf.ReadLength()
			if err != nil {
				return err
			}
			for i1:=0; i1<l1; i1++ {
				d, err := bf.ReadDataBuf(protocol.String)
				if err != nil {
					return err
				}
				_d, ok := d.(string)
				if !ok {
					return errors.NewStError(1004)
				}
				sl1 = append(sl1, _d)
			}
			st.Hobby = sl1
		case byte(2):
			mp1 := make(map[string]*Person)
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
			for i1:=0; i1<l1; i1++ {
				k, err := bf.ReadDataBuf(protocol.String)
				if err != nil {
					return err
				}
				_k, ok := k.(string)
				if !ok {
					return errors.NewStError(1004)
				}
				s := NewPerson()
				if err := s.ReadDataBuf(bf); err != nil {
					return err
				}
				mp1[_k] = s
			}
			st.Family = mp1
		}
	}

	return nil
}
func NewPerson() *Person {
	return &Person{
		Name: "",
		Hobby: make([]string, 0),
		Family: make(map[string]*Person),
	}
}
