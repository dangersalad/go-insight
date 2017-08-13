package insight

import (
	"net/url"
)

type Input struct {
}

type ScriptPubKey struct {
	Addresses []string `json:"addresses"`
}

type Output struct {
	Value        float64       `json:"value,string"`
	N            int           `json:"n"`
	ScriptPubKey *ScriptPubKey `json:"scriptPubKey,omitempty"`
}

type Transaction struct {
	TxID          string    `json:"txid"`
	Version       int       `json:"version"`
	Locktime      int64     `json:"locktime"`
	Vin           []*Input  `json:"vin"`
	Vout          []*Output `json:"vout"`
	Blockhash     string    `json:"blockhash"`
	Blockheight   int64     `json:"blockheight"`
	Confirmations int64     `json:"confirmations"`
	Time          int64     `json:"time"`
	BlockTime     int64     `json:"blocktime"`
	IsCoinbase    bool      `json:"isCoinbase"`
	ValueOut      float64   `json:"valueOut"`
	ValueIn       float64   `json:"valueIn,omitempty"`
	Fees          float64   `json:"fees,omitempty"`
	Size          int64     `json:"size"`
}

type TransactionsResponse struct {
	PagesTotal   int64          `json:"pagesTotal"`
	Transactions []*Transaction `json:"txs"`
}

func (c *Client) TransactionsByBlock(block string) (*TransactionsResponse, error) {
	params := make(url.Values)
	params.Add("block", block)
	resp := &TransactionsResponse{}
	err := c.doRequest("/txs", params, resp)
	return resp, err
}
