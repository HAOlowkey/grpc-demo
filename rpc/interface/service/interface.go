package service

const (
	Name = "HelloService"
)

type Service interface {
	Hello(string, *string) error
}
