package tickets

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	pb "github.com/postie-labs/proto/tickets"
	"github.com/spf13/cobra"
	"github.com/terra-money/terra.go/client"
	"github.com/terra-money/terra.go/msg"
)

var (
	NotValidBefore *int64
	NotValidAfter  *int64
	AttributesStr  *string
	Attributes     []*pb.Attribute
	RecipientAddr  cosmtypes.AccAddress
)

var IssueCmd = &cobra.Command{
	Use:   "issue <recipient>",
	Short: "issue a ticket",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		attributesBytes := []byte(*AttributesStr)
		err := json.Unmarshal(attributesBytes, &Attributes)
		if err != nil {
			return err
		}
		recipientAddr, err := cosmtypes.AccAddressFromBech32(args[0])
		if err != nil {
			return err
		}
		RecipientAddr = recipientAddr
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		// create extension
		extension := &pb.Extension{
			NotValidBefore: *NotValidBefore,
			NotValidAfter:  *NotValidAfter,
			Attributes:     Attributes,
		}
		extensionBytes, err := json.Marshal(extension)
		if err != nil {
			return err
		}
		tokenIdBytes := sha256.Sum256(extensionBytes)
		tokenIdStr := strings.ToUpper(hex.EncodeToString(tokenIdBytes[:]))
		execMsg := &pb.TxMint{
			Mint: &pb.Mint{
				Owner:     RecipientAddr.String(),
				TokenId:   tokenIdStr,
				TokenUri:  "",
				Extension: extension,
			},
		}
		execMsgBytes, err := json.Marshal(execMsg)
		if err != nil {
			return err
		}
		execMsgSize := len(execMsgBytes)
		fmt.Println("size:", execMsgSize)
		if execMsgSize > 4096 {
			panic(fmt.Errorf("exec_msg_size(%d) exceeds max_contract_msg_size(%d)", execMsgSize, 4096))
		}
		txOpts := client.CreateTxOptions{
			Msgs: []msg.Msg{
				msg.NewMsgExecuteContract(
					cosmtypes.AccAddress(LCDClient.PrivKey.PubKey().Address()),
					*ContractAddr,
					execMsgBytes,
					msg.NewCoins(),
				),
			},
			Memo:     "i made this",
			GasLimit: 400000,
		}

		// create and sign tx
		tx, err := LCDClient.CreateAndSignTx(ctx, txOpts)
		if err != nil {
			return err
		}

		// broadcast transaction
		respBroadcast, err := LCDClient.Broadcast(ctx, tx)
		if err != nil {
			return err
		}
		fmt.Println(respBroadcast)

		return nil
	},
}

func init() {
	now := time.Now()
	defaultNotValidBefore := now.Unix()
	defaultNotValidAfter := now.Add(8760 * time.Hour).Unix()
	defaultAttributesStr := `[{"key":"hello","value":"earth."}]`

	NotValidBefore = IssueCmd.Flags().Int64("not-valid-before", defaultNotValidBefore, "Not Valid Before")
	NotValidAfter = IssueCmd.Flags().Int64("not-valid-after", defaultNotValidAfter, "Not Valid After")
	AttributesStr = IssueCmd.Flags().String("attributes", defaultAttributesStr, "Attributes")
}
