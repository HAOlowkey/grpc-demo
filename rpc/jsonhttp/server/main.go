package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/HAOlowkey/grpc-demo/rpc/jsonhttp/service"
)

var _ (service.Service) = (*HelloService)(nil)

type HelloService struct {
}

type HTTPReaderWriterCloser struct {
	io.ReadCloser
	io.Writer
}

func newHttpReaderWriterCloser(w http.ResponseWriter, r *http.Request) *HTTPReaderWriterCloser {
	return &HTTPReaderWriterCloser{r.Body, w}
}

func (e *HelloService) Hello(req string, resp *string) error {
	*resp = fmt.Sprintf("Hello,%s", req)
	return nil
}

func main() {
	err := rpc.RegisterName(service.Name, new(HelloService))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		conn := newHttpReaderWriterCloser(w, r)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":8090", nil)
}
