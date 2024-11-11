package binder

import (
	"net/http"

	"github.com/afikrim/pot/binder/option"
)

type RequestEncoder struct {
	Opts    *option.Options
	Request *http.Request
}

type ResponseEncoder struct {
	ResponseWriter http.ResponseWriter
}

func NewRequestEncoder(r *http.Request, opts ...option.Option) *RequestEncoder {
	return &RequestEncoder{
		Opts:    option.New(opts...),
		Request: r,
	}
}
