package service

const (
	Name = "HelloService"
)

type Service interface {
	Hello(*Request, *Response) error
}
