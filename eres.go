package eres

import (
	"net/http"

	"github.com/ilgooz/httpres"
)

type Response struct {
	Message string `json:"message"`
	Fields  Fields `json:"fields"`

	w http.ResponseWriter
}

type Fields map[string]string

func New(w http.ResponseWriter) *Response {
	return &Response{
		Fields: make(Fields),
		w:      w,
	}
}

func (res *Response) SetMessage(message string) *Response {
	res.Message = message
	return res
}

func (res *Response) SetFields(fields Fields) *Response {
	res.Fields = fields
	return res
}

func (res *Response) AddField(key, val string) *Response {
	res.Fields[key] = val
	return res
}

func (res *Response) Send() {
	if res.Message == "" {
		res.Message = "Invalid Data"
	}
	httpres.Json(res.w, http.StatusBadRequest, res)
}

func (res *Response) WeakSend() bool {
	if !res.HasError() {
		return false
	}
	res.Send()
	return true
}

func (res *Response) HasError() bool {
	return len(res.Fields) > 0 || res.Message != ""
}
