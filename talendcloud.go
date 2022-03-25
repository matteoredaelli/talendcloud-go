package talendcloud

import (
	"bytes"
	"encoding/json"
	"fmt" //	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func init() {
}

// String converts string variable and literal to pointer
func String(s string) *string {
	return &s
}

// Client is main structure of the library, a requester to talendcloud.
type Client struct {
	BaseURL string
	apiKey  string
}

// NewClient is a constructor of Client
func NewClient(baseUrl, apiKey string) Client {
	client := Client{
		apiKey:  apiKey,
		BaseURL: baseUrl,
	}

	return client
}

func (x Client) Post(apiName string, input interface{}) (string, error) {
	rawData, err := json.Marshal(input)
	if err != nil {
		return "{}", err
	}

	uri := fmt.Sprintf("%s/%s/", x.BaseURL, apiName)

	client := &http.Client{}
	req, err := http.NewRequest("POST", uri, bytes.NewReader(rawData))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", `Bearer ${x.apiKey}`)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (x Client) Get(apiName string, values url.Values) (string, error) {
	var qs string
	if values != nil {
		qs = "?" + values.Encode()
	}

	uri := fmt.Sprintf("%s/%s%s", x.BaseURL, apiName, qs)

	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+x.apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
