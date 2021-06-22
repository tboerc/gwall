package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/tboerc/divert-go"
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
		return
	}

	defer dh.End()
	defer fmt.Print("Filtering done\n\n")

	fmt.Println("Filtering started...")

	go func() {
		for packet := range pc {
			for _, v := range *l {
				if v.PublicIP != nil && packet.SrcIP().Equal(v.PublicIP) {
					fmt.Println("Packet sent to", packet.SrcIP().String())
					dh.Send(packet.Raw, packet.Addr)
				} else if v.LocalIP != nil && packet.SrcIP().Equal(v.LocalIP) {
					fmt.Println("Packet sent to", packet.SrcIP().String())
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
		ip := services.PublicIP()
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

		fmt.Fprintln(w, "Index\tPublic IP\tLocal IP")
		for i, v := range config.Whitelist {
			fmt.Fprintf(w, "%d\t%s\t%s\n", i, v.PublicIP.String(), v.LocalIP.String())
		}

		return
	},
}

var add = &cli.Command{
	Name:      "add",
	ArgsUsage: "PUBLIC_IP LOCAL_IP",
	Usage:     "Add a IP to whitelist",
	Action: func(c *cli.Context) (err error) {
		if c.Args().Len() == 0 {
			return errNoIP
		}

		p := net.ParseIP(c.Args().Get(0))
		l := net.ParseIP(c.Args().Get(1))

		config := ReadConfig()
		config.Whitelist = append(config.Whitelist, &Allowed{PublicIP: p, LocalIP: l})

		err = WriteConfig(config)

		return
	},
}

var remove = &cli.Command{
	Name:      "remove",
	ArgsUsage: "INDEX",
	Usage:     "Removes an IP from whitelist based on its index. To get the index use \"gwall list\"",
	Action: func(c *cli.Context) (err error) {
		if c.Args().Len() == 0 {
			return errNoIndex
		}

		i, err := strconv.Atoi(c.Args().Get(0))
		if err != nil {
			return errInvalidIndex
		}

		config := ReadConfig()
		config.Whitelist = append(config.Whitelist[:i], config.Whitelist[i+1:]...)

		err = WriteConfig(config)

		return
	},
}

var solo = &cli.Command{
	Name:  "solo",
	Usage: "Start gwall in solo mode",
	Action: func(_ *cli.Context) (err error) {
		var config Config

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
	Usage: "Start gwall in whitelist mode",
	Action: func(_ *cli.Context) (err error) {
		config, _ := GetConfig()

		fmt.Print("Loaded whitelist: ", config.Whitelist.String(), "\n\n")

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
	Usage: "Stop WinDivert service",
	Action: func(_ *cli.Context) (err error) {
		err = services.StopDivert()
		return
	},
}
