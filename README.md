# go-fetch

a http client like node-fetch

# usage

```
go get -u github.com/zzfup/go-fetch
```

# API

## Fetch(url[, options])

- url: string. should be an absolute url, such as http://example.com/.
- options: map[string]string. the default values are here:

  ```golang
  Options{
  		Method: "GET",
  		Header: map[string]string{
  			"Accept-Encoding": "gzip,deflate",
  			"Accept":          "*/*",
  		},
        Body:    nil,
        Timeout: 20 * time.Millisecond,
  }
  ```

## NewDefaultOptions()

return a defautl options.

you can also create an Options, like this:

```golang
options := fetch.Options{
	Method: "POST",
    Header: headers,  // your custom header, its type is map[string]string
    Body: payload, // your request body, its type is []byte
}
```

## ToString()

convert the respone body to string

## BindJSON(i interface{})

convert the response body to a struct or a map

# example

```golang
options := fetch.Options{}
// option = fetch.NewDefaulOptions()
resp, err := fetch.Fetch("https://www.baidu.com", fetch.Options{})
fmt.Println(err)
fmt.Println(resp.ToString())
```

## GET

```golang
import "github.com/zzfup/go-fetch"
import "fmt"

// header can be just like this
var headers = map[string]string{
	"Accept":       "application/json, text/plain, */*",
	"Content-Type": "application/json",
}

func main(){
    url := "https://www.example.com"
    options := fetch.Options{
		Method: "GET",
        Header: headers,
        Timeout: 2 * time.Second,
	}

    resp, err := fetch.Fetch(url, options)
	if err != nil {
		return  err
	}

    fmt.Println(resp.StatusCode)
    fmt.Println(resp.ToString)
    var j struct{
        Test string `json:"test"`
    }
    err := resp.BindJSON(&j)
    fmt.Println(j)
}
```

## POST

```golang
import "github.com/zzfup/go-fetch"
import "fmt"

// header can be just like this
var headers = map[string]string{
	"Accept":       "application/json, text/plain, */*",
	"Content-Type": "application/json",
}

func main(){
    url := "https://www.example.com"

    payload, err := json.Marshal(a) // a can be a struct or a map

    options := fetch.Options{
		Method: "POST",
        // Header: headers,
        Body: payload
        // Timeout: 2 * time.Second,
	}

    resp, err := fetch.Fetch(url, options)
	if err != nil {
		return  err
	}

    fmt.Println(resp.StatusCode)
    fmt.Println(resp.ToString)
    var j struct{
        Test string `json:"test"`
    }
    err := resp.BindJSON(&j)
    fmt.Println(j)
}
```
