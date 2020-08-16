package portscanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

type ScanResult struct {
	Port  int
	State string
}

func GetPortStatus(protocol string, ip string, port int, waitGroup *sync.WaitGroup) ScanResult {
	result := ScanResult{Port: port}

	address := ip + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", address)

	if err != nil {
		result.State = "Closed"
		fmt.Printf("Port %d:\tClosed\n", port)
		waitGroup.Done()

		return result
	}

	defer conn.Close()

	result.State = "Open"
	fmt.Printf("Port %d:\tOpen\n", port)
	waitGroup.Done()

	return result
}

func StartScan(hostname string) []ScanResult {

	var results []ScanResult
	var waitGroup sync.WaitGroup

	for port := 0; port <= 1024; port++ {
		waitGroup.Add(1)
		results = append(results, GetPortStatus("tcp", hostname, port, &waitGroup))
	}

	waitGroup.Wait()

	return results
}
