package fetch

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// Fetch function
// url: Absolute url
// op: Fetch options
func Fetch(url string, op Options) (Resp, error) {
	defaultOp := NewDefaultOptions()
	if op.Method != "" {
		defaultOp.Method = op.Method
	}
	if op.Header != nil {
		defaultOp.Header = op.Header
	}
	if op.Body != nil {
		defaultOp.Body = op.Body
	}
	if int(op.Timeout) != 0 {
		defaultOp.Timeout = op.Timeout
	}

	// create a new http client
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableKeepAlives:     true,
		},
		Timeout: defaultOp.Timeout,
	}
	req, err := http.NewRequest(defaultOp.Method, url, bytes.NewReader(defaultOp.Body))
	if err != nil {
		return Resp{}, err
	}
	// set header
	for k, v := range defaultOp.Header {
		req.Header.Set(k, v)
	}
	// send request
	resp, err := client.Do(req)
	if err != nil {
		return Resp{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Resp{}, err
	}
	return Resp{
		StatusCode: resp.StatusCode,
		Body:       body,
	}, nil
}
