package cmd

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/mangelsnc/portscanner/portscanner"
	"github.com/spf13/cobra"
)

const (
	openPortString     = "Port: %d\t\t\033[1;34mOpen\033[0m\n"
	closedPortString   = "Port: %d\t\t\033[1;31mClosed\033[0m\n"
	filteredPortString = "Port: %d\t\t\033[1;33mClosed\033[0m\n"
)

var showClosedPorts bool

func init() {
	rootCmd.AddCommand(connectScanCmd)
	connectScanCmd.Flags().BoolVarP(&showClosedPorts, "show-closed-ports", "c", false, "Shows closed ports in the report")
}

var connectScanCmd = &cobra.Command{
	Use:   "connectScan <ip>",
	Short: "Runs a connect scan against the specified host and ports",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires specify a target IP")
		}
		if isValidIP(args[0]) {
			return nil
		}
		return fmt.Errorf("Invalid host specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Port Scan...")
		start := time.Now()
		results := portscanner.ConnectScan(args[0])
		elapsed := time.Since(start)

		for _, result := range results {
			if result.State == portscanner.OPEN {
				fmt.Printf(openPortString, result.Port)
			} else if showClosedPorts {
				fmt.Printf(closedPortString, result.Port)
			}
		}

		fmt.Printf("\nPort Scan took %s\n", elapsed)
	},
}

func isValidIP(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	}

	return true
}
