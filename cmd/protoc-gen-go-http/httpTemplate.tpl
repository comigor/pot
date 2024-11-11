{{$svcType := .ServiceType}}
{{$svcName := .ServiceName}}

const (
{{- range .MethodSets}}
  Operation_{{$svcType}}_{{.OriginalName}} = "/{{$svcName}}/{{.OriginalName}}"
{{- end}}
)

type {{$svcType}}HTTPServer interface {
{{- range .MethodSets}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(ctx context.Context, in *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

func Register{{$svcType}}HTTPServer(srv {{$svcType}}HTTPServer) http.Handler {
  return pot.RegisterService(&_{{$svcType}}_HTTP_ServiceDesc, srv)
}

{{range .Methods}}
func _{{$svcType}}_{{.Name}}{{.Num}}_HTTP_Handler(ctx context.Context, srv interface{}, dec pot.DecoderFunc, middleware pot.MiddlewareFunc) (interface{}, error) {
  in := new({{.Request}})
  if err := dec(in); err != nil {
    return nil, err
  }
  if middleware == nil {
    return srv.({{$svcType}}HTTPServer).{{.Name}}(ctx, in)
  }
  h := middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.({{$svcType}}HTTPServer).{{.Name}}(ctx, req.(*{{.Request}}))
  })
  return h(ctx, in)
}
{{end}}

var _{{$svcType}}_HTTP_ServiceDesc = pot.ServiceDescriptor{
  ServiceName: "{{$svcName}}",
  Methods: []pot.MethodDescriptor{
    {{- range .Methods}}
    {
      MethodName: "{{.Name}}",
      HttpMethod: "{{.Method}}",
      HttpPath: "{{.Path}}",
      Handler: _{{$svcType}}_{{.Name}}{{.Num}}_HTTP_Handler,
    },
    {{- end}}
  },
}

type {{$svcType}}HTTPClient interface {
{{- range .MethodSets}}
	{{.Name}}(ctx context.Context, in *{{.Request}}, opts ...option.Option) (*{{.Reply}}, error)
{{- end}}
}

type {{$svcType}}HTTPClientImpl struct{
	cc *http.Client
}

func New{{$svcType}}HTTPClient (client *http.Client) {{$svcType}}HTTPClient {
	return &{{$svcType}}HTTPClientImpl{client}
}

{{range .MethodSets}}
func (c *{{$svcType}}HTTPClientImpl) {{.Name}}(ctx context.Context, in *{{.Request}}, opts ...option.Option) (*{{.Reply}}, error) {
	out := new({{.Reply}})
	pattern := "{{.Path}}"
  req, err := http.NewRequest("{{.Method}}", pattern, nil)
  if err != nil {
    return nil, err
  }
  if err = binder.NewRequestEncoder(req, opts...).Bind(in); err != nil {
      return nil, err
  }
  req = req.WithContext(ctx)
  res, err := c.cc.Do(req)
	if err != nil {
		return nil, err
	}
  defer res.Body.Close()
	dec := &binder.ResponseDecoder{Response: res}
	if err := errors.ErrorMap[res.StatusCode]; err != nil {
		customErr := new(errors.Error)
    if err := dec.BindBody(customErr); err != nil {
      return nil, err
    }
		return nil, customErr
	}
	if err := dec.BindBody(out); err != nil {
    return nil, err
  }
	return out, nil
}
{{end}}
