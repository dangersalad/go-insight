package insight

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	url        string
	httpClient *http.Client
}

func NewClient(apiURL string) *Client {
	return &Client{
		url:        apiURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) doRequest(path string, params url.Values, res interface{}) error {

	if path[0] != '/' {
		path = "/" + path
	}

	urlParams := ""
	if params != nil {
		urlParams = "?" + params.Encode()
	}
	reqURL := fmt.Sprintf("%s%s%s", c.url, path, urlParams)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Error %d: %s", resp.StatusCode, body)
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(res)
	if err != nil {
		return err
	}
	return nil
}
