package option

import "time"

type Options struct {
	Headers     map[string]any
	ContentType ContentType
	Timeout     time.Duration
	Operation   string
}

type Option func(*Options)

func New(options ...Option) *Options {
	o := Options{
		Headers:     make(map[string]any),
		ContentType: ContentTypeApplicationJson,
		Timeout:     time.Second * 10,
	}

	for _, option := range options {
		option(&o)
	}

	return &o
}

func WithHeader(key string, val any) Option {
	return func(o *Options) {
		o.Headers[key] = val
	}
}

func WithHeaders(headers map[string]any) Option {
	return func(o *Options) {
		o.Headers = headers
	}
}

func WithContentType(contentType ContentType) Option {
	return func(o *Options) {
		o.ContentType = contentType
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func WithOperation(operation string) Option {
	return func(o *Options) {
		o.Operation = operation
	}
}
