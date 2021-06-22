package services

import (
	"io/ioutil"
	"net"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type response struct {
	IP net.IP `json:"query"`
}

func PublicIP() (i net.IP) {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return
	}
	defer req.Body.Close()

	b, _ := ioutil.ReadAll(req.Body)

	var p response
	jsoniter.Unmarshal(b, &p)

	return p.IP
}

func LocalIP() (i net.IP, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP, nil
}
