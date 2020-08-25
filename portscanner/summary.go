package portscanner

import (
	"fmt"
	"time"
)

const (
	openPortString     = "Port: %d\t\t\033[1;34mOpen\033[0m\n"
	closedPortString   = "Port: %d\t\t\033[1;31mClosed\033[0m\n"
	filteredPortString = "Port: %d\t\t\033[1;33mClosed\033[0m\n"
)

// DisplayStartScanner shows a message indicating the scan is starting
func DisplayStartScanner() {
	fmt.Println("Starting Port Scan...")
}

// DisplayResults show a summary ofthe results in a friendly way
func DisplayResults(showClosedPorts bool, results []ScanResult, elapsedTime time.Duration) {
	totalOpenPorts := 0
	for _, result := range results {
		if result.State == OPEN {
			fmt.Printf(openPortString, result.Port)
			totalOpenPorts++
		} else if showClosedPorts {
			fmt.Printf(closedPortString, result.Port)
		}
	}

	fmt.Printf("\nPort Scan took %s\n", elapsedTime)
	fmt.Printf("Total scanned ports: %d\n", len(results))
	fmt.Printf("Total open ports: %d\n", totalOpenPorts)
}
