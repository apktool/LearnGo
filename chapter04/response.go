package main

import (
	"errors"
	"fmt"
)

type ResponseData struct {
	Code    int64  `json:"code"`
	Message string `json:"data"`
	Data    any    `json:"message"`
}

func ToBadResponse(err error) (*ResponseData, error) {
	resp := ResponseData{Data: ""}
	if err == nil {
		return &resp, fmt.Errorf("can't found error info")
	}

	var errCode ErrorCode
	if errors.As(err, &errCode) {
		resp.Code = errCode.Code
		resp.Message = errCode.Message
	} else {
		resp.Code = CommonErr.Code
		resp.Message = err.Error()
	}

	return &resp, nil
}

func ToOkResponse(data any) (*ResponseData, error) {
	resp := ResponseData{
		Code:    Success.Code,
		Message: Success.Message,
		Data:    data,
	}

	return &resp, nil
}
