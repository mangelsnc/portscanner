package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/mangelsnc/portscanner/portscanner"
)

const (
	openPortString     = "Port: %d\t\t\033[1;34mOpen\033[0m\n"
	closedPortString   = "Port: %d\t\t\033[1;31mClosed\033[0m\n"
	filteredPortString = "Port: %d\t\t\033[1;33mClosed\033[0m\n"
)

func main() {
	host := getHost()
	fmt.Println("Starting Port Scan...")
	start := time.Now()
	results := portscanner.ConnectScan(host)
	elapsed := time.Since(start)

	for _, result := range results {
		if result.State == portscanner.OPEN {
			fmt.Printf(openPortString, result.Port)
		} else {
			fmt.Printf(closedPortString, result.Port)
		}
	}

	fmt.Printf("\nPort Scan took %s\n", elapsed)
}

func getHost() string {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "You must specify a host to scan\n\nUsage %s <host>|<ip>\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	return string(os.Args[1])
}
