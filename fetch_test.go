package fetch

import (
	"reflect"
	"testing"
)

func TestFetch(t *testing.T) {
	type args struct {
		url string
		op  Options
	}
	tests := []struct {
		name    string
		args    args
		want    Resp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fetch(tt.args.url, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}
