package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "go-telnet",
			Usage: "telnet client realization",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "timeout"},
			},
			Action: func(c *cli.Context) error {
				telnet := InitFlags(c.Args().Slice()[0], c.Args().Slice()[1], c.Int("timeout"))
				err := telnet.Run()
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

type Flags struct {
	host    string
	port    string
	timeout time.Duration
}

func InitFlags(host, port string, timeout int) *Flags {
	return &Flags{
		host:    host,
		port:    port,
		timeout: time.Duration(timeout),
	}
}

func (f *Flags) Run() error {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(f.host, f.port), f.timeout)
	if err != nil {
		time.Sleep(f.timeout)
		return err
	}
	chanOsSignals := make(chan os.Signal, 1)
	chanErr := make(chan error, 1)
	signal.Notify(chanOsSignals, syscall.SIGINT, syscall.SIGTERM)
	go req(conn, chanErr, chanOsSignals)
	go resp(conn, chanErr, chanOsSignals)

	select {
	case <-chanOsSignals:
		conn.Close()
	case err = <-chanErr:
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

		fmt.Fprintf(conn, text+"\n")
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
