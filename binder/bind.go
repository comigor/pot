package binder

import "net/http"

type Binder struct {
	Request *http.Request

	Pattern string
}

func (b *Binder) Bind(v interface{}) error {
	hasBody := hasBody(b.Request)

	b.BindHeader()
	if err := b.BindParams(b.Pattern, v); err != nil {
		return err
	}
	if err := b.BindQuery(v); err != nil {
		return err
	}

	if !hasBody {
		return nil
	}

	if err := b.BindBody(v); err != nil {
		return err
	}

	return nil
}
