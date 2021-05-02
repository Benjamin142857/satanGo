package transport

import (
	"satanGo/satan/errors"
	"satanGo/satan/protocol"
	"satanGo/satan/settings"
)

type ResponsePackage struct {
	Protocol   string `json:"Protocol"`
	Version    string `json:"Version"`
	TimeStamp  int64  `json:"TimeStamp"`
	Data       []byte `json:"Data"`
	Host       string `json:"Host"`
	Port       int    `json:"Port"`
	StatusCode int    `json:"StatusCode"`
	ErrCode    int    `json:"ErrCode"`
	ErrMsg     string `json:"ErrMsg"`
}

func (st *ResponsePackage) WriteToBytes() ([]byte, error) {
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
	if err := bf.WriteDataBuf(protocol.Long, st.TimeStamp); err != nil {
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
	if err := bf.WriteDataBuf(protocol.String, st.Host); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(5); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.Int); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.Int, st.Port); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(6); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.Int); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.Int, st.StatusCode); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(7); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.Int); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.Int, st.ErrCode); err != nil {
		return nil, err
	}

	if err := bf.WriteTag(8); err != nil {
		return nil, err
	}
	if err := bf.WriteDataType(protocol.String); err != nil {
		return nil, err
	}
	if err := bf.WriteDataBuf(protocol.String, st.ErrMsg); err != nil {
		return nil, err
	}

	return bf.Bytes(), nil
}
func (st *ResponsePackage) ReadFromBytes(bs []byte) error {
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
			st.TimeStamp = d1
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
			st.Host = d1
		case byte(5):
			_d1, err := bf.ReadDataBuf(protocol.Int)
			if err != nil {
				return err
			}
			d1, ok := _d1.(int)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.Port = d1
		case byte(6):
			_d1, err := bf.ReadDataBuf(protocol.Int)
			if err != nil {
				return err
			}
			d1, ok := _d1.(int)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.StatusCode = d1
		case byte(7):
			_d1, err := bf.ReadDataBuf(protocol.Int)
			if err != nil {
				return err
			}
			d1, ok := _d1.(int)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.ErrCode = d1
		case byte(8):
			_d1, err := bf.ReadDataBuf(protocol.String)
			if err != nil {
				return err
			}
			d1, ok := _d1.(string)
			if !ok {
				return errors.ErrDecodeBuf
			}
			st.ErrMsg = d1
		}

	}
	return nil
}
func NewResponsePackage() *ResponsePackage {
	return &ResponsePackage{
		Protocol:  "stprotocol",
		Version:   settings.StVersion,
		TimeStamp:  int64(0),
		Data:       make([]byte, 0),
		Host:       "",
		Port:       0,
		StatusCode: 0,
		ErrCode:    0,
		ErrMsg:     "",
	}
}
