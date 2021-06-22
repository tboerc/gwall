package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/tboerc/gwall/services"
)

type Allowed struct {
	LocalIP  net.IP `json:"local_ip,omitempty"`
	PublicIP net.IP `json:"public_ip,omitempty"`
}

type Whitelist []*Allowed

type Config struct {
	Whitelist Whitelist `json:"whitelist"`
}

func (w Whitelist) String() string {
	s := "["
	for i, v := range w {
		if i > 0 {
			s += ", "
		}
		if v.PublicIP != nil && v.LocalIP != nil {
			s += fmt.Sprintf("%s, %s", v.PublicIP.String(), v.LocalIP.String())
		} else if v.PublicIP != nil {
			s += v.PublicIP.String()
		} else if v.LocalIP != nil {
			s += v.LocalIP.String()
		}
	}
	return s + "]"
}

func ReadConfig() (c *Config) {
	file, _ := ioutil.ReadFile(cp)
	jsoniter.Unmarshal([]byte(file), &c)

	return
}

func WriteConfig(c *Config) (err error) {
	file, _ := jsoniter.MarshalIndent(&c, "", " ")

	err = ioutil.WriteFile(cp, file, 0644)
	if err != nil {
		return errConfigWrite
	}

	return
}

func GetConfig() (c *Config, err error) {
	c = ReadConfig()
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		l, er := services.LocalIP()
		if er != nil {
			err = er
		}
		c.Whitelist = append(c.Whitelist, &Allowed{LocalIP: l})
	}()
	go func() {
		defer wg.Done()
		p := services.PublicIP()
		if p != nil {
			c.Whitelist = append(c.Whitelist, &Allowed{PublicIP: p})
		}
	}()

	wg.Wait()

	return
}
