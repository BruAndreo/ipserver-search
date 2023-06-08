package app

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const FLAG string = "host"

var HOST_FLAG = cli.StringFlag{
	Name: FLAG,
}

var IP_COMMAND cli.Command = cli.Command{
	Name:   "ip",
	Usage:  "Search IPs by host",
	Flags:  []cli.Flag{HOST_FLAG},
	Action: searchIpByHost,
}

var SERVER_COMMAND cli.Command = cli.Command{
	Name:   "server",
	Usage:  "Search Server by host",
	Flags:  []cli.Flag{HOST_FLAG},
	Action: searchServerByHost,
}

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "IPSever-Search"
	app.Usage = "Search IP and Server of a domain"

	app.Commands = []cli.Command{IP_COMMAND, SERVER_COMMAND}

	return app
}

func searchIpByHost(context *cli.Context) {
	host := context.String(FLAG)

	if isInvalidHost(host) {
		fmt.Println("Invalid host passed")
		os.Exit(0)
	}

	ips := SearchIP(host)

	for _, ip := range ips {
		fmt.Println("=>", ip)
	}
}

func searchServerByHost(context *cli.Context) {
	host := context.String(FLAG)

	if isInvalidHost(host) {
		fmt.Println("Invalid host passed")
		os.Exit(0)
	}

	servers := SearchServer(host)

	for _, server := range servers {
		fmt.Println("=>", server)
	}
}

func isInvalidHost(host string) bool {
	return len(host) < 4
}
