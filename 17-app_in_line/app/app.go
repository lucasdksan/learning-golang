package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func search_ips(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func search_server(c *cli.Context) {
	host := c.String("host")
	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func Generate() *cli.App {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de IPs na Internet",
			Flags:  flags,
			Action: search_ips,
		},
		{
			Name:   "servers",
			Usage:  "Buscar os nomes dos servidores na internet",
			Flags:  flags,
			Action: search_server,
		},
	}

	return app
}
