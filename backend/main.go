package main

import (
	"fmt"
	"net/http"
	"teacket/database"
	"teacket/gen/api/v1/apiv1connect"
	"teacket/handler"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Database connected")

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

	path, ticketHandler := apiv1connect.NewTicketServiceHandler(
		handler.TicketHandler{
			DB: db,
		},
	)

	fmt.Println(path)

	mux.Handle(path, ticketHandler)

	s := &http.Server{
		Addr:      addr,
		Handler:   mux,
		Protocols: p,
	}

	fmt.Printf("server started on %s\n", addr)
	err = s.ListenAndServe()

	if err != nil {
		panic(fmt.Errorf("impossible to start the server %v\n", err))
	}
}
