package client

import (
	"context"
	"satanGo/satan/transport"
	"time"
)

type StClient struct {
	conf *StClientConf
}

type stInvoker struct {
	*StClientConf
}

func (sc *StClient) newStInvoker(opts ...StInvokeOption) (*stInvoker, error) {
	_cf := *sc.conf
	si := &stInvoker{&_cf}
	for _, opt := range opts {
		if err := opt(si); err != nil {
			return nil, err
		}
	}
	return si, nil
}

func (sc *StClient) StInvoke(reqPkg *transport.RequestPackage, rspPkg *transport.ResponsePackage, opts ...StInvokeOption) {
	si, err := sc.newStInvoker(opts...)

	// client conf error
	if err != nil {
		return
	}

	// get server endpoint from stKeeper
	if si.endpoint == "" {

	}

	// net.Dial
	tOut := time.Now().Add(si.timeout)
	reqPkg.Timeout = tOut.UnixNano() / 1000000
	ctx, cancel := context.WithDeadline(context.TODO(), tOut)
	transport.StProtocolDial(reqPkg, rspPkg, ctx, si.timeout, si.endpoint)
	cancel()
}