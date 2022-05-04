package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/postie-labs/go-tickets/qr"
	"github.com/terra-money/terra.go/client"
	"github.com/terra-money/terra.go/key"
	"github.com/terra-money/terra.go/msg"
)

const (
	Mnemonic               = "term wait vessel monitor rack oak found athlete lens mimic grow magnet spatial frown budget balance rebuild fossil acid vicious tiger avocado brand venture"
	LCDEndpoint            = "https://bombay-lcd.terra.dev"
	ChainID                = "testnet"
	ContractAddressBench32 = "terra1al87aagg7asjyceu9x8f4xj554ddzlk9q2t8ls"
	Owner                  = "terra1jrj9kjwv5pwrttrsdmg33w0kzj3szzc230g9z5"
)

func main() {
	// init
	ctx := context.Background()

	// generate privKey, pubKey, address
	privKeyBytes, err := key.DerivePrivKeyBz(Mnemonic, key.CreateHDPath(0, 0))
	if err != nil {
		panic(err)
	}
	privKey, err := key.PrivKeyGen(privKeyBytes)
	if err != nil {
		panic(err)
	}
	pubKey := privKey.PubKey()
	address := types.AccAddress(pubKey.Address())

	// create LCDClient
	LCDClient := client.NewLCDClient(
		LCDEndpoint,
		ChainID,
		msg.NewDecCoinFromDec("uusd", msg.NewDecFromIntWithPrec(msg.NewInt(15), 2)), // 0.15uusd
		msg.NewDecFromIntWithPrec(msg.NewInt(15), 1),
		privKey,
		10*time.Second,
	)

	// create transaction
	contractAddress, err := types.AccAddressFromBech32(ContractAddressBench32)
	if err != nil {
		panic(err)
	}
	// {
	// 	"mint": {
	// 		"owner": "terra1k05lru8us3ctq7ngc396sxesmd7dsd2a8ppfv7",
	// 		"token_id": "this is the first token issued by alice+bob",
	// 		"token_uri": "https://github.com/postie-labs/cw721-tickets",
	// 		"extension": {
	// 			"not_valid_before": 1650456000,
	// 			"not_valid_after": 1966075200,
	// 			"attributes": [
	// 				{
	// 					"key": "hello",
	// 					"value": "world"
	// 				},
	// 				{
	// 					"key": "location",
	// 					"value": "seoul"
	// 				},
	// 				{
	// 					"key": "date",
	// 					"value": "1650499200"
	// 				}
	// 			]
	//
	// 		}
	// 	}
	// }

	now := time.Now()
	extension := qr.Extension{
		NotValidBefore: now.Unix(),
		NotValidAfter:  now.Add(time.Hour * 3600).Unix(),
		Attributes: []qr.Attribute{
			qr.Attribute{
				Key:   "hello",
				Value: "world",
			},
		},
	}
	extensionBytes, err := json.Marshal(extension)
	if err != nil {
		panic(err)
	}
	tokenIdBytes := sha256.Sum256(extensionBytes)

	execMsg := qr.TxExecuteMint{
		Mint: qr.Mint{
			Owner:     Owner,
			TokenId:   hex.EncodeToString(tokenIdBytes[:]),
			TokenUri:  "",
			Extension: extension,
		},
	}
	execMsgBytes, err := json.Marshal(execMsg)
	if err != nil {
		panic(err)
	}
	_, err = LCDClient.CreateAndSignTx(ctx, client.CreateTxOptions{
		Msgs: []msg.Msg{
			msg.NewMsgExecuteContract(
				address,
				contractAddress,
				execMsgBytes,
				msg.NewCoins(),
			),
		},
	})
	if err != nil {
		panic(err)
	}
}

func init() {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount("terra", "terrapub")
	config.Seal()
}
