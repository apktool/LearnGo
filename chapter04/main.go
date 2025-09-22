package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func to1() (*ResponseData, error) {
	return ToOkResponse("Hello, I'm Ok")
}

func to2() (*ResponseData, error) {
	var code ErrorCode = FormatErr.Format("hello, I'm error")
	return ToBadResponse(code)
}

func to3() (*ResponseData, error) {
	var code = InternalServerErr.Wrap(errors.New("OS Error"))
	fmt.Println(code.ErrorStack)
	return ToBadResponse(code)
}

func main() {
	var res []byte
	var response *ResponseData
	var err error

	response, err = to1()
	res, err = json.Marshal(response)
	fmt.Printf("data=%v, err=%v\n", string(res), err)

	response, err = to2()
	res, err = json.Marshal(response)
	fmt.Printf("data=%v, err=%v\n", string(res), err)

	response, err = to3()
	res, err = json.Marshal(response)
	fmt.Printf("data=%v, err=%v\n", string(res), err)
}
