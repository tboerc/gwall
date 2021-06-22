package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/tboerc/divert-go"
	"github.com/tboerc/gwall/services"
	"github.com/urfave/cli/v2"
)

var (
	cp string
)

func onClose(c *chan os.Signal) {
	<-*c
	if dh.Open == divert.HandleOpen {
		dh.End()
	}
	services.StopDivert()
	os.Exit(0)
}

func main() {
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(errExecPath)
	}

	cp = filepath.Join(filepath.Dir(ex), "config.json")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go onClose(&c)

	app := &cli.App{
		Name:    "gwall",
		Version: "1.0.0",
		Usage:   "Firewall for public, solo or whitelisted, sessions on Grand Theft Auto V",
		Commands: []*cli.Command{
			ip, list, add, remove, stop,
			{
				Name:        "start",
				Usage:       "Start gwall in solo or whitelist mode",
				Subcommands: []*cli.Command{solo, whitelist},
				HideHelp:    true,
			},
		},
		HideHelp: true,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
