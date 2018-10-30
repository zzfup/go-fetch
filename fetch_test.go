package fetch

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	url := "https://www.baidu.com"

	options := Options{
		Method: "GET",
		// Header: headers,
		// Timeout: 1 * time.Second,
	}
	resp, err := Fetch(url, options)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	fmt.Println(resp.ToString())
}
