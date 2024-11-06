{{$svcType := .ServiceType}}
{{$svcName := .ServiceName}}

const (
{{- range .MethodSets}}
  Operation_{{$svcType}}_{{.OriginalName}} = "/{{$svcName}}/{{.OriginalName}}"
{{- end}}

{{- range .MethodSets}}
  {{$svcType}}_{{.OriginalName}}_FullPathName = "{{.Method}} {{.Path}}"
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

{{range .Methods}}
func _{{$svcType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv {{$svcType}}HTTPServer, middleware func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
    binder := &binder.Binder{
      Request: r,
      Pattern: "{{.Path}}",
    }

		var in {{.Request}}
		if err := binder.Bind(&in); err != nil {
			return err
		}

    *r = *r.WithContext(context.WithValue(ctx, "http_operation", Operation_{{$svcType}}_{{.OriginalName}}))
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.{{.Name}}(ctx, req.(*{{.Request}}))
		})
    h := middleware(func(rw http.ResponseWriter, r *http.Request) {
      
    })

		out, err := h(ctx, &in)
		if err != nil {
			return err
		}

		reply := out.(*{{.Reply}})

		return ctx.Result(200, reply{{.ResponseBody}})
	}
}
{{end}}

type {{$svcType}}HTTPClient interface {
{{- range .MethodSets}}
	{{.Name}}(ctx context.Context, req *{{.Request}}, opts ...http.CallOption) (rsp *{{.Reply}}, err error)
{{- end}}
}

type {{$svcType}}HTTPClientImpl struct{
	cc *http.Client
}

func New{{$svcType}}HTTPClient (client *http.Client) {{$svcType}}HTTPClient {
	return &{{$svcType}}HTTPClientImpl{client}
}

{{range .MethodSets}}
func (c *{{$svcType}}HTTPClientImpl) {{.Name}}(ctx context.Context, in *{{.Request}}, opts ...http.CallOption) (*{{.Reply}}, error) {
	var out {{.Reply}}
	pattern := "{{.Path}}"
	path := binding.EncodeURL(pattern, in, {{not .HasBody}})
	opts = append(opts, http.Operation(Operation{{$svcType}}{{.OriginalName}}))
	opts = append(opts, http.PathTemplate(pattern))
	{{if .HasBody -}}
	err := c.cc.Invoke(ctx, "{{.Method}}", path, in{{.Body}}, &out{{.ResponseBody}}, opts...)
	{{else -}}
	err := c.cc.Invoke(ctx, "{{.Method}}", path, nil, &out{{.ResponseBody}}, opts...)
	{{end -}}
	if err != nil {
		return nil, err
	}
	return &out, nil
}
{{end}}
