package fetch

import "time"

// Options http client option
type Options struct {
	Method   string
	Body     []byte
	Header   map[string]string
	Redirect string

	Follow   int
	Timeout  time.Duration
	Compress bool
	Size     int
	Agent    interface{}
}
