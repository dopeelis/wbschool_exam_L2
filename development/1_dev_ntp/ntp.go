package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	timeNow, _ := currentTime("0.beevik-ntp.pool.ntp.org")
	fmt.Println("Current time:", timeNow)

	timeExact, _ := exactTime("0.beevik-ntp.pool.ntp.org")
	fmt.Println("Exact time:", timeExact)
}

func exactTime(host string) (time.Time, error) {
	response, err := ntp.Query(host)
	if err != nil {
		log.Fatalln(err)
	}
	time := time.Now().Add(response.ClockOffset)

	return time, err
}

func currentTime(host string) (time.Time, error) {
	time, err := ntp.Time(host)
	if err != nil {
		log.Fatalln(err)
	}

	return time, err
}
