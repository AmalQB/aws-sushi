package api

import (
	"net/http"

	"github.com/microservices-today/aws-sushi/json"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write(json.Message("OK", "I'm Still alive! :)"))
}
