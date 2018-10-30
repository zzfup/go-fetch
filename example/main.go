package main

import (
	"fmt"

	"github.com/zzfup/go-fetch"
)

func main() {
	url := "http://www.baidu.com"
	resp, err := fetch.Fetch(url, fetch.Options{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.ToString())
}
