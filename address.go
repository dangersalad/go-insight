package insight

import (
	"fmt"
	"net/url"
)

type AddressSummariesResponse struct {
	Address                 string   `json:"addrStr"`
	Balance                 float64  `json:"balance"`
	BalanceSat              int64    `json:"balanceSat"`
	TotalReceived           float64  `json:"totalReceived"`
	TotalReceivedSat        int64    `json:"totalReceivedSat"`
	TotalSent               float64  `json:"totalSent"`
	TotalSentSat            int64    `json:"totalSentSat"`
	UnconfirmedBalance      float64  `json:"unconfirmedBalance"`
	UnconfirmedBalanceSat   int64    `json:"unconfirmedBalanceSat"`
	UnconfirmedTxApperances float64  `json:"unconfirmedTxApperances"`
	TxApperances            int64    `json:"txApperances"`
	Transactions            []string `json:"transactions,omitempty"`
}

func (c *Client) AddressSummaries(addr string, listTx bool) (*AddressSummariesResponse, error) {
	params := make(url.Values)
	if !listTx {
		params.Add("noTxList", "1")
	}
	resp := &AddressSummariesResponse{}
	url := fmt.Sprintf("/addr/%s", addr)
	err := c.doRequest(url, params, resp)
	return resp, err
}

func (c *Client) AddressBalance(addr string) (float64, error) {
	as, err := c.AddressSummaries(addr, false)
	if err != nil {
		return 0, err
	}
	return as.Balance, nil
}
