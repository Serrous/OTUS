package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

const ntpsrv = "0.pool.ntp.org"

func main() {

	curTime, err := ntp.Time(ntpsrv)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error(), "\n")
		os.Exit(1)
	} else {
		fmt.Fprint(os.Stdout, "Local time:\t", time.Now(), "\n", "NTP time:\t", curTime, "\n")
	}
}
