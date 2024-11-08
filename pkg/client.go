package pkg

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/fleimkeipa/case/model"
)

const BaseURL = "https://api.trendyol.com/sapigw/"

type Client struct {
	HTTPClient *http.Client
	APIKey     string
	APISecret  string
}

func NewHTTPClient(apiKey, apiSecret string) Client {
	return Client{
		HTTPClient: &http.Client{},
		APIKey:     apiKey,
		APISecret:  apiSecret,
	}
}

func (rc *Client) Do(req model.Request, resp interface{}) error {
	parsedURL, err := parseURL(req)
	if err != nil {
		return err
	}

	reqBody := new(bytes.Buffer)
	if req.Body != nil {
		marshalled, err := json.Marshal(req.Body)
		if err != nil {
			return err
		}

		reqBody = bytes.NewBuffer(marshalled)
	}

	newReq, err := http.NewRequest(req.Method, parsedURL.String(), reqBody)
	if err != nil {
		return err
	}

	for i, v := range req.Headers {
		newReq.Header.Add(i, v)
	}

	newReq.Header.Add("Authorization", basicAuth(rc.APIKey, rc.APISecret))

	response, err := rc.HTTPClient.Do(newReq)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response != nil {
		respBody, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(bytes.NewBuffer(respBody))

		if err := decoder.Decode(resp); err != nil {
			return err
		}
	}

	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	return fmt.Sprintf("Basic %s", encoded)
}

func parseURL(req model.Request) (*url.URL, error) {
	joinedURL, err := url.JoinPath(BaseURL, req.Paths...)
	if err != nil {
		return nil, err
	}

	parsedURL, err := url.Parse(joinedURL)
	if err != nil {
		return nil, err
	}

	return parsedURL, nil
}
