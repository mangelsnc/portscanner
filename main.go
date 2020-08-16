package main

import (
	"fmt"

	"github.com/mangelsnc/portscanner/portscanner"
)

func main() {
	fmt.Println("Port Scanning")
	results := portscanner.StartScan("localhost")
	fmt.Println(results)
}
