package errors

import "fmt"

var StErrMap = map[int]string{
	1001: "Invalid StProtocolType",
	1002: "StProtocol encode data type error",
	1003: "StProtocol encode buffer error",
	1004: "StProtocol decode buffer error",
	1005: "The maximum support length of StProtocol is exceeded",
	1006: "StProtocol Map key not support List, Map, Struct",
}

type StError struct {
	ErrCode         int
	ErrMsg          string
	ErrTrackBackMsg string
}

func (e *StError) Error() string {
	return fmt.Sprintf("StError[%v] %v%v", e.ErrCode, e.ErrMsg, e.ErrTrackBackMsg)
}

func NewStError(code int, tbErr ...error) *StError {
	tbMsg := ""
	for _, e := range tbErr {
		if e != nil {
			tbMsg += "\n" + e.Error()
		}
	}

	if errMsg := StErrMap[code]; errMsg == "" {
		return &StError{
			ErrCode: 0,
			ErrMsg: "unknown error",
			ErrTrackBackMsg: tbMsg,
		}
	}
	return &StError{
		ErrCode:         code,
		ErrMsg:          StErrMap[code],
		ErrTrackBackMsg: tbMsg,
	}
}