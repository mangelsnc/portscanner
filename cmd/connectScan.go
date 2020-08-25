package cmd

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/mangelsnc/portscanner/portscanner"
	"github.com/spf13/cobra"
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
		portscanner.DisplayStartScanner()
		start := time.Now()
		host := args[0]
		results := portscanner.ConnectScan(host)
		elapsed := time.Since(start)

		portscanner.DisplayResults(showClosedPorts, results, elapsed)
	},
}

func isValidIP(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	}

	return true
}
