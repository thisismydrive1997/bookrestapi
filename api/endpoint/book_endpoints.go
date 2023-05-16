package endpoint

import (
	"context"
	"example/restfulapi/models"
	"example/restfulapi/repository"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type BookEndpoints struct {
	FindBookAll  endpoint.Endpoint
	FindBookById endpoint.Endpoint
	CreateBook   endpoint.Endpoint
	UpdateBook   endpoint.Endpoint
	DeleteBook   endpoint.Endpoint
}

func NewBookEndpoints(b repository.BookRepo) BookEndpoints {
	return BookEndpoints{
		FindBookAll:  MakeFindBookAllEndpoint(b),
		FindBookById: MakeFindBookByIdEndpoint(b),
		CreateBook:   MakeCreateBookEndpoint(b),
		UpdateBook:   MakeUpdateBookEndpoint(b),
		DeleteBook:   MakeDeleteBookByIdEndpoint(b),
	}
}

func MakeFindBookAllEndpoint(b repository.BookRepo) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		resp, err := b.FindBookAll()
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func MakeFindBookByIdEndpoint(b repository.BookRepo) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		resp, err := b.FindBookById(fmt.Sprint(request))
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

func MakeCreateBookEndpoint(b repository.BookRepo) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.Book)
		err := b.CreateBook(req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func MakeUpdateBookEndpoint(b repository.BookRepo) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.Book)
		err := b.UpdateBook(req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func MakeDeleteBookByIdEndpoint(b repository.BookRepo) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := b.DeleteBookById(fmt.Sprint(request))
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
