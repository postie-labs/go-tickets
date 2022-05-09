package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/postie-labs/go-tickets/cli"
	pb "github.com/postie-labs/proto/tickets"
	"github.com/terra-money/terra.go/client"
	"github.com/terra-money/terra.go/key"
	"github.com/terra-money/terra.go/msg"
)

const (
	DefaultLCDEndpoint            = "https://bombay-lcd.terra.dev"
	DefaultChainID                = "bombay-12"
	DefaultContractAddressBench32 = "terra1al87aagg7asjyceu9x8f4xj554ddzlk9q2t8ls"
	DefaultOwnerBench32           = "terra1k05lru8us3ctq7ngc396sxesmd7dsd2a8ppfv7"
)

func main() {
	err := cli.RootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func gen() {
	// init
	ctx := context.Background()
	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		panic(fmt.Errorf("failed to read MNEMONIC envrionment variable"))
	}

	// generate privKey, pubKey, address
	privKeyBytes, err := key.DerivePrivKeyBz(mnemonic, key.CreateHDPath(0, 0))
	if err != nil {
		panic(err)
	}
	privKey, err := key.PrivKeyGen(privKeyBytes)
	if err != nil {
		panic(err)
	}
	pubKey := privKey.PubKey()
	address := cosmtypes.AccAddress(pubKey.Address())

	// create LCDClient
	LCDClient := client.NewLCDClient(
		DefaultLCDEndpoint,
		DefaultChainID,
		cosmtypes.NewDecCoinFromDec("uluna", cosmtypes.NewDecFromIntWithPrec(cosmtypes.NewInt(1133), 5)),
		cosmtypes.NewDecFromIntWithPrec(cosmtypes.NewInt(150), 2),
		privKey,
		10*time.Second,
	)

	// create transaction
	contractAddress, err := cosmtypes.AccAddressFromBech32(DefaultContractAddressBench32)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	extension := pb.Extension{
		NotValidBefore: now.Unix(),
		NotValidAfter:  now.Add(time.Hour * 3600).Unix(),
		Attributes: []*pb.Attribute{
			{
				Key:   "name",
				Value: "wedding invitation card",
			},
			{
				Key:   "invitor",
				Value: "alice and bob",
			},
			{
				Key:   "invitee",
				Value: "carol",
			},
			{
				Key:   "content",
				Value: "please come celebrate our wedding together",
			},
			{
				Key:   "image",
				Value: "data:image/webp;base64,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",
			},
		},
	}
	extensionBytes, err := json.Marshal(extension)
	if err != nil {
		panic(err)
	}
	tokenIdBytes := sha256.Sum256(extensionBytes)
	tokenIdStr := strings.ToUpper(hex.EncodeToString(tokenIdBytes[:]))

	execMsg := pb.TxMint{
		Mint: &pb.Mint{
			Owner:     DefaultOwnerBench32,
			TokenId:   tokenIdStr,
			TokenUri:  "",
			Extension: &extension,
		},
	}
	execMsgBytes, err := json.Marshal(execMsg)
	if err != nil {
		panic(err)
	}
	execMsgSize := len(execMsgBytes)
	fmt.Println("size:", execMsgSize)
	if execMsgSize > 4096 {
		panic(fmt.Errorf("exec_msg_size(%d) exceeds max_contract_msg_size(%d)", execMsgSize, 4096))
	}
	txOpts := client.CreateTxOptions{
		Msgs: []msg.Msg{
			msg.NewMsgExecuteContract(
				address,
				contractAddress,
				execMsgBytes,
				msg.NewCoins(),
			),
		},
		Memo:     "i made this",
		GasLimit: 400000,
	}
	tx, err := LCDClient.CreateAndSignTx(ctx, txOpts)
	if err != nil {
		panic(err)
	}

	// broadcast transaction
	respBroadcast, err := LCDClient.Broadcast(ctx, tx)
	if err != nil {
		panic(err)
	}
	fmt.Println(respBroadcast)
}

func init() {
	config := cosmtypes.NewConfig()
	config.SetBech32PrefixForAccount("terra", "terrapub")
	config.Seal()
}
