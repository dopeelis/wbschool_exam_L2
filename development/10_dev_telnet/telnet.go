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
	"time"
)

type TelnetClient struct {
	host    string
	port    string
	timeout time.Duration
	conn    net.Conn
}

func NewTelnetClient(host string, port string, timeout time.Duration) *TelnetClient {
	return &TelnetClient{
		host:    host,
		port:    port,
		timeout: timeout,
	}
}

func (t *TelnetClient) Connection() error {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(t.host, t.port), t.timeout)
	if err != nil {
		return fmt.Errorf("conncetion error: %v", err)
	}
	t.conn = conn
	fmt.Println("connected to ", t.host+":"+t.port)
	return nil
}

func (t *TelnetClient) Receive() error {
	reader := bufio.NewReader(t.conn)
	text, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			err = fmt.Errorf("error: closed")
		}
		return err
	}
	_, err = fmt.Print(text)
	if err != nil {
		return err
	}
	return nil
}

func (t *TelnetClient) Send(text string) error {
	_, err := t.conn.Write([]byte(text))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	app := &cli.App{}
	app.Commands = []*cli.Command{
		{
			Name:  "go-telnet",
			Usage: "telnet client realization",
			Flags: []cli.Flag{
				&cli.DurationFlag{Name: "timeout"},
			},
			Action: func(c *cli.Context) error {
				client := NewTelnetClient(c.Args().Slice()[0], c.Args().Slice()[1], c.Duration("timeout"))
				Run(client)
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func Run(client *TelnetClient) {
	err := client.Connection()
	if err != nil {
		log.Fatalf("connecting error: %v\n", err)
	}
	defer func() {
		err := client.conn.Close()
		if err != nil {
			log.Fatalf("closing connection error: %v\n", err)
		}
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	errors := make(chan error, 1)

	go func() {
		for {
			err := client.Receive()
			if err != nil {
				errors <- err
				return
			}
		}
	}()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				errors <- err
				return
			}
			err = client.Send(text)
			if err != nil {
				errors <- err
				return
			}
		}
	}()

	exit := make(chan interface{})
	go func() {
		defer close(exit)
		for {
			select {
			case <-osSignals:
				return
			case err := <-errors:
				log.Println(err)
				if err != nil {
					return
				}
			default:
				continue
			}
		}
	}()

	<-exit
}
