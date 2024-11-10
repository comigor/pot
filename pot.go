package pot

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/afikrim/pot/binder"
	"github.com/afikrim/pot/errors"
	"github.com/go-chi/chi/v5"
)

type (
	HandlerFunc       func(ctx context.Context, req interface{}) (interface{}, error)
	MiddlewareFunc    func(HandlerFunc) (interface{}, error)
	DecoderFunc       func(req interface{}) error
	MethodHandlerFunc func(ctx context.Context, srv interface{}, dec DecoderFunc, middleware MiddlewareFunc) (interface{}, error)

	MethodDescriptor struct {
		MethodName string

		HttpMethod string
		HttpPath   string
		Handler    MethodHandlerFunc
	}

	ServiceDescriptor struct {
		ServiceName string
		HandlerType interface{}
		Method      []MethodDescriptor
	}
)

func NewDecoderFunc(r *http.Request) DecoderFunc {
	dec := binder.Decoder{Request: r}

	return func(req interface{}) error {
		return dec.Bind(req)
	}
}

func httpHandlerWrapper(impl interface{}, handler MethodHandlerFunc) http.HandlerFunc {
	type ErrResp struct {
		Message string `json:"message"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		decoder := NewDecoderFunc(r)
		out, err := handler(r.Context(), impl, decoder, nil)
		if err != nil {
			status, err := errors.ParseErr(err)
			rw.WriteHeader(status)
			json.NewEncoder(rw).Encode(ErrResp{Message: err.Error()})
			return
		}

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(out)
	}
}

func RegisterService(desc *ServiceDescriptor, impl interface{}) http.Handler {
	if impl != nil {
		ht := reflect.TypeOf(desc.HandlerType).Elem()
		st := reflect.TypeOf(impl)
		if !st.Implements(ht) {
			log.Fatalf("pot: RegisterService found the handler of type %v that does not satisfy %v", st, ht)
		}
	}

	router := chi.NewRouter()
	for _, method := range desc.Method {
		switch method.HttpMethod {
		case http.MethodGet:
			router.Get(method.HttpPath, httpHandlerWrapper(impl, method.Handler))
		case http.MethodPost:
			router.Post(method.HttpPath, httpHandlerWrapper(impl, method.Handler))
		case http.MethodPut:
			router.Put(method.HttpPath, httpHandlerWrapper(impl, method.Handler))
		case http.MethodPatch:
			router.Patch(method.HttpPath, httpHandlerWrapper(impl, method.Handler))
		case http.MethodDelete:
			router.Delete(method.HttpPath, httpHandlerWrapper(impl, method.Handler))
		default:
			panic("pot: RegisterService found unsupported HTTP method: " + method.HttpMethod)
		}
	}

	return router
}
