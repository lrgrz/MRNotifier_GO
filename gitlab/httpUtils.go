package gitlab

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

type httpUtil struct {
	client   *http.Client
	apiToken string
}

func (h *httpUtil) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if h.apiToken != "" {
		req.Header.Add("private-token", h.apiToken)
	}

	client := http.DefaultClient
	if h.client != nil {
		client = h.client
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected http status: %d. message: %s", resp.StatusCode, body)
	}

	return body, nil
}

func (h *httpUtil) configure(verifySSL bool, apiToken string) {
	h.client = createClient(verifySSL)
	h.apiToken = apiToken
}

func createClient(verifySSL bool) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: !verifySSL,
		},
	}
	return &http.Client{Transport: transport}
}
