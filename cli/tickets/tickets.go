package tickets

import (
	"fmt"
	"os"
	"time"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/postie-labs/go-tickets/util"
	"github.com/spf13/cobra"
	"github.com/terra-money/terra.go/client"
	"github.com/terra-money/terra.go/key"
)

var (
	LCDClient   *client.LCDClient
	LCDEndpoint *string
	ChainID     *string
	Timeout     *int64
)

func init() {
	cobra.OnInitialize(InitClient)

	config := cosmtypes.NewConfig()
	config.SetBech32PrefixForAccount("terra", "terrapub")
	config.Seal()
}

func InitClient() {
	// derive privKey
	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		util.HandleError(fmt.Errorf("failed to read MNEMONIC envrionment variable"))
	}
	privKeyBytes, err := key.DerivePrivKeyBz(mnemonic, key.CreateHDPath(0, 0))
	if err != nil {
		util.HandleError(err)
	}
	privKey, err := key.PrivKeyGen(privKeyBytes)
	if err != nil {
		util.HandleError(err)
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
