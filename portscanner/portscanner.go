package portscanner

import (
	"net"
	"sort"
	"strconv"
	"sync"
)

// ScanResult models a scan result
type ScanResult struct {
	Port  int
	State string
}

const (
	// CLOSED means closed
	CLOSED = "Closed"
	// OPEN means open
	OPEN = "Open"
)

// ScanPort Gets the port status
func ScanPort(protocol string, ip string, port int, waitGroup *sync.WaitGroup, portResultChannel chan ScanResult) {
	scanResult := ScanResult{
		Port: port,
	}

	address := ip + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", address)

	if err != nil {
		scanResult.State = CLOSED
		portResultChannel <- scanResult
		waitGroup.Done()
		return
	}

	defer conn.Close()

	scanResult.State = OPEN
	portResultChannel <- scanResult
	waitGroup.Done()
	return
}

// PingScan starts a ping scan
func PingScan(hostname string) []ScanResult {
	var results []ScanResult
	var portResultChannel = make(chan ScanResult)
	defer close(portResultChannel)

	var waitGroup sync.WaitGroup

	for port := 1; port <= 1024; port++ {
		waitGroup.Add(1)
		go ScanPort("tcp", hostname, port, &waitGroup, portResultChannel)
		portResult := <-portResultChannel
		results = append(results, portResult)
	}

	waitGroup.Wait()
	sortResults(results)

	return results
}

func sortResults(results []ScanResult) []ScanResult {
	sort.Slice(results, func(i, j int) bool {
		return results[i].Port < results[j].Port
	})

	return results
}
