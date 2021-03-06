package fetch

import "encoding/json"

// Resp custom http response
type Resp struct {
	Body       []byte
	StatusCode int
}

// BindJSON convert body to s
// s can be a map or a struct
func (resp Resp) BindJSON(s interface{}) error {
	if err := json.Unmarshal(resp.Body, &s); err != nil {
		return err
	}
	return nil
}

// ToString convert response body to string
func (resp Resp) ToString() string {
	return string(resp.Body)
}
