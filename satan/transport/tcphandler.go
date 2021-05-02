package transport

import (
	"context"
	"net"
	"satanGo/satan/errors"
	"time"
)

func StProtocolDial(
	reqPkg *RequestPackage, rspPkg *ResponsePackage, ctx context.Context,
	timeout time.Duration,
	endpoint string,
) {
	// tcp dial
	conn, err := net.DialTimeout("tcp", endpoint, timeout)
	if err != nil {
		_err := errors.StErrorAppendMsg(errors.ErrTcpDial, err.Error())
		rspPkg.StatusCode = 1
		rspPkg.ErrCode = _err.Code()
		rspPkg.ErrMsg = _err.Error()
		return
	}
	defer func() { _ = conn.Close() }()

	// write buf
	bs, err := reqPkg.WriteToBytes()
	if err != nil {
		rspPkg.StatusCode = 1
		rspPkg.ErrCode = err.Code()
		rspPkg.ErrMsg = err.Error()
		return
	}
	if _, err := conn.Write(bs); err != nil {
		_err := errors.StErrorAppendMsg(errors.ErrTcpDial, err.Error())
		rspPkg.StatusCode = 1
		rspPkg.ErrCode = _err.Code()
		rspPkg.ErrMsg = _err.Error()
		return
	}
}
