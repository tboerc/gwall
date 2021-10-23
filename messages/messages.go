package messages

import "errors"

var (
	ErrFilter               = errors.New("error: no admin privileges")
	ErrLocalIP              = errors.New("error: could not get local ip")
	ErrPublicIP             = errors.New("error: could not get local ip")
	ErrInvalidIP            = errors.New("error: ip is not valid")
	ErrNoIP                 = errors.New("error: no ip provided")
	ErrConfigWrite          = errors.New("error: config file write error")
	ErrExecPath             = errors.New("error: could not get executable path")
	ErrWindDivertStop       = errors.New("error: could not stop windivert service")
	ErrWindDivertNotRunning = errors.New("error: windivert service is not running")
	ErrServiceConnect       = errors.New("error: could not connect to service control manager")
)
