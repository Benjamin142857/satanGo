package errors

var (
	// protocol
	ErrInvalidType  = NewStError(1001, "Invalid StProtocol DataType")
	ErrEncodeType   = NewStError(1002, "StProtocol encode data type error")
	ErrEncodeBuf    = NewStError(1003, "StProtocol encode buffer error")
	ErrDecodeBuf    = NewStError(1004, "StProtocol decode buffer error")
	ErrExceedMaxLen = NewStError(1005, "The maximum support length of StProtocol is exceeded")
	ErrMapKey       = NewStError(1006, "StProtocol Map key not support List, Map, Struct")

	// Client
	ErrWithTimeout  = NewStError(4001, "WithTimeout invalid params")
	ErrWithEndPoint = NewStError(4002, "WithEndPoint invalid params")
	ErrTcpDial      = NewStError(4011, "Tcp Dial error")
)
