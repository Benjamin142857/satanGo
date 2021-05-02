package errors

import (
	"fmt"
	"runtime"
)

type StError struct {
	errCode int
	errMsg  string
	tbRoot  bool
}

func (e *StError) Error() string {
	if e.tbRoot {
		return fmt.Sprintf("StError[%v] %v", e.errCode, e.errMsg)
	}
	return e.errMsg
}

func (e *StError) Code() int {
	return e.errCode
}

func (e *StError) Equal(_e *StError) bool {
	return e.errCode == _e.errCode
}

func NewStError(code int, msg string) *StError {
	return &StError{
		errCode: code,
		errMsg:  msg,
		tbRoot: true,
	}
}

func StErrorAppendMsg(e *StError, msg string) *StError {
	return NewStError(e.errCode, fmt.Sprintf("%v, %v", e.errMsg, msg))
}

func TrackBackStError(e *StError) *StError {
	var tb string
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		tb = "[get runtime.Caller error]"
	}
	pFunc := runtime.FuncForPC(pc)
	tb = fmt.Sprintf("File \"%v\", line %v, in <%v>", file, line, pFunc.Name())
	return &StError{
		errCode: e.errCode,
		errMsg: fmt.Sprintf("%v\n%v", tb, e.Error()),
		tbRoot: false,
	}
}

func TrackBackStErrorWithMsg(e *StError, msg string) *StError {
	var tb string
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		tb = "[get runtime.Caller error]"
	}
	pFunc := runtime.FuncForPC(pc)

	tb = fmt.Sprintf("File \"%v\", line %v, in <%v>\n\t%v", file, line, pFunc.Name(), msg)
	return &StError{
		errCode: e.errCode,
		errMsg: fmt.Sprintf("%v\n%v", tb, e.Error()),
		tbRoot: false,
	}
}
