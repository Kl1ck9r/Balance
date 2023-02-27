package error

import (
	"encoding/json"
	"fmt"
	"io"
)

type JsonError struct {
	Ok     bool   `json:"ok" default:"false"`
	Err    string `json:"error" default:"Something  went wrong"`
	Status int    `json:"status" default:"501"`
}

func (JsonError) WriteJsError(wrt io.Writer, err error, status int) (int, error) {
	msg := err.Error()
	jsonBuf := &JsonError{
		Ok:     false,
		Err:    msg,
		Status: status,
	}

	decode, err := json.Marshal(&jsonBuf)
	if err != nil {
		return 0, fmt.Errorf("Failed marshal structure in json format: %v", err.Error())
	}

	return wrt.Write(decode)
}

func (JsonError) WriteJsString(wrt io.Writer, msg string, status int) (int, error) {
	jsonBuf := &JsonError{
		Ok:     true,
		Err:    msg,
		Status: status,
	}

	decode, err := json.Marshal(&jsonBuf)
	if err != nil {
		return 0, fmt.Errorf("Failed marshal structure in json format: %v", err.Error())
	}

	return wrt.Write(decode)
}
