package binder

import (
	"context"
	"fmt"
)

func (b *Binder) BindHeader() {
	params := b.Request.Header

	ctx := b.Request.Context()
	for key := range params {
		ctx = context.WithValue(ctx, fmt.Sprintf("%s%s", contextHeaderPrefix, key), params.Get(key))
	}

	*b.Request = *b.Request.WithContext(ctx)
}
