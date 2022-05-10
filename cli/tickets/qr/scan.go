package qr

import "github.com/spf13/cobra"

var ScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan QR code",
}
