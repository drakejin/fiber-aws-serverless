package todo

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/drakejin/fiber-aws-serverless/internal/container"
	"github.com/drakejin/fiber-aws-serverless/model"
)

type Service struct {
	Container *container.Container
}

func New(c *container.Container) *Service {
	return &Service{
		Container: c,
	}
}

type In struct {
	Todo *model.Todo `json:"Todo"`
}

func (i *In) JSON() ([]byte, error) {
	return jsoniter.Marshal(i)
}

type Out struct {
	Todo  *model.Todo   `json:"Todo,omitempty"`
	Todos []*model.Todo `json:"Todos,omitempty"`
}

func (o *Out) JSON() ([]byte, error) {
	return jsoniter.Marshal(o)
}
