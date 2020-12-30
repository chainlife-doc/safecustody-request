package test

import (
	"fmt"
	req "safecustody_request"
	"testing"
)

func TestPost(t *testing.T) {
	r, err := req.Http().Post("http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(r)
	}
	fmt.Println(r.Body)
}
