package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "netcat",
			Usage: "netcat client realization",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "u"},
			},
			Action: func(c *cli.Context) error {
				if len(c.Args().Slice()) < 2 {
					log.Fatalln("usage: netcat [--u] host port")
				}
				client := NewNetCatClient(c.Args().Slice()[0], c.Args().Slice()[1], c.Bool("u"))
				err := Run(client)
				if err != nil {
					fmt.Println(err)
				}
				return err
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

type NetCatClient struct {
	host     string
	port     string
	protocol bool
}

func NewNetCatClient(host, port string, u bool) *NetCatClient {
	return &NetCatClient{
		host:     host,
		port:     port,
		protocol: u,
	}
}

func Run(nc *NetCatClient) error {
	var protocol string
	if nc.protocol {
		protocol = "udp"
	} else {
		protocol = "tcp"
	}
	fmt.Println("protocol:", protocol)
	conn, err := net.Dial(protocol, net.JoinHostPort(nc.host, nc.port))
	if err != nil {
		return err
	}
	osSignals := make(chan os.Signal, 1)
	errors := make(chan error, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	go Send(conn, errors, osSignals)

	select {
	case <-osSignals:
		err := conn.Close()
		if err != nil {
			return err
		}
	case err = <-errors:
		if err != nil {
			return err
		}
	}
	return nil
}

func Send(conn net.Conn, errors chan<- error, osSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				osSignals <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			errors <- err
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
