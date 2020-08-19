package portscanner

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
