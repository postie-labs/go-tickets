package tickets

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pb "github.com/postie-labs/proto/tickets"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <ticket_id>",
	Short: "get a ticket's information",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ticketId := args[0]
		query := pb.QueryAllInfo{AllNftInfo: &pb.QueryAllInfo_AllInfo{TokenId: ticketId}}
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
		var queryResp pb.QueryAllInfoResponse
		err = json.Unmarshal(data, &queryResp)
		if err != nil {
			return err
		}
		queryResult := queryResp.GetQueryResult()
		access := queryResult.GetAccess()
		info := queryResult.GetInfo()
		extension := info.GetExtension()
		attributes := extension.GetAttributes()
		fmt.Println("owner:", access.Owner)
		if len(access.Approvals) > 0 {
			fmt.Println("approvals:", access.Approvals)
			for _, approval := range access.Approvals {
				fmt.Printf("  - %s\n", approval)
			}
		}
		fmt.Println("not valid before:", extension.NotValidBefore)
		fmt.Println("not valid after:", extension.NotValidAfter)
		if len(attributes) > 0 {
			fmt.Println("attributes:")
			for _, attribute := range attributes {
				fmt.Printf("  - %s: %s\n", attribute.GetKey(), attribute.GetValue())
			}
		}
		return nil
	},
}
