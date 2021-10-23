package services

import (
	"io/ioutil"
	"net"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/tboerc/gwall/messages"
)

type response struct {
	IP net.IP `json:"query"`
}

func PublicIP() (net.IP, error) {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, messages.ErrPublicIP
	}
	defer req.Body.Close()

	b, _ := ioutil.ReadAll(req.Body)

	var p response
	jsoniter.Unmarshal(b, &p)

	return p.IP, nil
}

func LocalIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, messages.ErrLocalIP
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP, nil
}
