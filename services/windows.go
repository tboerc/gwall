package services

import (
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func StopDivert() (err error) {
	m, err := mgr.Connect()
	if err != nil {
		return
	}
	defer m.Disconnect()

	s, err := m.OpenService("WinDivert")
	if err != nil {
		return nil
	}
	defer s.Close()

	_, err = s.Control(svc.Stop)
	if err != nil {
		return
	}

	return
}
