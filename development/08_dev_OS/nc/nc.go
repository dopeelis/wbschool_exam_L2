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
				&cli.StringFlag{Name: "protocol"},
			},
			Action: func(c *cli.Context) error {
				client := NewNetCatClient(c.Args().Slice()[0], c.Args().Slice()[1], c.String("protocol"))
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
	protocol string
}

func NewNetCatClient(host, port string, protocol string) *NetCatClient {
	return &NetCatClient{
		host:     host,
		port:     port,
		protocol: protocol,
	}
}

func Run(nc *NetCatClient) error {
	conn, err := net.Dial(nc.protocol, net.JoinHostPort(nc.host, nc.port))
	if err != nil {
		return err
	}
	osSignals := make(chan os.Signal, 1)
	errors := make(chan error, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	go req(conn, errors, osSignals)
	go resp(conn, errors, osSignals)

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

func req(conn net.Conn, chanErr chan<- error, chanOsSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				chanOsSignals <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			chanErr <- err
		}

		_, err = fmt.Fprintf(conn, text+"\n")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func resp(conn net.Conn, chanErr chan<- error, chanOsSignals chan<- os.Signal) {
	for {
		reader := bufio.NewReader(conn)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				chanOsSignals <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			chanErr <- err
		}

		fmt.Print(text)
	}
}
