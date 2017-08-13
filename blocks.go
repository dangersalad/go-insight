package insight

import (
	"net/url"
	"strconv"
	"time"
)

type PoolInfo struct {
	PoolName string `json:"poolName"`
	URL      string `json:"url"`
}

type BlockSummary struct {
	Height   int64     `json:"height"`
	Size     int64     `json:"size"`
	Hash     string    `json:"hash"`
	Time     int64     `json:"time"`
	TxLength int64     `json:"txlength"`
	PollInfo *PoolInfo `json:"poolInfo"`
}

type BlockSummariesResponse struct {
	Blocks []*BlockSummary `json:"blocks"`
	Length int             `json:"length"`
}

func (c *Client) BlockSummaries(limit int, blockDate time.Time) (*BlockSummariesResponse, error) {
	params := make(url.Values)
	if limit > 0 {
		params.Add("limit", strconv.Itoa(limit))
	}
	params.Add("blockDate", blockDate.Format("2006-01-02"))
	resp := &BlockSummariesResponse{}
	err := c.doRequest("/blocks", params, resp)
	return resp, err
}
