package transport

import (
	"context"
	"example/restfulapi/api/endpoint"
	"example/restfulapi/models"
	"example/restfulapi/repository"
	"fmt"
	"net/http"

	"encoding/json"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewBookHttpHandler(b repository.BookRepo) http.Handler {
	router := mux.NewRouter()

	ep := endpoint.NewBookEndpoints(b)

	router.Methods(http.MethodGet).Path("/api/v1/findbookall").Handler(httptransport.NewServer(
		ep.FindBookAll,
		decodeFindBookAllRequest,
		encodeResponse,
	))

	router.Methods(http.MethodGet).Path("/api/v1/findbookbyid/{id}").Handler(httptransport.NewServer(
		ep.FindBookById,
		decodeFindBookByIdRequest,
		encodeResponse,
	))

	router.Methods(http.MethodPost).Path("/api/v1/createbook").Handler(httptransport.NewServer(
		ep.CreateBook,
		decodeCreateBookRequest,
		encodeResponse,
	))

	router.Methods(http.MethodPost).Path("/api/v1/updatebook").Handler(httptransport.NewServer(
		ep.UpdateBook,
		decodeUpdateBookRequest,
		encodeResponse,
	))

	router.Methods(http.MethodGet).Path("/api/v1/deletebookbyid/{id}").Handler(httptransport.NewServer(
		ep.DeleteBook,
		decodeDeleteBookByIdRequest,
		encodeResponse,
	))

	router.Use(middleware)

	return router
}

func decodeFindBookAllRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req models.GetReq
	if r.ContentLength == 0 {
		return req, nil
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeFindBookByIdRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	id := mux.Vars(r)["id"]

	return id, nil
}

func decodeCreateBookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req models.Book
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUpdateBookRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req models.Book
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeDeleteBookByIdRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	id := mux.Vars(r)["id"]

	return id, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		return nil
	}
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before handler")
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Running after handler")
	})
}
