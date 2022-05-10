package tickets

import (
	"github.com/postie-labs/go-tickets/cli/tickets/qr"
	"github.com/spf13/cobra"
)

var QRCmd = &cobra.Command{
	Use:   "qr",
	Short: "QR related operations",
}

func init() {
	cobra.EnableCommandSorting = false
	QRCmd.AddCommand(
		qr.GenerateCmd,
		qr.ScanCmd,
	)
}
