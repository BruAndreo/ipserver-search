package cli

import (
	"fmt"
	"ipserver-search/internal/app/domain"
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

	err, ips := domain.SearchIpByHost(host)
	if err != nil {
		fmt.Println("Erro:", err.Error())
		os.Exit(0)
	}

	display(ips)
}

func searchServerByHost(context *cli.Context) {
	host := context.String(FLAG)

	err, servers := domain.SearchServerByHost(host)
	if err != nil {
		fmt.Println("Erro:", err.Error())
		os.Exit(0)
	}

	display(servers)
}

func display(itens []string) {
	for _, item := range itens {
		fmt.Println("=>", item)
	}
}
