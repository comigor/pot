package binder

import (
	"fmt"
	"io"

	"github.com/afikrim/pot/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (b *Binder) BindBody(v interface{}) error {
	// check content type of request
	switch b.Request.Header.Get("Content-Type") {
	case MimeApplicationJSON:
		body, err := io.ReadAll(b.Request.Body)
		if err != nil {
			return err
		}

		return protojson.Unmarshal(body, v.(protoreflect.ProtoMessage))
	default:
		return fmt.Errorf("content-type is not supported, %w", errors.ErrGeneralUnsupportedMediaType)
	}
}
