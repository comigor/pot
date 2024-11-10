package binder

import (
	"context"
	"fmt"
)

func (d *Decoder) BindHeader() {
	params := d.Request.Header

	ctx := d.Request.Context()
	for key := range params {
		ctx = context.WithValue(ctx, fmt.Sprintf("%s%s", contextHeaderPrefix, key), params.Get(key))
	}

	*d.Request = *d.Request.WithContext(ctx)
}

func (e *Encoder) BindHeader() {
	for key, val := range e.Opts.Headers {
		e.Request.Header.Set(key, fmt.Sprintf("%v", val))
	}
}
