package main

import (
	"io/ioutil"
	"net"
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/tboerc/gwall/messages"
	"github.com/tboerc/gwall/services"
)

type Allowed struct {
	IP net.IP `json:"public_ip,omitempty"`
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
		s += v.IP.String()
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
		return messages.ErrConfigWrite
	}

	return
}

func GetConfig() (c *Config, err error) {
	c = ReadConfig()
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		l, err := services.LocalIP()
		if err == nil {
			c.Whitelist = append(c.Whitelist, &Allowed{IP: l})
		}
	}()
	go func() {
		defer wg.Done()
		p, err := services.PublicIP()
		if err == nil {
			c.Whitelist = append(c.Whitelist, &Allowed{IP: p})
		}
	}()

	wg.Wait()

	return
}
