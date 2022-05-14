package tickets

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	pb "github.com/postie-labs/proto/tickets"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list issued tickets",
	RunE: func(cmd *cobra.Command, args []string) error {
		ownerAddr := cosmtypes.AccAddress(LCDClient.PrivKey.PubKey().Address())
		query := pb.QueryList{AllTokens: &pb.QueryList_List{Owner: ownerAddr.String()}}
		queryBytes, err := json.Marshal(query)
		if err != nil {
			return err
		}
		queryStr := base64.StdEncoding.EncodeToString(queryBytes)
		urlPath := fmt.Sprintf("%s/%s/%s/%s?query_msg=%s",
			*LCDEndpoint,
			"terra/wasm/v1beta1/contracts",
			ContractAddr,
			"store",
			queryStr,
		)
		resp, err := http.Get(urlPath)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var queryResp pb.QueryListResponse
		err = json.Unmarshal(data, &queryResp)
		if err != nil {
			return err
		}
		for _, ticket := range queryResp.GetQueryResult().GetTokens() {
			fmt.Println(ticket)
		}
		return nil
	},
}
