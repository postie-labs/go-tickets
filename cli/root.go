package cli

import (
	"fmt"
	"os"
	"time"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/postie-labs/go-tickets/cli/tickets"
	"github.com/spf13/cobra"
	"github.com/terra-money/terra.go/client"
	"github.com/terra-money/terra.go/key"
)

const (
	DefaultLCDEndpoint            = "https://bombay-lcd.terra.dev"
	DefaultChainID                = "bombay-12"
	DefaultTimeout                = 10 // sec
	DefaultContractAddressBench32 = "terra1al87aagg7asjyceu9x8f4xj554ddzlk9q2t8ls"
)

var (
	LCDClient   *client.LCDClient
	LCDEndpoint *string
	ChainID     *string
	Timeout     *int64
)

var RootCmd = &cobra.Command{
	Use:   "tickets",
	Short: "'tickets' is a CLI tool to interact with the smart contract `cw721-tickets`",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}

func init() {
	cobra.OnInitialize(
		InitClient,
	)
	cobra.EnableCommandSorting = false
	LCDEndpoint = RootCmd.PersistentFlags().String("endpoint", DefaultLCDEndpoint, "LCDEndpoint")
	ChainID = RootCmd.PersistentFlags().String("chain-id", DefaultChainID, "Chain ID")
	Timeout = RootCmd.PersistentFlags().Int64("timeout", DefaultTimeout, "Timeout")
	RootCmd.AddCommand(
		tickets.IssueCmd,
		tickets.TransferCmd,
		tickets.ListCmd,
		tickets.GetCmd,
		tickets.QRCmd,
	)
}

func InitClient() {
	// derive privKey
	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		handleError(fmt.Errorf("failed to read MNEMONIC envrionment variable"))
	}
	privKeyBytes, err := key.DerivePrivKeyBz(mnemonic, key.CreateHDPath(0, 0))
	if err != nil {
		handleError(err)
	}
	privKey, err := key.PrivKeyGen(privKeyBytes)
	if err != nil {
		handleError(err)
	}

	// create LCDClient
	LCDClient = client.NewLCDClient(
		*LCDEndpoint,
		*ChainID,
		cosmtypes.NewDecCoinFromDec("uluna", cosmtypes.NewDecFromIntWithPrec(cosmtypes.NewInt(1133), 5)),
		cosmtypes.NewDecFromIntWithPrec(cosmtypes.NewInt(150), 2),
		privKey,
		time.Duration(*Timeout)*time.Second,
	)
}

func handleError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
