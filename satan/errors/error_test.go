package errors

import (
	"fmt"
	"testing"
)

func f1() *StError {
	return ErrEncodeType
}

func f2() *StError {
	if err := f1(); err != nil {
		return TrackBackStErrorWithMsg(err, "f2 occur error.")
	}
	return nil
}

func f3() *StError {
	if err := f2(); err != nil {
		return TrackBackStErrorWithMsg(err, "f3 occur error.")
	}
	return nil
}

func fTest() *StError {
	if err := f3(); err != nil {
		return TrackBackStError(err)
	}
	return nil
}

func TestNewError(t *testing.T) {
	if err := fTest(); err != nil {
		fmt.Println(err)
	}
}
