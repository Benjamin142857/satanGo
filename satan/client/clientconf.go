package client

import (
	"regexp"
	"satanGo/satan/errors"
	"time"
)

var regEndPoint = regexp.MustCompile("^(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])\\.(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])\\.(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])\\.(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5]):([0-9]|[1-9]\\d|[1-9]\\d{2}|[1-9]\\d{3}|[1-5]\\d{4}|6[0-4]\\d{3}|65[0-4]\\d{2}|655[0-2]\\d|6553[0-5])$")

type StClientConf struct {
	timeout  time.Duration
	endpoint string
}

type StInvokeOption func(si *stInvoker) error

func (sc *StClient) WithTimeout(t time.Duration) StInvokeOption {
	return func(si *stInvoker) error {
		if t<0 {
			return errors.ErrWithTimeout
		}
		si.timeout = t
		return nil
	}
}

func (sc *StClient) WithEndPoint(s string) StInvokeOption {
	return func(si *stInvoker) error {
		if !regEndPoint.MatchString(s) {
			return errors.ErrWithEndPoint
		}
		si.endpoint = s
		return nil
	}
}