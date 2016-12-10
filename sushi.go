package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/microservices-today/aws-sushi/api"
	"github.com/microservices-today/aws-sushi/conf"
)

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		log.Println(req.URL)
		h.ServeHTTP(rw, req)
	})
}

func main() {
	http.HandleFunc("/health", api.Health)
	http.HandleFunc("/", api.Post)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.SushiPort), logHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}