package main

import (
	"fmt"
	"net"
	"os"
	"text/tabwriter"
	"time"

	"github.com/tboerc/divert-go"
	"github.com/tboerc/gwall/messages"
	"github.com/tboerc/gwall/services"
	"github.com/urfave/cli/v2"
)

const packetFilter = "(udp.SrcPort == 6672 or udp.DstPort == 6672) and ip"

var dh *divert.Handle

func onDivert() (pc chan *divert.Packet, err error) {
	dh, err = divert.Open(packetFilter, divert.LayerNetwork, divert.PriorityDefault, divert.FlagDefault)
	if err != nil {
		return
	}

	pc, err = dh.Packets()
	if err != nil {
		return
	}

	return
}

func filter(l *Whitelist) (err error) {
	pc, err := onDivert()
	if err != nil {
		return messages.ErrFilter
	}

	defer dh.End()

	go func() {
		for packet := range pc {
			for _, v := range *l {
				if v.IP != nil && packet.SrcIP().Equal(v.IP) {
					dh.Send(packet.Raw, packet.Addr)
				}
			}
		}
	}()

	time.Sleep(15 * time.Second)

	return
}

var ip = &cli.Command{
	Name:  "ip",
	Usage: "Returns your public IP",
	Action: func(_ *cli.Context) (err error) {
		ip, err := services.PublicIP()
		if err != nil {
			return
		}

		fmt.Println(ip)

		return
	},
}

var list = &cli.Command{
	Name:  "list",
	Usage: "List all IPs",
	Action: func(_ *cli.Context) (err error) {
		config := ReadConfig()
		if len(config.Whitelist) == 0 {
			fmt.Println("No entries. To add one use \"gwall add\"")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', 0)
		defer w.Flush()

		fmt.Fprintln(w, "Index\tIP")
		for i, v := range config.Whitelist {
			fmt.Fprintf(w, "%d\t%s\n", i, v.IP.String())
		}

		return
	},
}

var add = &cli.Command{
	Name:      "add",
	ArgsUsage: "IP",
	Usage:     "Add a IP to whitelist",
	Action: func(c *cli.Context) (err error) {
		if c.Args().Len() == 0 {
			return messages.ErrNoIP
		}

		p := net.ParseIP(c.Args().Get(0))
		if p == nil {
			return messages.ErrInvalidIP
		}

		config := ReadConfig()
		config.Whitelist = append(config.Whitelist, &Allowed{IP: p})

		err = WriteConfig(config)
		if err != nil {
			return
		}

		fmt.Printf("Successfully added %s to whitelist\n", p.String())

		return
	},
}

var remove = &cli.Command{
	Name:      "remove",
	ArgsUsage: "IP",
	Usage:     "Removes an IP from whitelist. To see the whitelist use \"gwall list\"",
	Action: func(c *cli.Context) (err error) {
		if c.Args().Len() == 0 {
			return messages.ErrNoIP
		}

		r := net.ParseIP(c.Args().Get(0))
		if r == nil {
			return messages.ErrInvalidIP
		}

		config := ReadConfig()
		var w Whitelist

		for _, p := range config.Whitelist {
			if !p.IP.Equal(r) {
				w = append(w, p)
			}
		}

		config.Whitelist = w

		err = WriteConfig(config)
		if err != nil {
			return
		}

		fmt.Printf("Successfully removed %s from whitelist\n", r.String())

		return
	},
}

var solo = &cli.Command{
	Name:  "solo",
	Usage: "Start gwall in solo mode. Need admin privileges",
	Action: func(_ *cli.Context) (err error) {
		var config Config

		fmt.Println("Press Ctrl+C to stop...")

		for {
			err = filter(&config.Whitelist)
			if err != nil {
				return
			}

			time.Sleep(15 * time.Second)
		}
	},
}

var whitelist = &cli.Command{
	Name:  "whitelist",
	Usage: "Start gwall in whitelist mode. Need admin privileges",
	Action: func(_ *cli.Context) (err error) {
		config, err := GetConfig()
		if err != nil {
			return
		}

		fmt.Println("Loaded whitelist: ", config.Whitelist.String())
		fmt.Println("Press Ctrl+C to stop...")

		for {
			err = filter(&config.Whitelist)
			if err != nil {
				return
			}

			time.Sleep(15 * time.Second)
		}
	},
}

var stop = &cli.Command{
	Name:  "stop",
	Usage: "Stop WinDivert service. Need admin privileges",
	Action: func(_ *cli.Context) (err error) {
		err = services.StopDivert()
		if err != nil {
			return
		}

		fmt.Println("WinDivert service stoped")

		return
	},
}
