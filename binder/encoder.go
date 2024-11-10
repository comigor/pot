package binder

import (
	"net/http"

	"github.com/afikrim/pot/binder/option"
)

type Encoder struct {
	Opts    *option.Options
	Request *http.Request
}

func NewEncoder(r *http.Request, opts ...option.Option) *Encoder {
	return &Encoder{
		Opts:    option.New(opts...),
		Request: r,
	}
}
