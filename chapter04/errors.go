package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type ErrorCode struct {
	Code       int64  `json:"code"`
	Message    string `json:"message,omitempty"`
	ErrorStack error  `json:"err_stack,omitempty"`
}

func (e ErrorCode) Error() string {
	data := make(map[string]any)
	data["code"] = e.Code
	data["message"] = e.Message
	if e.ErrorStack != nil {
		if _, ok := e.ErrorStack.(json.Marshaler); ok {
			data["err_stack"] = e.ErrorStack
		} else {
			data["err_stack"] = e.ErrorStack.Error()
		}
	}

	message, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("json marshal failed, %v", err)
	}

	return string(message)
}

func (e ErrorCode) Format(m ...any) ErrorCode {
	eo := e
	eo.Message = fmt.Sprintf(eo.Message, m...)
	return eo
}

func (e ErrorCode) Wrap(err error, m ...any) ErrorCode {
	eo := e.Format(m...)
	eo.ErrorStack = err
	return eo
}

func (e ErrorCode) Is(err error) bool {
	var eo ErrorCode
	ok := errors.As(err, &eo)
	return ok && e.Code == eo.Code
}

func (e ErrorCode) IsNil() bool {
	return e == ErrorCode{}
}
