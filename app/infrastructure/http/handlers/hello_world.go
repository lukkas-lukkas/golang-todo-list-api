package infrastructure_http_handlers

import "net/http"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from infra handlers!"))
}
