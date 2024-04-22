package main

import (
	"fmt"
	"net/http"
)

type Storage interface {
	Get(id int) (any, error)
	Put(val int, n any) error
}
type FooStorage struct {
}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}
func (s *FooStorage) Put(id int, n any) error {
	return nil
}

type Server struct {
	store Storage
}

func main() {
	s := &Server{
		store: &FooStorage{},
	}
	s.store.Get(1)
	http.HandleFunc("/", handleprint)
	http.ListenAndServe(":9090", nil)

}
func handleprint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello woorld")

}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello woorld")
// })
// http.ListenAndServe(":9090", nil)
