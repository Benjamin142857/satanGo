package transport

import (
	"satanGo/satan/errors"
	"satanGo/satan/protocol"
	"satanGo/satan/settings"
)

type RequestPackage struct {
	Protocol  string `json:"Protocol"`
	Version   string `json:"Version"`
	Timeout   int64  `json:"Timeout"`
	Data      []byte `json:"Data"`
	Origin    string `json:"Origin"`
	Function  string `json:"Function"`
	Host      string `json:"Host"`
	Port      int    `json:"Port"`
	ProxyInfo string `json:"ProxyInfo"`
}

func (st *RequestPackage) WriteToBytes() ([]byte, error) {
	bf := protocol.NewStBuffer([]byte{})
	
	if err := bf.WriteStructLength(9); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(0); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.String, st.Protocol); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(1); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.String, st.Version); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(2); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.Long); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.Long, st.Timeout); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(3); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.List); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.Byte); err != nil {
		return nil, err
	}
	if err := bf.WriteLength(len(st.Data)); err != nil {
		return nil, err
	}
	if err := bf.WriteBytes(st.Data); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(4); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.String, st.Origin); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(5); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.String, st.Function); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(6); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.String, st.Host); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(7); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.Int); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.Int, st.Port); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(8); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.String, st.ProxyInfo); err != nil {
		return nil, err
	}

	return bf.Bytes(), nil
}
func (st *RequestPackage) ReadFromBytes(bs []byte) error {
	bf := protocol.NewStBuffer(bs)

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
			st.Protocol = d1
		case byte(1):
			_d1, err := bf.ReadDataBuf(protocol.String)
			if err != nil {
				return err
			}
			d1, ok := _d1.(string)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.Version = d1
		case byte(2):
			_d1, err := bf.ReadDataBuf(protocol.Long)
			if err != nil {
				return err
			}
			d1, ok := _d1.(int64)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.Timeout = d1
		case byte(3):
			if _, err := bf.ReadDataType(); err != nil {
				return err
			}
			l1, err := bf.ReadLength()
			if err != nil {
				return err
			}
			d1, err := bf.ReadBytes(l1)
			if err != nil {
				return err
			}
			st.Data = d1
		case byte(4):
			_d1, err := bf.ReadDataBuf(protocol.String)
			if err != nil {
				return err
			}
			d1, ok := _d1.(string)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.Origin = d1
		case byte(5):
			_d1, err := bf.ReadDataBuf(protocol.String)
			if err != nil {
				return err
			}
			d1, ok := _d1.(string)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.Function = d1
		case byte(6):
			_d1, err := bf.ReadDataBuf(protocol.String)
			if err != nil {
				return err
			}
			d1, ok := _d1.(string)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.Host = d1
		case byte(7):
			_d1, err := bf.ReadDataBuf(protocol.Int)
			if err != nil {
				return err
			}
			d1, ok := _d1.(int)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.Port = d1
		case byte(8):
			_d1, err := bf.ReadDataBuf(protocol.String)
			if err != nil {
				return err
			}
			d1, ok := _d1.(string)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.ProxyInfo = d1
		}
	}
	return nil
}
func NewRequestPackage() *RequestPackage {
	return &RequestPackage{
		Protocol:  "stprotocol",
		Version:   settings.StVersion,
		Timeout:   int64(0),
		Data:      make([]byte, 0),
		Origin:    "",
		Function:  "",
		Host:      "",
		Port:      0,
		ProxyInfo: "",
	}
}
