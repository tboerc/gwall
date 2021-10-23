package services

import (
	"github.com/tboerc/gwall/messages"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func StopDivert() (err error) {
	m, err := mgr.Connect()
	if err != nil {
		return messages.ErrServiceConnect
	}
	defer m.Disconnect()

	s, err := m.OpenService("WinDivert")
	if err != nil {
		return messages.ErrWindDivertNotRunning
	}
	defer s.Close()

	_, err = s.Control(svc.Stop)
	if err != nil {
		return messages.ErrWindDivertStop
	}

	return
}
