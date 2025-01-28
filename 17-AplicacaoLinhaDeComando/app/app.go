package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na interne",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscaIps,
		},
	}

	return app
}

func buscaIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
