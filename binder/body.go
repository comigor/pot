package binder

import (
	"bytes"
	"fmt"
	"io"

	"github.com/afikrim/pot/binder/option"
	"github.com/afikrim/pot/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (d *Decoder) BindBody(v interface{}) error {
	// check content type of request
	switch d.Request.Header.Get("Content-Type") {
	case option.ContentTypeApplicationJson.String():
		body, err := io.ReadAll(d.Request.Body)
		if err != nil {
			return err
		}

		return protojson.Unmarshal(body, v.(protoreflect.ProtoMessage))
	default:
		return fmt.Errorf("content-type is not supported, %w", errors.ErrGeneralUnsupportedMediaType)
	}
}

func (e *Encoder) BindBody(v interface{}) error {
	var content []byte
	switch e.Opts.ContentType {
	case option.ContentTypeApplicationJson:
		var err error
		content, err = protojson.Marshal(v.(protoreflect.ProtoMessage))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("content-type is not supported, %w", errors.ErrGeneralUnsupportedMediaType)
	}

	e.Request.Body = io.NopCloser(bytes.NewBuffer(content))
	return nil
}
