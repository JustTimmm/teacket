package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := "localhost:8080"
	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			panic(fmt.Errorf("impossible to write the response %v\n", err))
		}
	})

	s := &http.Server{
		Addr:      addr,
		Handler:   mux,
		Protocols: p,
	}

	fmt.Printf("server started on %s\n", addr)
	err := s.ListenAndServe()

	if err != nil {
		panic(fmt.Errorf("impossible to start the server %v\n", err))
	}
}
