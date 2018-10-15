package fetch

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func newDefaultOptions() Options {
	return Options{
		Method: "GET",
		Header: map[string]string{
			"Accept-Encoding": "gzip,deflate",
			"Accept":          "*/*",
		},
		Body:     nil,
		Redirect: "follow",
		Timeout:  20 * time.Millisecond,
		Compress: true,
		Size:     0,
		Agent:    nil,
	}
}

// Fetch request url
func Fetch(url string, option Options) (Resp, error) {
	defaultOptions := newDefaultOptions()
	defaultOptions.Method = option.Method
	defaultOptions.Header = option.Header
	defaultOptions.Body = option.Body
	defaultOptions.Redirect = option.Redirect
	defaultOptions.Timeout = option.Timeout
	defaultOptions.Compress = option.Compress
	defaultOptions.Size = option.Size
	defaultOptions.Agent = option.Agent

	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   defaultOptions.Timeout,
				KeepAlive: defaultOptions.Timeout,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       defaultOptions.Timeout,
			TLSHandshakeTimeout:   defaultOptions.Timeout,
			ExpectContinueTimeout: defaultOptions.Timeout,
		},
		Timeout: defaultOptions.Timeout,
	}

	req, err := http.NewRequest(defaultOptions.Method, url, bytes.NewReader(defaultOptions.Body))
	if err != nil {
		return Resp{}, err
	}
	for k, v := range defaultOptions.Header {
		req.Header.Set(k, v)
	}

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
